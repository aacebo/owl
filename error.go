package owl

import (
	"encoding/json"
)

type Error struct {
	Key     string  `json:"key"`
	Message string  `json:"message,omitempty"`
	Errors  []error `json:"errors,omitempty"`
}

func newError(key string, message string) Error {
	return Error{
		Key:     key,
		Message: message,
		Errors:  []error{},
	}
}

func (self Error) Add(err error) Error {
	self.Errors = append(self.Errors, err)
	return self
}

func (self Error) Error() string {
	b, _ := json.MarshalIndent(self, "", "  ")
	return string(b)
}

func (self Error) String() string {
	b, _ := json.Marshal(self)
	return string(b)
}
