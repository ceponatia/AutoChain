package blockchain

import "sync"

var (
	transactionPool []*Transaction
	poolMutex       sync.Mutex
)

// AddTransactionToPool adds a transaction to the transaction pool.
func AddTransactionToPool(tx *Transaction) {
	poolMutex.Lock()
	defer poolMutex.Unlock()

	transactionPool = append(transactionPool, tx)
}

// GetTransactionsFromPool retrieves a set of transactions from the pool for a new block.
func GetTransactionsFromPool(maxTransactions int) []*Transaction {
	poolMutex.Lock()
	defer poolMutex.Unlock()

	if len(transactionPool) < maxTransactions {
		maxTransactions = len(transactionPool)
	}

	selectedTransactions := transactionPool[:maxTransactions]
	transactionPool = transactionPool[maxTransactions:]

	return selectedTransactions
}
