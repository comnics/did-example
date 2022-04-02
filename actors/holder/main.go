package main

import (
	"context"
	"crypto/ecdsa"
	"crypto/x509"
	"fmt"
	"github.com/comnics/did-example/core"
	"github.com/comnics/did-example/protos"
	"github.com/multiformats/go-multibase"
	"google.golang.org/grpc"
	"log"
	"os"
	"time"
)

const ISSUER_PB_KEY = "zaSq9DsNNvGhYxYyqA9wd2eduEAZ5AXWgJTbTJJNuMGAQke9NvqhBfPFLmeegjEDkQzjWsFTd8BSYiVCwiSGmNRTeT2trqAmWs8EnoETvpijUYvPT764vGoTC3TWV"

type client struct {
	protos.SimpleIssuerClient

	kms         *core.ECDSAManager
	did         *core.DID
	didDocument *core.DIDDocument
}

func (holder *client) generateDID() {
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

func registerDid(did string, document *core.DIDDocument) {

}

func main() {
	conn, err := grpc.Dial("localhost:1021", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := protos.NewSimpleIssuerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := c.IssueSimpleVC(ctx, &protos.MsgIssueVC{
		Did:   "76f5d6b9-9429-477e-9bc5-91bf2769ebbe",
		Nonce: "100001",
		ReqVP: "{'a': 'b'}",
	})
	if err != nil {
		log.Fatalf("could not request: %v", err)
	}

	fmt.Printf("Result: %s\n", res)

	_, bytePubKey, err := multibase.Decode(ISSUER_PB_KEY)
	pbKey, err := x509.ParsePKIXPublicKey(bytePubKey)
	if err != nil {
		log.Fatalf("key is not valid.")
		os.Exit(0)
	}

	verify, _ := core.VerifyJwt(res.GetVc(), pbKey.(*ecdsa.PublicKey))
	if verify {
		fmt.Println("VC is verified.")
	} else {
		fmt.Println("VC is NOT verified.")
	}

}
