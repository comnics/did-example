package main

import (
	"github.com/comnics/did-example/core"
	"github.com/comnics/did-example/protos"
	"testing"
)

// Test ECDSA Keypair generation.
func TestGenerateKeypair(t *testing.T) {
	var ecdsa core.ECDSAManager // ecdsa := new(core.ECDSAManager)
	err := ecdsa.Generate()

	if err != nil {
		t.Error("GenerateKeypair is Fail!")
	}
}

func TestGRPCDid(t *testing.T) {
	testDid := protos.Did{
		Id:   "test",
		Desc: "Dest Test",
	}

	if testDid.Id == "" {
		t.Error("Error.")
	}
}
