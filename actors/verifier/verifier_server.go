package main

import (
	"context"
	"github.com/comnics/did-example/core"
	"github.com/comnics/did-example/protos"
	"log"
)

type VerifierServer struct {
	protos.UnimplementedVerifierServer

	verifier *Verifier
}

func (server *VerifierServer) SubmitVP(ctx context.Context, req *protos.SubmitVPRequest) (*protos.SubmitVPResponse, error) {
	log.Printf("VP: %s\n", req.Vp)

	verify, _, err := core.ParseAndVerifyJwtForVP(req.Vp)

	res := &protos.SubmitVPResponse{Result: "fail"}
	if verify && err == nil {
		res.Result = "ok"
	}

	return res, nil
}
