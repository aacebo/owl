package owl

import "reflect"

type Rule struct {
	Key     string `json:"key"`
	Value   any    `json:"value"`
	Resolve RuleFn `json:"-"`
}

type RuleFn func(value reflect.Value) (any, error)
