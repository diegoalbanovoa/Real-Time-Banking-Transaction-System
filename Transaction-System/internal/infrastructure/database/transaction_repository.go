package database

import (
	"Transaction-System/internal/domain/transaction"
	"database/sql"
)

// TransactionRepository es un repositorio para interactuar con las transacciones en la base de datos
type TransactionRepository struct {
	db *sql.DB
}

// Aseguramos que TransactionRepository implementa la interfaz transaction.Repository
var _ transaction.Repository = &TransactionRepository{}

// NewTransactionRepository crea una nueva instancia del repositorio
func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

// Save guarda una transacción en la base de datos
func (r *TransactionRepository) Save(t *transaction.Transaction) error {
	_, err := r.db.Exec("INSERT INTO transactions (account_id, amount, transaction_type, created_at) VALUES (?, ?, ?, ?)",
		t.AccountID, t.Amount, t.TransactionType, t.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}
