package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type block struct {
	Index        int
	Timestamp    string
	Transactions []*Transaction
	PrevHash     string
	Hash         string
	Validator    string
}

var (
	chain           []*block
	transactionPool []*block.Transaction
	mtx             sync.RWMutex
)

// Creates the genesis block
func init() {
	genesisBlock := block.CreateBlock([]*block.Transaction{}, "", "GenesisValidator")
	chain = append(chain, genesisBlock)
}

// AddTransactionToPool prepares a transaction to be added to the next block
func AddTransactionToPool(transaction *block.Transaction) {
	mtx.Lock()
	transactionPool = append(transactionPool, transaction)
	mtx.Unlock()
}

// Collect transactions and create a new block
func CreateNewBlock(validatorAddress string) *block.Block {
	mtx.Lock()
	defer mtx.Unlock()

	prevBlock := chain[len(chain)-1]
	newBlock := block.CreateBlock(transactionPool, prevBlock.Hash, validatorAddress)
	transactionPool = []*block.Transaction{}

	return newBlock
}

// Append a new block to the chain
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

// Get the current blockchain
func GetChain() []*block.Block {
	mtx.RLock()
	defer mtx.RUnlock()
	return chain
}

type blockManager struct{}

// newBlockManager creates a new instance of blockManager
func newBlockManager() BlockchainManager {
	return &blockManager{}
}

// Block represents each 'item' in the blockchain
type Block struct {
	Index        int
	Timestamp    string
	Transactions []*Transaction
	PrevHash     string
	Hash         string
	Validator    string
}

// Transaction represents the data being stored in the blockchain
type Transaction struct {
	From   string
	To     string
	Amount float64
	Fee    float64
}

// createBlock creates a new block with the given transactions, previous hash, and validator.
func (bm *blockManager) CreateBlock(index int, transactions []*Transaction, prevHash string, validator string) *Block {
	block := &Block{Index: index, Timestamp: time.Now().String(), Transactions: transactions, PrevHash: prevHash, Validator: validator}
	block.Hash = bm.ComputeHash(block)
	return block
}

// computeHash computes the hash for the given block.
func (bm *blockManager) ComputeHash(block *Block) string {
	data := strconv.Itoa(block.Index) + block.Timestamp + block.PrevHash + block.Validator
	for _, tx := range block.Transactions {
		data += tx.From + tx.To + strconv.FormatFloat(tx.Amount, 'f', -1, 64) + strconv.FormatFloat(tx.Fee, 'f', -1, 64)
	}
	hash := sha256.New()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}

// createGenesisBlock creates the genesis block for the blockchain.
func (bm *blockManager) CreateGenesisBlock() *Block {
	transactions := []*Transaction{}
	return bm.CreateBlock(0, transactions, "0", "")
}

func TransactionHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		transaction := &block.Transaction{}
		err := json.NewDecoder(r.Body).Decode(transaction)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if transaction.From == "" || transaction.To == "" || transaction.Amount <= 0 {
			http.Error(w, "invalid transaction", http.StatusBadRequest)
			return
		}
		AddTransactionToPool(transaction)
		fmt.Fprintf(w, "Transaction added to pool")
	}
}
