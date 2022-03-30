package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"fmt"

	"github.com/comnics/did-example/core"
)

func main() {
	rsam, err := core.NewRsa()

	fmt.Printf("PbKey: %x\n", rsam.PublicKey)

	pbKeyByte := x509.MarshalPKCS1PublicKey(rsam.PublicKey)
	fmt.Printf("PbKey byte: %s\n", pbKeyByte)
	pbKeyBase64 := base64.StdEncoding.EncodeToString(pbKeyByte)
	fmt.Printf("PbKey: %s\n", pbKeyBase64)

	pbKeyByte2, err := base64.StdEncoding.DecodeString(pbKeyBase64)
	fmt.Printf("PbKey byte: %s\n", pbKeyByte2)

	pub, err := x509.ParsePKCS1PublicKey(pbKeyByte2)
	fmt.Printf("PbKey: %x\n", pub)

	message := "Hello World."
	hash := sha256.New()
	hash.Write([]byte(message))
	digest := hash.Sum(nil)

	var h1 crypto.Hash
	//개인키로 서명.
	signature, err := rsa.SignPKCS1v15(
		rand.Reader,
		rsam.PrivateKey,
		h1,
		digest,
	)

	fmt.Println(base64.StdEncoding.EncodeToString(signature))

	var h2 crypto.Hash
	// 공개키로 서명 검증.
	err = rsa.VerifyPKCS1v15(
		pub,
		h2,
		digest,
		signature,
	)

	if err != nil {
		fmt.Println("Fail verification.")
	} else {
		fmt.Println("Success verification")
	}
}
