package owl

var owl = New()

func AddRule(name string, rule Rule) {
	owl.AddRule(name, rule)
}

func Validate(v any) []Error {
	return owl.Validate(v)
}
