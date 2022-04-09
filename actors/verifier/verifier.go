package main

import (
	"github.com/comnics/did-example/core"
	"github.com/comnics/did-example/protos"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Verifier struct {
	kms         *core.ECDSAManager
	did         *core.DID
	didDocument *core.DIDDocument
}

func main() {
	// New Issuer
	verifier := new(Verifier)
	//issuer.generateDID()

	lis, err := net.Listen("tcp", "0.0.0.0:1022")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	verifierServer := VerifierServer{}
	verifierServer.verifier = verifier

	s := grpc.NewServer()
	protos.RegisterVerifierServer(s, &verifierServer)

	log.Printf("Issuer Server is listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
