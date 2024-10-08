package ordered_map

import (
	"bytes"
	"encoding/json"
	"fmt"
	"slices"
)

type Item[K comparable, V any] struct {
	Key   K
	Value V
}

type Map[K comparable, V any] []Item[K, V]

func (self Map[K, V]) Has(key K) bool {
	i := slices.IndexFunc(self, func(v Item[K, V]) bool {
		return v.Key == key
	})

	return i > -1
}

func (self Map[K, V]) Get(key K) (V, bool) {
	var value V
	exists := false

	i := slices.IndexFunc(self, func(v Item[K, V]) bool {
		return v.Key == key
	})

	if i > -1 {
		value = self[i].Value
		exists = true
	}

	return value, exists
}

func (self *Map[K, V]) Set(key K, value V) {
	i := slices.IndexFunc(*self, func(v Item[K, V]) bool {
		return v.Key == key
	})

	if i > -1 {
		(*self)[i].Key = key
		(*self)[i].Value = value
		return
	}

	*self = append(*self, Item[K, V]{
		Key:   key,
		Value: value,
	})
}

func (self Map[K, V]) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	buf.Write([]byte{'{'})

	for i, item := range self {
		b, err := json.Marshal(item.Value)

		if err != nil {
			return nil, err
		}

		buf.WriteString(fmt.Sprintf("%q:", fmt.Sprintf("%v", item.Key)))
		buf.Write(b)

		if i < len(self)-1 {
			buf.Write([]byte{','})
		}
	}

	buf.Write([]byte{'}'})
	return buf.Bytes(), nil
}
