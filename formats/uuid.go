package formats

import (
	"errors"
	"fmt"
	"regexp"
)

var expr = regexp.MustCompile("[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}")

func UUID(input string) error {
	ok := expr.MatchString(input)

	if !ok {
		return errors.New(fmt.Sprintf(
			`"%s" does not match uuid format`,
			input,
		))
	}

	return nil
}
