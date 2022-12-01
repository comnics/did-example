package main

import (
	"github.com/comnics/did-example/actors/issuer"
	"github.com/comnics/did-example/config"
	protos "github.com/comnics/did-example/protos"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

var VcCustomFilePath string

// for Custom VC of type json
// go run actors/issuer/cmd/issuerd.go custom_vc.json
func main() {
	argsWithoutProg := os.Args[1:]

	// New Issuer
	issr := new(issuer.Issuer)
	issr.GenerateDID()

	if len(argsWithoutProg) > 0 {
		issr.CredentialSubjectJsonFilePath = argsWithoutProg[0]
		//loadJson(vcCustomFilePath)
	}

	lis, err := net.Listen("tcp", config.SystemConfig.IssuerAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	issuerServer := issuer.Server{}
	issuerServer.Issuer = issr

	s := grpc.NewServer()
	protos.RegisterSimpleIssuerServer(s, &issuerServer)

	log.Printf("Issuer Server is listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
