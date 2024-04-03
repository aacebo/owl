package owl

import "reflect"

type Rule interface {
	Select(parent reflect.Value, value reflect.Value) bool
	Validate(config string, parent reflect.Value, value reflect.Value) []error
}
