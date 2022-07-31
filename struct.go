package owl

import (
	"reflect"
)

type StructKeys map[string]Schema

type StructSchema struct {
	BaseSchema

	keys StructKeys
}

func Struct(keys StructKeys) *StructSchema {
	v := StructSchema{}
	v.name = "Struct"
	v.keys = keys
	v.operations = []*Operation{NewOperation("must be a structure", func(v any) (any, bool) {
		value, typeOf := getTypeValue(v)
		return value.Interface(), typeOf.Kind() == reflect.Struct
	})}

	return &v
}

func (self *StructSchema) Message(message string) *StructSchema {
	self.BaseSchema.Message(message)
	return self
}

func (self *StructSchema) Required() *StructSchema {
	self.BaseSchema.Required()
	return self
}

func (self *StructSchema) Validate(v any) (any, []*Error) {
	v, errors := self.BaseSchema.Validate(v)

	if len(errors) > 0 {
		return v, errors
	}

	value, typeOf := getTypeValue(v)

	for i := 0; i < typeOf.NumField(); i++ {
		fieldName := typeOf.Field(i).Name
		schema, ok := self.keys[fieldName]

		if ok {
			_, errs := schema.Validate(value.Field(i).Interface())
			errors = append(errors, errs...)
		}
	}

	return v, errors
}
