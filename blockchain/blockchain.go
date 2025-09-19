package blockchain

type Blockchain struct {
	Blocks     []Block
	Mempool    []Transaction
	Balances   map[string]int
	Difficulty int
}

func NewBlockchain() *Blockchain {
	genesis := NewBlock(0, []Transaction{}, "0", 2)
	return &Blockchain{
		Blocks:     []Block{genesis},
		Mempool:    []Transaction{},
		Balances:   make(map[string]int),
		Difficulty: 2,
	}
}

func (bc *Blockchain) AddTransaction(tx Transaction) {
	bc.Mempool = append(bc.Mempool, tx)
}

func (bc *Blockchain) MineBlock(miner string) Block {
	newBlock := NewBlock(len(bc.Blocks), bc.Mempool, bc.Blocks[len(bc.Blocks)-1].Hash, bc.Difficulty)

	// Reward miner
	reward := Transaction{From: "network", To: miner, Amount: 10}
	newBlock.Transactions = append(newBlock.Transactions, reward)

	// Update balances
	for _, tx := range newBlock.Transactions {
		bc.Balances[tx.From] -= tx.Amount
		bc.Balances[tx.To] += tx.Amount
	}

	bc.Blocks = append(bc.Blocks, newBlock)
	bc.Mempool = []Transaction{}
	return newBlock
}

func (bc *Blockchain) GetBalance(addr string) int {
	return bc.Balances[addr]
}
