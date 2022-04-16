package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
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

func verify(msg string, signature []byte, pbKey *rsa.PublicKey) bool {
	hash := sha256.New()
	hash.Write([]byte(msg))
	digest := hash.Sum(nil)

	var h crypto.Hash
	err := rsa.VerifyPKCS1v15(
		pbKey,
		h,
		digest,
		signature,
	)
	if err != nil {
		return false
	}

	return true
}

func main() {
	pvKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Printf("Failed to Generate RSA.")
	}

	pbKey := &pvKey.PublicKey

	msg := "Hello World."
	sig, err := sign(msg, pvKey)
	if err != nil {
		log.Printf("Failed to Sign.")
	}

	bVerify := verify(msg, sig, pbKey)
	if bVerify {
		fmt.Println("Verify.")
	} else {
		fmt.Println("Not Verify.")
	}

	// 다른 key로 검증 시도.
	pvKey2, err := rsa.GenerateKey(rand.Reader, 2048)
	pbKey2 := &pvKey2.PublicKey
	bVerify = verify(msg, sig, pbKey2)
	if bVerify {
		fmt.Println("Verify.")
	} else {
		fmt.Println("Not Verify.")
	}
}
