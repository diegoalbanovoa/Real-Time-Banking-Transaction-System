package application

import (
	"Transaction-System/internal/domain/account"
	"Transaction-System/internal/domain/transaction"
	"fmt"
)

type TransactionService struct {
	accountRepo     account.Repository
	transactionRepo transaction.Repository
}

// NewTransactionService crea un nuevo servicio de transacciones
func NewTransactionService(aRepo account.Repository, tRepo transaction.Repository) *TransactionService {
	return &TransactionService{
		accountRepo:     aRepo,
		transactionRepo: tRepo,
	}
}

// ProcessTransaction procesa una transacción de depósito o retiro
func (s *TransactionService) ProcessTransaction(accountID int, amount float64, transactionType string) error {
	// Obtener la cuenta por ID
	acc, err := s.accountRepo.FindByID(accountID)
	if err != nil {
		return err
	}

	// Procesar la transacción según el tipo
	switch transactionType {
	case "deposit":
		acc.Deposit(amount)
	case "withdrawal":
		if err := acc.Withdraw(amount); err != nil {
			return err
		}
	default:
		return fmt.Errorf("tipo de transacción no válido")
	}

	// Guardar la transacción en la base de datos
	tr := transaction.New(accountID, amount, transactionType)
	return s.transactionRepo.Save(tr)
}
