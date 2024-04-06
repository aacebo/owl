package owl

import (
	"reflect"

	"github.com/aacebo/owl/rules"
)

var validator = New()

func AddRule(name string, rule rules.Rule, dependsOn ...string) *owl {
	return validator.AddRule(name, rule, dependsOn...)
}

func AddTransform(t reflect.Type, fn Transform) *owl {
	return validator.AddTransform(t, fn)
}

func AddFormat(name string, formatter Formatter) *owl {
	return validator.AddFormat(name, formatter)
}

func Validate(v any) []Error {
	return validator.Validate(v)
}
