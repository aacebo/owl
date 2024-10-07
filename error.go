package owl

import (
	"encoding/json"
)

type Error struct {
	Rule    string `json:"rule,omitempty"`
	Key     string `json:"key,omitempty"`
	Message string `json:"message,omitempty"`
}

func NewError(rule string, key string, message string) Error {
	return Error{
		Rule:    rule,
		Key:     key,
		Message: message,
	}
}

func (self Error) Error() string {
	b, _ := json.Marshal(self)
	return string(b)
}

func (self Error) String() string {
	b, _ := json.Marshal(self)
	return string(b)
}
