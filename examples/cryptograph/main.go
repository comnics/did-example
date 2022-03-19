package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"errors"
	"fmt"
	"log"

	"github.com/btcsuite/btcutil/base58"
)

// Simple KMS
type Keypair struct {
	privateKey ecdsa.PrivateKey //interface{}
	publicKey  ecdsa.PublicKey  //interface{}
}

const (
	keyType = "ecdsa"
)

// Main
func main() {
	fmt.Println("### Start Main() ###")

	// Generate ECDSA keypair
	myKeypair, err := GenerateKeypair(keyType)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	// Print
	fmt.Printf("Private Key: %x\n", myKeypair.privateKey)
	fmt.Printf("Private Key(Base58): %x\n", EncodeBase58_PrivateKey(myKeypair.privateKey))
	fmt.Printf("Public Key: %x\n", myKeypair.publicKey)
	fmt.Printf("Public Key(Base58): %x\n", EncodeBase58_PublicKey(myKeypair.publicKey))

}

// GenerateKeypair : Generate a keypair (ECDSA)
func GenerateKeypair(keyType string) (Keypair, error) {
	pvKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	if err != nil {
		fmt.Println("Key generation error.", "Key Type is [", keyType, "]")
		return Keypair{}, errors.New("ECDSA Keypair generation was Fail!")
	}

	return Keypair{*pvKey, pvKey.PublicKey}, nil
}

// EncodeBase58_PrivateKey : Encode a PrivateKey as Base58.
func EncodeBase58_PrivateKey(privateKey ecdsa.PrivateKey) string {
	var privateKeyBase58 string

	privateKeyBytes, err := x509.MarshalECPrivateKey(&privateKey)
	if err != nil {
		fmt.Printf("error occured: %v", err.Error())
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
