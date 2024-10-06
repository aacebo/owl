package owl

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

type TimeSchema struct {
	schema *AnySchema
	layout string
}

func Time() *TimeSchema {
	return &TimeSchema{Any(), time.RFC3339}
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
	return self.validate("<root>", reflect.Indirect(reflect.ValueOf(value)))
}

func (self TimeSchema) validate(key string, value reflect.Value) error {
	if value.IsValid() && value.Kind() != reflect.String && value.Type() != reflect.TypeFor[time.Time]() {
		return newError(key, "must be a string or time.Time")
	}

	if value.IsValid() && value.Kind() == reflect.String {
		parsed, err := time.Parse(self.layout, value.String())

		if err != nil {
			return newError(key, err.Error())
		}

		value = reflect.ValueOf(parsed)
	}

	if err := self.schema.validate(key, value); err != nil {
		return err
	}

	return nil
}
