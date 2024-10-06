package owl_test

import (
	"testing"

	"github.com/aacebo/owl"
)

func Test_Float(t *testing.T) {
	t.Run("required", func(t *testing.T) {
		t.Run("should succeed", func(t *testing.T) {
			err := owl.Float().Required().Validate(1)

			if err != nil {
				t.Fatal(err.Error())
			}
		})

		t.Run("should fail", func(t *testing.T) {
			err := owl.Float().Required().Validate(nil)

			if err == nil {
				t.Fatal()
			}
		})
	})

	t.Run("enum", func(t *testing.T) {
		t.Run("should succeed", func(t *testing.T) {
			err := owl.Float().Enum(1.0).Validate(1)

			if err != nil {
				t.Fatal(err.Error())
			}
		})

		t.Run("should fail", func(t *testing.T) {
			err := owl.Float().Enum(1.0).Validate(2)

			if err == nil {
				t.Fatal()
			}
		})
	})

	t.Run("min", func(t *testing.T) {
		t.Run("should succeed", func(t *testing.T) {
			err := owl.Float().Min(5).Validate(6)

			if err != nil {
				t.Fatal(err.Error())
			}
		})

		t.Run("should fail", func(t *testing.T) {
			err := owl.Float().Min(5).Validate(4)

			if err == nil {
				t.Fatal()
			}
		})
	})

	t.Run("max", func(t *testing.T) {
		t.Run("should succeed", func(t *testing.T) {
			err := owl.Float().Max(5).Validate(4)

			if err != nil {
				t.Fatal(err.Error())
			}
		})

		t.Run("should fail", func(t *testing.T) {
			err := owl.Float().Max(5).Validate(6)

			if err == nil {
				t.Fatal()
			}
		})
	})
}
