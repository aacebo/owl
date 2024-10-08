package owl

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"slices"
)

type AnySchema []Rule

func Any() *AnySchema {
	self := &AnySchema{}
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
	i := slices.IndexFunc(*self, func(rule Rule) bool {
		return rule.Key == key
	})

	if i > -1 {
		(*self)[i] = Rule{
			Key:     key,
			Value:   value,
			Resolve: rule,
		}
	} else {
		*self = append(*self, Rule{
			Key:     key,
			Value:   value,
			Resolve: rule,
		})
	}

	return self
}

func (self *AnySchema) Message(message string) *AnySchema {
	(*self)[len(*self)-1].Message = message
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
	buf := &bytes.Buffer{}
	buf.Write([]byte{'{'})

	for i, item := range self {
		b, err := json.Marshal(item.Value)

		if err != nil {
			return nil, err
		}

		buf.WriteString(fmt.Sprintf("%q:", fmt.Sprintf("%v", item.Key)))
		buf.Write(b)

		if i < len(self)-1 {
			buf.Write([]byte{','})
		}
	}

	buf.Write([]byte{'}'})
	return buf.Bytes(), nil
}

func (self AnySchema) Validate(value any) error {
	return self.validate("", reflect.Indirect(reflect.ValueOf(value)))
}

func (self AnySchema) validate(key string, value reflect.Value) error {
	err := NewEmptyError("", key)

	for _, rule := range self {
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
