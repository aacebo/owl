package owl

type Condition struct {
	message string
	fn      func(v any) (any, bool)
}

func NewCondition(message string, fn func(v any) (any, bool)) *Condition {
	v := Condition{message, fn}
	return &v
}

func (self *Condition) Eval(v any) (any, bool) {
	return self.fn(v)
}
