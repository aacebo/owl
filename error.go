package owl

import (
	"encoding/json"
)

type Error struct {
	Rule    string  `json:"rule,omitempty"`
	Key     string  `json:"key,omitempty"`
	Message string  `json:"message,omitempty"`
	Errors  []error `json:"errors,omitempty"`
}

func NewError(rule string, key string, message string) Error {
	return Error{
		Rule:    rule,
		Key:     key,
		Message: message,
		Errors:  []error{},
	}
}

func NewEmptyError(rule string, key string) Error {
	return Error{
		Rule:   rule,
		Key:    key,
		Errors: []error{},
	}
}

func (self Error) Add(err error) Error {
	self.Errors = append(self.Errors, err)
	return self
}

func (self Error) Error() string {
	b, _ := json.Marshal(self)
	return string(b)
}

func (self Error) String() string {
	b, _ := json.Marshal(self)
	return string(b)
}
