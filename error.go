package owl

import "fmt"

type Error struct {
	Path    string `json:"path"`
	Keyword string `json:"keyword"`
	Message string `json:"message"`
}

func (self Error) Error() string {
	return fmt.Sprintf("[%s/%s]: %s", self.Path, self.Keyword, self.Message)
}
