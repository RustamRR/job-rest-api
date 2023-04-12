package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
)

func GenerateKeys() error {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}

	privateKeyPem := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}

	f, err := os.Create("data/keys/private.pem")
	if err != nil {
		return err
	}

	defer f.Close()

	if err = pem.Encode(f, privateKeyPem); err != nil {
		return err
	}

	publicKeyPem := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(&privateKey.PublicKey),
	}

	f, err = os.Create("data/keys/public.pem")
	if err != nil {
		return err
	}

	defer f.Close()

	if err = pem.Encode(f, publicKeyPem); err != nil {
		return err
	}

	return nil
}
