package owl

import (
	"encoding/json"
	"errors"
	"reflect"
)

type BoolSchema struct {
	schema *AnySchema
}

func Bool() *BoolSchema {
	self := &BoolSchema{Any()}
	self.Rule("type", self.Type(), func(value reflect.Value) (any, error) {
		if !value.IsValid() {
			return nil, nil
		}

		if value.Kind() != reflect.Bool {
			return value.Interface(), errors.New("must be a bool")
		}

		return value.Interface(), nil
	})

	return self
}

func (self BoolSchema) Type() string {
	return "bool"
}

func (self *BoolSchema) Rule(key string, value any, rule RuleFn) *BoolSchema {
	self.schema.Rule(key, value, rule)
	return self
}

func (self *BoolSchema) Message(message string) *BoolSchema {
	self.schema.Message(message)
	return self
}

func (self *BoolSchema) Required() *BoolSchema {
	self.schema.Required()
	return self
}

func (self *BoolSchema) Enum(values ...bool) *BoolSchema {
	newValues := make([]any, len(values))

	for i, value := range values {
		newValues[i] = value
	}

	self.schema.Enum(newValues...)
	return self
}

func (self BoolSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(self.schema)
}

func (self BoolSchema) Validate(value any) error {
	return self.validate("", reflect.ValueOf(value))
}

func (self BoolSchema) validate(key string, value reflect.Value) error {
	return self.schema.validate(key, value)
}
