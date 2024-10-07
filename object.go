package owl

import (
	"encoding/json"
	"errors"
	"reflect"
)

type ObjectSchema struct {
	schema *AnySchema
	fields map[string]Schema
}

func Object() *ObjectSchema {
	self := &ObjectSchema{Any(), map[string]Schema{}}
	self.Rule("fields", self.fields, nil)
	self.Rule("type", self.Type(), func(value reflect.Value) (any, error) {
		if !value.IsValid() {
			return nil, nil
		}

		if value.Kind() != reflect.Struct && value.Kind() != reflect.Map {
			return value.Interface(), errors.New("must be an object")
		}

		return value.Interface(), nil
	})

	return self
}

func (self ObjectSchema) Type() string {
	return "object"
}

func (self *ObjectSchema) Rule(key string, value any, rule RuleFn) *ObjectSchema {
	self.schema.Rule(key, value, rule)
	return self
}

func (self *ObjectSchema) Required() *ObjectSchema {
	self.schema.Required()
	return self
}

func (self *ObjectSchema) Field(key string, schema Schema) *ObjectSchema {
	self.fields[key] = schema
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
	err := newErrorGroup(key)

	for name, schema := range self.fields {
		k := reflect.ValueOf(name)
		v := reflect.Indirect(value.MapIndex(k))

		if v.Kind() == reflect.Interface {
			v = v.Elem()
		}

		if e := schema.validate(name, v); e != nil {
			err = err.Add(e)
		}
	}

	if len(err.Errors) > 0 {
		return err
	}

	return nil
}

func (self ObjectSchema) validateStruct(key string, value reflect.Value) error {
	err := newErrorGroup(key)

	for name, schema := range self.fields {
		fieldName, exists := self.getStructFieldByName(name, value)

		if !exists {
			continue
		}

		field := value.FieldByName(fieldName)

		if e := schema.validate(name, field); e != nil {
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
