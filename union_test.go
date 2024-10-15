package owl_test

import (
	"encoding/json"
	"testing"

	"github.com/aacebo/owl"
)

func TestUnion(t *testing.T) {
	t.Run("union", func(t *testing.T) {
		t.Run("should succeed", func(t *testing.T) {
			err := owl.Union(
				owl.String().Required(),
				owl.Int().Required(),
			).Validate("test")

			if err != nil {
				t.Fatal(err.Error())
			}
		})

		t.Run("should fail", func(t *testing.T) {
			err := owl.Union(
				owl.String().Required(),
				owl.Int().Required(),
			).Validate(true)

			if err == nil {
				t.Fatal()
			}
		})
	})

	t.Run("message", func(t *testing.T) {
		t.Run("should have custom error message", func(t *testing.T) {
			err := owl.Union(
				owl.String().Required(),
				owl.Int().Required(),
			).Message("a test message").Validate(true)

			if err == nil {
				t.FailNow()
			}

			if err.Error() != `{"errors":[{"rule":"type","message":"a test message"}]}` {
				t.Errorf(
					"expected `%s`, received `%s`",
					`{"errors":[{"rule":"type","message":"required"}]}`,
					err.Error(),
				)
			}
		})
	})

	t.Run("json", func(t *testing.T) {
		t.Run("serialize", func(t *testing.T) {
			schema := owl.Union(
				owl.String().Required(),
				owl.Int().Required(),
			)

			b, err := json.Marshal(schema)

			if err != nil {
				t.Error(err)
			}

			if string(b) != `{"type":"union[string,int]"}` {
				t.Errorf(
					"expected `%s`, received `%s`",
					`{"type":"union[string,int]"}`,
					string(b),
				)
			}
		})
	})
}

func ExampleUnion() {
	schema := owl.Union(
		owl.String().Required(),
		owl.Int().Required(),
	)

	if err := schema.Validate("test"); err != nil { // nil
		panic(err)
	}

	if err := schema.Validate(1); err != nil { // nil
		panic(err)
	}

	if err := schema.Validate(true); err != nil { // error
		panic(err)
	}
}
