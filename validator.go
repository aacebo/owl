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
	errs         []Error
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

func (self *owl) Validate(v any) []Error {
	self.errs = []Error{}
	value := reflect.Indirect(reflect.ValueOf(v))

	self.validateStruct(
		"",
		value,
		value,
	)

	return self.errs
}

func (self *owl) validateStruct(path string, root reflect.Value, value reflect.Value) {
	if value.Kind() != reflect.Struct || !value.IsValid() {
		return
	}

	for i := 0; i < value.NumField(); i++ {
		field := value.Type().Field(i)
		self.validateField(
			fmt.Sprintf("%s/%s", path, self.getFieldName(field)),
			root,
			value,
			field,
			value.Field(i),
		)
	}
}

func (self *owl) validateField(path string, root reflect.Value, parent reflect.Value, field reflect.StructField, value reflect.Value) {
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
						self.validate(path, ctx)
						visited[dep] = true
					}
				}
			}

			ctx.rule = key
			self.validate(path, ctx)
			visited[key] = true
		}
	}

	value = ctx.CoerceValue()

	if value.Kind() == reflect.Struct {
		self.validateStruct(
			path,
			root,
			value,
		)
	}
}

func (self *owl) validate(path string, ctx *context) {
	rule, ok := self.rules[ctx.rule]

	if !ok {
		self.errs = append(self.errs, Error{
			Path:    path,
			Keyword: ctx.rule,
			Message: "not found",
		})

		return
	}

	if transform, ok := self.transforms[ctx.value.Type()]; ok {
		ctx.value = transform(ctx.value)
	}

	for _, err := range rule(ctx) {
		self.errs = append(self.errs, Error{
			Path:    path,
			Keyword: ctx.rule,
			Message: err.Error(),
		})
	}
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
		param := ""

		if len(ruleParts) == 2 {
			param = ruleParts[1]
		}

		schema[key] = param
	}

	return schema
}
