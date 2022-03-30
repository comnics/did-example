// github.com/comnics/did-example/core/vc.go

package core

import (
	"crypto/ecdsa"
	"crypto/x509"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"github.com/google/uuid"
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

func (vc *VC) GenerateString(pvKey *ecdsa.PrivateKey) string {
	return ""
}

// VC를 JTW로 생성하고 string으로 반환한다.
// JTW의 경우 JWS로 증명되기에 Proofs를 빼고, JWT와 중복되는 properties를 제거한다.
func (vc *VC) GenerateJWT(pvKey *ecdsa.PrivateKey) string {
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
	token.Header["kid"] = vc.Proof.VerificationMethod

	tokenString, err := token.SignedString(pvKey)

	if err != nil {

	}

	return tokenString
}

func ParseJwt(tokenString string) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		did := token.Header["kid"].(string)
		pbKeyBase58 := did // getPbKey(did, "") //DID를 통해 DID-Document의 pbKey를 구한다.
		pbKey, _ := x509.ParsePKIXPublicKey(base58.Decode(pbKeyBase58))

		return pbKey, nil
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
