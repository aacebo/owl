package rules

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

var floatType = reflect.TypeFor[float64]()

type Min struct{}

func (self Min) Select(schema map[string]string, parent reflect.Value, value reflect.Value) bool {
	return value.CanFloat() || value.CanConvert(floatType) || value.Kind() == reflect.String
}

func (self Min) Validate(schema map[string]string, parent reflect.Value, value reflect.Value) []error {
	errs := []error{}
	config, ok := schema["min"]

	if !ok {
		errs = append(errs, errors.New("must be greater than or equal to 0"))
		return errs
	}

	if value.Kind() == reflect.String {
		return self.validateString(config, value)
	}

	return self.validateNumber(config, value)
}

func (self Min) validateNumber(config string, value reflect.Value) []error {
	errs := []error{}
	min, err := strconv.ParseFloat(config, 64)

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

func (self Min) validateString(config string, value reflect.Value) []error {
	errs := []error{}
	min, err := strconv.ParseInt(config, 10, 64)

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
