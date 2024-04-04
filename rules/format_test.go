package rules_test

import (
	"testing"

	"github.com/aacebo/owl"
)

func Test_Format(t *testing.T) {
	t.Run("should error on email", func(t *testing.T) {
		errs := owl.Validate(struct {
			A string `json:"a" owl:"format=email"`
		}{"test"})

		if len(errs) == 0 {
			t.Error("should have error")
		}
	})

	t.Run("should error on invalid format", func(t *testing.T) {
		errs := owl.Validate(struct {
			A string `json:"a" owl:"format=test"`
		}{"123a"})

		if len(errs) == 0 {
			t.Error("should have error")
		}
	})

	t.Run("should succeed on email", func(t *testing.T) {
		str := "test@test.com"
		errs := owl.Validate(struct {
			A *string `json:"a" owl:"format=email"`
		}{&str})

		if len(errs) > 0 {
			t.Error(errs)
		}
	})
}
