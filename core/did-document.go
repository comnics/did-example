package core

type didDocumentInterface interface {
	produce(doc DIDDocument) string
}

type DIDDocument struct {
	// @context
	// Mendatory
	Context []string `json:"@context"`

	Id                   string                       `json:"id"`
	AlsoKnownAs          []string                     `json:"alsoKnownAs"`
	Controller           string                       `json:"controller"`
	VerificationMethod   []VerificationMethodProperty `json:"verificationMethod"`
	Authentication       []AuthenticationProperty     `json:"authentication"`
	AssertionMethod      string                       `json:"assertionMethod"`
	KeyAgreement         string                       `json:"keyAgreement"`
	CapabilityInvocation string                       `json:"capabilityInvocation"`
	CapabilityDelegation string                       `json:"capabilityDelegation"`
	Service              []ServiceProperty            `json:"service"`
}

type VerificationMethodProperty struct {
	Id         string `json:"id"`
	Type       string `json:"type"`
	Controller string `json:"controller"`
}

type AuthenticationProperty struct {
	Id              string `json:"id"`
	Type            string `json:"type"`
	Controller      string `json:"controller"`
	PublicKeyBase58 string `json:"publicKeyBase58"`
}

type ServiceProperty struct {
	Id              string `json:"id"`
	Type            string `json:"type"`
	ServiceEndpoint string `json:"serviceEndpoint"`
}

func (doc *DIDDocument) produce() string {

	return ""
}
