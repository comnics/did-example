package main

import (
	"fmt"
	"github.com/comnics/did-example/actors/holder"
	"github.com/comnics/did-example/core"
	"github.com/comnics/did-example/util"
	"os"
)

// Holder's Scenario
// 1. Generate DID from ECDSA(secp256k1).
// 2. Register Holder's DID(with DID Document).
// 3. Request VC to Issuer.(after connecting Issuer.)
// 4. Verify VC(issued by Issuer)
// 5. Send VP to Verifier.
// 6. Receive Verifier's response.
func main() {
	fmt.Println("### Start HOLDER's Wallet ###")
	// New Holder
	hldr := new(holder.Holder)

	util.PressKey("1. DID를 생성합니다. [아무키나 입력하세요.]")
	hldr.GenerateDID()
	fmt.Printf("DID: %s\n", hldr.Did.String())
	fmt.Printf("DID Document: %+v\n", hldr.DidDocument)

	util.PressKey("2. DID를 VDR에 등록합니다. [아무키나 입력하세요.]")
	holder.RegisterDid(hldr.Did.String(), hldr.DidDocument)

	util.PressKey("3. Issuer에게 VC를 요청합니다. [아무키나 입력하세요.]")
	vcToken := holder.RequestVC(hldr.Did.String())
	fmt.Printf("VC Token: %s\n", vcToken)

	util.PressKey("4. 전달 받은 VC를 검증합니다. [아무키나 입력하세요.]")
	verify, _, _ := core.ParseAndVerifyJwtForVC(vcToken)

	if verify {
		fmt.Println("VC is verified.")
	}

	util.PressKey("5. 전달 받은 VC를 VP에 넣어서 Verifier에게 전달합니다. [아무키나 입력하세요.]")
	vcList := []string{}
	vcList = append(vcList, vcToken)

	vp, err := core.NewVP("", []string{"", ""}, "", vcList)
	if err != nil {
		fmt.Println("ERROR")
		os.Exit(0)
	}

	vpToken := vp.GenerateJWT(hldr.DidDocument.VerificationMethod[0].Id, hldr.Kms.PrivateKey)
	fmt.Printf("VP Token: %s\n", vpToken)

	holder.SubmitVP(vpToken)

}
