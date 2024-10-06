package owl

import "reflect"

type Schema interface {
	Type() string
	Validate(value any) error
	validate(key string, value reflect.Value) error
}
