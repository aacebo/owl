package owl_test

import (
	"encoding/json"
	"testing"

	"github.com/aacebo/owl"
)

func TestBool(t *testing.T) {
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

	t.Run("message", func(t *testing.T) {
		t.Run("should have custom error message", func(t *testing.T) {
			err := owl.Bool().Required().Message("a test message").Validate(nil)

			if err == nil {
				t.FailNow()
			}

			if err.Error() != `{"errors":[{"rule":"required","message":"a test message"}]}` {
				t.Errorf(
					"expected `%s`, received `%s`",
					`{"errors":[{"rule":"required","message":"required"}]}`,
					err.Error(),
				)
			}
		})
	})

	t.Run("json", func(t *testing.T) {
		t.Run("serialize", func(t *testing.T) {
			schema := owl.Bool()
			b, err := json.Marshal(schema)

			if err != nil {
				t.Error(err)
			}

			if string(b) != `{"type":"bool"}` {
				t.Errorf("expected `%s`, received `%s`", `{"type":"bool"}`, string(b))
			}
		})
	})
}

func BenchmarkBool(b *testing.B) {
	b.Run("bool", func(b *testing.B) {
		schema := owl.Bool()

		for i := 0; i < b.N; i++ {
			err := schema.Validate(true)

			if err != nil {
				b.Fatal(err)
			}
		}
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
