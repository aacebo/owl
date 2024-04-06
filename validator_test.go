package owl_test

import (
	"testing"

	"github.com/aacebo/owl"
)

type Address struct {
	Street string `json:"street" owl:"required"`
	City   string `json:"city" owl:"required"`
	State  string `json:"state" owl:"required"`
	Zip    int    `json:"zip" owl:"required,min=00501,max=89049"`
}

func Test_Validator(t *testing.T) {
	t.Run("should succeed with nested complex structure", func(t *testing.T) {
		errs := owl.Validate(struct {
			Name    string   `json:"name" owl:"required,min=3"`
			Address *Address `json:"address"`
		}{
			Name: "hii",
			Address: &Address{
				Street: "111 My Street Way",
				City:   "New York",
				State:  "New York",
				Zip:    11249,
			},
		})

		if len(errs) > 0 {
			t.Error(errs)
		}
	})
}

func Benchmark_Validator(b *testing.B) {
	b.Run("should succeed with nested complex structure", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			errs := owl.Validate(struct {
				Name    string   `json:"name" owl:"required,min=3"`
				Address *Address `json:"address"`
			}{
				Name: "hii",
				Address: &Address{
					Street: "111 My Street Way",
					City:   "New York",
					State:  "New York",
					Zip:    11249,
				},
			})

			if len(errs) > 0 {
				b.Error(errs)
			}
		}
	})
}
