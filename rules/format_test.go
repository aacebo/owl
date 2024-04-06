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

	t.Run("should succeed on nil", func(t *testing.T) {
		errs := owl.Validate(struct {
			A *string `json:"a" owl:"format=email"`
		}{})

		if len(errs) > 0 {
			t.Error(errs)
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

func Benchmark_Format(b *testing.B) {
	b.Run("should succeed on email", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			str := "test@test.com"
			errs := owl.Validate(struct {
				A *string `json:"a" owl:"format=email"`
			}{&str})

			if len(errs) > 0 {
				b.Error(errs)
			}
		}
	})
}
