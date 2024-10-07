package owl_test

import (
	"encoding/json"
	"regexp"
	"testing"

	"github.com/aacebo/owl"
)

func Test_String(t *testing.T) {
	t.Run("required", func(t *testing.T) {
		t.Run("should succeed", func(t *testing.T) {
			err := owl.String().Required().Validate("test")

			if err != nil {
				t.Fatal(err.Error())
			}
		})

		t.Run("should fail", func(t *testing.T) {
			err := owl.String().Required().Validate(nil)

			if err == nil {
				t.Fatal()
			}
		})
	})

	t.Run("enum", func(t *testing.T) {
		t.Run("should succeed", func(t *testing.T) {
			err := owl.String().Enum("test").Validate("test")

			if err != nil {
				t.Fatal(err.Error())
			}
		})

		t.Run("should fail", func(t *testing.T) {
			err := owl.String().Enum("test").Validate("tester")

			if err == nil {
				t.Fatal()
			}
		})
	})

	t.Run("min", func(t *testing.T) {
		t.Run("should succeed when nil", func(t *testing.T) {
			err := owl.String().Min(5).Validate(nil)

			if err != nil {
				t.Fatal(err.Error())
			}
		})

		t.Run("should succeed when length gt min", func(t *testing.T) {
			err := owl.String().Min(5).Validate("tester")

			if err != nil {
				t.Fatal(err.Error())
			}
		})

		t.Run("should fail when length lt min", func(t *testing.T) {
			err := owl.String().Min(5).Validate("test")

			if err == nil {
				t.Fatal()
			}
		})
	})

	t.Run("max", func(t *testing.T) {
		t.Run("should succeed when nil", func(t *testing.T) {
			err := owl.String().Max(5).Validate(nil)

			if err != nil {
				t.Fatal(err.Error())
			}
		})

		t.Run("should succeed when length lt max", func(t *testing.T) {
			err := owl.String().Max(5).Validate("test")

			if err != nil {
				t.Fatal(err.Error())
			}
		})

		t.Run("should fail when length gt max", func(t *testing.T) {
			err := owl.String().Max(5).Validate("tester")

			if err == nil {
				t.Fatal()
			}
		})
	})

	t.Run("regex", func(t *testing.T) {
		t.Run("should succeed when nil", func(t *testing.T) {
			err := owl.String().Regex(regexp.MustCompile("^[0-9a-zA-Z_-]+$")).Validate(nil)

			if err != nil {
				t.Fatal(err.Error())
			}
		})

		t.Run("should succeed when matches", func(t *testing.T) {
			err := owl.String().Regex(regexp.MustCompile("^[0-9a-zA-Z_-]+$")).Validate("test")

			if err != nil {
				t.Fatal(err.Error())
			}
		})

		t.Run("should fail when not matches", func(t *testing.T) {
			err := owl.String().Regex(regexp.MustCompile("^[0-9a-zA-Z_-]+$")).Validate("a test")

			if err == nil {
				t.Fatal()
			}
		})
	})

	t.Run("email", func(t *testing.T) {
		t.Run("should succeed when nil", func(t *testing.T) {
			err := owl.String().Email().Validate(nil)

			if err != nil {
				t.Fatal(err.Error())
			}
		})

		t.Run("should succeed when email", func(t *testing.T) {
			err := owl.String().Email().Validate("test@gmail.com")

			if err != nil {
				t.Fatal(err.Error())
			}
		})

		t.Run("should fail when not email", func(t *testing.T) {
			err := owl.String().Email().Validate("test")

			if err == nil {
				t.Fatal()
			}
		})
	})

	t.Run("uuid", func(t *testing.T) {
		t.Run("should succeed when nil", func(t *testing.T) {
			err := owl.String().UUID().Validate(nil)

			if err != nil {
				t.Fatal(err.Error())
			}
		})

		t.Run("should succeed when uuid", func(t *testing.T) {
			err := owl.String().UUID().Validate("afefc1ab-b8f2-4803-8e9a-60515854141a")

			if err != nil {
				t.Fatal(err.Error())
			}
		})

		t.Run("should fail when not uuid", func(t *testing.T) {
			err := owl.String().UUID().Validate("test")

			if err == nil {
				t.Fatal()
			}
		})
	})

	t.Run("url", func(t *testing.T) {
		t.Run("should succeed when nil", func(t *testing.T) {
			err := owl.String().URL().Validate(nil)

			if err != nil {
				t.Fatal(err.Error())
			}
		})

		t.Run("should succeed when url", func(t *testing.T) {
			err := owl.String().URL().Validate("https://www.google.com")

			if err != nil {
				t.Fatal(err.Error())
			}
		})

		t.Run("should fail when not url", func(t *testing.T) {
			err := owl.String().URL().Validate("test")

			if err == nil {
				t.Fatal()
			}
		})
	})

	t.Run("json", func(t *testing.T) {
		t.Run("serialize", func(t *testing.T) {
			schema := owl.String().Min(1).Max(5).Email()
			b, err := json.Marshal(schema)

			if err != nil {
				t.Error(err)
			}

			if string(b) != `{"email":true,"max":5,"min":1,"type":"string"}` {
				t.Errorf(
					"expected `%s`, received `%s`",
					`{"email":true,"max":5,"min":1,"type":"string"}`,
					string(b),
				)
			}
		})
	})
}

func ExampleString() {
	schema := owl.String()

	if err := schema.Validate("test"); err != nil { // nil
		panic(err)
	}

	if err := schema.Validate(true); err != nil { // error
		panic(err)
	}
}

func ExampleStringSchema_Min() {
	schema := owl.String().Min(5)

	if err := schema.Validate("tester"); err != nil { // nil
		panic(err)
	}

	if err := schema.Validate("test"); err != nil { // error
		panic(err)
	}
}

func ExampleStringSchema_Max() {
	schema := owl.String().Max(5)

	if err := schema.Validate("test"); err != nil { // nil
		panic(err)
	}

	if err := schema.Validate("tester"); err != nil { // error
		panic(err)
	}
}

func ExampleStringSchema_Regex() {
	schema := owl.String().Regex(regexp.MustCompile("^[0-9a-zA-Z_-]+$"))

	if err := schema.Validate("test"); err != nil { // nil
		panic(err)
	}

	if err := schema.Validate("hello world"); err != nil { // error
		panic(err)
	}
}

func ExampleStringSchema_Email() {
	schema := owl.String().Email()

	if err := schema.Validate("test@gmail.com"); err != nil { // nil
		panic(err)
	}

	if err := schema.Validate("test"); err != nil { // error
		panic(err)
	}
}

func ExampleStringSchema_UUID() {
	schema := owl.String().UUID()

	if err := schema.Validate("afefc1ab-b8f2-4803-8e9a-60515854141a"); err != nil { // nil
		panic(err)
	}

	if err := schema.Validate("test"); err != nil { // error
		panic(err)
	}
}

func ExampleStringSchema_URL() {
	schema := owl.String().URL()

	if err := schema.Validate("https://www.google.com"); err != nil { // nil
		panic(err)
	}

	if err := schema.Validate("test"); err != nil { // error
		panic(err)
	}
}
