package rules

import (
	"errors"
	"fmt"
	"reflect"
)

func Format(ctx Context) []error {
	errs := []error{}
	value := ctx.CoerceValue()
	param := ctx.Param()

	if value.Kind() == reflect.Invalid {
		return errs
	}

	if value.Kind() != reflect.String {
		errs = append(errs, errors.New(`can only be used on "string" type`))
		return errs
	}

	if param == "" {
		errs = append(errs, errors.New("param is required"))
		return errs
	}

	if !ctx.HasFormat(param) {
		errs = append(errs, errors.New(fmt.Sprintf(
			`format "%s" not found`,
			param,
		)))

		return errs
	}

	if err := ctx.Format(param, value.String()); err != nil {
		errs = append(errs, err)
	}

	return errs
}
