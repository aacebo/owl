package owl

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type UnionSchema struct {
	schema *AnySchema
	anyOf  []Schema
}

func Union(anyOf ...Schema) *UnionSchema {
	self := &UnionSchema{Any(), anyOf}
	self.Rule("type", self.Type(), func(value reflect.Value) (any, error) {
		for _, schema := range self.anyOf {
			e := schema.Validate(value.Interface())

			if e == nil {
				return value.Interface(), nil
			}
		}

		return value.Interface(), errors.New("must match one or more types in union")
	})

	return self
}

func (self UnionSchema) Type() string {
	anyOf := make([]string, len(self.anyOf))

	for i := 0; i < len(anyOf); i++ {
		anyOf[i] = self.anyOf[i].Type()
	}

	return fmt.Sprintf("union[%v]", strings.Join(anyOf, ","))
}

func (self *UnionSchema) Rule(key string, value any, rule RuleFn) *UnionSchema {
	self.schema.Rule(key, value, rule)
	return self
}

func (self *UnionSchema) Message(message string) *UnionSchema {
	self.schema.Message(message)
	return self
}

func (self UnionSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(self.schema)
}

func (self UnionSchema) Validate(value any) error {
	return self.validate("", reflect.Indirect(reflect.ValueOf(value)))
}

func (self UnionSchema) validate(key string, value reflect.Value) error {
	return self.schema.validate(key, value)
}
