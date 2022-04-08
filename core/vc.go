// github.com/comnics/did-example/core/vc.go

package core

import (
	"crypto/ecdsa"
	"crypto/x509"
	"fmt"
	"github.com/google/uuid"
	"github.com/multiformats/go-multibase"
	"log"
	"strings"
	"time"

	"github.com/getlantern/deepcopy"
	"github.com/golang-jwt/jwt"
)

// https://www.w3.org/TR/vc-data-model
// https://www.w3.org/TR/vc-data-model/#json-web-token
type VC struct {
	// Mendatory
	Context []string `json:"@context"`

	Id                string                 `json:"id,omitempty"`
	Type              []string               `json:"type,omitempty"`
	Issuer            string                 `json:"issuer,omitempty"`
	IssuanceDate      string                 `json:"issuanceDate,omitempty"`
	CredentialSubject map[string]interface{} `json:"credentialSubject,omitempty"`
	Proof             *Proof                 `json:"proof,omitempty"`
}

type Proof struct {
	Type               string `json:"type,omitempty"`
	Created            string `json:"created,omitempty"`
	ProofPurpose       string `json:"proofPurpose,omitempty"`
	VerificationMethod string `json:"verificationMethod,omitempty"`
	ProofValue         string `json:"proofValue,omitempty"`
	Jws                string `json:"jws,omitempty"`
}

// JWT를 위한 claim
type JwtClaims struct {
	jwt.StandardClaims

	Nonce string
	Vc    VC `json:"vc,omitempty"`
}

type VCInterface interface {
	Sign()
	Verify() bool
}

func NewVC(id string, typ []string, issuer string, credentialSubject map[string]interface{}) (*VC, error) {
	newVC := &VC{
		Context: []string{
			"https://www.w3.org/2018/credentials/v1",
			"https://www.w3.org/2018/credentials/v2",
		},
		Id:                id,
		Type:              typ,
		Issuer:            issuer,
		IssuanceDate:      time.Now().Format(time.RFC3339), //"2010-01-01T19:23:24Z",
		CredentialSubject: credentialSubject,
	}
	return newVC, nil
}

func (vc *VC) GenerateString(pvKey *ecdsa.PrivateKey) string {
	return ""
}

// VC를 JTW로 생성하고 string으로 반환한다.
// JTW의 경우 JWS로 증명되기에 Proofs를 빼고, JWT와 중복되는 properties를 제거한다.
func (vc *VC) GenerateJWT(verificationId string, pvKey *ecdsa.PrivateKey) string {
	aud := ""
	exp := time.Now().Add(time.Minute * 5).Unix()       //만료 시간. 현재 + 5분
	jti := uuid.NewString()                             // JWT ID
	t, err := time.Parse(time.RFC3339, vc.IssuanceDate) //unixtime으로 바꾸기 위해.
	iat := t.Unix()
	nbf := iat
	iss := vc.Issuer
	sub := "Verifiable Credential"

	// Proof를 제거하고 JWT를 만들기 위해 복제한다.
	vcTmp := new(VC)
	deepcopy.Copy(vcTmp, vc)
	vcTmp.Proof = nil

	jwtClaims := JwtClaims{
		jwt.StandardClaims{
			Audience:  aud,
			ExpiresAt: exp,
			Id:        jti,
			IssuedAt:  iat,
			Issuer:    iss,
			NotBefore: nbf,
			Subject:   sub,
		},
		"qwasd!234",
		*vcTmp,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwtClaims)
	token.Header["kid"] = verificationId

	tokenString, err := token.SignedString(pvKey)

	if err != nil {

	}

	return tokenString
}

// VC의 claim에 Proof를 생성한다.
func (vc *VC) GenerateProof() string {

	vc.Proof = &Proof{}
	return ""
}

// jwt를 검증한다.
// header에서 kid를 추출하고 DID를 추출한다.
// DID를 Resolve해서 DID Document를 받아온다.
// DID도큐먼트의 key ID를 기준으로 public key의 값을 가져와야 하나,
// 여기서는 1개만 존재한다고 가정하고 첫번째를 사용해서 public key를 만들어 사용한다.
func Deprecated_ParseJwt(tokenString string) (bool, *JwtClaims, error) {
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
	} else {
		fmt.Println(err)
		return false, nil, err
	}

}

func (vc *VC) VerifyJwt(token string, pbKey *ecdsa.PublicKey) (bool, error) {
	return VerifyJwt(token, pbKey)
}
