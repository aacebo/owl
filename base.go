package owl

import "reflect"

type Schema interface {
	Kind() reflect.Kind
	IsRequired() bool
	Validate(v any) (any, []*Error)
}

type BaseSchema struct {
	required   bool
	conditions []*Condition
}

func (self *BaseSchema) Kind() reflect.Kind {
	return reflect.Invalid
}

func (self *BaseSchema) Condition(message string, fn func(v any) (any, bool)) *BaseSchema {
	self.conditions = append(self.conditions, NewCondition(message, fn))
	return self
}

func (self *BaseSchema) Required() *BaseSchema {
	self.required = true
	return self
}

func (self *BaseSchema) IsRequired() bool {
	return self.required
}

func (self *BaseSchema) Message(message string) *BaseSchema {
	if length := len(self.conditions); length > -1 {
		self.conditions[length-1].message = message
	}

	return self
}

func (self *BaseSchema) Validate(v any) (any, []*Error) {
	errors := []*Error{}
	valid := false

	for _, con := range self.conditions {
		if v, valid = con.Eval(v); !valid {
			errors = append(errors, NewError(self.Kind().String(), con.message, []string{}))
		}
	}

	return v, errors
}
