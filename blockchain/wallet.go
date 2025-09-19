package blockchain

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

type Wallet struct {
	PrivateKey *ecdsa.PrivateKey
	Address    string
}

func NewWallet() *Wallet {
	priv, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	pubBytes, _ := x509.MarshalPKIXPublicKey(&priv.PublicKey)
	address := fmt.Sprintf("%x", pubBytes)[:40]
	return &Wallet{PrivateKey: priv, Address: address}
}

func ExportPrivateKey(priv *ecdsa.PrivateKey) string {
	x509Encoded, _ := x509.MarshalECPrivateKey(priv)
	pemEncoded := pem.EncodeToMemory(&pem.Block{Type: "EC PRIVATE KEY", Bytes: x509Encoded})
	return string(pemEncoded)
}
