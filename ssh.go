package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
)

func ReadSSHKeyFrom(file, password string) (*ssh.PublicKeys, error) {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return nil, fmt.Errorf("private key file %s does not exist", file)
	}

	keys, err := ssh.NewPublicKeysFromFile("git", file, password)
	if err != nil {
		return nil, fmt.Errorf("failed to read private key file %s: %w", file, err)
	}

	return keys, nil
}

func GenerateSSHKey(path string, name string) error {
	if err := os.MkdirAll(filepath.Join(path, gitxFolder), 0755); err != nil {
		return fmt.Errorf("failed to create .gitx directory: %w", err)
	}

	privateKeyFile, err := os.Create(filepath.Join(path, gitxFolder, name+"prv.pem"))
	if err != nil {
		return fmt.Errorf("failed to create private key file: %w", err)
	}
	defer privateKeyFile.Close()

	publicKeyFile, err := os.Create(filepath.Join(path, gitxFolder, name+"pub.pem"))
	if err != nil {
		return fmt.Errorf("failed to create public key file: %w", err)
	}
	defer publicKeyFile.Close()

	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return fmt.Errorf("failed to generate private key: %w", err)
	}

	privateKeyDer := x509.MarshalPKCS1PrivateKey(privateKey)
	if err := pem.Encode(privateKeyFile, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyDer,
	}); err != nil {
		return fmt.Errorf("failed to write private key: %w", err)
	}

	publicKeyDer := x509.MarshalPKCS1PublicKey(&privateKey.PublicKey)
	if err := pem.Encode(publicKeyFile, &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKeyDer,
	}); err != nil {
		return fmt.Errorf("failed to write public key: %w", err)
	}

	return nil
}
