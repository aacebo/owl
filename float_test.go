package owl_test

import (
	"encoding/json"
	"testing"

	"github.com/aacebo/owl"
)

func Test_Float(t *testing.T) {
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

	t.Run("min", func(t *testing.T) {
		t.Run("should succeed", func(t *testing.T) {
			err := owl.Float().Min(5).Validate(6)

			if err != nil {
				t.Fatal(err.Error())
			}
		})

		t.Run("should fail", func(t *testing.T) {
			err := owl.Float().Min(5).Validate(4)

			if err == nil {
				t.Fatal()
			}
		})
	})

	t.Run("max", func(t *testing.T) {
		t.Run("should succeed", func(t *testing.T) {
			err := owl.Float().Max(5).Validate(4)

			if err != nil {
				t.Fatal(err.Error())
			}
		})

		t.Run("should fail", func(t *testing.T) {
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

			if string(b) != `{"max":5,"min":1,"type":"float"}` {
				t.Errorf(
					"expected `%s`, received `%s`",
					`{"max":5,"min":1,"type":"float"}`,
					string(b),
				)
			}
		})
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
