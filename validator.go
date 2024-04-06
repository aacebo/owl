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
	rules      map[string]rules.Rule
	formats    map[string]Formatter
	transforms map[reflect.Type]Transform
}

func New() *owl {
	return &owl{
		rules: map[string]rules.Rule{
			"default":  rules.Default,
			"required": rules.Required,
			"pattern":  rules.Pattern,
			"format":   rules.Format,
			"min":      rules.Min,
			"max":      rules.Max,
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
	}
}

func (self *owl) AddRule(name string, rule rules.Rule) *owl {
	self.rules[name] = rule
	return self
}

func (self *owl) AddType(t reflect.Type, fn Transform) *owl {
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
		schema := self.tagToSchema(tag)
		ctx.schema = schema

		for key := range schema {
			ctx.rule = key
			validate, ok := self.rules[key]

			if !ok {
				errs = append(errs, Error{
					Path:    path,
					Keyword: key,
					Message: "not found",
				})

				continue
			}

			if transform, ok := self.transforms[value.Type()]; ok {
				ctx.value = transform(value)
			}

			_errs := validate(ctx)

			for _, err := range _errs {
				errs = append(errs, Error{
					Path:    path,
					Keyword: key,
					Message: err.Error(),
				})
			}
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
