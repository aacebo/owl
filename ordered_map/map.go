package ordered_map

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Item[K comparable, V any] struct {
	Key   K
	Value V
}

type Map[K comparable, V any] []Item[K, V]

func (self *Map[K, V]) Set(key K, value V) {
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
