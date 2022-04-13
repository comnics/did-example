package core

import "testing"

func TestNewRsa(t *testing.T) {
	_, err := NewRsa()

	if err != nil {
		t.Error("Failed to generate RAS.")
	}
}

func TestRsaSignAndVerify(t *testing.T) {
	rsam, err := NewRsa()
	if err != nil {
		t.Error("Failed to generate RAS.")
	}

	msg := "Hello World."
	sig, err := rsam.Sign(msg)
	if err != nil {
		t.Error("Failed to Sign.")
	}

	res := rsam.Verify(msg, sig)

	if !res {
		t.Error("Failed to Verify.")
	}

}
