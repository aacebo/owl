package formats_test

import (
	"testing"

	"github.com/aacebo/owl/formats"
)

func Test_URI(t *testing.T) {
	t.Run("should error", func(t *testing.T) {
		err := formats.URI("hello/world")

		if err == nil {
			t.Log("should have error")
			t.FailNow()
		}
	})

	t.Run("should error on no scheme", func(t *testing.T) {
		err := formats.URI("google.com")

		if err == nil {
			t.Log("should have error")
			t.FailNow()
		}
	})

	t.Run("should succeed", func(t *testing.T) {
		err := formats.URI("https://google.com")

		if err != nil {
			t.Log(err)
			t.FailNow()
		}
	})
}
