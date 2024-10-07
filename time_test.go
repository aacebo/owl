package owl_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/aacebo/owl"
)

func Test_Time(t *testing.T) {
	t.Run("required", func(t *testing.T) {
		t.Run("should succeed", func(t *testing.T) {
			err := owl.Time().Required().Validate(time.Now().Format(time.RFC3339))

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

	t.Run("type", func(t *testing.T) {
		t.Run("should succeed when `time.Time`", func(t *testing.T) {
			err := owl.Time().Validate(time.Now())

			if err != nil {
				t.Fatal(err.Error())
			}
		})

		t.Run("should succeed when time.Time string", func(t *testing.T) {
			err := owl.Time().Validate(time.Now().Format(time.RFC3339))

			if err != nil {
				t.Fatal(err.Error())
			}
		})

		t.Run("should fail when invalid string", func(t *testing.T) {
			err := owl.Time().Validate(time.Now().String())

			if err == nil {
				t.FailNow()
			}
		})

		t.Run("should fail when not string/time.Time", func(t *testing.T) {
			err := owl.Time().Validate(true)

			if err == nil {
				t.FailNow()
			}
		})
	})

	t.Run("min", func(t *testing.T) {
		t.Run("should succeed when nil", func(t *testing.T) {
			err := owl.Time().Min(time.Now().AddDate(-1, 0, 0)).Validate(nil)

			if err != nil {
				t.Fatal(err.Error())
			}
		})

		t.Run("should succeed when gt min", func(t *testing.T) {
			err := owl.Time().Min(time.Now().AddDate(-1, 0, 0)).Validate(time.Now().Format(time.RFC3339))

			if err != nil {
				t.Fatal(err.Error())
			}
		})

		t.Run("should fail when lt min", func(t *testing.T) {
			err := owl.Time().Min(time.Now()).Validate(time.Now().AddDate(-1, 0, 0).Format(time.RFC3339))

			if err == nil {
				t.Fatal()
			}
		})
	})

	t.Run("max", func(t *testing.T) {
		t.Run("should succeed when nil", func(t *testing.T) {
			err := owl.Time().Max(time.Now().AddDate(1, 0, 0)).Validate(nil)

			if err != nil {
				t.Fatal(err.Error())
			}
		})

		t.Run("should succeed when lt max", func(t *testing.T) {
			err := owl.Time().Max(time.Now().AddDate(1, 0, 0)).Validate(time.Now())

			if err != nil {
				t.Fatal(err.Error())
			}
		})

		t.Run("should fail when gt max", func(t *testing.T) {
			err := owl.Time().Max(time.Now()).Validate(time.Now().AddDate(1, 0, 0))

			if err == nil {
				t.Fatal()
			}
		})
	})

	t.Run("json", func(t *testing.T) {
		t.Run("serialize", func(t *testing.T) {
			schema := owl.Time()
			b, err := json.Marshal(schema)

			if err != nil {
				t.Error(err)
			}

			if string(b) != `{"type":"time"}` {
				t.Errorf(
					"expected `%s`, received `%s`",
					`{"type":"time"}`,
					string(b),
				)
			}
		})
	})
}

func ExampleTime() {
	schema := owl.Time()

	if err := schema.Validate(time.Now()); err != nil { // nil
		panic(err)
	}

	if err := schema.Validate(time.Now().Format(time.RFC3339)); err != nil { // nil
		panic(err)
	}

	if err := schema.Validate("test"); err != nil { // error
		panic(err)
	}
}

func ExampleTimeSchema_Min() {
	schema := owl.Time().Min(time.Now())

	if err := schema.Validate(time.Now().AddDate(-1, 0, 0)); err != nil { // error
		panic(err)
	}
}

func ExampleTimeSchema_Max() {
	schema := owl.Time().Max(time.Now())

	if err := schema.Validate(time.Now().AddDate(1, 0, 0)); err != nil { // error
		panic(err)
	}
}
