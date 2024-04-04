package rules_test

import (
	"testing"

	"github.com/aacebo/owl"
)

func Test_Min(t *testing.T) {
	t.Run("float", func(t *testing.T) {
		t.Run("should error", func(t *testing.T) {
			errs := owl.Validate(struct {
				A float64 `json:"a" owl:"min=3"`
			}{2})

			if len(errs) == 0 {
				t.Error("should have error")
			}
		})

		t.Run("should succeed", func(t *testing.T) {
			errs := owl.Validate(struct {
				A float32 `json:"a" owl:"min=3"`
			}{3})

			if len(errs) > 0 {
				t.Error(errs)
			}
		})
	})

	t.Run("int", func(t *testing.T) {
		t.Run("should error", func(t *testing.T) {
			errs := owl.Validate(struct {
				A int `json:"a" owl:"min=3"`
			}{2})

			if len(errs) == 0 {
				t.Error("should have error")
			}
		})

		t.Run("should succeed", func(t *testing.T) {
			errs := owl.Validate(struct {
				A int32 `json:"a" owl:"min=3"`
			}{3})

			if len(errs) > 0 {
				t.Error(errs)
			}
		})
	})

	t.Run("string", func(t *testing.T) {
		t.Run("should error", func(t *testing.T) {
			errs := owl.Validate(struct {
				A string `json:"a" owl:"min=3"`
			}{"ab"})

			if len(errs) == 0 {
				t.Error("should have error")
			}
		})

		t.Run("should succeed", func(t *testing.T) {
			errs := owl.Validate(struct {
				A string `json:"a" owl:"min=3"`
			}{"abc"})

			if len(errs) > 0 {
				t.Error(errs)
			}
		})
	})
}
