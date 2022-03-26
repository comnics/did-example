package core

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestVCStruct2(t *testing.T) {
	//credentialSubject := make(map[string]interface{})

	myVC := &VC{
		Context: []string{
			"https://www.w3.org/2018/credentials/v1",
			"https://www.w3.org/2018/credentials/v2",
		},
		Id:           "http://example.edu/credentials/1872",
		Type:         []string{"VerifiableCredential", "AlumniCredential"},
		Issuer:       "https://example.edu/issuers/565049",
		IssuanceDate: "2010-01-01T19:23:24Z",
		CredentialSubject: map[string]interface{}{
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
		},
		Proof: Proof{
			Type:               "RsaSignature2018",
			Created:            "2017-06-18T21:19:10Z",
			ProofPurpose:       "assertionMethod",
			VerificationMethod: "https://example.edu/issuers/565049#key-1",
			Jws:                "eyJhbGciOiJSUzI1NiIsImI2NCI6ZmFsc2UsImNyaXQiOlsiYjY0Il19..TCYt5XsITJX1CxPCT8yAV-TVkIEq_PbChOMqsLfRoPsnsgw5WEuts01mq-pQy7UJiN5mgRxD-WUcX16dUEMGlv50aqzpqh4Qktb3rk-BuQy72IFLOqV0G_zS245-kronKb78cPN25DGlcTwLtjPAYuNzVBAh4vGHSrQyHUdBBPM",
		},
	}

	jsonVC, _ := json.Marshal(myVC)
	fmt.Printf("My VC: %s\n", jsonVC)
}
