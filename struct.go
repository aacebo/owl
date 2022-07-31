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
		if typeOf := reflect.TypeOf(v).Kind(); typeOf == reflect.Pointer {
			return *v.(*any), reflect.TypeOf(*v.(*any)).Kind() == reflect.Struct
		}

		return v, reflect.TypeOf(v).Kind() == reflect.Struct
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

func (self *StructSchema) Validate(v any) []*Error {
	errors := self.BaseSchema.Validate(&v)

	if len(errors) > 0 {
		return errors
	}

	typeOf := reflect.TypeOf(v)
	var typeValue reflect.Value

	if typeOf.Kind() != reflect.Pointer {
		typeValue = reflect.ValueOf(v)
	} else {
		typeValue = reflect.Indirect(reflect.ValueOf(v))
	}

	rr := reflect.TypeOf(typeValue.Interface())

	for i := 0; i < rr.NumField(); i++ {
		fieldName := rr.Field(i).Name
		schema, ok := self.keys[fieldName]

		if ok {
			errors = append(errors, schema.Validate(typeValue.Field(i).Interface())...)
		}
	}

	return errors
}
