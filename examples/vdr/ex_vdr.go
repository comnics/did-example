package main

import (
	"fmt"
	"github.com/comnics/did-example/core"
	"log"
)

func main() {
	var method = "comnic"

	kms := new(core.ECDSAManager)
	kms.Generate()

	did, err := core.NewDID(method, kms.PublicKeyMultibase())

	if err != nil {
		log.Printf("Failed to generate DID, error: %v\n", err)
	}

	// DID Document 생성.
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

	fmt.Println("### New DID ###")
	fmt.Printf("did => %s\n", did)
	fmt.Printf("did document => %+v\n", didDocument)

	RegisterDid(did.String(), didDocument)

	//Resolve한다.
	didDocumentStr, err := core.ResolveDid(did.String())
	if err != nil {
		log.Printf("Failed to Resolve DID.\nError: %x\n", err)
	}

	fmt.Printf("did document => %+v\n", didDocumentStr)
}

func RegisterDid(did string, document *core.DIDDocument) error {
	err := core.RegisterDid(did, document.String())
	if err != nil {
		return err
	}
	return nil
}
