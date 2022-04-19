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

	// New DID Document
	didDocument := NewDIDDocument(did.String(), nil)

	// Add VerificationMethod
	didDocument.AddVerificationMethod(verificationId, VERIFICATION_KEY_TYPE_SECP256K1, did.String(), kms.PublicKeyMultibase())

	if didDocument == nil {
		t.Error("failed to generate did-document.")
	}

	if len(didDocument.GetVerificationMethod()) != 1 {
		t.Error("VerificationMethod Count must be 1.")
	}
}
