package rules

import (
	"errors"
	"reflect"
)

type Required struct{}

func (self Required) Select(schema map[string]string, parent reflect.Value, value reflect.Value) bool {
	return true
}

func (self Required) Validate(schema map[string]string, parent reflect.Value, value reflect.Value) (reflect.Value, []error) {
	errs := []error{}

	if !value.IsValid() || value.IsZero() {
		errs = append(errs, errors.New("required"))
	}

	return value, errs
}
