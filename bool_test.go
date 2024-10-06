package owl_test

import (
	"testing"

	"github.com/aacebo/owl"
)

func Test_Bool(t *testing.T) {
	t.Run("required", func(t *testing.T) {
		t.Run("should succeed", func(t *testing.T) {
			err := owl.Bool().Required().Validate(true)

			if err != nil {
				t.Fatal(err.Error())
			}
		})

		t.Run("should fail", func(t *testing.T) {
			err := owl.Bool().Required().Validate(nil)

			if err == nil {
				t.Fatal()
			}
		})
	})

	t.Run("enum", func(t *testing.T) {
		t.Run("should succeed", func(t *testing.T) {
			err := owl.Bool().Enum(true).Validate(true)

			if err != nil {
				t.Fatal(err.Error())
			}
		})

		t.Run("should fail", func(t *testing.T) {
			err := owl.Bool().Enum(true).Validate(false)

			if err == nil {
				t.Fatal()
			}
		})
	})
}

func ExampleBool() {
	schema := owl.Bool()

	if err := schema.Validate(true); err != nil { // nil
		panic(err)
	}

	if err := schema.Validate(false); err != nil { // nil
		panic(err)
	}

	if err := schema.Validate("test"); err != nil { // error
		panic(err)
	}
}
