// core/rsa.go

package core

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"log"
)

type RSAManager struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
	PvKeyStr   string
	PbKeyStr   string
}

type RSAInterface interface {
	Encrypt(plainText string) (cipherText string)
	Decrypt(cipherText string) (plainText string)
	Sign(msg string) (signature []byte, err error)
	Verify(msg string, signature []byte) (result bool)
}

func NewRsa() (rsam *RSAManager, err error) {
	rsam = new(RSAManager)
	pvKey, err := rsa.GenerateKey(rand.Reader, 2048) // 개인 키와 공개 키 생성
	if err != nil {
		log.Printf("Fail to RSA Generate.")
		return nil, err
	}

	rsam.PrivateKey = pvKey
	rsam.PublicKey = &pvKey.PublicKey

	pbKeyByte := x509.MarshalPKCS1PublicKey(rsam.PublicKey)
	rsam.PbKeyStr = base64.StdEncoding.EncodeToString(pbKeyByte)

	pvKeyByte := x509.MarshalPKCS1PrivateKey(rsam.PrivateKey)
	rsam.PvKeyStr = base64.StdEncoding.EncodeToString(pvKeyByte)

	return rsam, nil
}

func (rsam *RSAManager) Sign(msg string) ([]byte, error) {
	hash := sha256.New()
	hash.Write([]byte(msg))
	digest := hash.Sum(nil)

	var h crypto.Hash
	signature, err := rsa.SignPKCS1v15(
		rand.Reader,
		rsam.PrivateKey,
		h,
		digest,
	)

	return signature, err
}

func (rsam *RSAManager) Verify(msg string, signature []byte) bool {
	hash := sha256.New()
	hash.Write([]byte(msg))
	digest := hash.Sum(nil)

	var h crypto.Hash
	err := rsa.VerifyPKCS1v15(
		rsam.PublicKey,
		h,
		digest,
		signature,
	)
	if err != nil {
		return false
	}

	return true
}
