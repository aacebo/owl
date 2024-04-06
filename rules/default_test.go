package rules_test

import (
	"testing"

	"github.com/aacebo/owl"
)

func Test_Default(t *testing.T) {
	t.Run("should error on no param", func(t *testing.T) {
		errs := owl.Validate(struct {
			A *string `json:"a" owl:"default"`
		}{})

		if len(errs) == 0 {
			t.Error("should have error")
		}
	})

	t.Run("should error on type mismatch", func(t *testing.T) {
		errs := owl.Validate(struct {
			A *int `json:"a" owl:"default=abc"`
		}{})

		if len(errs) == 0 {
			t.Error("should have error")
		}
	})

	t.Run("should succeed with int", func(t *testing.T) {
		v := struct {
			A *int `json:"a" owl:"default=123"`
		}{}

		errs := owl.Validate(&v)

		if len(errs) > 0 {
			t.Error(errs)
		}

		if v.A == nil || *v.A != 123 {
			t.Error("should set default value")
		}
	})

	t.Run("should succeed with int64", func(t *testing.T) {
		v := struct {
			A *int64 `json:"a" owl:"default=123"`
		}{}

		errs := owl.Validate(&v)

		if len(errs) > 0 {
			t.Error(errs)
		}

		if v.A == nil || *v.A != 123 {
			t.Error("should set default value")
		}
	})

	t.Run("should succeed with float32", func(t *testing.T) {
		v := struct {
			A *float32 `json:"a" owl:"default=123"`
		}{}

		errs := owl.Validate(&v)

		if len(errs) > 0 {
			t.Error(errs)
		}

		if v.A == nil || *v.A != 123 {
			t.Error("should set default value")
		}
	})

	t.Run("should succeed with float64", func(t *testing.T) {
		v := struct {
			A *float64 `json:"a" owl:"default=123"`
		}{}

		errs := owl.Validate(&v)

		if len(errs) > 0 {
			t.Error(errs)
		}

		if v.A == nil || *v.A != 123 {
			t.Error("should set default value")
		}
	})

	t.Run("should succeed with uint", func(t *testing.T) {
		v := struct {
			A *uint `json:"a" owl:"default=123"`
		}{}

		errs := owl.Validate(&v)

		if len(errs) > 0 {
			t.Error(errs)
		}

		if v.A == nil || *v.A != 123 {
			t.Error("should set default value")
		}
	})

	t.Run("should succeed with string", func(t *testing.T) {
		v := struct {
			A *string `json:"a" owl:"default=123"`
		}{}

		errs := owl.Validate(&v)

		if len(errs) > 0 {
			t.Error(errs)
		}

		if v.A == nil || *v.A != "123" {
			t.Error("should set default value")
		}
	})

	t.Run("should succeed with required", func(t *testing.T) {
		v := struct {
			A *string `json:"a" owl:"default=123,required"`
		}{}

		errs := owl.Validate(&v)

		if len(errs) > 0 {
			t.Error(errs)
		}

		if v.A == nil || *v.A != "123" {
			t.Error("should set default value")
		}
	})
}

func Benchmark_Default(b *testing.B) {
	b.Run("should succeed", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			v := struct {
				A *string `json:"a" owl:"default=123"`
			}{}

			errs := owl.Validate(&v)

			if len(errs) > 0 {
				b.Error(errs)
			}

			if v.A == nil || *v.A != "123" {
				b.Error("should set default value")
			}
		}
	})
}
