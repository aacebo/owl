package owl

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

type IntSchema struct {
	schema *AnySchema
}

func Int() *IntSchema {
	self := &IntSchema{Any()}
	self.Rule("type", self.Type(), func(rule Rule, value reflect.Value) (any, error) {
		if !value.IsValid() {
			return nil, nil
		}

		if value.CanConvert(reflect.TypeFor[int]()) {
			value = value.Convert(reflect.TypeFor[int]())
		}

		if value.Kind() != reflect.Int {
			return value.Interface(), errors.New("must be an int")
		}

		return value.Interface(), nil
	})

	return self
}

func (self IntSchema) Type() string {
	return "int"
}

func (self *IntSchema) Rule(key string, value any, rule RuleFn) *IntSchema {
	self.schema.Rule(key, value, rule)
	return self
}

func (self *IntSchema) Message(message string) *IntSchema {
	self.schema.Message(message)
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
	return self.Rule("min", min, func(rule Rule, value reflect.Value) (any, error) {
		if !value.IsValid() {
			return nil, nil
		}

		if value.Int() < int64(min) {
			return value.Interface(), fmt.Errorf("must have value of at least %d", min)
		}

		return value.Interface(), nil
	})
}

func (self *IntSchema) Max(max int) *IntSchema {
	return self.Rule("max", max, func(rule Rule, value reflect.Value) (any, error) {
		if !value.IsValid() {
			return nil, nil
		}

		if value.Int() > int64(max) {
			return value.Interface(), fmt.Errorf("must have value of at most %d", max)
		}

		return value.Interface(), nil
	})
}

func (self IntSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(self.schema)
}

func (self IntSchema) Validate(value any) error {
	return self.validate("", reflect.ValueOf(value))
}

func (self IntSchema) validate(key string, value reflect.Value) error {
	return self.schema.validate(key, value)
}
