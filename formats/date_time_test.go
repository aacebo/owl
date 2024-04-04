package formats_test

import (
	"testing"

	"github.com/aacebo/owl/formats"
)

func Test_DateTime(t *testing.T) {
	t.Run("should error", func(t *testing.T) {
		err := formats.DateTime("3/9/2024, 3:39:46 PM")

		if err == nil {
			t.Log("should have error")
			t.FailNow()
		}
	})

	t.Run("should succeed", func(t *testing.T) {
		err := formats.DateTime("2024-03-09T20:39:33.335Z")

		if err != nil {
			t.Log(err)
			t.FailNow()
		}
	})
}
