package core

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEcdsaGenerate(t *testing.T) {
	var ecdsa ECDSAManager // ecdsa := new(core.ECDSAManager)
	err := ecdsa.Generate()

	require.NoError(t, err)
}
