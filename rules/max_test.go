package rules_test

import (
	"testing"

	"github.com/aacebo/owl"
)

func Test_Max(t *testing.T) {
	t.Run("config", func(t *testing.T) {
		t.Run("should error on string", func(t *testing.T) {
			errs := owl.Validate(struct {
				A float64 `json:"a" owl:"max=abc"`
			}{2})

			if len(errs) == 0 {
				t.Error("should have error")
			}
		})

		t.Run("should error on negative", func(t *testing.T) {
			errs := owl.Validate(struct {
				A float64 `json:"a" owl:"max=-1"`
			}{2})

			if len(errs) == 0 {
				t.Error("should have error")
			}
		})

		t.Run("should error on less than min", func(t *testing.T) {
			errs := owl.Validate(struct {
				A float64 `json:"a" owl:"min=5,max=3"`
			}{2})

			if len(errs) == 0 {
				t.Error("should have error")
			}
		})
	})

	t.Run("float", func(t *testing.T) {
		t.Run("should error", func(t *testing.T) {
			errs := owl.Validate(struct {
				A float64 `json:"a" owl:"max=3"`
			}{4})

			if len(errs) == 0 {
				t.Error("should have error")
			}
		})

		t.Run("should succeed", func(t *testing.T) {
			errs := owl.Validate(struct {
				A float32 `json:"a" owl:"max=3"`
			}{3})

			if len(errs) > 0 {
				t.Error(errs)
			}
		})
	})

	t.Run("int", func(t *testing.T) {
		t.Run("should error", func(t *testing.T) {
			errs := owl.Validate(struct {
				A int `json:"a" owl:"max=3"`
			}{4})

			if len(errs) == 0 {
				t.Error("should have error")
			}
		})

		t.Run("should succeed", func(t *testing.T) {
			errs := owl.Validate(struct {
				A int32 `json:"a" owl:"max=3"`
			}{3})

			if len(errs) > 0 {
				t.Error(errs)
			}
		})
	})

	t.Run("string", func(t *testing.T) {
		t.Run("should error", func(t *testing.T) {
			errs := owl.Validate(struct {
				A string `json:"a" owl:"max=3"`
			}{"abcd"})

			if len(errs) == 0 {
				t.Error("should have error")
			}
		})

		t.Run("should succeed", func(t *testing.T) {
			errs := owl.Validate(struct {
				A string `json:"a" owl:"max=3"`
			}{"abc"})

			if len(errs) > 0 {
				t.Error(errs)
			}
		})
	})
}
