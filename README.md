Go Blockchain Template
A simple, modular blockchain implementation in Go that you can use as a learning project or customize for your own dApp.

📂 Project Structure
The project is organized to be easy to understand and extend.

api/: Contains the HTTP API endpoints that let you interact with the blockchain.

blockchain/: The core logic of the blockchain, including blocks, transactions, wallets, and the Proof-of-Work algorithm.

cmd/node/: The entry point for running the blockchain node.

🚀 Getting Started
Follow these steps to get the blockchain node up and running.

Clone the repository:

Bash

git clone https://github.com/yourusername/go-blockchain-template.git
cd go-blockchain-template
Initialize Go modules:

Bash

go mod init go-blockchain-template
go mod tidy
Run the node:

Bash

go run ./cmd/node
You'll see a message that the node is running on http://localhost:8080.

🌐 API Endpoints
The blockchain provides a simple REST API for common actions.

Endpoint	Method	Description
/wallet/new	GET	Generates a new wallet (address + private key).
/wallet/balance?addr=ADDRESS	GET	Gets the balance for a specific address.
/tx/new	POST	Creates a new transaction.
/mine?miner=ADDRESS	GET	Mines a new block with pending transactions.
/chain	GET	Displays the entire blockchain.
/mempool	GET	Shows pending transactions waiting to be mined.

Export to Sheets
⚡ Key Features
This template includes fundamental blockchain concepts to get you started:

Proof-of-Work (PoW): A basic consensus mechanism.

Wallets: Uses ECDSA keys for secure transactions.

Signed Transactions: Ensures transactions are authentic and tamper-proof.

Mining Rewards: A simple reward system for miners.

In-memory data: The current state is stored in memory. Note that data will not persist after the application closes.

📌 Next Steps & Customization
This project is a great starting point. To build a more robust application, consider these enhancements:

Data Persistence: Use a database like BoltDB or LevelDB to save the blockchain to disk.

P2P Networking: Implement a peer-to-peer network for node synchronization.

Double-Spend Prevention: Add logic for nonce or gas to prevent duplicate transactions.

Smart Contracts: Integrate a virtual machine (e.g., WASM or EVM) for executing smart contracts.

✍️ Author: Eddy Mugaruka 📖 License: MIT
