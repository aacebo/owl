package formats

import (
	"errors"
	"fmt"
	"net/mail"
)

func Email(input string) error {
	_, err := mail.ParseAddress(input)

	if err != nil {
		return errors.New(fmt.Sprintf(
			`"%s" does not match email format`,
			input,
		))
	}

	return nil
}
