package core

type VC struct {
	// Mendatory
	Context []string `json:"@context"`

	Id                string            `json:"id"`
	Type              []string          `json:"type"`
	Issuer            string            `json:"issuer"`
	IssuanceDate      string            `json:"issuanceDate"`
	CredentialSubject CredentialSubject `json:"credentialSubject"`
	Proof             Proof             `json:"proof"`
}

type CredentialSubject struct {
	Id       string   `json:"id"`
	AlumniOf AlumniOf `json:"alumniOf"`
}

type AlumniOf struct {
	Id   string `json:"id"`
	Name []Name `json:"name"`
}

type Name struct {
	Value string `json:"value"`
	Lang  string `json:"lang"`
}

type Proof struct {
	Type               string `json:"type"`
	Created            string `json:"created"`
	ProofPurpose       string `json:"proofPurpose"`
	VerificationMethod string `json:"verificationMethod"`
	Jws                string `json:"jws"`
}
