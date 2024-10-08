package owl

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"slices"

	"github.com/aacebo/owl/ordered_map"
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
	i := slices.IndexFunc(self.rules, func(rule Rule) bool {
		return rule.Key == key
	})

	if i > -1 {
		self.rules[i] = Rule{
			Key:     key,
			Value:   value,
			Resolve: rule,
		}
	} else {
		self.rules = append(self.rules, Rule{
			Key:     key,
			Value:   value,
			Resolve: rule,
		})
	}

	return self
}

func (self *AnySchema) Message(message string) *AnySchema {
	self.rules[len(self.rules)-1].Message = message
	return self
}

func (self *AnySchema) Required() *AnySchema {
	return self.Rule("required", true, func(value reflect.Value) (any, error) {
		if !value.IsValid() {
			return nil, errors.New("required")
		}

		return value.Interface(), nil
	})
}

func (self *AnySchema) Enum(values ...any) *AnySchema {
	return self.Rule("enum", values, func(value reflect.Value) (any, error) {
		if !value.IsValid() {
			return nil, nil
		}

		for _, v := range values {
			if value.Equal(reflect.Indirect(reflect.ValueOf(v))) {
				return value.Interface(), nil
			}
		}

		return nil, fmt.Errorf("must be one of %v", values)
	})
}

func (self AnySchema) MarshalJSON() ([]byte, error) {
	data := ordered_map.Map[string, any]{}

	for _, rule := range self.rules {
		data.Set(rule.Key, rule.Value)
	}

	return json.Marshal(data)
}

func (self AnySchema) Validate(value any) error {
	return self.validate("", reflect.Indirect(reflect.ValueOf(value)))
}

func (self AnySchema) validate(key string, value reflect.Value) error {
	err := NewEmptyError("", key)

	for _, rule := range self.rules {
		if rule.Resolve == nil {
			continue
		}

		v, e := rule.Resolve(value)

		if e != nil {
			if group, ok := e.(ErrorGroup); ok {
				for _, subErr := range group {
					err = err.Add(subErr)
				}
			} else {
				message := e.Error()

				if rule.Message != "" {
					message = rule.Message
				}

				err = err.Add(NewError(rule.Key, key, message))
				continue
			}
		}

		value = reflect.ValueOf(v)
	}

	if len(err.Errors) > 0 {
		return err
	}

	return nil
}
