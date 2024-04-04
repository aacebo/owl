package owl

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/aacebo/owl/formats"
	"github.com/aacebo/owl/rules"
)

type owl struct {
	rules   map[string]Rule
	formats map[string]Formatter
}

func New() *owl {
	self := &owl{
		rules: map[string]Rule{
			"required": rules.Required{},
			"pattern":  rules.Pattern{},
			"min":      rules.Min{},
			"max":      rules.Max{},
		},
		formats: map[string]Formatter{
			"date_time": formats.DateTime,
			"email":     formats.Email,
			"ipv4":      formats.IPv4,
			"ipv6":      formats.IPv6,
			"uri":       formats.URI,
			"uuid":      formats.UUID,
		},
	}

	self.rules["format"] = rules.Format{
		HasFormat: self.HasFormat,
		Format:    self.Format,
	}

	return self
}

func (self *owl) AddRule(name string, rule Rule) *owl {
	self.rules[name] = rule
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
	return self.validate(
		"",
		"",
		reflect.ValueOf(nil),
		reflect.ValueOf(v),
	)
}

func (self owl) validate(path string, tag string, parent reflect.Value, value reflect.Value) []Error {
	errs := []Error{}
	value = reflect.Indirect(value)

	if value.Kind() == reflect.Pointer || value.Kind() == reflect.Interface {
		value = value.Elem()
	}

	if tag != "" {
		parts := strings.Split(tag, ",")

		for _, part := range parts {
			ruleParts := strings.SplitN(part, "=", 2)
			key := ruleParts[0]
			config := ""

			if len(ruleParts) == 2 {
				config = ruleParts[1]
			}

			rule, ok := self.rules[key]

			if !ok {
				errs = append(errs, Error{
					Path:    path,
					Keyword: key,
					Message: "not found",
				})

				continue
			}

			if !rule.Select(parent, value) {
				continue
			}

			_errs := rule.Validate(config, parent, value)

			for _, err := range _errs {
				errs = append(errs, Error{
					Path:    path,
					Keyword: key,
					Message: err.Error(),
				})
			}
		}
	}

	if value.Kind() == reflect.Struct {
		for i := 0; i < value.NumField(); i++ {
			field := value.Type().Field(i)

			_errs := self.validate(
				fmt.Sprintf("%s/%s", path, self.getFieldName(field)),
				field.Tag.Get("owl"),
				value,
				value.Field(i),
			)

			if len(_errs) > 0 {
				errs = append(errs, _errs...)
			}
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
