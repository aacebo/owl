package formats_test

import (
	"testing"

	"github.com/aacebo/owl/formats"
)

func Test_UUID(t *testing.T) {
	t.Run("should error", func(t *testing.T) {
		err := formats.UUID("test")

		if err == nil {
			t.Log("should have error")
			t.FailNow()
		}
	})

	t.Run("should succeed", func(t *testing.T) {
		err := formats.UUID("09946696-054f-46ce-8f87-865f71e6b035")

		if err != nil {
			t.Log(err)
			t.FailNow()
		}
	})
}
