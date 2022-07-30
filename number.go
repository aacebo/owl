package owl

import (
	"fmt"

	"github.com/aacebo/owl/types"
)

type NumberSchema[T types.Number] struct {
	operations []*Operation[T]
	index      int
}

func Number[T types.Number]() *NumberSchema[T] {
	v := NumberSchema[T]{index: -1}

	return &v
}

func (self *NumberSchema[T]) Min(min T) *NumberSchema[T] {
	self.operations = append(
		self.operations,
		NewOperation(fmt.Sprintf("must be at least %v", min), func(v T) bool {
			return v >= min
		}),
	)

	self.index++
	return self
}

func (self *NumberSchema[T]) Max(max T) *NumberSchema[T] {
	self.operations = append(
		self.operations,
		NewOperation(fmt.Sprintf("must be at most %v", max), func(v T) bool {
			return v <= max
		}),
	)

	self.index++
	return self
}

func (self *NumberSchema[T]) Message(v string) *NumberSchema[T] {
	if self.index > -1 {
		self.operations[self.index].Message = v
	}

	return self
}

func (self *NumberSchema[T]) Validate(v T) []*Error {
	errors := []*Error{}

	for _, op := range self.operations {
		if valid := op.Eval(v); !valid {
			errors = append(errors, NewError("Number", op.Message))
		}
	}

	return errors
}
