package owl_test

import (
	"testing"

	"github.com/aacebo/owl"
)

func Test_Object(t *testing.T) {
	t.Run("required", func(t *testing.T) {
		t.Run("should succeed", func(t *testing.T) {
			err := owl.Object().Field(
				"hello",
				owl.String().Enum("world").Required(),
			).Required().Validate(map[string]any{
				"hello": "world",
			})

			if err != nil {
				t.Fatal(err)
			}
		})

		t.Run("should fail", func(t *testing.T) {
			err := owl.Object().Required().Validate(nil)

			if err == nil {
				t.Fatal()
			}
		})
	})

	t.Run("map", func(t *testing.T) {
		t.Run("should succeed", func(t *testing.T) {
			err := owl.Object().Field(
				"hello",
				owl.Object().Field("hello", owl.String().Enum("world")).Required(),
			).Required().Validate(map[string]any{
				"hello": map[string]any{
					"hello": "world",
				},
			})

			if err != nil {
				t.Fatal(err)
			}
		})

		t.Run("should fail when schema not found", func(t *testing.T) {
			err := owl.Object().Field(
				"hello",
				owl.Object().Field("hello", owl.String().Enum("world")).Required(),
			).Required().Validate(map[string]any{
				"hello": map[string]any{
					"hello1": "world",
				},
			})

			if err == nil {
				t.Fatal()
			}
		})

		t.Run("should fail when invalid schema", func(t *testing.T) {
			err := owl.Object().Field(
				"hello",
				owl.Object().Field("hello", owl.String().Enum("world")).Required(),
			).Required().Validate(map[string]any{
				"hello": "world",
			})

			if err == nil {
				t.Fatal()
			}
		})
	})

	t.Run("struct", func(t *testing.T) {
		type HelloWorld struct {
			Hello string `json:"hello"`
		}

		type Hello struct {
			Hello HelloWorld `json:"hello"`
		}

		t.Run("should succeed", func(t *testing.T) {
			err := owl.Object().Field(
				"hello",
				owl.Object().Field("hello", owl.String().Enum("world")).Required(),
			).Required().Validate(Hello{
				Hello: HelloWorld{"world"},
			})

			if err != nil {
				t.Fatal(err)
			}
		})

		t.Run("should fail when invalid schema", func(t *testing.T) {
			err := owl.Object().Field(
				"hello",
				owl.Object().Field("hello", owl.String().Enum("world")).Required(),
			).Required().Validate(HelloWorld{
				Hello: "world",
			})

			if err == nil {
				t.Fatal()
			}
		})
	})
}
