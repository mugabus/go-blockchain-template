package api

import (
	"encoding/json"
	"net/http"

	"go-blockchain-template/blockchain"
)

type API struct {
	BC *blockchain.Blockchain
}

func NewAPI(bc *blockchain.Blockchain) *API {
	return &API{BC: bc}
}

func (api *API) writeJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func (api *API) HandleNewWallet(w http.ResponseWriter, r *http.Request) {
	wallet := blockchain.NewWallet()
	api.writeJSON(w, map[string]string{
		"address":     wallet.Address,
		"private_key": blockchain.ExportPrivateKey(wallet.PrivateKey),
	})
}

func (api *API) HandleGetBalance(w http.ResponseWriter, r *http.Request) {
	addr := r.URL.Query().Get("addr")
	bal := api.BC.GetBalance(addr)
	api.writeJSON(w, map[string]interface{}{"address": addr, "balance": bal})
}

func (api *API) HandleNewTransaction(w http.ResponseWriter, r *http.Request) {
	var tx blockchain.Transaction
	if err := json.NewDecoder(r.Body).Decode(&tx); err != nil {
		http.Error(w, "invalid transaction", http.StatusBadRequest)
		return
	}
	api.BC.AddTransaction(tx)
	api.writeJSON(w, map[string]string{"status": "transaction added"})
}

func (api *API) HandleMine(w http.ResponseWriter, r *http.Request) {
	miner := r.URL.Query().Get("miner")
	block := api.BC.MineBlock(miner)
	api.writeJSON(w, block)
}

func (api *API) HandleChain(w http.ResponseWriter, r *http.Request) {
	api.writeJSON(w, api.BC.Blocks)
}

func (api *API) HandleMempool(w http.ResponseWriter, r *http.Request) {
	api.writeJSON(w, api.BC.Mempool)
}
