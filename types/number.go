package types

import (
	"reflect"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

func IsInt(v any) bool {
	kind := reflect.TypeOf(v).Kind()

	return kind == reflect.Int ||
		kind == reflect.Int8 ||
		kind == reflect.Int16 ||
		kind == reflect.Int32 ||
		kind == reflect.Int64 ||
		kind == reflect.Uint ||
		kind == reflect.Uint8 ||
		kind == reflect.Uint16 ||
		kind == reflect.Uint32 ||
		kind == reflect.Uint64
}

func IsFloat(v any) bool {
	kind := reflect.TypeOf(v).Kind()

	return kind == reflect.Float32 ||
		kind == reflect.Float64
}

func IsNumber(v any) bool {
	return IsInt(v) || IsFloat(v)
}
