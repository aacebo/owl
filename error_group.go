package owl

import "encoding/json"

type ErrorGroup []error

func (self ErrorGroup) Error() string {
	b, _ := json.Marshal(self)
	return string(b)
}

func (self ErrorGroup) String() string {
	b, _ := json.Marshal(self)
	return string(b)
}
