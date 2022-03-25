package main

import (
	"context"
	protos "github.com/comnics/did-example/protos"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	protos.UnimplementedSimpleIssuerServer
}

func (issuer *server) IssueSimpleVC(_ context.Context, msg *protos.MsgIssueVC) (*protos.MsgIssueVCResponse, error) {

	response := new(protos.MsgIssueVCResponse)

	log.Printf("IssueSimpleVC MSG: %+v \n", msg)

	response.Vc = "Test VC"

	return response, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:1021")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	protos.RegisterSimpleIssuerServer(s, &server{})
	log.Printf("Issuer Server is listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
