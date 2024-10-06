package owl

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

type AnySchema struct {
	rules []Rule
}

func Any() *AnySchema {
	self := &AnySchema{[]Rule{}}
	self.Rule("type", self.Type(), func(value reflect.Value) (any, error) {
		if !value.IsValid() {
			return nil, nil
		}

		return value.Interface(), nil
	})

	return self
}

func (self AnySchema) Type() string {
	return "any"
}

func (self *AnySchema) Rule(key string, value any, rule RuleFn) *AnySchema {
	self.rules = append(self.rules, Rule{
		Key:     key,
		Value:   value,
		Resolve: rule,
	})

	return self
}

func (self *AnySchema) Required() *AnySchema {
	return self.Rule("required", true, func(value reflect.Value) (any, error) {
		if !value.IsValid() {
			return value, errors.New("required")
		}

		return value, nil
	})
}

func (self *AnySchema) Enum(values ...any) *AnySchema {
	return self.Rule("enum", values, func(value reflect.Value) (any, error) {
		for _, v := range values {
			if value.Equal(reflect.Indirect(reflect.ValueOf(v))) {
				return value, nil
			}
		}

		return value, fmt.Errorf("must be one of %v", values)
	})
}

func (self AnySchema) MarshalJSON() ([]byte, error) {
	data := map[string]any{}

	for _, rule := range self.rules {
		data[rule.Key] = rule.Value
	}

	return json.Marshal(data)
}

func (self AnySchema) Validate(value any) error {
	return self.validate("<root>", reflect.Indirect(reflect.ValueOf(value)))
}

func (self AnySchema) validate(key string, value reflect.Value) error {
	err := newError(key, "")

	for _, rule := range self.rules {
		v, e := rule.Resolve(value)

		if e != nil {
			err = err.Add(e)
			continue
		}

		if value.CanSet() {
			value.Set(reflect.ValueOf(v))
		}
	}

	if len(err.Errors) > 0 {
		return err
	}

	return nil
}