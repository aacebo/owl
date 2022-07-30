package owl

type Operation[T any] struct {
	Message string
	fn      func(v T) bool
}

func NewOperation[T any](message string, fn func(v T) bool) *Operation[T] {
	v := Operation[T]{
		Message: message,
		fn:      fn,
	}

	return &v
}

func (self *Operation[T]) Eval(v T) bool {
	return self.fn(v)
}
