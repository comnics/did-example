package util

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/btcsuite/btcutil/base58"
)

func MakeHash(plain string) []byte {
	digest := sha256.Sum256([]byte(plain))
	return digest[:]
}

func MakeHashBase58(plain string) string {
	return base58.Encode(MakeHash(plain))
}

func MakeHashHex(plain string) string {
	return hex.EncodeToString(MakeHash(plain))
}

func MakeSha256(plainText string) string {
	hash := sha256.New()
	hash.Write([]byte(plainText))

	md := hash.Sum(nil)
	return base58.Encode(md)
}
