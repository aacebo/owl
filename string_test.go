package owl_test

import (
	"encoding/json"
	"regexp"
	"testing"

	"github.com/aacebo/owl"
)

func TestString(t *testing.T) {
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

	t.Run("message", func(t *testing.T) {
		t.Run("should have custom error message", func(t *testing.T) {
			err := owl.String().Required().Message("a test message").Validate(nil)

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

			if string(b) != `{"type":"string","min":1,"max":5,"email":true}` {
				t.Errorf(
					"expected `%s`, received `%s`",
					`{"type":"string","min":1,"max":5,"email":true}`,
					string(b),
				)
			}
		})
	})
}

func BenchmarkString(b *testing.B) {
	b.Run("string", func(b *testing.B) {
		schema := owl.String()

		for i := 0; i < b.N; i++ {
			err := schema.Validate("test")

			if err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("min", func(b *testing.B) {
		schema := owl.String().Min(4)

		for i := 0; i < b.N; i++ {
			err := schema.Validate("test")

			if err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("max", func(b *testing.B) {
		schema := owl.String().Max(4)

		for i := 0; i < b.N; i++ {
			err := schema.Validate("test")

			if err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("regex", func(b *testing.B) {
		schema := owl.String().Regex(regexp.MustCompile("^[0-9a-zA-Z_-]+$"))

		for i := 0; i < b.N; i++ {
			err := schema.Validate("test")

			if err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("email", func(b *testing.B) {
		schema := owl.String().Email()

		for i := 0; i < b.N; i++ {
			err := schema.Validate("test@gmail.com")

			if err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("uuid", func(b *testing.B) {
		schema := owl.String().UUID()

		for i := 0; i < b.N; i++ {
			err := schema.Validate("bdc8ffad-a82a-4a03-bd8c-e3ddd6ed98de")

			if err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("url", func(b *testing.B) {
		schema := owl.String().URL()

		for i := 0; i < b.N; i++ {
			err := schema.Validate("https://www.google.com")

			if err != nil {
				b.Fatal(err)
			}
		}
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
