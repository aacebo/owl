package owl_test

import (
	"testing"

	"github.com/aacebo/owl"
)

func Test_Union(t *testing.T) {
	t.Run("union", func(t *testing.T) {
		t.Run("should succeed", func(t *testing.T) {
			err := owl.Union(
				owl.String().Required(),
				owl.Int().Required(),
			).Validate("test")

			if err != nil {
				t.Fatal(err.Error())
			}
		})

		t.Run("should fail", func(t *testing.T) {
			err := owl.Union(
				owl.String().Required(),
				owl.Int().Required(),
			).Validate(true)

			if err == nil {
				t.Fatal()
			}
		})
	})
}
