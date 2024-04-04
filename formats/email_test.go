package formats_test

import (
	"testing"

	"github.com/aacebo/owl/formats"
)

func Test_Email(t *testing.T) {
	t.Run("should error", func(t *testing.T) {
		err := formats.Email("test")

		if err == nil {
			t.Log("should have error")
			t.FailNow()
		}
	})

	t.Run("should succeed", func(t *testing.T) {
		err := formats.Email("test@test.com")

		if err != nil {
			t.Log(err)
			t.FailNow()
		}
	})
}
