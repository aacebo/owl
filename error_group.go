package owl

import "encoding/json"

type ErrorGroup struct {
	Key    string  `json:"key,omitempty"`
	Errors []error `json:"errors,omitempty"`
}

func newErrorGroup(key string) ErrorGroup {
	return ErrorGroup{
		Key:    key,
		Errors: []error{},
	}
}

func (self ErrorGroup) Add(err error) ErrorGroup {
	self.Errors = append(self.Errors, err)
	return self
}

func (self ErrorGroup) Error() string {
	b, _ := json.MarshalIndent(self, "", "  ")
	return string(b)
}

func (self ErrorGroup) String() string {
	b, _ := json.Marshal(self)
	return string(b)
}
