package rules

import (
	"errors"
	"reflect"
	"strconv"
)

func Default(ctx Context) []error {
	errs := []error{}
	value := ctx.Value()
	param := ctx.Param()

	if param == "" {
		errs = append(errs, errors.New("param is required"))
		return errs
	}

	if !value.CanAddr() {
		errs = append(errs, errors.New("cannot be addressed"))
		return errs
	}

	if value.Kind() == reflect.Pointer && value.IsZero() {
		t := value.Type()
		ptr := reflect.New(t.Elem())

		if t.Elem().ConvertibleTo(reflect.TypeFor[int]()) {
			v, err := strconv.ParseInt(param, 10, 64)

			if err != nil {
				errs = append(errs, errors.New("param must be an int"))
				return errs
			}

			ptr.Elem().Set(reflect.ValueOf(v).Convert(t.Elem()))
		}

		if t.Elem().ConvertibleTo(reflect.TypeFor[float64]()) {
			v, err := strconv.ParseFloat(param, 64)

			if err != nil {
				errs = append(errs, errors.New("param must be an float"))
				return errs
			}

			ptr.Elem().Set(reflect.ValueOf(v).Convert(t.Elem()))
		}

		if t.Elem().Kind() == reflect.Bool {
			v, err := strconv.ParseBool(param)

			if err != nil {
				errs = append(errs, errors.New("param must be a bool"))
				return errs
			}

			ptr.Elem().SetBool(v)
		}

		if t.Elem().Kind() == reflect.String {
			ptr.Elem().SetString(param)
		}

		value.Set(ptr)
	}

	return errs
}
