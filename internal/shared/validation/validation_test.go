package validation

import "testing"

func TestPhone(t *testing.T) {
	if err := Phone("0900123456"); err != nil {
		t.Fatalf("expected valid phone: %v", err)
	}
	if err := Phone("abc"); err == nil {
		t.Fatalf("expected invalid phone")
	}
}
