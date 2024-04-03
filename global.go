package owl

var validator = New()

func AddRule(name string, rule Rule) {
	validator.AddRule(name, rule)
}

func Validate(v any) []Error {
	return validator.Validate(v)
}
