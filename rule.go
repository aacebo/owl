package owl

import "reflect"

type Rule interface {
	Select(
		schema map[string]string,
		parent reflect.Value,
		value reflect.Value,
	) bool

	Validate(
		schema map[string]string,
		parent reflect.Value,
		// _type reflect.Type,
		value reflect.Value,
	) (reflect.Value, []error)
}
