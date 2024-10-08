package owl

import (
	"encoding/json"
	"errors"
	"reflect"

	"github.com/aacebo/owl/ordered_map"
)

type ObjectSchema struct {
	schema *AnySchema
	fields ordered_map.Map[string, Schema]
}

func Object() *ObjectSchema {
	self := &ObjectSchema{Any(), ordered_map.Map[string, Schema]{}}
	self.Rule("type", self.Type(), func(value reflect.Value) (any, error) {
		if !value.IsValid() {
			return nil, nil
		}

		if value.Kind() != reflect.Struct && value.Kind() != reflect.Map {
			return value.Interface(), errors.New("must be an object")
		}

		return value.Interface(), nil
	})

	self.Rule("fields", &self.fields, nil)
	return self
}

func (self ObjectSchema) Type() string {
	return "object"
}

func (self *ObjectSchema) Rule(key string, value any, rule RuleFn) *ObjectSchema {
	self.schema.Rule(key, value, rule)
	return self
}

func (self *ObjectSchema) Message(message string) *ObjectSchema {
	self.schema.Message(message)
	return self
}

func (self *ObjectSchema) Required() *ObjectSchema {
	self.schema.Required()
	return self
}

func (self *ObjectSchema) Field(key string, schema Schema) *ObjectSchema {
	self.fields.Set(key, schema)
	return self
}

func (self *ObjectSchema) Fields(fields map[string]Schema) *ObjectSchema {
	for key, schema := range fields {
		self.fields.Set(key, schema)
	}

	return self
}

func (self ObjectSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(self.schema)
}

func (self ObjectSchema) Validate(value any) error {
	return self.validate("", reflect.Indirect(reflect.ValueOf(value)))
}

func (self ObjectSchema) validate(key string, value reflect.Value) error {
	if err := self.schema.validate(key, value); err != nil {
		return err
	}

	if !value.IsValid() {
		return nil
	}

	if value.Kind() == reflect.Interface {
		value = value.Elem()
	}

	if value.Kind() == reflect.Map {
		return self.validateMap(key, value)
	}

	return self.validateStruct(key, value)
}

func (self ObjectSchema) validateMap(key string, value reflect.Value) error {
	err := NewEmptyError("fields", key)

	for _, item := range self.fields {
		k := reflect.ValueOf(item.Key)
		v := reflect.Indirect(value.MapIndex(k))

		if v.Kind() == reflect.Interface {
			v = v.Elem()
		}

		if e := item.Value.validate(item.Key, v); e != nil {
			err = err.Add(e)
		}
	}

	if len(err.Errors) > 0 {
		return err
	}

	return nil
}

func (self ObjectSchema) validateStruct(key string, value reflect.Value) error {
	err := NewEmptyError("fields", key)

	for _, item := range self.fields {
		fieldName, exists := self.getStructFieldByName(item.Key, value)

		if !exists {
			continue
		}

		field := reflect.Indirect(value.FieldByName(fieldName))

		if field.Kind() == reflect.Interface {
			field = field.Elem()
		}

		if e := item.Value.validate(item.Key, field); e != nil {
			err = err.Add(e)
		}
	}

	if len(err.Errors) > 0 {
		return err
	}

	return nil
}

func (self ObjectSchema) getStructFieldByName(name string, object reflect.Value) (string, bool) {
	if !object.IsValid() {
		return "", false
	}

	for i := 0; i < object.NumField(); i++ {
		field := object.Type().Field(i)
		tag := field.Tag.Get("json")

		if tag == "" {
			tag = field.Name
		}

		if tag == "" || tag == "-" {
			continue
		}

		if tag == name {
			return field.Name, true
		}
	}

	return "", false
}
