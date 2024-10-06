package owl_test

import (
	"testing"

	"github.com/aacebo/owl"
)

func Test_Any(t *testing.T) {
	t.Run("required", func(t *testing.T) {
		t.Run("should succeed", func(t *testing.T) {
			err := owl.Any().Required().Validate("test")

			if err != nil {
				t.Fatal(err.Error())
			}
		})

		t.Run("should fail", func(t *testing.T) {
			err := owl.Any().Required().Validate(nil)

			if err == nil {
				t.Fatal()
			}
		})
	})

	t.Run("enum", func(t *testing.T) {
		t.Run("should succeed", func(t *testing.T) {
			err := owl.Any().Enum("test", 1, false).Validate(1)

			if err != nil {
				t.Fatal(err.Error())
			}
		})

		t.Run("should fail", func(t *testing.T) {
			err := owl.Any().Enum("test", 1, false).Validate(true)

			if err == nil {
				t.Fatal()
			}
		})
	})
}

func ExampleAny() {
	schema := owl.Any()

	if err := schema.Validate("..."); err != nil { // nil
		panic(err)
	}

	if err := schema.Validate(1); err != nil { // nil
		panic(err)
	}

	if err := schema.Validate(true); err != nil { // nil
		panic(err)
	}
}
