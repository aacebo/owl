package owl_test

import (
	"encoding/json"
	"regexp"
	"testing"

	"github.com/aacebo/owl"
)

func Test_Object(t *testing.T) {
	t.Run("required", func(t *testing.T) {
		t.Run("should succeed", func(t *testing.T) {
			err := owl.Object().Field(
				"hello", owl.String().Enum("world").Required(),
			).Required().Validate(map[string]any{
				"hello": "world",
			})

			if err != nil {
				t.Fatal(err)
			}
		})

		t.Run("should succeed when field nil and not required", func(t *testing.T) {
			err := owl.Object().Field(
				"hello", owl.String().Enum("world"),
			).Required().Validate(map[string]any{})

			if err != nil {
				t.Fatal(err)
			}
		})

		t.Run("should fail when nil", func(t *testing.T) {
			err := owl.Object().Required().Validate(nil)

			if err == nil {
				t.Fatal()
			}
		})

		t.Run("should fail when required field nil", func(t *testing.T) {
			err := owl.Object().Fields(map[string]owl.Schema{
				"hello": owl.String().Required(),
			}).Required().Validate(map[string]any{})

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
				owl.Object().Field("hello", owl.String().Enum("world").Required()).Required(),
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

		t.Run("should fail when missing required field", func(t *testing.T) {
			err := owl.Object().Field(
				"hello",
				owl.String().Required(),
			).Field(
				"world",
				owl.Bool().Required(),
			).Validate(map[string]any{
				"hello": "test",
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

	t.Run("json", func(t *testing.T) {
		t.Run("serialize", func(t *testing.T) {
			schema := owl.Object().Fields(map[string]owl.Schema{
				"username":     owl.String().Regex(regexp.MustCompile("^[0-9a-zA-Z_-]+$")).Required(),
				"password":     owl.String().Min(5).Max(20).Required(),
				"staySignedIn": owl.Bool(),
			})

			b, err := json.Marshal(schema)

			if err != nil {
				t.Error(err)
			}

			if string(b) != `{"fields":{"password":{"max":20,"min":5,"required":true,"type":"string"},"staySignedIn":{"type":"bool"},"username":{"regex":"^[0-9a-zA-Z_-]+$","required":true,"type":"string"}},"type":"object"}` {
				t.Errorf(
					"expected `%s`, received `%s`",
					`{"fields":{"password":{"max":20,"min":5,"required":true,"type":"string"},"staySignedIn":{"type":"bool"},"username":{"regex":"^[0-9a-zA-Z_-]+$","required":true,"type":"string"}},"type":"object"}`,
					string(b),
				)
			}
		})
	})
}

func ExampleObject() {
	schema := owl.Object().Field(
		"email", owl.String().Email().Required(),
	).Field(
		"password", owl.String().Min(5).Max(20).Required(),
	)

	if err := schema.Validate(map[string]any{
		"email":    "test@test.com",
		"password": "mytestpassword",
	}); err != nil { // nil
		panic(err)
	}

	if err := schema.Validate(map[string]any{
		"email": "test@test.com",
	}); err != nil { // error
		panic(err)
	}
}
