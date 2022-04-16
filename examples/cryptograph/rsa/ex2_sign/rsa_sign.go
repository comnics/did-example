package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"log"
)

func sign(msg string, pvKey *rsa.PrivateKey) ([]byte, error) {
	hash := sha256.New()
	hash.Write([]byte(msg))
	digest := hash.Sum(nil)

	var h crypto.Hash
	signature, err := rsa.SignPKCS1v15(
		rand.Reader,
		pvKey,
		h,
		digest,
	)

	return signature, err
}

func main() {
	pvKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Printf("Failed to Generate RSA.")
	}

	msg := "Hello World."
	sig, err := sign(msg, pvKey)
	if err != nil {
		log.Printf("Failed to Sign.")
	}

	fmt.Printf("Signature: %v\n", sig)

	sigBase58 := base58.Encode(sig)
	fmt.Printf("Signature(base58): %v\n", sigBase58)
}
