package owl

import (
	"fmt"
	"reflect"
	"strings"
)

type owl struct {
	rules map[string]Rule
}

func New() *owl {
	return &owl{
		rules: map[string]Rule{
			"required": Required{},
			"pattern":  Pattern{},
		},
	}
}

func (self *owl) AddRule(name string, rule Rule) *owl {
	self.rules[name] = rule
	return self
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
