package transaction

import "time"

// Transaction representa una transacción bancaria
type Transaction struct {
	ID              int
	AccountID       int
	Amount          float64
	TransactionType string
	CreatedAt       time.Time
}

// New crea una nueva transacción
func New(accountID int, amount float64, transactionType string) *Transaction {
	return &Transaction{
		AccountID:       accountID,
		Amount:          amount,
		TransactionType: transactionType,
		CreatedAt:       time.Now(),
	}
}
