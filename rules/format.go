package rules

import (
	"errors"
	"fmt"
	"reflect"
)

type Format struct {
	HasFormat func(string) bool
	Format    func(string, string) error
}

func (self Format) Select(parent reflect.Value, value reflect.Value) bool {
	return value.Kind() == reflect.String
}

func (self Format) Validate(config string, parent reflect.Value, value reflect.Value) []error {
	errs := []error{}

	if config == "" {
		errs = append(errs, errors.New("invalid empty format"))
		return errs
	}

	if !self.HasFormat(config) {
		errs = append(errs, errors.New(fmt.Sprintf(
			`format "%s" not found`,
			config,
		)))

		return errs
	}

	if err := self.Format(config, value.String()); err != nil {
		errs = append(errs, err)
	}

	return errs
}
