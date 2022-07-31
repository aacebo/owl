package owl

import "reflect"

type StructSchema struct {
	operations []*Operation
	index      int
}

func Struct() *StructSchema {
	v := StructSchema{
		index: 0,
		operations: []*Operation{
			NewOperation("must be a structure", func(v any) (any, bool) {
				kind := reflect.TypeOf(v).Kind()
				return v, kind == reflect.Struct
			}),
		},
	}

	return &v
}
