package owl

import "testing"

type ATestStruct struct {
	A int
	B int
}

type BTestStruct struct {
	A int
	B CTestStruct
}

type CTestStruct struct {
	One int
}

func TestStruct(t *testing.T) {
	s := Struct(StructKeys{
		"A": Int().Min(0).Max(100).Required(),
		"B": Struct(StructKeys{
			"One": Int().Equal(1000).Message("oops wrong answer"),
		}),
	})

	_, errors := s.Validate(ATestStruct{
		A: 50,
		B: 1000,
	})

	if len(errors) != 1 || errors[0].Message != "must be a structure" {
		t.Fatal("should have one error")
	}

	_, errors = s.Validate(&BTestStruct{
		A: 50,
		B: CTestStruct{
			One: 50,
		},
	})

	if len(errors) != 1 || errors[0].Message != "oops wrong answer" {
		for _, err := range errors {
			t.Log(err.Message)
		}

		t.Fatal("should have one error")
	}
}
