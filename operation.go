package owl

type Operation struct {
	message string
	fn      func(v any) (any, bool)
}

func NewOperation(message string, fn func(v any) (any, bool)) *Operation {
	v := Operation{message, fn}
	return &v
}

func (self *Operation) Eval(v any) (any, bool) {
	return self.fn(v)
}
