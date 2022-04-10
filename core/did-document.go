// github.com/comnics/did-example/core/did-document.go

package core

import "encoding/json"

type didDocumentInterface interface {
	produce(doc DIDDocument) string
	consume(str string) DIDDocument
}

// DID Document
// Decentralized Identifiers (DIDs) v1.0 https://www.w3.org/TR/did-core/
// Simple Example: https://www.w3.org/TR/did-core/#a-simple-example
type DIDDocument struct {
	// @context
	// Mendatory
	Context []string `json:"@context"`

	Id                   string               `json:"id"`
	AlsoKnownAs          []string             `json:"alsoKnownAs,omitempty"`
	Controller           string               `json:"controller,omitempty"`
	VerificationMethod   []VerificationMethod `json:"verificationMethod,omitempty"`
	Authentication       []Authentication     `json:"authentication,omitempty"`
	AssertionMethod      string               `json:"assertionMethod,omitempty"`
	KeyAgreement         string               `json:"keyAgreement,omitempty"`
	CapabilityInvocation string               `json:"capabilityInvocation,omitempty"`
	CapabilityDelegation string               `json:"capabilityDelegation,omitempty"`
	Service              []Service            `json:"service,omitempty"`
}

type VerificationMethod struct {
	Id                 string `json:"id"`
	Type               string `json:"type"`
	Controller         string `json:"controller"`
	PublicKeyMultibase string `json:"PublicKeyMultibase,omitempty"`
}

type Authentication struct {
	Id                 string `json:"id"`
	Type               string `json:"type"`
	Controller         string `json:"controller"`
	PublicKeyBase58    string `json:"publicKeyBase58,omitempty"`
	PublicKeyMultibase string `json:"PublicKeyMultibase,omitempty"`
}

type Service struct {
	Id              string `json:"id"`
	Type            string `json:"type"`
	ServiceEndpoint string `json:"serviceEndpoint"`
}

func NewDIDDocument(did string, verificationMethod []VerificationMethod) (doc *DIDDocument) {
	var docTmp = new(DIDDocument)
	docTmp.Context = []string{"https://www.w3.org/ns/did/v1"}
	docTmp.Id = did
	docTmp.VerificationMethod = verificationMethod

	return docTmp
}

func NewDIDDocumentForString(docStr string) (didDoc *DIDDocument, err error) {
	didDoc = new(DIDDocument)
	e := json.Unmarshal([]byte(docStr), didDoc)
	if e != nil {
		return nil, e
	}

	return didDoc, nil
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
