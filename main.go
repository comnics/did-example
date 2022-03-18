package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"errors"
	"github.com/comnics/did-example/core"
	"github.com/comnics/did-example/util"
	"log"

	"github.com/btcsuite/btcutil/base58"

	pb "github.com/comnics/did-example/protos"
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
	logger := util.GetLogger()

	logger.Info.Println("### Start Main() ###")

	// Generate ECDSA keypair
	myKeypair, err := GenerateKeypair(keyType)
	if err != nil {
		logger.Error.Println(err)
	}

	// Print
	logger.Info.Printf("private Key : %s", EncodeBase58_PrivateKey(myKeypair.privateKey))
	logger.Info.Printf("public Key : %s", EncodeBase58_PublicKey(myKeypair.publicKey))

	method := "comnic"
	did := core.MakeDID(method, EncodeBase58_PublicKey(myKeypair.publicKey))
	logger.Info.Printf("DID : [%s]", did)

	didDocument := MakeDIDDocument(method, did)
	logger.Info.Printf("DID Document : [%s]", didDocument)

	testDid := pb.Did{Id: "12342", Desc: "abcde"}

	logger.Info.Printf("Test Did: %s", testDid.GetId())

}

// GenerateKeypair : Generate a keypair (ECDSA)
func GenerateKeypair(keyType string) (Keypair, error) {
	pvKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	if err != nil {
		logger.Error.Println("Key generation error.", "Key Type is [", keyType, "]")
		return Keypair{}, errors.New("ECDSA Keypair generation was Fail!")
	}

	return Keypair{*pvKey, pvKey.PublicKey}, nil
}

// EncodeBase58_PrivateKey : Encode a PrivateKey as Base58.
func EncodeBase58_PrivateKey(privateKey ecdsa.PrivateKey) string {
	var privateKeyBase58 string

	privateKeyBytes, err := x509.MarshalECPrivateKey(&privateKey)
	if err != nil {
		log.Printf("error occured: %v", err.Error())
		return ""
	} else {
		privateKeyBase58 = base58.Encode(privateKeyBytes)
	}

	return privateKeyBase58
}

// EncodeBase58_PrivateKey : Encode a PublicKey as Base58.
func EncodeBase58_PublicKey(publicKey ecdsa.PublicKey) string {
	var publicKeyBase58 string

	publicKeyBytes, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		log.Printf("error occured: %v", err.Error())
		return ""
	} else {
		publicKeyBase58 = base58.Encode(publicKeyBytes)
	}

	return publicKeyBase58
}

func MakeDIDDocument(method string, did string) core.DIDDocument {

	doc := core.DIDDocument{Id: did}

	return doc
}
