package owl_test

import (
	"testing"

	"github.com/aacebo/owl"
)

func Test_Pattern(t *testing.T) {
	t.Run("should error on numeric", func(t *testing.T) {
		errs := owl.Validate(struct {
			A string `json:"a" owl:"pattern=^[0-9]*$"`
		}{"123a"})

		if len(errs) == 0 {
			t.Error("should have error")
		}
	})

	t.Run("should error on invalid pattern", func(t *testing.T) {
		errs := owl.Validate(struct {
			A string `json:"a" owl:"pattern=["`
		}{"123a"})

		if len(errs) == 0 {
			t.Error("should have error")
		}
	})

	t.Run("should succeed on numeric", func(t *testing.T) {
		str := "123"
		errs := owl.Validate(struct {
			A *string `json:"a" owl:"pattern=^[0-9]*$"`
		}{&str})

		if len(errs) > 0 {
			t.Error(errs)
		}
	})

	t.Run("should succeed on empty", func(t *testing.T) {
		errs := owl.Validate(struct {
			A string `json:"a" owl:"pattern="`
		}{"123"})

		if len(errs) > 0 {
			t.Error(errs)
		}
	})
}
