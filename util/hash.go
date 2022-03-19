package util

import (
	"crypto/sha512"
	"encoding/hex"
)

func StringHash(input string) (str string) {
	hash := sha512.New()
	str = hex.EncodeToString(hash.Sum([]byte(input)))
	return
}
