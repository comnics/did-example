// github.com/comnics/did-example/core/vc.go

package core

import (
	"crypto/ecdsa"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

// https://www.w3.org/TR/vc-data-model
// https://www.w3.org/TR/vc-data-model/#json-web-token
type VC struct {
	// Mendatory
	Context []string `json:"@context"`

	Id                string                 `json:"id"`
	Type              []string               `json:"type"`
	Issuer            string                 `json:"issuer"`
	IssuanceDate      string                 `json:"issuanceDate"`
	CredentialSubject map[string]interface{} `json:"credentialSubject"`
	Proof             Proof                  `json:"proof"`
}

type Proof struct {
	Type               string `json:"type"`
	Created            string `json:"created"`
	ProofPurpose       string `json:"proofPurpose"`
	VerificationMethod string `json:"verificationMethod"`
	Jws                string `json:"jws"`
}

// JWT를 위한 claim
type JwtClaims struct {
	jwt.StandardClaims

	Nonce string
	Vc    VC `json:"vc"`
}

type VCInterface interface {
	Sign()
	Verify() bool
}

// VC를 JTW로 생성하고 string으로 반환한다.
func (vc *VC) CreateJWT(pvKey *ecdsa.PrivateKey) string {

	aud := ""
	exp := time.Now().Add(time.Minute * 5).Unix()
	jti := "089a411f-0d88-450f-8cc0-1a3acfebecd3"
	iat := time.Now().Unix()
	nbf := iat
	iss := "http://google.com/issuer"
	sub := ""

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
		*vc,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwtClaims)

	tokenString, err := token.SignedString(pvKey)

	if err != nil {

	}

	return tokenString
}

func ParseJwt(tokenString string) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})

	if claims, ok := token.Claims.(*JwtClaims); ok && token.Valid {
		fmt.Printf("%v %v", claims.Vc, claims.Issuer)
	} else {
		fmt.Println(err)
	}

}

func VerifyJwt(token string, pbKey *ecdsa.PublicKey) (bool, error) {
	parts := strings.Split(token, ".")
	err := jwt.SigningMethodES256.Verify(strings.Join(parts[0:2], "."), parts[2], pbKey)
	if err != nil {
		return false, nil
	}
	return true, nil
}
