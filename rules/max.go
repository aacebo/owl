package rules

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

type Max struct{}

func (self Max) Select(parent reflect.Value, value reflect.Value) bool {
	return value.CanFloat() || value.CanConvert(floatType) || value.Kind() == reflect.String
}

func (self Max) Validate(config string, parent reflect.Value, value reflect.Value) []error {
	errs := []error{}

	if config == "" {
		errs = append(errs, errors.New("empty config"))
		return errs
	}

	if value.Kind() == reflect.String {
		return self.validateString(config, value)
	}

	return self.validateNumber(config, value)
}

func (self Max) validateNumber(config string, value reflect.Value) []error {
	errs := []error{}
	max, err := strconv.ParseFloat(config, 64)

	if err != nil {
		errs = append(errs, err)
		return errs
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

func (self Max) validateString(config string, value reflect.Value) []error {
	errs := []error{}
	max, err := strconv.ParseInt(config, 10, 64)

	if err != nil {
		errs = append(errs, err)
		return errs
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
