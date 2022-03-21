// examples/cryptograph/ecdsa/main.go

package main

import (
	"crypto/sha256"
	"fmt"
	"github.com/comnics/did-example/core"
)

func main() {
	fmt.Println("### Start Main() ###")

	// ECDSA secp256k1
	var ecdsa core.ECDSAManager // ecdsa := new(core.ECDSAManager)
	ecdsa.Generate()

	fmt.Printf("########## Key pair ##########\n")
	fmt.Printf("===== Private Key =====\n")
	fmt.Printf("Private Key: %x\n", ecdsa.PrivateKey.D)
	fmt.Printf("===== Public Key(X, Y) =====\n")
	fmt.Printf("X=%s Y=%s\n", ecdsa.PublicKey.X, ecdsa.PublicKey.Y)
	fmt.Printf("  Hex: X=%x Y=%x\n\n", ecdsa.PublicKey.X.Bytes(), ecdsa.PublicKey.Y.Bytes())

	msg := "Hello World."
	digest := sha256.Sum256([]byte(msg))
	signature, err := ecdsa.Sign(digest[:])
	if err != nil {
		fmt.Printf("Fail to sign to msg.")
	}

	fmt.Printf("########## Sign ##########\n")
	fmt.Printf("===== Message =====\n")
	fmt.Printf("Msg: %s\n", msg)
	fmt.Printf("Digest: %x\n", digest)
	fmt.Printf("R: %s, S: %s\n\n", signature.R, signature.S)

	ret := ecdsa.Verify(signature, digest[:])

	fmt.Printf("########## Verification ##########\n")
	if ret {
		fmt.Printf("Signature verifies")
	} else {
		fmt.Printf("Signature does not verify")
	}

}
