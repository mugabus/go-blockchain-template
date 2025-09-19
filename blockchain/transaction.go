package blockchain

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"math/big"
)

type Transaction struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Amount int    `json:"amount"`
	SigR   string `json:"sig_r"`
	SigS   string `json:"sig_s"`
}

func NewTransaction(from, to string, amount int, priv *ecdsa.PrivateKey) Transaction {
	tx := Transaction{From: from, To: to, Amount: amount}
	hash := tx.Hash()
	r, s, _ := ecdsa.Sign(rand.Reader, priv, hash[:])
	tx.SigR, tx.SigS = r.String(), s.String()
	return tx
}

func (tx *Transaction) Hash() [32]byte {
	data := []byte(tx.From + tx.To + string(rune(tx.Amount)))
	return sha256.Sum256(data)
}

func (tx *Transaction) Verify(pub *ecdsa.PublicKey) bool {
	r := new(big.Int)
	s := new(big.Int)
	r.SetString(tx.SigR, 10)
	s.SetString(tx.SigS, 10)
	hash := tx.Hash()
	return ecdsa.Verify(pub, hash[:], r, s)
}
