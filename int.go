package owl

import (
	"fmt"
	"reflect"
)

type IntSchema struct {
	BaseSchema
}

func Int() *IntSchema {
	v := IntSchema{}
	v.conditions = []*Condition{}

	v.Condition("must be an integer", func(v any) (any, bool) {
		value, ok := v.(int)
		return value, ok
	})

	return &v
}

func (self *IntSchema) Kind() reflect.Kind {
	return reflect.String
}

func (self *IntSchema) Min(min int) *IntSchema {
	self.Condition(fmt.Sprintf("must be at least %d", min), func(v any) (any, bool) {
		return v, v.(int) >= min
	})

	return self
}

func (self *IntSchema) Max(max int) *IntSchema {
	self.Condition(fmt.Sprintf("must be at most %d", max), func(v any) (any, bool) {
		return v, v.(int) <= max
	})

	return self
}

func (self *IntSchema) Equal(equal int) *IntSchema {
	self.Condition(fmt.Sprintf("must be equal to %d", equal), func(v any) (any, bool) {
		return v, v.(int) == equal
	})

	return self
}
