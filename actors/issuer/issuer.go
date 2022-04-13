package issuer

import (
	"context"
	"errors"
	"fmt"
	"github.com/comnics/did-example/core"
	"github.com/comnics/did-example/protos"
	"log"
)

type Server struct {
	protos.UnimplementedSimpleIssuerServer

	Issuer *Issuer
}

type Issuer struct {
	kms         *core.ECDSAManager
	did         *core.DID
	didDocument *core.DIDDocument
}

func (server *Server) IssueSimpleVC(_ context.Context, msg *protos.MsgIssueVC) (*protos.MsgIssueVCResponse, error) {
	log.Printf("IssueSimpleVC MSG: %+v \n", msg)

	response := new(protos.MsgIssueVCResponse)

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
			Type:               "EcdsaSecp256k1VerificationKey2019",
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
	// VC 생성.
	vc, err := core.NewVC(
		"1234567890",
		[]string{"VerifiableCredential", "AlumniCredential"},
		issuer.did.String(),
		map[string]interface{}{
			"id": "1234567890",
			"name": map[string]interface{}{
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
