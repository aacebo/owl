package owl_test

import (
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

	t.Run("min", func(t *testing.T) {
		t.Run("should succeed", func(t *testing.T) {
			err := owl.Time().Min(time.Now().AddDate(-1, 0, 0)).Validate(time.Now().Format(time.RFC3339))

			if err != nil {
				t.Fatal(err.Error())
			}
		})

		t.Run("should fail", func(t *testing.T) {
			err := owl.Time().Min(time.Now()).Validate(time.Now().AddDate(-1, 0, 0).Format(time.RFC3339))

			if err == nil {
				t.Fatal()
			}
		})
	})

	t.Run("max", func(t *testing.T) {
		t.Run("should succeed", func(t *testing.T) {
			err := owl.Time().Max(time.Now().AddDate(1, 0, 0)).Validate(time.Now())

			if err != nil {
				t.Fatal(err.Error())
			}
		})

		t.Run("should fail", func(t *testing.T) {
			err := owl.Time().Max(time.Now()).Validate(time.Now().AddDate(1, 0, 0))

			if err == nil {
				t.Fatal()
			}
		})
	})
}
