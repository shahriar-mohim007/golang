package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// SignatureGenerate generates a base64-encoded RSA signature for the input data
func SignatureGenerate(requestData string) (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		log.Println("Failed to get current directory:", err)
		return "", err
	}

	// Construct the path to the key file
	relativePath := filepath.Join(cwd, "upay_private_key.pem")
	privateKey, err := LoadPEMKey(relativePath)
	if err != nil {
		log.Println("Failed to load private key:", err)
		return "", err
	}

	// Hash the data with SHA-256
	hasher := sha256.New()
	hasher.Write([]byte(requestData))
	hashed := hasher.Sum(nil)

	// Sign the hashed data with RSA
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed)
	if err != nil {
		log.Println("Failed to sign data:", err)
		return "", err
	}

	// Encode the signature in base64
	encryptedSignatureDataStr := base64.StdEncoding.EncodeToString(signature)
	return encryptedSignatureDataStr, nil
}

// LoadPEMKey loads an RSA private key from a PEM file
func LoadPEMKey(path string) (*rsa.PrivateKey, error) {
	keyData, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(keyData)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing private key")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %w", err)
	}

	return privateKey, nil
}

func main() {
	clientID := "21"
	userID := "upay2pbl"
	password := "ZFhCaGVUSndZbXd4TWpNMA=="

	// Concatenate clientID, userID, and password
	requestData := clientID + userID + password

	// Generate signature
	dataSign, err := SignatureGenerate(requestData)
	if err != nil {
		log.Fatalf("Failed to generate data sign: %v", err)
	}

	// Output the datasign
	fmt.Printf("Data Sign: %s\n", dataSign)
}
