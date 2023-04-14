package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

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
