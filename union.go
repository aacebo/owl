package owl

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

type UnionSchema struct {
	schema *AnySchema
	anyOf  []Schema
}

func Union(anyOf ...Schema) *UnionSchema {
	return &UnionSchema{Any(), anyOf}
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

func (self UnionSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(self.schema)
}

func (self UnionSchema) Validate(value any) error {
	return self.validate("<root>", reflect.Indirect(reflect.ValueOf(value)))
}

func (self UnionSchema) validate(key string, value reflect.Value) error {
	if err := self.schema.validate(key, value); err != nil {
		return err
	}

	err := newError(key, "must match one or more types in union")

	for _, schema := range self.anyOf {
		e := schema.validate(key, value)

		if e == nil {
			return nil
		}

		err = err.Add(e)
	}

	return err
}
