package rules_test

import (
	"testing"

	"github.com/aacebo/owl"
)

func Test_Required(t *testing.T) {
	t.Run("should error on nil", func(t *testing.T) {
		errs := owl.Validate(struct {
			A *string `json:"a" owl:"required"`
		}{})

		if len(errs) == 0 {
			t.Error("should have error")
		}
	})

	t.Run("should error on zero", func(t *testing.T) {
		errs := owl.Validate(struct {
			A string `json:"a" owl:"required"`
		}{})

		if len(errs) == 0 {
			t.Error("should have error")
		}
	})
}
