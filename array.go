package owl

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

type ArraySchema struct {
	schema *AnySchema
	of     Schema
}

func Array(schema Schema) *ArraySchema {
	self := &ArraySchema{Any(), schema}
	self.Rule("type", self.Type(), func(value reflect.Value) (any, error) {
		if !value.IsValid() {
			return nil, nil
		}

		if value.Kind() != reflect.Array && value.Kind() != reflect.Slice {
			return value.Interface(), errors.New("must be an array/slice")
		}

		return value.Interface(), nil
	})

	self.Rule("items", schema, nil)
	return self
}

func (self ArraySchema) Type() string {
	return fmt.Sprintf("array[%s]", self.of.Type())
}

func (self *ArraySchema) Rule(key string, value any, rule RuleFn) *ArraySchema {
	self.schema.Rule(key, value, rule)
	return self
}

func (self *ArraySchema) Message(message string) *ArraySchema {
	self.schema.Message(message)
	return self
}

func (self *ArraySchema) Required() *ArraySchema {
	self.schema.Required()
	return self
}

func (self *ArraySchema) Min(min int) *ArraySchema {
	return self.Rule("min", min, func(value reflect.Value) (any, error) {
		if !value.IsValid() {
			return nil, nil
		}

		if value.Len() < min {
			return value.Interface(), fmt.Errorf("must have length of at least %d", min)
		}

		return value.Interface(), nil
	})
}

func (self *ArraySchema) Max(max int) *ArraySchema {
	return self.Rule("max", max, func(value reflect.Value) (any, error) {
		if !value.IsValid() {
			return nil, nil
		}

		if value.Len() > max {
			return value.Interface(), fmt.Errorf("must have length of at most %d", max)
		}

		return value.Interface(), nil
	})
}

func (self ArraySchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(self.schema)
}

func (self ArraySchema) Validate(value any) error {
	return self.validate("", reflect.Indirect(reflect.ValueOf(value)))
}

func (self ArraySchema) validate(key string, value reflect.Value) error {
	if err := self.schema.validate(key, value); err != nil {
		return err
	}

	if !value.IsValid() {
		return nil
	}

	if value.Kind() == reflect.Interface {
		value = value.Elem()
	}

	err := NewEmptyError("items", key)

	for i := 0; i < value.Len(); i++ {
		item := reflect.Indirect(value.Index(i))

		if item.Kind() == reflect.Interface {
			item = item.Elem()
		}

		if e := self.of.validate(strconv.Itoa(i), item); e != nil {
			err = err.Add(e)
		}
	}

	if len(err.Errors) > 0 {
		return err
	}

	return nil
}
