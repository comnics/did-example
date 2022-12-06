// examples/proto/server/proto_server.go
package main

import (
	"context"
	"fmt"
	"github.com/comnics/did-example/protos"
	"google.golang.org/grpc"
	"log"
	"net"
)

type simpleDidServer struct {
	protos.UnimplementedSimpleDidServer
}

func (server *simpleDidServer) GetDid(ctx context.Context, req *protos.Did) (*protos.Did, error) {
	fmt.Println("Call GetDid")
	fmt.Printf("	Id: %s\n", req.Id)
	fmt.Printf("	Desc: %s\n", req.Desc)

	return &protos.Did{
		Id:   "12345",
		Desc: "Description.",
	}, nil
}

func main() {
	fmt.Println("### Start SimpleDID ###")
	lis, err := net.Listen("tcp", "0.0.0.0:1004")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := simpleDidServer{}
	s := grpc.NewServer()
	protos.RegisterSimpleDidServer(s, &server)

	log.Printf("Simple DID Server is listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
