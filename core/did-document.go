package core

import "encoding/json"

type didDocumentInterface interface {
	produce(doc DIDDocument) string
	consume(str string) DIDDocument
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

func NewDIDDocument(did string) (doc *DIDDocument) {
	var docTmp = new(DIDDocument)
	docTmp.Context = []string{"https://www.w3.org/ns/did/v1"}
	docTmp.Id = did

	return docTmp
}

func (doc *DIDDocument) Produce() string {
	str, err := json.Marshal(doc)
	if err != nil {
		// fmt.Println("Fail to marshal: ", err)
		return ""
	}
	return string(str)
}

func (doc *DIDDocument) Consume(str string) {
	err := json.Unmarshal([]byte(str), doc)
	if err != nil {

	}
}

func (doc *DIDDocument) String() string {
	return doc.Produce()
}
