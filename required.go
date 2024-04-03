package owl

import (
	"errors"
	"reflect"
)

type Required struct{}

func (self Required) Select(parent reflect.Value, value reflect.Value) bool {
	return true
}

func (self Required) Validate(config string, parent reflect.Value, value reflect.Value) []error {
	errs := []error{}

	if !value.IsValid() || value.IsZero() {
		errs = append(errs, errors.New("required"))
	}

	return errs
}
