package core

import (
	"crypto/sha256"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEcdsaGenerate(t *testing.T) {
	var ecdsa ECDSAManager // ecdsa := new(core.ECDSAManager)
	err := ecdsa.Generate()

	require.NoError(t, err)
}

func TestEcdsaSign(t *testing.T) {
	ecdsa := NewEcdsa()

	msg := "Hello World."
	digest := sha256.Sum256([]byte(msg))
	_, err := ecdsa.Sign(digest[:])
	if err != nil {
		t.Error("Fail to Sign.")
	}
}

func ExampleEcdsa() {
	var ecdsa ECDSAManager // ecdsa := new(core.ECDSAManager)
	err := ecdsa.Generate()

	msg := "Hello World."
	digest := sha256.Sum256([]byte(msg))
	signature, err := ecdsa.Sign(digest[:])
	if err != nil {
		fmt.Printf("Fail to sign to msg.")
	}

	ret := ecdsa.Verify(signature, digest[:])

	if ret {
		fmt.Printf("Signature verifies")
	} else {
		fmt.Printf("Signature does not verify")
	}

	// Output: Signature verifies
}
