package rules

import "reflect"

type Context interface {
	// get the root struct
	Root() reflect.Value

	// get the parent struct
	Parent() reflect.Value

	// get the field value
	Value() reflect.Value

	// get coerced field value, indirecting pointers and interfaces
	CoerceValue() reflect.Value

	// get the field "json" name if exists, otherwise get the struct field name
	Name() string

	// get the struct field name
	FieldName() string

	// get the current rule name
	Rule() string

	// get the current rules parameter
	Param() string

	// get a rules parameter
	RuleParam(name string) (string, bool)

	// check if instance has formatter
	HasFormat(name string) bool

	// format text using formatter
	Format(name string, text string) error
}
