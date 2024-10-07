package owl_test

import (
	"testing"

	"github.com/aacebo/owl"
)

func Test_ErrorGroup(t *testing.T) {
	t.Run("should serialize", func(t *testing.T) {
		schema := owl.Object().Fields(map[string]owl.Schema{
			"a": owl.String().Required(),
			"b": owl.Bool(),
			"c": owl.Int(),
		})

		err := schema.Validate(map[string]any{
			"b": 1.0,
			"c": true,
		})

		if err == nil {
			t.FailNow()
		}

		if len(err.Error()) != 507 {
			t.Errorf(
				"expected `%d`, received `%d`",
				507,
				len(err.Error()),
			)
		}
	})
}
