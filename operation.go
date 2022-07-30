package owl

type Operation struct {
	Message string
	fn      func(v any) (any, bool)
}

func NewOperation(message string, fn func(v any) (any, bool)) *Operation {
	v := Operation{
		Message: message,
		fn:      fn,
	}

	return &v
}

func (self *Operation) Eval(v any) (any, bool) {
	return self.fn(v)
}
