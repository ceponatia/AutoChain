package blockchain

type Transaction struct {
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
	Fee    float64 `json:"fee"`
}

func NewTransaction(from string, to string, amount float64, fee float64) *Transaction {
	return &Transaction{From: from, To: to, Amount: amount, Fee: fee}
}
