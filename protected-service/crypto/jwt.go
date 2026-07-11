package crypto

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
)

func LoadPublicKey(
	path string,
) (*rsa.PublicKey, error) {

	data, err :=
		os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	block, _ :=
		pem.Decode(data)

	key, err :=
		x509.ParsePKIXPublicKey(
			block.Bytes,
		)

	if err != nil {
		return nil, err
	}

	return key.(*rsa.PublicKey), nil
}
