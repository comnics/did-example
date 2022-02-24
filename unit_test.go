package main

import "testing"

// Test ECDSA Keypair generation.
func TestGenerateKeypair(t *testing.T) {
	_, err := GenerateKeypair(keyType)

	if err != nil {
		t.Error("GenerateKeypair is Fail!")
	}
}

func TestEncodeBase58(t *testing.T) {
	keypair, err := GenerateKeypair(keyType)
	if err != nil {
		t.Error("GenerateKeypair is Fail!")
	}

	pvKeyBase58 := EncodeBase58_PrivateKey(keypair.privateKey)
	if pvKeyBase58 == "" {
		t.Error("Fail to encode the PrivateKey.")
	}
	pbKeyBase58 := EncodeBase58_PublicKey(keypair.publicKey)
	if pbKeyBase58 == "" {
		t.Error("Fail to encode the PublicKey.")
	}

}
