package owl

type Error struct {
	Source  string
	Message string
	Path    []string
}

func NewError(source string, message string, path []string) *Error {
	v := Error{
		Source:  source,
		Message: message,
		Path:    path,
	}

	return &v
}
