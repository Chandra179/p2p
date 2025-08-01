package cryptoutils

import (
	"encoding/base64"
	"fmt"

	"github.com/libp2p/go-libp2p/core/crypto"
)

// Decode base64 key (Ed25519, RSA, ECDSA, etc...)
func DecodeBase64Key(base64Key string) (crypto.PrivKey, error) {
	privBytes, err := base64.StdEncoding.DecodeString(base64Key)
	if err != nil {
		return nil, err
	}
	privKey, err := crypto.UnmarshalPrivateKey(privBytes)
	if err != nil {
		return nil, err
	}
	return privKey, nil
}

// GenerateEd25519Key generates a new Ed25519 private key and returns it as a base64-encoded string.
func GenerateEd25519Key() (string, error) {
	priv, _, err := crypto.GenerateKeyPair(crypto.Ed25519, -1)
	if err != nil {
		return "", fmt.Errorf("failed to generate key: %w", err)
	}

	bytes, err := crypto.MarshalPrivateKey(priv)
	if err != nil {
		return "", fmt.Errorf("failed to marshal private key: %w", err)
	}
	return base64.StdEncoding.EncodeToString(bytes), nil
}
