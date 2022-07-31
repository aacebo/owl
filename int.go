package owl

import (
	"fmt"
)

type IntSchema struct {
	BaseSchema
}

func Int() *IntSchema {
	v := IntSchema{}
	v.name = "Int"
	v.operations = []*Operation{NewOperation("must be an integer", func(v any) (any, bool) {
		value, ok := v.(int)
		return value, ok
	})}

	return &v
}

func (self *IntSchema) Min(min int) *IntSchema {
	self.AddOperation(NewOperation(fmt.Sprintf("must be at least %d", min), func(v any) (any, bool) {
		return v, v.(int) >= min
	}))

	return self
}

func (self *IntSchema) Max(max int) *IntSchema {
	self.AddOperation(NewOperation(fmt.Sprintf("must be at most %d", max), func(v any) (any, bool) {
		return v, v.(int) <= max
	}))

	return self
}

func (self *IntSchema) Equal(equal int) *IntSchema {
	self.AddOperation(NewOperation(fmt.Sprintf("must be equal to %d", equal), func(v any) (any, bool) {
		return v, v.(int) == equal
	}))

	return self
}

func (self *IntSchema) Message(message string) *IntSchema {
	self.BaseSchema.Message(message)
	return self
}
