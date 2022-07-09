package main

import (
	"github.com/comnics/did-example/actors/issuer"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLoadJson(t *testing.T) {
	data := issuer.LoadJson("custom_vc.json")
	expectData := map[string]interface{}(map[string]interface{}{"employment": map[string]interface{}{"birth": "2000/01/01", "join": "2020/03/01", "name": "홍길동", "salary": "100000000"}, "id": "1234567890"})

	require.Equal(t, expectData, data)
}
