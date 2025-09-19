package main

import (
	"log"
	"net/http"

	"go-blockchain-template/api"
	"go-blockchain-template/blockchain"
)

func main() {
	log.Println("🚀 Starting Go Blockchain Node on :8080")

	// Create blockchain instance
	bc := blockchain.NewBlockchain()

	// Register API routes
	api := api.NewAPI(bc)

	http.HandleFunc("/wallet/new", api.HandleNewWallet)
	http.HandleFunc("/wallet/balance", api.HandleGetBalance)
	http.HandleFunc("/tx/new", api.HandleNewTransaction)
	http.HandleFunc("/mine", api.HandleMine)
	http.HandleFunc("/chain", api.HandleChain)
	http.HandleFunc("/mempool", api.HandleMempool)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
