// core/rsa.go

package core

import "crypto/rsa"

type RSAManager struct {
	PrivateKey rsa.PrivateKey
	PublicKey  rsa.PublicKey
}

type RSAInterface interface {
	GenerateKey()
	Encrypt(plainText string) (cipherText string)
	Decrypt(cipherText string) (plainText string)
	Sign() (signature string)
	Verify() (result bool)
}
