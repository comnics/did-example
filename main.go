package main

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/comnics/did-example/core"
	"github.com/comnics/did-example/util"
)

// Simple KMS
type Keypair struct {
	privateKey ecdsa.PrivateKey //interface{}
	publicKey  ecdsa.PublicKey  //interface{}
}

const (
	keyType = "ecdsa"
)

var (
	logger util.Logger
)

// Main
func main() {
	//logger := util.GetLogger()

	fmt.Println("### Start Main() ###")

	// ECDSA secp256k1
	var ecdsa core.ECDSAManager // ecdsa := new(core.ECDSAManager)
	ecdsa.Generate()

	did := core.NewDID("comnic", ecdsa.PublicKeyBase58())
	fmt.Printf("DID : [%s]\n", did)

	didDocument := core.NewDIDDocument(did.String())
	fmt.Printf("DID Document: %s", didDocument)

	//// protobuf test
	//testDid := pb.Did{Id: "12342", Desc: "abcde"}
	//
	//fmt.Printf("Test Did: %s\n", testDid.GetId())

}

func MakeDIDDocument(method string, did string) core.DIDDocument {

	doc := core.DIDDocument{Id: did}

	return doc
}
