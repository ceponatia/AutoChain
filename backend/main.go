package main

import (
	"fmt"
	"time"

	"github.com/ceponatia/autochain/backend/blockchain/block"
	"github.com/ceponatia/autochain/backend/consensus/pos"
	"github.com/ceponatia/autochain/backend/consensus/weightedstaking"
)

func main() {
	// Create a few transactions
	t1 := &block.Transaction{From: "Alice", To: "Bob", Amount: 10, Fee: 0.1}
	t2 := &block.Transaction{From: "Bob", To: "Charlie", Amount: 5, Fee: 0.05}
	t3 := &block.Transaction{From: "Charlie", To: "Alice", Amount: 3, Fee: 0.03}

	// Add transactions to the pool
	blockchain.AddTransactionToPool(t1)
	blockchain.AddTransactionToPool(t2)
	blockchain.AddTransactionToPool(t3)

	// Example validators
	validators := []*pos.Validator{
		{Address: "Validator1", Stake: 1000, StakingPeriod: 30 * 24 * time.Hour},
		{Address: "Validator2", Stake: 2000, StakingPeriod: 60 * 24 * time.Hour},
		{Address: "Validator3", Stake: 1500, StakingPeriod: 45 * 24 * time.Hour},
	}

	// Select the next validator using weighted staking
	posSystem := weightedstaking.New(validators)
	selectedValidator := posSystem.SelectValidator()

	// Add a new block to the chain with transactions from the pool
	newBlock := blockchain.CreateNewBlock(selectedValidator.Address)

	// Add the new block to the chain
	err := blockchain.AddBlockToChain(newBlock)
	if err != nil {
		fmt.Println("Error adding block to chain:", err)
		return
	}

	// Print the current blockchain
	chain := blockchain.GetChain()
	for _, block := range chain {
		fmt.Printf("Block index: %d, Timestamp: %s, Validator: %s\n", block.Index, block.Timestamp, block.Validator)
		for _, transaction := range block.Transactions {
			fmt.Printf("  Transaction: From %s To %s Amount %f Fee %f\n", transaction.From, transaction.To, transaction.Amount, transaction.Fee)
		}
	}
}
