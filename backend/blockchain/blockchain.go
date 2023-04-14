package blockchain

// BlockchainManager is an interface to interact with blockchain-related operations
type BlockchainManager interface {
	CreateGenesisBlock() *Block
	CreateBlock(index int, transactions []*Transaction, prevHash string, validator string) *Block
	ComputeHash(block *Block) string
}
