package owl

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"time"
)

type TimeSchema struct {
	schema *AnySchema
	layout string
}

func Time() *TimeSchema {
	self := &TimeSchema{Any(), time.RFC3339}
	self.Rule("type", self.Type(), func(value reflect.Value) (any, error) {
		if !value.IsValid() {
			return nil, nil
		}

		if value.Kind() != reflect.String && value.Type() != reflect.TypeFor[time.Time]() {
			return value.Interface(), errors.New("must be a string or time.Time")
		}

		if value.Kind() == reflect.String {
			parsed, err := time.Parse(self.layout, value.String())

			if err != nil {
				return value.Interface(), err
			}

			value = reflect.ValueOf(parsed)
		}

		return value.Interface(), nil
	})

	return self
}

func (self TimeSchema) Type() string {
	return "time"
}

func (self *TimeSchema) Layout(layout string) *TimeSchema {
	self.layout = layout
	return self
}

func (self *TimeSchema) Rule(key string, value any, rule RuleFn) *TimeSchema {
	self.schema.Rule(key, value, rule)
	return self
}

func (self *TimeSchema) Message(message string) *TimeSchema {
	self.schema.Message(message)
	return self
}

func (self *TimeSchema) Required() *TimeSchema {
	self.schema.Required()
	return self
}

func (self *TimeSchema) Min(min time.Time) *TimeSchema {
	return self.Rule("min", min, func(value reflect.Value) (any, error) {
		if !value.IsValid() {
			return nil, nil
		}

		parsed := value.Interface().(time.Time)

		if parsed.Before(min) {
			return parsed, fmt.Errorf("must have value of at least %s", min.String())
		}

		return parsed, nil
	})
}

func (self *TimeSchema) Max(max time.Time) *TimeSchema {
	return self.Rule("max", max, func(value reflect.Value) (any, error) {
		if !value.IsValid() {
			return nil, nil
		}

		parsed := value.Interface().(time.Time)

		if parsed.After(max) {
			return parsed, fmt.Errorf("must have value of at most %s", max.String())
		}

		return parsed, nil
	})
}

func (self TimeSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(self.schema)
}

func (self TimeSchema) Validate(value any) error {
	return self.validate("", reflect.Indirect(reflect.ValueOf(value)))
}

func (self TimeSchema) validate(key string, value reflect.Value) error {
	return self.schema.validate(key, value)
}
