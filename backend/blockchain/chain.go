package chain

import (
	"errors"
	"sync"

	"github.com/ceponatia/autochain/backend/blockchain/block"
)

var (
	chain           []*block.Block
	transactionPool []*block.Transaction
	mtx             sync.RWMutex
)

func init() {
	genesisBlock := block.CreateBlock([]*block.Transaction{}, "", "GenesisValidator")
	chain = append(chain, genesisBlock)
}

func AddTransactionToPool(transaction *block.Transaction) {
	mtx.Lock()
	transactionPool = append(transactionPool, transaction)
	mtx.Unlock()
}

func CreateNewBlock(validatorAddress string) *block.Block {
	mtx.Lock()
	defer mtx.Unlock()

	prevBlock := chain[len(chain)-1]
	newBlock := block.CreateBlock(transactionPool, prevBlock.Hash, validatorAddress)
	transactionPool = []*block.Transaction{}

	return newBlock
}

func AddBlockToChain(newBlock *block.Block) error {
	mtx.Lock()
	defer mtx.Unlock()

	prevBlock := chain[len(chain)-1]

	if newBlock.PrevHash != prevBlock.Hash {
		return errors.New("invalid previous block hash")
	}

	chain = append(chain, newBlock)
	return nil
}

func GetChain() []*block.Block {
	mtx.RLock()
	defer mtx.RUnlock()
	return chain
}
