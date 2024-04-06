package transforms_test

import (
	"database/sql/driver"
	"testing"

	"github.com/aacebo/owl"
)

type TestValuer struct {
	Valid bool
	Other string
}

func (self TestValuer) Value() (driver.Value, error) {
	return self.Other, nil
}

func Test_Valuer(t *testing.T) {
	t.Run("should transform value", func(t *testing.T) {
		v := struct {
			A driver.Valuer `json:"a" owl:"min=3"`
		}{TestValuer{true, "hii"}}

		errs := owl.Validate(&v)

		if len(errs) > 0 {
			t.Error(errs)
		}
	})
}
