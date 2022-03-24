package core

type VC struct {
	// Mendatory
	Context []string `json:"@context"`

	Id           string   `json:"id"`
	Type         []string `json:"type"`
	Issuer       string   `json:"issuer"`
	IssuanceDate string   `json:"issuanceDate"`

	CredentialSubject struct {
		Id       string `json:"id"`
		AlumniOf struct {
			Id   string `json:"id"`
			Name []struct {
				value string `json:"value"`
				lang  string `json:"lang"`
			} `json:"name"`
		}
	} `json:"credentialSubject"`

	Proof struct {
		Type               string `json:"type"`
		Created            string `json:"created"`
		ProofPurpose       string `json:"proofPurpose"`
		VerificationMethod string `json:"verificationMethod"`
		Jws                string `json:"jws"`
	} `json:"proof"`
}
