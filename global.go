package owl

import (
	"reflect"

	"github.com/aacebo/owl/rules"
	"github.com/aacebo/owl/types"
)

var validator = New()

func AddRule(name string, rule rules.Rule) *owl {
	return validator.AddRule(name, rule)
}

func AddType(t reflect.Type, fn types.Type) *owl {
	return validator.AddType(t, fn)
}

func AddFormat(name string, formatter Formatter) *owl {
	return validator.AddFormat(name, formatter)
}

func Validate(v any) []Error {
	return validator.Validate(v)
}
