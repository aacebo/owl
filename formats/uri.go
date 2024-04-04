package formats

import (
	"errors"
	"fmt"
	"net/url"
)

func URI(input string) error {
	_, err := url.ParseRequestURI(input)

	if err != nil {
		return errors.New(fmt.Sprintf(
			`"%s" does not match uri format`,
			input,
		))
	}

	return nil
}
