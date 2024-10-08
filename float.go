package owl

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

type FloatSchema struct {
	schema *AnySchema
}

func Float() *FloatSchema {
	self := &FloatSchema{Any()}
	self.Rule("type", self.Type(), func(value reflect.Value) (any, error) {
		if !value.IsValid() {
			return nil, nil
		}

		if value.CanConvert(reflect.TypeFor[float64]()) {
			value = value.Convert(reflect.TypeFor[float64]())
		}

		if value.Kind() != reflect.Float64 {
			return value.Interface(), errors.New("must be a float")
		}

		return value.Interface(), nil
	})

	return self
}

func (FloatSchema) Type() string {
	return "float"
}

func (self *FloatSchema) Rule(key string, value any, rule RuleFn) *FloatSchema {
	self.schema.Rule(key, value, rule)
	return self
}

func (self *FloatSchema) Message(message string) *FloatSchema {
	self.schema.Message(message)
	return self
}

func (self *FloatSchema) Required() *FloatSchema {
	self.schema.Required()
	return self
}

func (self *FloatSchema) Enum(values ...float64) *FloatSchema {
	newValues := make([]any, len(values))

	for i, value := range values {
		newValues[i] = value
	}

	self.schema.Enum(newValues...)
	return self
}

func (self *FloatSchema) Min(min float64) *FloatSchema {
	return self.Rule("min", min, func(value reflect.Value) (any, error) {
		if !value.IsValid() {
			return nil, nil
		}

		if value.Float() < min {
			return value.Interface(), fmt.Errorf("must have value of at least %f", min)
		}

		return value.Interface(), nil
	})
}

func (self *FloatSchema) Max(max float64) *FloatSchema {
	return self.Rule("max", max, func(value reflect.Value) (any, error) {
		if !value.IsValid() {
			return nil, nil
		}

		if value.Float() > max {
			return value.Interface(), fmt.Errorf("must have value of at most %f", max)
		}

		return value.Interface(), nil
	})
}

func (self FloatSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(self.schema)
}

func (self FloatSchema) Validate(value any) error {
	return self.validate("", reflect.ValueOf(value))
}

func (self FloatSchema) validate(key string, value reflect.Value) error {
	return self.schema.validate(key, value)
}
