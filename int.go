package owl

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type IntSchema struct {
	schema *AnySchema
}

func Int() *IntSchema {
	return &IntSchema{Any()}
}

func (self IntSchema) Type() string {
	return "int"
}

func (self *IntSchema) Rule(key string, value any, rule RuleFn) *IntSchema {
	self.schema.Rule(key, value, rule)
	return self
}

func (self *IntSchema) Required() *IntSchema {
	self.schema.Required()
	return self
}

func (self *IntSchema) Enum(values ...int) *IntSchema {
	newValues := make([]any, len(values))

	for i, value := range values {
		newValues[i] = value
	}

	self.schema.Enum(newValues...)
	return self
}

func (self *IntSchema) Min(min int) *IntSchema {
	return self.Rule("min", min, func(value reflect.Value) (any, error) {
		if !value.IsValid() {
			return value, nil
		}

		if value.Int() < int64(min) {
			return value, fmt.Errorf("must have value of at least %d", min)
		}

		return value, nil
	})
}

func (self *IntSchema) Max(max int) *IntSchema {
	return self.Rule("max", max, func(value reflect.Value) (any, error) {
		if !value.IsValid() {
			return value, nil
		}

		if value.Int() > int64(max) {
			return value, fmt.Errorf("must have value of at most %d", max)
		}

		return value, nil
	})
}

func (self IntSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(self.schema)
}

func (self IntSchema) Validate(value any) error {
	return self.validate("<root>", reflect.Indirect(reflect.ValueOf(value)))
}

func (self IntSchema) validate(key string, value reflect.Value) error {
	if value.IsValid() && value.CanConvert(reflect.TypeFor[int]()) {
		value = value.Convert(reflect.TypeFor[int]())
	}

	if err := self.schema.validate(key, value); err != nil {
		return err
	}

	if value.IsValid() && value.Kind() != reflect.Int {
		return newError(key, "must be an integer")
	}

	return nil
}
