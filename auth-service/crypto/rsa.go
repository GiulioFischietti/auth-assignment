package crypto

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func LoadPrivateKey(path string) (*rsa.PrivateKey, error) {

	data, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(data)

	if block == nil {
		return nil, fmt.Errorf("invalid PEM file")
	}

	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)

	if err != nil {
		return nil, err
	}

	privateKey, ok := key.(*rsa.PrivateKey)

	if !ok {
		return nil, fmt.Errorf("not an RSA private key")
	}

	return privateKey, nil
}
