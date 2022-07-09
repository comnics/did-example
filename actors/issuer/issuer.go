package issuer

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/comnics/did-example/core"
	"github.com/comnics/did-example/protos"
	"io/ioutil"
	"log"
	"os"
)

type Server struct {
	protos.UnimplementedSimpleIssuerServer

	Issuer *Issuer
}

type Issuer struct {
	kms         *core.ECDSAManager
	did         *core.DID
	didDocument *core.DIDDocument

	CredentialSubjectJsonFilePath string
}

const (
	VC_MODE_TEST   = 1
	VC_MODE_CUSTOM = 2
)

var (
	vcMode = VC_MODE_TEST
)

type VC_CUSTOM_CLAIM struct {
	data map[string]interface{}
}

func (server *Server) IssueSimpleVC(_ context.Context, msg *protos.MsgIssueVC) (*protos.MsgIssueVCResponse, error) {
	log.Printf("IssueSimpleVC MSG: %+v \n", msg)

	response := new(protos.MsgIssueVCResponse)

	server.Issuer.CredentialSubjectJsonFilePath = "custom_vc.json"

	vcToken, err := server.Issuer.GenerateSampleVC()
	if err != nil {

	}
	response.Vc = vcToken

	return response, nil
}

func (issuer *Issuer) GenerateDID() {
	// 키생성(ECDSA) - 향후 KMS로 대체.
	issuer.kms = core.NewEcdsa()

	// DID 생성.
	issuerDid, _ := core.NewDID("comnic", issuer.kms.PublicKeyBase58())

	issuer.did = issuerDid

	// DID Document 생성.
	verificationId := fmt.Sprintf("%s#keys-1", issuerDid)
	verificationMethod := []core.VerificationMethod{
		{
			Id:                 verificationId,
			Type:               core.VERIFICATION_KEY_TYPE_SECP256K1,
			Controller:         issuerDid.String(),
			PublicKeyMultibase: issuer.kms.PublicKeyMultibase(),
		},
	}
	didDocument := core.NewDIDDocument(issuerDid.String(), verificationMethod)
	issuer.didDocument = didDocument

	fmt.Printf("검증용 issuer pbKey: %s\n", issuer.didDocument.VerificationMethod[0].PublicKeyMultibase)

	RegisterDid(issuerDid.String(), didDocument)
}

func (issuer *Issuer) GenerateSampleVC() (string, error) {

	var credentialSubject map[string]interface{}

	if issuer.CredentialSubjectJsonFilePath == "" {

	} else {
		credentialSubject = LoadJson(issuer.CredentialSubjectJsonFilePath) // "custom_vc.json"
	}

	// VC 생성.
	vc, err := core.NewVC(
		"1234567890",
		[]string{"VerifiableCredential", "CertificationOfEmployee"},
		issuer.did.String(),
		credentialSubject,
	)

	if err != nil {
		return "", errors.New("Failed creation VC.")
	}

	// VC에 Issuer의 private key로 서명한다.(JWT 사용)
	token := vc.GenerateJWT(issuer.didDocument.VerificationMethod[0].Id, issuer.kms.PrivateKey)

	return token, nil
}

func RegisterDid(did string, document *core.DIDDocument) error {
	err := core.RegisterDid(did, document.String())
	if err != nil {
		return err
	}
	return nil

}

func LoadJson(path string) map[string]interface{} {
	jsonData := make(map[string]interface{})

	data, err := os.Open(path)
	if err != nil {
		return nil
	}

	vcMode = VC_MODE_CUSTOM
	byteValue, _ := ioutil.ReadAll(data)

	json.Unmarshal(byteValue, &jsonData)

	return jsonData
}
