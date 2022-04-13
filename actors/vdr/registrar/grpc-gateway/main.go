package main

import (
	"context"
	"fmt"
	"github.com/comnics/did-example/config"
	"github.com/comnics/did-example/protos"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

func main() {
	ctx := context.Background()
	mux := runtime.NewServeMux()

	options := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	fmt.Println("### Start Registrar ###")

	if err := protos.RegisterRegistrarHandlerFromEndpoint(
		ctx,
		mux,
		config.SystemConfig.RegistrarAddr,
		options,
	); err != nil {
		log.Fatalf("failed to register gRPC gateway: %v", err)
	}

	log.Printf("start HTTP server on %s", config.SystemConfig.RegistrarGatewayAddr)
	if err := http.ListenAndServe(config.SystemConfig.RegistrarGatewayAddr, mux); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
