package owl_test

import (
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
