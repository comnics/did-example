package core

import (
	"crypto/ecdsa"
	"fmt"
	"testing"
)

func NewSimpleVC(verificationId string, privateKey *ecdsa.PrivateKey) string {
	credentialSubject := map[string]interface{}{
		"id": "1234567890",
		"alumniOf": map[string]interface{}{
			"id": "1234567",
			"name": []map[string]string{
				{
					"value": "Example University",
					"lang":  "en",
				}, {
					"value": "Exemple d'Université",
					"lang":  "fr",
				},
			},
		},
	}

	vc, err := NewVC("id", []string{"type1", "type2"}, "issuer", credentialSubject)
	if err != nil {

	}

	token := vc.GenerateJWT(verificationId, privateKey)

	return token

}
func TestNewVP(t *testing.T) {

	kms := NewEcdsa()

	vcToken := NewSimpleVC("", kms.PrivateKey)

	vcList := []string{}
	vcList = append(vcList, vcToken)

	vp, err := NewVP("", []string{"", ""}, "", vcList)
	if err != nil {
		t.Error("Failed create VP.")
	}

	vpToken := vp.GenerateJWT("", kms.PrivateKey)

	res, _ := vp.VerifyJwt(vpToken, kms.PublicKey)
	if res {
		fmt.Println("VP is valid.")
	} else {
		fmt.Println("VP is not valid.")
		t.Error("VP is not valid.")
	}

}
