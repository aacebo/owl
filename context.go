package owl

import (
	"reflect"
	"strings"
)

type context struct {
	rule      string
	schema    map[string]string
	root      reflect.Value
	parent    reflect.Value
	value     reflect.Value
	field     reflect.StructField
	hasFormat func(string) bool
	format    func(string, string) error
}

func (self context) Root() reflect.Value {
	return self.root
}

func (self context) Parent() reflect.Value {
	return self.parent
}

func (self context) Rule() string {
	return self.rule
}

func (self context) Param() string {
	return self.schema[self.rule]
}

func (self context) RuleParam(name string) (string, bool) {
	v, ok := self.schema[name]
	return v, ok
}

func (self context) Name() string {
	name := self.field.Name

	if tag := self.field.Tag.Get("json"); tag != "" {
		parts := strings.Split(tag, ",")

		if parts[0] != "-" {
			name = parts[0]
		}
	}

	return name
}

func (self context) FieldName() string {
	return self.field.Name
}

func (self context) Value() reflect.Value {
	return self.value
}

func (self context) CoerceValue() reflect.Value {
	return reflect.Indirect(self.value)
}

func (self context) HasFormat(name string) bool {
	return self.hasFormat(name)
}

func (self context) Format(name string, text string) error {
	return self.format(name, text)
}
