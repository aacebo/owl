package formats_test

import (
	"testing"

	"github.com/aacebo/owl/formats"
)

func Test_IPv4(t *testing.T) {
	t.Run("should error", func(t *testing.T) {
		err := formats.IPv4("test")

		if err == nil {
			t.Log("should have error")
			t.FailNow()
		}
	})

	t.Run("should error on ipv6", func(t *testing.T) {
		err := formats.IPv4("2603:7000:873a:2a48:3d6c:bce7:fe8d:3b8")

		if err == nil {
			t.Log("should have error")
			t.FailNow()
		}
	})

	t.Run("should succeed", func(t *testing.T) {
		err := formats.IPv4("127.0.0.1")

		if err != nil {
			t.Log(err)
			t.FailNow()
		}
	})
}
