package main

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"github.com/consensys/gnark-crypto/ecc/twistededwards"
	"github.com/consensys/gnark-crypto/hash"
	"github.com/consensys/gnark-crypto/signature/eddsa"
)

func main() {
	// instantiate hash function
	hFunc := hash.MIMC_BN254.New()

	// create a eddsa key pair
	privateKey, _ := eddsa.New(twistededwards.BN254, rand.Reader)
	publicKey := privateKey.Public()

	// note that the message is on 4 bytes
	msg := "Hello, World"
	digest := sha256.Sum256([]byte(msg))

	// sign the message
	signature, _ := privateKey.Sign(digest[:], hFunc)

	fmt.Println("PvKey: ", privateKey)
	fmt.Println("PbKey: ", publicKey)

	fmt.Println("signature: ", signature)

	// verifies signature
	isValid, _ := publicKey.Verify(signature, digest[:], hFunc)
	if !isValid {
		fmt.Println("1. invalid signature")
	} else {
		fmt.Println("1. valid signature")
	}

}
