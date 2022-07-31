package owl

import "reflect"

func getTypeValue(v any) (reflect.Value, reflect.Type) {
	var value reflect.Value

	if reflect.TypeOf(v).Kind() != reflect.Pointer {
		value = reflect.ValueOf(v)
	} else {
		value = reflect.Indirect(reflect.ValueOf(v))
	}

	return value, reflect.TypeOf(value.Interface())
}
