package owl

import "reflect"

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
