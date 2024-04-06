package owl

import (
	"database/sql/driver"
	"fmt"
	"reflect"
	"strings"

	"github.com/aacebo/owl/formats"
	"github.com/aacebo/owl/rules"
	"github.com/aacebo/owl/transforms"
)

type owl struct {
	rules        map[string]rules.Rule
	formats      map[string]Formatter
	transforms   map[reflect.Type]Transform
	dependencies map[string][]string
}

func New() *owl {
	return &owl{
		rules: map[string]rules.Rule{
			"default":  rules.Default,
			"pattern":  rules.Pattern,
			"format":   rules.Format,
			"min":      rules.Min,
			"max":      rules.Max,
			"required": rules.Required,
		},
		formats: map[string]Formatter{
			"date_time": formats.DateTime,
			"email":     formats.Email,
			"ipv4":      formats.IPv4,
			"ipv6":      formats.IPv6,
			"uri":       formats.URI,
			"uuid":      formats.UUID,
		},
		transforms: map[reflect.Type]Transform{
			reflect.TypeFor[driver.Valuer](): transforms.Valuer,
		},
		dependencies: map[string][]string{
			"required": {"default"},
		},
	}
}

func (self *owl) AddRule(name string, rule rules.Rule, dependsOn ...string) *owl {
	self.rules[name] = rule
	self.dependencies[name] = dependsOn
	return self
}

func (self *owl) AddTransform(t reflect.Type, fn Transform) *owl {
	self.transforms[t] = fn
	return self
}

func (self *owl) AddFormat(name string, formatter Formatter) *owl {
	self.formats[name] = formatter
	return self
}

func (self owl) HasFormat(name string) bool {
	_, ok := self.formats[name]
	return ok
}

func (self owl) Format(name string, input string) error {
	handler := self.formats[name]
	return handler(input)
}

func (self owl) Validate(v any) []Error {
	value := reflect.Indirect(reflect.ValueOf(v))

	return self.validateStruct(
		"",
		value,
		value,
	)
}

func (self owl) validateStruct(path string, root reflect.Value, value reflect.Value) []Error {
	errs := []Error{}

	if value.Kind() != reflect.Struct || !value.IsValid() {
		return errs
	}

	for i := 0; i < value.NumField(); i++ {
		field := value.Type().Field(i)
		_errs := self.validateField(
			fmt.Sprintf("%s/%s", path, self.getFieldName(field)),
			root,
			value,
			field,
			value.Field(i),
		)

		if len(_errs) > 0 {
			errs = append(errs, _errs...)
		}
	}

	return errs
}

func (self owl) validateField(path string, root reflect.Value, parent reflect.Value, field reflect.StructField, value reflect.Value) []Error {
	errs := []Error{}
	tag := field.Tag.Get("owl")
	ctx := &context{
		root:      root,
		parent:    parent,
		value:     value,
		field:     field,
		hasFormat: self.HasFormat,
		format:    self.Format,
	}

	if tag != "" {
		visited := map[string]bool{}
		schema := self.tagToSchema(tag)
		ctx.schema = schema

		for key := range schema {
			// run dependencies first
			if dependsOn, ok := self.dependencies[key]; ok {
				for _, dep := range dependsOn {
					if _, ok := schema[dep]; ok && !visited[dep] {
						ctx.rule = dep
						_errs := self.validate(path, ctx)

						if len(_errs) > 0 {
							errs = append(errs, _errs...)
						}

						visited[dep] = true
					}
				}
			}

			ctx.rule = key
			_errs := self.validate(path, ctx)

			if len(_errs) > 0 {
				errs = append(errs, _errs...)
			}

			visited[key] = true
		}
	}

	value = ctx.CoerceValue()

	if value.Kind() == reflect.Struct {
		_errs := self.validateStruct(
			path,
			root,
			value,
		)

		if len(_errs) > 0 {
			errs = append(errs, _errs...)
		}
	}

	return errs
}

func (self owl) validate(path string, ctx *context) []Error {
	errs := []Error{}
	rule, ok := self.rules[ctx.rule]

	if !ok {
		errs = append(errs, Error{
			Path:    path,
			Keyword: ctx.rule,
			Message: "not found",
		})

		return errs
	}

	if transform, ok := self.transforms[ctx.value.Type()]; ok {
		ctx.value = transform(ctx.value)
	}

	_errs := rule(ctx)

	for _, err := range _errs {
		errs = append(errs, Error{
			Path:    path,
			Keyword: ctx.rule,
			Message: err.Error(),
		})
	}

	return errs
}

func (self owl) getFieldName(field reflect.StructField) string {
	name := field.Name

	if tag := field.Tag.Get("json"); tag != "" {
		parts := strings.Split(tag, ",")

		if parts[0] != "-" {
			name = parts[0]
		}
	}

	return name
}

func (self owl) tagToSchema(tag string) map[string]string {
	schema := map[string]string{}

	if tag == "" {
		return schema
	}

	parts := strings.Split(tag, ",")

	for _, part := range parts {
		ruleParts := strings.SplitN(part, "=", 2)
		key := ruleParts[0]
		config := ""

		if len(ruleParts) == 2 {
			config = ruleParts[1]
		}

		schema[key] = config
	}

	return schema
}
