// github.com/comnics/did-example/core/vc.go

package core

// https://www.w3.org/TR/vc-data-model
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
