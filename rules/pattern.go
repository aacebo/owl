package rules

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
)

type Pattern struct{}

func (self Pattern) Select(schema map[string]string, parent reflect.Value, value reflect.Value) bool {
	return value.Kind() == reflect.String
}

func (self Pattern) Validate(schema map[string]string, parent reflect.Value, value reflect.Value) []error {
	errs := []error{}
	config, ok := schema["pattern"]

	if !ok {
		errs = append(errs, errors.New("must be a regular expression"))
		return errs
	}

	expr, err := regexp.Compile(config)

	if err != nil {
		errs = append(errs, errors.New("invalid regular expression"))
		return errs
	}

	if !expr.MatchString(value.String()) {
		errs = append(errs, errors.New(fmt.Sprintf(
			`"%s" does not match "%s"`,
			value.String(),
			config,
		)))
	}

	return errs
}
