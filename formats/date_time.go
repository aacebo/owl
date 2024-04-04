package formats

import (
	"errors"
	"fmt"
	"time"
)

func DateTime(input string) error {
	_, err := time.Parse(time.RFC3339, input)

	if err != nil {
		return errors.New(fmt.Sprintf(
			`"%s" does not match date-time format "YYYY-MM-DDTmm:ss:ms"`,
			input,
		))
	}

	return nil
}
