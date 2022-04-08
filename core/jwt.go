package core

import (
	"crypto/ecdsa"
	"crypto/x509"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/multiformats/go-multibase"
	"log"
	"strings"
)

func VerifyJwt(token string, pbKey *ecdsa.PublicKey) (bool, error) {
	parts := strings.Split(token, ".")
	err := jwt.SigningMethodES256.Verify(strings.Join(parts[0:2], "."), parts[2], pbKey)
	if err != nil {
		return false, nil
	}
	return true, nil
}

// Parse VC JWT Claim and Verify VC JWT.
//
// header에서 kid를 추출하고 DID를 추출한다.
// DID를 Resolve해서 DID Document를 받아온다.
// DID도큐먼트의 key ID를 기준으로 public key의 값을 가져와야 하나,
// 여기서는 1개만 존재한다고 가정하고 첫번째를 사용해서 public key를 만들어 사용한다.
func ParseAndVerifyJwtForVC(tokenString string) (bool, *JwtClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		//DID를 추출한다.
		parts := strings.Split(token.Header["kid"].(string), "#")
		did := parts[0]

		//Resolve한다.
		didDocumentStr, err := ResolveDid(did)
		if err != nil {
			log.Printf("Failed to Resolve DID.\nError: %x\n", err)
		}

		//Json string을 DID Document 객체로 생성한다.
		didDocument, err := NewDIDDocumentForString(didDocumentStr)
		if err != nil {
			log.Printf("Failed generate DID Document from string.\nError: %x\n", err)
		}
		// 첫 번째를 사용한다고 가정한다.
		// TODO: 키 ID(위의 kid)에 해당하는 키 값 구하기.
		pbKeyBaseMultibase := didDocument.VerificationMethod[0].PublicKeyMultibase
		_, bytePubKey, err := multibase.Decode(pbKeyBaseMultibase)
		pbKey, err := x509.ParsePKIXPublicKey(bytePubKey)

		return pbKey, nil
	})

	if claims, ok := token.Claims.(*JwtClaims); ok && token.Valid {
		//fmt.Printf("%v %v", claims.Vc, claims.Issuer)
		return true, claims, nil
	}

	return false, nil, err
}
