package transforms

import (
	"database/sql/driver"
	"reflect"
)

func Valuer(value reflect.Value) reflect.Value {
	v, err := value.Interface().(driver.Valuer).Value()

	if err != nil {
		panic(err)
	}

	return reflect.ValueOf(v)
}
