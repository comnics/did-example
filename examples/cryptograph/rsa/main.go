package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"fmt"
)

func main() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048) // 개인 키와 공개 키 생성
	if err != nil {
		fmt.Println(err)
		return
	}
	publicKey := &privateKey.PublicKey // 개인 키 변수 안에 공개 키가 들어있음
	fmt.Printf("PbKey: %x\n", publicKey)

	pbkey_byte := x509.MarshalPKCS1PublicKey(publicKey)
	fmt.Printf("PbKey byte: %s\n", pbkey_byte)
	pbkey_base64 := base64.StdEncoding.EncodeToString(pbkey_byte)
	fmt.Printf("PbKey: %s\n", pbkey_base64)

	pbkey_byte2, err := base64.StdEncoding.DecodeString(pbkey_base64)
	fmt.Printf("PbKey byte: %s\n", pbkey_byte2)

	pub, err := x509.ParsePKCS1PublicKey(pbkey_byte2)
	fmt.Printf("PbKey: %x\n", pub)

	message := "안녕하세요. Go 언어"
	hash := sha256.New()        // 해시 인스턴스 생성
	hash.Write([]byte(message)) // 해시 인스턴스에 문자열 추가
	digest := hash.Sum(nil)     // 문자열의 MD5 해시 값 추출

	var h1 crypto.Hash
	signature, err := rsa.SignPKCS1v15( // 개인 키로 서명
		rand.Reader,
		privateKey, // 개인 키
		h1,
		digest, //
	)

	fmt.Println(base64.StdEncoding.EncodeToString(signature))

	var h2 crypto.Hash
	err = rsa.VerifyPKCS1v15( // 공개 키로 서명 검증
		pub, // 공개 키
		h2,
		digest,    //
		signature, // 서명 값
	)

	if err != nil {
		fmt.Println("검증 실패")
	} else {
		fmt.Println("검증 성공")
	}
}
