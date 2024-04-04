package rules

import (
	"errors"
	"reflect"
	"strconv"
)

type Default struct{}

func (self Default) Select(schema map[string]string, parent reflect.Value, value reflect.Value) bool {
	return true
}

func (self Default) Validate(schema map[string]string, parent reflect.Value, value reflect.Value) (reflect.Value, []error) {
	errs := []error{}
	config, ok := schema["default"]

	if !ok || config == "" {
		errs = append(errs, errors.New("config is required"))
		return value, errs
	}

	if value.Kind() == reflect.Pointer && value.IsZero() {
		t := value.Type()
		ptr := reflect.New(t.Elem())

		if t.Elem().ConvertibleTo(reflect.TypeFor[int]()) {
			v, err := strconv.ParseInt(config, 10, 64)

			if err != nil {
				errs = append(errs, errors.New("config must be an int"))
			}

			ptr.Elem().Set(reflect.ValueOf(v).Convert(t.Elem()))
			return ptr, errs
		}

		if t.Elem().ConvertibleTo(reflect.TypeFor[float64]()) {
			v, err := strconv.ParseFloat(config, 64)

			if err != nil {
				errs = append(errs, errors.New("config must be an float"))
			}

			ptr.Elem().Set(reflect.ValueOf(v).Convert(t.Elem()))
			return ptr, errs
		}

		if t.Elem().Kind() == reflect.Bool {
			v, err := strconv.ParseBool(config)

			if err != nil {
				errs = append(errs, errors.New("config must be a bool"))
			}

			ptr.Elem().SetBool(v)
			return ptr, errs
		}

		if t.Elem().Kind() == reflect.String {
			ptr.Elem().SetString(config)
			return ptr, errs
		}

		errs = append(errs, errors.New("default can only be used on primitive types"))
	}

	return value, errs
}
