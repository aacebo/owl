package owl_test

import (
	"testing"

	"github.com/aacebo/owl"
)

func Test_Int(t *testing.T) {
	t.Run("required", func(t *testing.T) {
		t.Run("should succeed", func(t *testing.T) {
			err := owl.Int().Required().Validate(1)

			if err != nil {
				t.Fatal(err.Error())
			}
		})

		t.Run("should fail", func(t *testing.T) {
			err := owl.Int().Required().Validate(nil)

			if err == nil {
				t.Fatal()
			}
		})
	})

	t.Run("enum", func(t *testing.T) {
		t.Run("should succeed", func(t *testing.T) {
			err := owl.Int().Enum(1).Validate(1)

			if err != nil {
				t.Fatal(err.Error())
			}
		})

		t.Run("should fail", func(t *testing.T) {
			err := owl.Int().Enum(1).Validate(2)

			if err == nil {
				t.Fatal()
			}
		})
	})

	t.Run("min", func(t *testing.T) {
		t.Run("should succeed", func(t *testing.T) {
			err := owl.Int().Min(5).Validate(6)

			if err != nil {
				t.Fatal(err.Error())
			}
		})

		t.Run("should fail", func(t *testing.T) {
			err := owl.Int().Min(5).Validate(4)

			if err == nil {
				t.Fatal()
			}
		})
	})

	t.Run("max", func(t *testing.T) {
		t.Run("should succeed", func(t *testing.T) {
			err := owl.Int().Max(5).Validate(4)

			if err != nil {
				t.Fatal(err.Error())
			}
		})

		t.Run("should fail", func(t *testing.T) {
			err := owl.Int().Max(5).Validate(6)

			if err == nil {
				t.Fatal()
			}
		})
	})
}

func ExampleInt() {
	schema := owl.Int()

	if err := schema.Validate(1); err != nil { // nil
		panic(err)
	}

	if err := schema.Validate(1); err != nil { // nil
		panic(err)
	}

	if err := schema.Validate("test"); err != nil { // error
		panic(err)
	}
}

func ExampleIntSchema_Min() {
	schema := owl.Int().Min(5)

	if err := schema.Validate(5); err != nil { // nil
		panic(err)
	}

	if err := schema.Validate(4); err != nil { // error
		panic(err)
	}
}

func ExampleIntSchema_Max() {
	schema := owl.Int().Max(5)

	if err := schema.Validate(5); err != nil { // nil
		panic(err)
	}

	if err := schema.Validate(6); err != nil { // error
		panic(err)
	}
}
