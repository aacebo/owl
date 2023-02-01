package owl

import "testing"

func TestBool(t *testing.T) {
	s := Bool()

	if _, errors := s.Validate("test"); len(errors) == 0 {
		t.Fatal("should validate boolean")
	}
}
