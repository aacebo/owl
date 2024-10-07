package owl_test

import (
	"testing"

	"github.com/aacebo/owl"
)

func Test_Error(t *testing.T) {
	t.Run("should serialize", func(t *testing.T) {
		schema := owl.String()
		err := schema.Validate(1)

		if err == nil {
			t.FailNow()
		}

		if len(err.Error()) != 57 {
			t.Errorf(
				"expected `%d`, received `%d`",
				57,
				len(err.Error()),
			)
		}
	})
}
