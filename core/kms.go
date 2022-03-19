package core

import "crypto/ecdsa"

type Keypair struct {
	privateKey ecdsa.PrivateKey //interface{}
	publicKey  ecdsa.PublicKey  //interface{}
}

type KeyManagementSystem struct {
	Keys Keypair
}
