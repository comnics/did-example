package main

import "testing"

// Test ECDSA Keypair generation.
func TestGenerateKeypair(t *testing.T) {
	_, err := GenerateKeypair(keyType)

	if err != nil {
		t.Error("GenerateKeypair is Fail!")
	}
}
