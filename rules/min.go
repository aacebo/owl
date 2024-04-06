package rules

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

var floatType = reflect.TypeFor[float64]()

func Min(ctx Context) []error {
	errs := []error{}
	value := ctx.CoerceValue()

	if !value.IsValid() || value.IsZero() {
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
		return minString(ctx)
	}

	return minNumber(ctx)
}

func minNumber(ctx Context) []error {
	errs := []error{}
	value := ctx.CoerceValue()
	min, err := strconv.ParseFloat(ctx.Param(), 64)

	if err != nil {
		errs = append(errs, err)
		return errs
	}

	if min < 0 {
		errs = append(errs, errors.New("must be greater than or equal to 0"))
		return errs
	}

	if value.Kind() != reflect.Float64 && value.CanConvert(floatType) {
		value = value.Convert(floatType)
	}

	if min > value.Float() {
		errs = append(errs, errors.New(fmt.Sprintf(
			`%v is less than minimum %v`,
			value.Float(),
			min,
		)))
	}

	return errs
}

func minString(ctx Context) []error {
	errs := []error{}
	value := ctx.CoerceValue()
	min, err := strconv.ParseInt(ctx.Param(), 10, 64)

	if err != nil {
		errs = append(errs, err)
		return errs
	}

	if min < 0 {
		errs = append(errs, errors.New("config must be greater than or equal to 0"))
		return errs
	}

	if min > int64(value.Len()) {
		errs = append(errs, errors.New(fmt.Sprintf(
			`"%s" has a length of %v, which is less than minimum %v`,
			value.String(),
			value.Len(),
			min,
		)))
	}

	return errs
}
