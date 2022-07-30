package owl

import "testing"

func TestIntMin(t *testing.T) {
	s := Int().Min(0)

	if errors := s.Validate(0); len(errors) > 0 {
		t.Fatal(errors[0].Message)
	}

	if errors := s.Validate(int32(-1)); len(errors) == 0 {
		t.Fatal("should validate min of 0")
	}
}

func TestIntMessage(t *testing.T) {
	s := Int().Min(0).Message("a failure message")
	errors := s.Validate(-1)

	if len(errors) == 0 {
		t.Fatal("should validate min of 0")
	}

	if errors[0].Message != "a failure message" {
		t.Fatal("failure message does not match expected")
	}
}

func TestIntMinMax(t *testing.T) {
	s := Int().Min(0).Max(100)

	if errors := s.Validate(1); len(errors) > 0 {
		t.Fatal("should not fail validation")
	}

	if errors := s.Validate(-1); len(errors) == 0 {
		t.Fatal("should be invalid")
	}

	if errors := s.Validate(101); len(errors) == 0 {
		t.Fatal("should be invalid")
	}
}

func TestIntEqual(t *testing.T) {
	s := Int().Equal(1000)

	if errors := s.Validate(1000); len(errors) > 0 {
		for _, err := range errors {
			t.Log(err.Message)
		}

		t.Fatal("should not fail validation")
	}

	if errors := s.Validate(float32(50)); len(errors) == 0 {
		t.Fatal("should be invalid")
	}
}
