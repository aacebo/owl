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
	i := value.MapRange()

	for i.Next() {
		k := i.Key()
		v := reflect.Indirect(i.Value())

		if v.Kind() == reflect.Interface {
			v = v.Elem()
		}

		if e := self.validateMapField(k, v); e != nil {
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

	for i := 0; i < value.NumField(); i++ {
		field := value.Type().Field(i)

		if e := self.validateStructField(field, value.Field(i)); e != nil {
			err = err.Add(e)
		}
	}

	if len(err.Errors) > 0 {
		return err
	}

	return nil
}

func (self ObjectSchema) validateMapField(key reflect.Value, value reflect.Value) error {
	schema, exists := self.fields[key.String()]

	if !exists {
		return errors.New("schema not found")
	}

	if err := schema.validate(key.String(), value); err != nil {
		return err
	}

	return nil
}

func (self ObjectSchema) validateStructField(field reflect.StructField, value reflect.Value) error {
	name := field.Tag.Get("json")

	if name == "" {
		name = field.Name
	}

	schema, exists := self.fields[name]

	if !exists {
		return errors.New("schema not found")
	}

	if err := schema.validate(name, value); err != nil {
		return err
	}

	return nil
}
