package main

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/comnics/did-example/core"
	"github.com/comnics/did-example/util"
	"time"
)

// Simple KMS
type Keypair struct {
	privateKey ecdsa.PrivateKey //interface{}
	publicKey  ecdsa.PublicKey  //interface{}
}

const (
	keyType = "ecdsa"
)

var (
	logger util.Logger
)

// Main
func main() {
	//logger := util.GetLogger()

	fmt.Println("### Start Main() ###")

	// ECDSA secp256k1
	//var ecdsa *core.ECDSAManager // ecdsa := new(core.ECDSAManager)
	ecdsa := core.NewEcdsa()

	did := core.NewDID("comnic", ecdsa.PublicKeyBase58())
	fmt.Printf("DID : [%s]\n", did)

	verificationId := fmt.Sprintf("%s#keys-1", did)
	verificationMethod := []core.VerificationMethod{
		{
			Id:                 verificationId,
			Type:               "EcdsaSecp256k1VerificationKey2019",
			Controller:         did.String(),
			PublicKeyMultibase: ecdsa.PublicKeyMultibase(),
		},
	}
	didDocument := core.NewDIDDocument(did.String(), verificationMethod)
	fmt.Printf("DID Document: %s\n", didDocument)

	// @@@@@@@@@@@@@@@@@@@

	myVC := &core.VC{
		Context: []string{
			"https://www.w3.org/2018/credentials/v1",
			"https://www.w3.org/2018/credentials/v2",
		},
		Id:           "http://example.edu/credentials/1872",
		Type:         []string{"VerifiableCredential", "AlumniCredential"},
		Issuer:       "https://example.edu/issuers/565049",
		IssuanceDate: time.Now().Format(time.RFC3339), //"2010-01-01T19:23:24Z",
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
		Proof: &core.Proof{
			Type:               "RsaSignature2018",
			Created:            "2017-06-18T21:19:10Z",
			ProofPurpose:       "assertionMethod",
			VerificationMethod: verificationId,
			Jws:                "eyJhbGciOiJSUzI1NiIsImI2NCI6ZmFsc2UsImNyaXQiOlsiYjY0Il19..TCYt5XsITJX1CxPCT8yAV-TVkIEq_PbChOMqsLfRoPsnsgw5WEuts01mq-pQy7UJiN5mgRxD-WUcX16dUEMGlv50aqzpqh4Qktb3rk-BuQy72IFLOqV0G_zS245-kronKb78cPN25DGlcTwLtjPAYuNzVBAh4vGHSrQyHUdBBPM",
		},
	}
	token := myVC.GenerateJWT(ecdsa.PrivateKey)
	fmt.Printf("\nVC JWT Token: %s\n", token)

	res, _ := core.VerifyJwt(token, ecdsa.PublicKey)

	if res {
		fmt.Println("VC is verified.")
	} else {
		fmt.Println("VC is Not verified.")
	}

	//// protobuf test
	//testDid := pb.Did{Id: "12342", Desc: "abcde"}
	//
	//fmt.Printf("Test Did: %s\n", testDid.GetId())

}

func MakeDIDDocument(method string, did string) core.DIDDocument {

	doc := core.DIDDocument{Id: did}

	return doc
}
