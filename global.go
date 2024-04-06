package owl

import "github.com/aacebo/owl/rules"

var validator = New()

func AddRule(name string, rule rules.Rule) {
	validator.AddRule(name, rule)
}

func AddFormat(name string, formatter Formatter) {
	validator.AddFormat(name, formatter)
}

func Validate(v any) []Error {
	return validator.Validate(v)
}
