package owl

type Error struct {
	Source  string
	Message string
}

func NewError(source string, message string) *Error {
	v := Error{
		Source:  source,
		Message: message,
	}

	return &v
}
