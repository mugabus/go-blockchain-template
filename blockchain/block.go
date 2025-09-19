package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Block struct {
	Index        int           `json:"index"`
	Timestamp    int64         `json:"timestamp"`
	Transactions []Transaction `json:"transactions"`
	PrevHash     string        `json:"prev_hash"`
	Nonce        int64         `json:"nonce"`
	Hash         string        `json:"hash"`
}

func NewBlock(index int, txs []Transaction, prevHash string, difficulty int) Block {
	var nonce int64
	var hash string
	timestamp := time.Now().Unix()

	for {
		b := Block{Index: index, Timestamp: timestamp, Transactions: txs, PrevHash: prevHash, Nonce: nonce}
		h := b.CalculateHash()
		if strings.HasPrefix(h, strings.Repeat("0", difficulty)) {
			hash = h
			break
		}
		nonce++
	}

	return Block{Index: index, Timestamp: timestamp, Transactions: txs, PrevHash: prevHash, Nonce: nonce, Hash: hash}
}

func (b *Block) CalculateHash() string {
	txBytes, _ := json.Marshal(b.Transactions)
	record := fmt.Sprintf("%d%d%s%s%d", b.Index, b.Timestamp, string(txBytes), b.PrevHash, b.Nonce)
	h := sha256.Sum256([]byte(record))
	return hex.EncodeToString(h[:])
}
