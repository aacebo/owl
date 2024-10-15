package owl_test

import (
	"encoding/json"
	"testing"

	"github.com/aacebo/owl"
)

func TestFloat(t *testing.T) {
	t.Run("required", func(t *testing.T) {
		t.Run("should succeed", func(t *testing.T) {
			err := owl.Float().Required().Validate(1)

			if err != nil {
				t.Fatal(err.Error())
			}
		})

		t.Run("should fail", func(t *testing.T) {
			err := owl.Float().Required().Validate(nil)

			if err == nil {
				t.Fatal()
			}
		})
	})

	t.Run("enum", func(t *testing.T) {
		t.Run("should succeed", func(t *testing.T) {
			err := owl.Float().Enum(1.0).Validate(1)

			if err != nil {
				t.Fatal(err.Error())
			}
		})

		t.Run("should fail", func(t *testing.T) {
			err := owl.Float().Enum(1.0).Validate(2)

			if err == nil {
				t.Fatal()
			}
		})
	})

	t.Run("message", func(t *testing.T) {
		t.Run("should have custom error message", func(t *testing.T) {
			err := owl.Float().Required().Message("a test message").Validate(nil)

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

	t.Run("type", func(t *testing.T) {
		t.Run("should succeed when float", func(t *testing.T) {
			err := owl.Float().Validate(1.5)

			if err != nil {
				t.Fatal(err.Error())
			}
		})

		t.Run("should succeed when int", func(t *testing.T) {
			err := owl.Float().Validate(1)

			if err != nil {
				t.Fatal(err.Error())
			}
		})

		t.Run("should fail when not int/float", func(t *testing.T) {
			err := owl.Float().Validate(true)

			if err == nil {
				t.FailNow()
			}
		})
	})

	t.Run("min", func(t *testing.T) {
		t.Run("should succeed when nil", func(t *testing.T) {
			err := owl.Float().Min(5).Validate(nil)

			if err != nil {
				t.Fatal(err.Error())
			}
		})

		t.Run("should succeed when gt min", func(t *testing.T) {
			err := owl.Float().Min(5).Validate(6)

			if err != nil {
				t.Fatal(err.Error())
			}
		})

		t.Run("should fail when lt min", func(t *testing.T) {
			err := owl.Float().Min(5).Validate(4)

			if err == nil {
				t.Fatal()
			}
		})
	})

	t.Run("max", func(t *testing.T) {
		t.Run("should succeed when nil", func(t *testing.T) {
			err := owl.Float().Max(5).Validate(nil)

			if err != nil {
				t.Fatal(err.Error())
			}
		})

		t.Run("should succeed when lt max", func(t *testing.T) {
			err := owl.Float().Max(5).Validate(4)

			if err != nil {
				t.Fatal(err.Error())
			}
		})

		t.Run("should fail when gt max", func(t *testing.T) {
			err := owl.Float().Max(5).Validate(6)

			if err == nil {
				t.Fatal()
			}
		})
	})

	t.Run("json", func(t *testing.T) {
		t.Run("serialize", func(t *testing.T) {
			schema := owl.Float().Min(1).Max(5)
			b, err := json.Marshal(schema)

			if err != nil {
				t.Error(err)
			}

			if string(b) != `{"type":"float","min":1,"max":5}` {
				t.Errorf(
					"expected `%s`, received `%s`",
					`{"type":"float","min":1,"max":5}`,
					string(b),
				)
			}
		})
	})
}

func BenchmarkFloat(b *testing.B) {
	b.Run("float32", func(b *testing.B) {
		schema := owl.Float()

		for i := 0; i < b.N; i++ {
			err := schema.Validate(float32(1.5))

			if err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("float64", func(b *testing.B) {
		schema := owl.Float()

		for i := 0; i < b.N; i++ {
			err := schema.Validate(float64(1.5))

			if err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("min", func(b *testing.B) {
		schema := owl.Float().Min(5)

		for i := 0; i < b.N; i++ {
			err := schema.Validate(5)

			if err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("max", func(b *testing.B) {
		schema := owl.Float().Max(5)

		for i := 0; i < b.N; i++ {
			err := schema.Validate(5)

			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func ExampleFloat() {
	schema := owl.Float()

	if err := schema.Validate(1.0); err != nil { // nil
		panic(err)
	}

	if err := schema.Validate(1); err != nil { // nil
		panic(err)
	}

	if err := schema.Validate("test"); err != nil { // error
		panic(err)
	}
}

func ExampleFloatSchema_Min() {
	schema := owl.Float().Min(5.0)

	if err := schema.Validate(5); err != nil { // nil
		panic(err)
	}

	if err := schema.Validate(4.5); err != nil { // error
		panic(err)
	}
}

func ExampleFloatSchema_Max() {
	schema := owl.Float().Max(5.0)

	if err := schema.Validate(5); err != nil { // nil
		panic(err)
	}

	if err := schema.Validate(5.5); err != nil { // error
		panic(err)
	}
}
