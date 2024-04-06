package rules

import (
	"errors"
)

func Required(ctx Context) []error {
	errs := []error{}
	value := ctx.CoerceValue()

	if !value.IsValid() || value.IsZero() {
		errs = append(errs, errors.New("required"))
	}

	return errs
}
