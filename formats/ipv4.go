package formats

import (
	"errors"
	"fmt"
	"net"
)

func IPv4(input string) error {
	ip := net.ParseIP(input)

	if ip == nil {
		return errors.New(fmt.Sprintf(
			`"%s" does not match ipv4 format`,
			input,
		))
	}

	if ip.To4() == nil {
		return errors.New(fmt.Sprintf(
			`"%s" does not match ipv4 format`,
			input,
		))
	}

	return nil
}
