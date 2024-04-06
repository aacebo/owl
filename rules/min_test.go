package rules_test

import (
	"testing"

	"github.com/aacebo/owl"
)

func Test_Min(t *testing.T) {
	t.Run("param", func(t *testing.T) {
		t.Run("should error on string", func(t *testing.T) {
			errs := owl.Validate(struct {
				A float64 `json:"a" owl:"min=abc"`
			}{2})

			if len(errs) == 0 {
				t.Error("should have error")
			}
		})

		t.Run("should error on negative", func(t *testing.T) {
			errs := owl.Validate(struct {
				A float64 `json:"a" owl:"min=-1"`
			}{2})

			if len(errs) == 0 {
				t.Error("should have error")
			}
		})
	})

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

func Benchmark_Min(b *testing.B) {
	b.Run("float", func(b *testing.B) {
		b.Run("should succeed", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				errs := owl.Validate(struct {
					A float32 `json:"a" owl:"min=3"`
				}{3})

				if len(errs) > 0 {
					b.Error(errs)
				}
			}
		})
	})

	b.Run("int", func(b *testing.B) {
		b.Run("should succeed", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				errs := owl.Validate(struct {
					A int32 `json:"a" owl:"min=3"`
				}{3})

				if len(errs) > 0 {
					b.Error(errs)
				}
			}
		})
	})

	b.Run("string", func(b *testing.B) {
		b.Run("should succeed", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				errs := owl.Validate(struct {
					A string `json:"a" owl:"min=3"`
				}{"abc"})

				if len(errs) > 0 {
					b.Error(errs)
				}
			}
		})
	})
}
