package rules

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

func Max(ctx Context) []error {
	errs := []error{}
	value := ctx.CoerceValue()

	if value.Kind() == reflect.Invalid {
		return errs
	}

	if !value.CanFloat() && !value.CanConvert(floatType) && value.Kind() != reflect.String {
		errs = append(errs, errors.New("invalid type"))
		return errs
	}

	if ctx.Param() == "" {
		errs = append(errs, errors.New("must be greater than or equal to 0"))
		return errs
	}

	if value.Kind() == reflect.String {
		return maxString(ctx)
	}

	return maxNumber(ctx)
}

func maxNumber(ctx Context) []error {
	errs := []error{}
	value := ctx.CoerceValue()
	max, err := strconv.ParseFloat(ctx.Param(), 64)

	if err != nil {
		errs = append(errs, err)
		return errs
	}

	if max < 0 {
		errs = append(errs, errors.New("must be greater than or equal to 0"))
		return errs
	}

	if v, ok := ctx.RuleParam("min"); ok {
		min, err := strconv.ParseFloat(v, 64)

		if err == nil && max < min {
			errs = append(errs, errors.New("must be greater than or equal to min"))
		}
	}

	if value.Kind() != reflect.Float64 && value.CanConvert(floatType) {
		value = value.Convert(floatType)
	}

	if max < value.Float() {
		errs = append(errs, errors.New(fmt.Sprintf(
			`%v is greater than maximum %v`,
			value.Float(),
			max,
		)))
	}

	return errs
}

func maxString(ctx Context) []error {
	errs := []error{}
	value := ctx.CoerceValue()
	max, err := strconv.ParseInt(ctx.Param(), 10, 64)

	if err != nil {
		errs = append(errs, err)
		return errs
	}

	if max < 0 {
		errs = append(errs, errors.New("config must be greater than or equal to 0"))
		return errs
	}

	if v, ok := ctx.RuleParam("min"); ok {
		min, err := strconv.ParseInt(v, 10, 64)

		if err == nil && max < min {
			errs = append(errs, errors.New("must be greater than or equal to min"))
		}
	}

	if max < int64(value.Len()) {
		errs = append(errs, errors.New(fmt.Sprintf(
			`"%s" has a length of %v, which is greater than maximum %v`,
			value.String(),
			value.Len(),
			max,
		)))
	}

	return errs
}
