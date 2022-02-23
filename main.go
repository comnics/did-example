package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/btcsuite/btcutil/base58"
)

// Simple KMS
type Keypair struct {
	privateKey ecdsa.PrivateKey //interface{}
	publicKey  ecdsa.PublicKey  //interface{}
}

const (
	keyType = "ecdsa"
)

// Logger
type Logger struct {
	Trace *log.Logger
	Warn  *log.Logger
	Info  *log.Logger
	Error *log.Logger
}

var logger Logger

// Main
func main() {
	InitLog(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)

	logger.Info.Println("### Start Main() ###")

	// Generate ECDSA keypair
	myKeypair, err := GenerateKeypair(keyType)
	if err != nil {
		logger.Error.Println(err)
	}

	// Print
	logger.Info.Printf("private Key : %s", EncodeBase58_PrivateKey(myKeypair.privateKey))
	logger.Info.Printf("public Key : %s", EncodeBase58_PublicKey(myKeypair.publicKey))

}

// InitLog: Initialize the log
func InitLog(traceHandle io.Writer, infoHandle io.Writer, warningHandle io.Writer, errorHandle io.Writer) {
	logger.Trace = log.New(traceHandle, "[TRACE] ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Info = log.New(infoHandle, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Warn = log.New(warningHandle, "[WARNING] ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Error = log.New(errorHandle, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)
}

// GenerateKeypair : Generate a keypair (ECDSA)
func GenerateKeypair(keyType string) (Keypair, error) {
	pvKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	if err != nil {
		logger.Error.Println("Key generation error.")
		return Keypair{}, errors.New("ECDSA Keypair generation was Fail!")
	}

	return Keypair{*pvKey, pvKey.PublicKey}, nil
}

// EncodeBase58_PrivateKey : Encode a PrivateKey as Base58.
func EncodeBase58_PrivateKey(privateKey ecdsa.PrivateKey) string {
	var privateKeyBase58 string

	privateKeyBytes, err := x509.MarshalECPrivateKey(&privateKey)
	if err != nil {
		log.Printf("error occured: %v", err.Error())
	} else {
		privateKeyBase58 = base58.Encode(privateKeyBytes)
	}

	return privateKeyBase58
}

// EncodeBase58_PrivateKey : Encode a PublicKey as Base58.
func EncodeBase58_PublicKey(publicKey ecdsa.PublicKey) string {
	var publicKeyBase58 string

	publicKeyBytes, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		log.Printf("error occured: %v", err.Error())
	} else {
		publicKeyBase58 = base58.Encode(publicKeyBytes)
	}

	return publicKeyBase58
}
