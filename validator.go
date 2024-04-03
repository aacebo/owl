package owl

import (
	"fmt"
	"reflect"
	"strings"
)

type Owl struct {
	rules map[string]Rule
}

func New() *Owl {
	return &Owl{
		rules: map[string]Rule{
			"pattern": Pattern{},
		},
	}
}

func (self *Owl) AddRule(name string, rule Rule) *Owl {
	self.rules[name] = rule
	return self
}

func (self Owl) Validate(v any) []Error {
	return self.validate(
		"",
		"",
		reflect.ValueOf(nil),
		reflect.ValueOf(v),
	)
}

func (self Owl) validate(path string, tag string, parent reflect.Value, value reflect.Value) []Error {
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

func (self Owl) getFieldName(field reflect.StructField) string {
	name := field.Name

	if tag := field.Tag.Get("json"); tag != "" {
		parts := strings.Split(tag, ",")

		if parts[0] != "-" {
			name = parts[0]
		}
	}

	return name
}
