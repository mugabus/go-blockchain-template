Go Blockchain Template

A simple, modular blockchain implementation in Go that you can use as a learning project or customize for your own dApp.

📂 Project Structure
go-blockchain-template/
├── api/
│   └── http.go          # HTTP API endpoints
│
├── blockchain/
│   ├── blockchain.go    # Blockchain manager
│   ├── block.go         # Block + Proof of Work
│   ├── transaction.go   # Transactions + signing
│   └── wallet.go        # Wallets + keys
│
├── cmd/
│   └── node/
│       └── main.go      # Entry point (runs node)
│
├── go.mod
└── READ

🚀 Getting Started

1. Clone & Setup
   git clone https://github.com/yourusername/go-blockchain-template.git
cd go-blockchain-template
go mod init go-blockchain-template
go mod tidy

go run ./cmd/node

🚀 Starting Go Blockchain Node on :8080

🌐 API Endpoints

| Endpoint                       | Method | Description                                     |
| ------------------------------ | ------ | ----------------------------------------------- |
| `/wallet/new`                  | GET    | Generate a new wallet (address + private key)   |
| `/wallet/balance?addr=ADDRESS` | GET    | Get balance of an address                       |
| `/tx/new`                      | POST   | Create a new transaction (from → to, amount)    |
| `/mine?miner=ADDRESS`          | GET    | Mine pending transactions, reward goes to miner |
| `/chain`                       | GET    | Show full blockchain                            |
| `/mempool`                     | GET    | Show pending transactions                       |

🛠️ Usage Examples (cURL)
1. Create a Wallet

   curl http://localhost:8080/wallet/new
Response:

{
  "address": "04d2c13...",
  "private_key": "-----BEGIN EC PRIVATE KEY-----\n..."
}
2. Check Balance

curl "http://localhost:8080/wallet/balance?addr=04d2c13..."

3. Create a Transaction

   curl -X POST http://localhost:8080/tx/new \
  -H "Content-Type: application/json" \
  -d '{"from":"ADDR1","to":"ADDR2","amount":5}'

4. Mine a Block

   curl "http://localhost:8080/mine?miner=ADDR1"
5. View Blockchain

   curl http://localhost:8080/chain

⚡ Features

Basic Proof-of-Work consensus

Wallets with ECDSA keys

Signed transactions

In-memory balances

Mining rewards

HTTP JSON API

📌 Next Steps / Customization

This template is intentionally simple. You can extend it with:

Persistence (BoltDB/LevelDB) so data survives restarts

P2P networking (so multiple nodes sync)

Nonce / gas handling (prevent double-spends)

Smart contract execution (WASM/EVM)

REST or gRPC APIs for your dApp frontend

✍️ Author: Eddy Mugaruka
📖 License: MIT


ME.md
