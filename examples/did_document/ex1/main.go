// examples/did_document/ex1/main.go

package main

import (
	"fmt"
	"github.com/comnics/did-example/core"
	"log"
)

func main() {
	var method = "ssikr"

	// 1. 1. 키생성(ECDSA)
	kms := new(core.ECDSAManager)
	kms.Generate()

	// 2. DID 생성.
	did, err := core.NewDID(method, kms.PublicKeyMultibase())
	if err != nil {
		log.Printf("Failed to generate DID, error: %v\n", err)
	}

	// 3. DID Document 생성.
	verificationId := fmt.Sprintf("%s#keys-1", did)
	verificationMethod := []core.VerificationMethod{
		{
			Id:                 verificationId,
			Type:               core.VERIFICATION_KEY_TYPE_SECP256K1,
			Controller:         did.String(),
			PublicKeyMultibase: kms.PublicKeyMultibase(),
		},
	}
	didDocument := core.NewDIDDocument(did.String(), verificationMethod)
	verificationId2 := fmt.Sprintf("%s#keys-2", did)
	didDocument.AddVerificationMethod(verificationId2, core.VERIFICATION_KEY_TYPE_SECP256K1, did.String(), "Second pub-key")

	fmt.Println("### Generate DID & DID Document ###")
	fmt.Printf("did => %s\n", did)
	fmt.Printf("did document => %+v\n", didDocument)

	vm := didDocument.FindVerificationMethod(verificationId)
	if vm != nil {
		fmt.Println("VerificaionMethod: ", vm)
	}

	vm2 := didDocument.FindVerificationMethod(verificationId2)
	if vm2 != nil {
		fmt.Println("VerificaionMethod: ", vm2)
	}

}
