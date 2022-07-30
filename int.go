package owl

import (
	"fmt"
)

type IntSchema struct {
	operations []*Operation
	index      int
}

func Int() *IntSchema {
	v := IntSchema{
		index: 0,
		operations: []*Operation{
			NewOperation("must be an integer", func(v any) (any, bool) {
				value, ok := v.(int)
				return value, ok
			}),
		},
	}

	return &v
}

func (self *IntSchema) Min(min int) *IntSchema {
	self.operations = append(
		self.operations,
		NewOperation(fmt.Sprintf("must be at least %d", min), func(v any) (any, bool) {
			return v, v.(int) >= min
		}),
	)

	self.index++
	return self
}

func (self *IntSchema) Max(max int) *IntSchema {
	self.operations = append(
		self.operations,
		NewOperation(fmt.Sprintf("must be at most %d", max), func(v any) (any, bool) {
			return v, v.(int) <= max
		}),
	)

	self.index++
	return self
}

func (self *IntSchema) Equal(equal int) *IntSchema {
	self.operations = append(
		self.operations,
		NewOperation(fmt.Sprintf("must be equal to %d", equal), func(v any) (any, bool) {
			return v, v.(int) == equal
		}),
	)

	self.index++
	return self
}

func (self *IntSchema) Message(v string) *IntSchema {
	if self.index > -1 {
		self.operations[self.index].Message = v
	}

	return self
}

func (self *IntSchema) Validate(v any) []*Error {
	errors := []*Error{}
	valid := false

	for _, op := range self.operations {
		if v, valid = op.Eval(v); !valid {
			errors = append(errors, NewError("Number", op.Message))
		}
	}

	return errors
}
