package main

import (
	"fmt"
	"time"

	"./blockchain"
)

func main() {
	// Create a few transactions
	t1 := &blockchain.Transaction{From: "Alice", To: "Bob", Amount: 10, Fee: 0.1}
	t2 := &blockchain.Transaction{From: "Bob", To: "Charlie", Amount: 5, Fee: 0.05}
	t3 := &blockchain.Transaction{From: "Charlie", To: "Alice", Amount: 3, Fee: 0.03}

	// Add transactions to the pool
	blockchain.AddTransactionToPool(t1)
	blockchain.AddTransactionToPool(t2)
	blockchain.AddTransactionToPool(t3)

	// Example validators
	validators := []*blockchain.Validator{
		{Address: "Validator1", Stake: 1000, StakingPeriod: 30 * 24 * time.Hour},
		{Address: "Validator2", Stake: 2000, StakingPeriod: 60 * 24 * time.Hour},
		{Address: "Validator3", Stake: 1500, StakingPeriod: 45 * 24 * time.Hour},
	}

	// Select the next validator
	selectedValidator := blockchain.SelectValidator(validators)

	// Add a new block to the chain with transactions from the pool
	block := blockchain.CreateNewBlock(selectedValidator.Address)
	err := blockchain.AddBlockToChain(block)
	if err != nil {
		fmt.Println("Failed to add block:", err)
	} else {
		fmt.Println("Block added successfully")
	}

	// Print the current state of the blockchain
	for _, b := range blockchain.GetChain() {
		fmt.Printf("Block %d: %s\n", b.Index, b.Hash)
	}
}
