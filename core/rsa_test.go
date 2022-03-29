package core

import "testing"

func TestNewRsa(t *testing.T) {
	_, err := NewRsa()

	if err != nil {
		t.Error("Fail to New RAS.")
	}
}

func TestRsaSignAndVerify(t *testing.T) {
	rsam, err := NewRsa()
	if err != nil {
		t.Error("Fail to New RAS.")
	}

	msg := "Hello World."
	sig, err := rsam.Sign(msg)
	if err != nil {
		t.Error("Fail to Sign.")
	}

	res := rsam.Verify(msg, sig)

	if !res {
		t.Error("Fail to Verify.")
	}

}
