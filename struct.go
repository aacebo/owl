package owl

import (
	"fmt"
	"reflect"
)

type StructKeys map[string]Schema

type StructSchema struct {
	BaseSchema

	keys StructKeys
}

func Struct(keys StructKeys) *StructSchema {
	v := StructSchema{}
	v.keys = keys
	v.conditions = []*Condition{}

	v.Condition("must be a structure", func(v any) (any, bool) {
		value, typeOf := getTypeValue(v)
		return value.Interface(), typeOf.Kind() == reflect.Struct
	})

	return &v
}

func (self *StructSchema) Kind() reflect.Kind {
	return reflect.Struct
}

func (self *StructSchema) Validate(v any) (any, []*Error) {
	v, errors := self.BaseSchema.Validate(v)

	if len(errors) > 0 {
		return v, errors
	}

	value, _ := getTypeValue(v)

	for key, schema := range self.keys {
		field := value.FieldByName(key)

		if !field.IsZero() && field.IsValid() {
			_, errs := schema.Validate(field.Interface())
			errors = append(errors, errs...)
		} else if schema.IsRequired() {
			errors = append(
				errors,
				NewError(self.Kind().String(), fmt.Sprintf("%s is a required field", key), []string{key}),
			)
		}
	}

	return v, errors
}
