package owl_test

import (
	"encoding/json"
	"testing"
	"time"

	"math/rand"

	"github.com/aacebo/owl"
)

func TestAny(t *testing.T) {
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

	t.Run("message", func(t *testing.T) {
		t.Run("should have custom error message", func(t *testing.T) {
			err := owl.Any().Required().Message("a test message").Validate(nil)

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
			schema := owl.Any().Enum(1, true, "hi").Required()
			b, err := json.Marshal(schema)

			if err != nil {
				t.Error(err)
			}

			if string(b) != `{"type":"any","enum":[1,true,"hi"],"required":true}` {
				t.Errorf(
					"expected `%s`, received `%s`",
					`{"type":"any","enum":[1,true,"hi"],"required":true}`,
					string(b),
				)
			}
		})
	})
}

func BenchmarkAny(b *testing.B) {
	b.Run("any", func(b *testing.B) {
		schema := owl.Any()

		for i := 0; i < b.N; i++ {
			err := schema.Validate(1)

			if err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("enum", func(b *testing.B) {
		enum := []any{"test", 1, true}
		schema := owl.Any().Enum(enum...)
		s := rand.NewSource(time.Now().Unix())
		r := rand.New(s)

		for i := 0; i < b.N; i++ {
			err := schema.Validate(enum[r.Intn(3)])

			if err != nil {
				b.Fatal(err)
			}
		}
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
