package main

import (
	"context"
	"fmt"
	"github.com/comnics/did-example/config"
	"github.com/comnics/did-example/protos"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
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

	fmt.Println("### Start Resover RESTful ###")

	if err := protos.RegisterRegistrarHandlerFromEndpoint(
		ctx,
		mux,
		config.SystemConfig.ResolverAddr,
		options,
	); err != nil {
		log.Fatalf("failed to register gRPC gateway: %v", err)
	}

	log.Printf("start HTTP server on %s", config.SystemConfig.ResolverGatewayAddr)
	if err := http.ListenAndServe(config.SystemConfig.ResolverGatewayAddr, mux); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
