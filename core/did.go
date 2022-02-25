package core

import (
	"crypto/sha256"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
)

func MakeDID(method string, pbKey string) (newDID string) {
	specificIdentifier := MakeSha256(pbKey)
	return fmt.Sprintf("did:%s:%s", method, specificIdentifier)
}

func MakeSha256(plainText string) string {
	hash := sha256.New()
	hash.Write([]byte(plainText))

	md := hash.Sum(nil)
	return base58.Encode(md)
}
