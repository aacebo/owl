package rules

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

type Max struct{}

func (self Max) Select(schema map[string]string, parent reflect.Value, value reflect.Value) bool {
	return value.CanFloat() || value.CanConvert(floatType) || value.Kind() == reflect.String
}

func (self Max) Validate(schema map[string]string, parent reflect.Value, value reflect.Value) (reflect.Value, []error) {
	errs := []error{}

	if _, ok := schema["max"]; !ok {
		errs = append(errs, errors.New("must be greater than or equal to 0"))
		return value, errs
	}

	if value.Kind() == reflect.String {
		return self.validateString(schema, value)
	}

	return self.validateNumber(schema, value)
}

func (self Max) validateNumber(schema map[string]string, value reflect.Value) (reflect.Value, []error) {
	errs := []error{}
	max, err := strconv.ParseFloat(schema["max"], 64)

	if err != nil {
		errs = append(errs, err)
		return value, errs
	}

	if max < 0 {
		errs = append(errs, errors.New("must be greater than or equal to 0"))
		return value, errs
	}

	if v, ok := schema["min"]; ok {
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

	return value, errs
}

func (self Max) validateString(schema map[string]string, value reflect.Value) (reflect.Value, []error) {
	errs := []error{}
	max, err := strconv.ParseInt(schema["max"], 10, 64)

	if err != nil {
		errs = append(errs, err)
		return value, errs
	}

	if max < 0 {
		errs = append(errs, errors.New("config must be greater than or equal to 0"))
		return value, errs
	}

	if v, ok := schema["min"]; ok {
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

	return value, errs
}
