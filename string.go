package owl

import (
	"encoding/json"
	"fmt"
	"net/mail"
	"net/url"
	"reflect"
	"regexp"
)

type StringSchema struct {
	schema *AnySchema
}

func String() *StringSchema {
	return &StringSchema{Any()}
}

func (self StringSchema) Type() string {
	return "string"
}

func (self *StringSchema) Rule(key string, value any, rule RuleFn) *StringSchema {
	self.schema.Rule(key, value, rule)
	return self
}

func (self *StringSchema) Required() *StringSchema {
	self.schema.Required()
	return self
}

func (self *StringSchema) Enum(values ...string) *StringSchema {
	newValues := make([]any, len(values))

	for i, value := range values {
		newValues[i] = value
	}

	self.schema.Enum(newValues...)
	return self
}

func (self *StringSchema) Min(min int) *StringSchema {
	return self.Rule("min", min, func(value reflect.Value) (any, error) {
		if !value.IsValid() {
			return value, nil
		}

		if value.Len() < min {
			return value, fmt.Errorf("must have length of at least %d", min)
		}

		return value, nil
	})
}

func (self *StringSchema) Max(max int) *StringSchema {
	return self.Rule("max", max, func(value reflect.Value) (any, error) {
		if !value.IsValid() {
			return value, nil
		}

		if value.Len() > max {
			return value, fmt.Errorf("must have length of at most %d", max)
		}

		return value, nil
	})
}

func (self *StringSchema) Regex(re *regexp.Regexp) *StringSchema {
	return self.Rule("regex", re.String(), func(value reflect.Value) (any, error) {
		if !value.IsValid() {
			return value, nil
		}

		if !re.MatchString(value.String()) {
			return value, fmt.Errorf("must match regex %s", re.String())
		}

		return value, nil
	})
}

func (self *StringSchema) Email() *StringSchema {
	return self.Rule("email", true, func(value reflect.Value) (any, error) {
		if !value.IsValid() {
			return value, nil
		}

		if _, err := mail.ParseAddress(value.String()); err != nil {
			return value, fmt.Errorf(
				`"%s" does not match email format`,
				value.String(),
			)
		}

		return value, nil
	})
}

func (self *StringSchema) UUID() *StringSchema {
	return self.Rule("uuid", true, func(value reflect.Value) (any, error) {
		if !value.IsValid() {
			return value, nil
		}

		if !uuid.MatchString(value.String()) {
			return value, fmt.Errorf(
				`"%s" does not match uuid format`,
				value.String(),
			)
		}

		return value, nil
	})
}

func (self *StringSchema) URL() *StringSchema {
	return self.Rule("url", true, func(value reflect.Value) (any, error) {
		if !value.IsValid() {
			return value, nil
		}

		if _, err := url.ParseRequestURI(value.String()); err != nil {
			return value, fmt.Errorf(
				`"%s" does not match url format`,
				value.String(),
			)
		}

		return value, nil
	})
}

func (self StringSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(self.schema)
}

func (self StringSchema) Validate(value any) error {
	return self.validate("<root>", reflect.Indirect(reflect.ValueOf(value)))
}

func (self StringSchema) validate(key string, value reflect.Value) error {
	if err := self.schema.validate(key, value); err != nil {
		return err
	}

	if value.IsValid() && value.Kind() != reflect.String {
		return newError(key, "must be a string")
	}

	return nil
}
