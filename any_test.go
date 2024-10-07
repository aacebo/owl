package owl_test

import (
	"encoding/json"
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

	t.Run("json", func(t *testing.T) {
		t.Run("serialize", func(t *testing.T) {
			schema := owl.Any().Enum(1, true, "hi").Required()
			b, err := json.Marshal(schema)

			if err != nil {
				t.Error(err)
			}

			if string(b) != `{"enum":[1,true,"hi"],"required":true,"type":"any"}` {
				t.Errorf(
					"expected `%s`, received `%s`",
					`{"enum":[1,true,"hi"],"required":true,"type":"any"}`,
					string(b),
				)
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
