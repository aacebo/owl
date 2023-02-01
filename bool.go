package owl

import "reflect"

type BoolSchema struct {
	BaseSchema
}

func Bool() *BoolSchema {
	v := BoolSchema{}
	v.conditions = []*Condition{}

	v.Condition("must be a boolean", func(v any) (any, bool) {
		value, ok := v.(bool)
		return value, ok
	})

	return &v
}

func (self *BoolSchema) Kind() reflect.Kind {
	return reflect.Bool
}
