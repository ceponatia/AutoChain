package blockchain

import (
	"errors"
	"sync"
)

var (
	chain           []*Block
	chainMutex      sync.Mutex
	transactionPool []*Transaction
	blockManager    BlockchainManager
)

func init() {
	blockManager = NewBlockManager()
	// Initialize the blockchain with a genesis block.
	genesisBlock := blockManager.CreateGenesisBlock()
	chain = append(chain, genesisBlock)
}

// AddBlockToChain adds a new block to the chain after validating it.
func AddBlockToChain(block *Block) error {
	chainMutex.Lock()
	defer chainMutex.Unlock()

	if isValidBlock(block, chain[len(chain)-1]) {
		chain = append(chain, block)
		return nil
	}

	return errors.New("invalid block")
}

// isValidBlock validates the new block against the previous block in the chain.
func isValidBlock(newBlock, prevBlock *Block)
