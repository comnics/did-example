package core

import (
	"fmt"
	"testing"
)

func TestDidDocument(t *testing.T) {
	method := "comnic"
	// 키생성(ECDSA) - 향후 KMS로 대체.
	kms := NewEcdsa()

	// DID 생성.
	did, _ := NewDID(method, kms.PublicKeyBase58())

	// DID Document 생성.
	verificationId := fmt.Sprintf("%s#keys-1", did)
	verificationMethod := []VerificationMethod{
		{
			Id:                 verificationId,
			Type:               "EcdsaSecp256k1VerificationKey2019",
			Controller:         did.String(),
			PublicKeyMultibase: kms.PublicKeyMultibase(),
		},
	}
	didDocument := NewDIDDocument(did.String(), verificationMethod)

	if didDocument == nil {
		t.Error("failed to generate did-document.")
	}
}
