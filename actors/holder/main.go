package main

import (
	"context"
	"fmt"
	"github.com/comnics/did-example/core"
	"github.com/comnics/did-example/protos"
	"google.golang.org/grpc"
	"log"
	"os"
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

func SubmitVP(vpToken string) error {
	conn, err := grpc.Dial("localhost:1022", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Verifier not connect: %v", err)
		return nil
	}
	defer conn.Close()
	c := protos.NewVerifierClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Issuer로 부터 VC를 요청한다.
	res, err := c.SubmitVP(ctx, &protos.SubmitVPRequest{
		Vp: vpToken,
	})
	if err != nil {
		log.Fatalf("could not request: %v", err)
		return nil
	}

	fmt.Printf("Verifier's response: %s\n", res.Result)

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

	// VC를 검증한다.
	vcToken := res.GetVc()
	verify, claims, _ := core.ParseAndVerifyJwtForVC(vcToken)

	if !verify {
		log.Fatal("VC is NOT verified.")
	}

	fmt.Println("VC is verified.")
	fmt.Println("claims Issuer: ", claims.Issuer)

	// send VP to Verifier.

	vcList := []string{}
	vcList = append(vcList, vcToken)

	vp, err := core.NewVP("", []string{"", ""}, "", vcList)
	if err != nil {
		fmt.Println("ERROR")
		os.Exit(0)
	}

	vpToken := vp.GenerateJWT(holder.didDocument.VerificationMethod[0].Id, holder.kms.PrivateKey)
	fmt.Printf("VP: %s\n", vpToken)

	SubmitVP(vpToken)

}
