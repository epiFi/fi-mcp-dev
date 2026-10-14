package pkg

import "testing"

func TestMobileNumberAllowed(t *testing.T) {
	allowed := []string{"1111111111", "9999999999"}

	if !MobileNumberAllowed("1111111111", allowed) {
		t.Fatal("known persona was rejected")
	}
	if !MobileNumberAllowed(" 9999999999 ", allowed) {
		t.Fatal("surrounding spaces should be ignored")
	}
	if MobileNumberAllowed("0000000000", allowed) {
		t.Fatal("unknown phone was accepted")
	}
	if MobileNumberAllowed("1111111111 ", []string{}) {
		t.Fatal("empty allow list should reject every phone")
	}
	if MobileNumberAllowed("   ", allowed) {
		t.Fatal("blank phone was accepted")
	}
}
