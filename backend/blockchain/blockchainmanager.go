package blockchain

import (
	"github.com/ceponatia/autochain/backend/blockchain/block"
	"github.com/ceponatia/autochain/backend/blockchain/chain"
)

func AddTransactionToPool(transaction *block.Transaction) {
	chain.AddTransactionToPool(transaction)
}

func CreateNewBlock(validatorAddress string) *block.Block {
	return chain.CreateNewBlock(validatorAddress)
}

func AddBlockToChain(block *block.Block) error {
	return chain.AddBlockToChain(block)
}

func GetChain() []*block.Block {
	return chain.GetChain()
}
