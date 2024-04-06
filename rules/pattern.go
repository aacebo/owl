package rules

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
)

func Pattern(ctx Context) []error {
	errs := []error{}
	value := ctx.CoerceValue()
	param := ctx.Param()

	if !value.IsValid() || value.IsZero() {
		return errs
	}

	if value.Kind() != reflect.String {
		errs = append(errs, errors.New("invalid type"))
		return errs
	}

	expr, err := regexp.Compile(param)

	if err != nil {
		errs = append(errs, errors.New("invalid regular expression"))
		return errs
	}

	if !expr.MatchString(value.String()) {
		errs = append(errs, errors.New(fmt.Sprintf(
			`"%s" does not match "%s"`,
			value.String(),
			param,
		)))
	}

	return errs
}
