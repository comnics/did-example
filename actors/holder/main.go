package main

import (
	"context"
	"fmt"
	"github.com/comnics/did-example/core"
	"github.com/comnics/did-example/protos"
	"google.golang.org/grpc"
	"log"
	"time"
)

const ISSUER_PB_KEY = "zaSq9DsNNvGhYxYyqA9wd2eduEAZ5AXWgJTbTJJNuMGAQke9NvqhBfPFLmeegjEDkQzjWsFTd8BSYiVCwiSGmNRTeT2trqAmWs8EnoETvpijUYvPT764vGoTC3TWV"

type client struct {
	protos.SimpleIssuerClient
}

type Holder struct {
	kms         *core.ECDSAManager
	did         *core.DID
	didDocument *core.DIDDocument
}

func (holder *Holder) generateDID() {
	// 키생성(ECDSA) - 향후 KMS로 대체.
	holder.kms = core.NewEcdsa()

	// DID 생성.
	issuerDid := core.NewDID("comnic", holder.kms.PublicKeyBase58())

	holder.did = issuerDid

	// DID Document 생성.
	verificationId := fmt.Sprintf("%s#keys-1", issuerDid)
	verificationMethod := []core.VerificationMethod{
		{
			Id:                 verificationId,
			Type:               "EcdsaSecp256k1VerificationKey2019",
			Controller:         issuerDid.String(),
			PublicKeyMultibase: holder.kms.PublicKeyMultibase(),
		},
	}
	didDocument := core.NewDIDDocument(issuerDid.String(), verificationMethod)
	holder.didDocument = didDocument

	registerDid(issuerDid.String(), didDocument)
}

func registerDid(did string, document *core.DIDDocument) error {
	err := core.RegisterDid(did, document.String())
	if err != nil {
		return err
	}
	return nil
}

func main() {

	// New Holder
	holder := new(Holder)
	holder.generateDID()

	// Test
	//didStr, _ := core.ResolveDid(holder.did.String())
	//fmt.Printf("Did Resolve: %s\n", didStr)
	//
	//didDoc := new(core.DIDDocument)
	//didDoc.Consume(didStr)
	//fmt.Printf("Did VerificationMethod: %+v\n", didDoc.VerificationMethod)

	// Connect Issuer
	conn, err := grpc.Dial("localhost:1021", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Issuer not connect: %v", err)
	}
	defer conn.Close()
	c := protos.NewSimpleIssuerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Issuer로 부터 VC를 요청한다.
	res, err := c.IssueSimpleVC(ctx, &protos.MsgIssueVC{
		Did:   holder.did.String(),
		Nonce: "100001",
		ReqVP: "{'Korea eID Credential': '12345'}",
	})
	if err != nil {
		log.Fatalf("could not request: %v", err)
	}

	fmt.Printf("Result: %s\n", res)

	// Verify VC
	//_, bytePubKey, err := multibase.Decode(ISSUER_PB_KEY)
	//pbKey, err := x509.ParsePKIXPublicKey(bytePubKey)
	//if err != nil {
	//	log.Fatalf("key is not valid.")
	//	os.Exit(0)
	//}

	//verify, _ := core.VerifyJwt(res.GetVc(), pbKey.(*ecdsa.PublicKey))
	verify, _, _ := core.ParseJwt(res.GetVc())

	if verify {
		fmt.Println("VC is verified.")
	} else {
		fmt.Println("VC is NOT verified.")
	}
}
