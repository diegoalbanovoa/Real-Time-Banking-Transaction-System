package http_test

import (
	"Transaction-System/internal/application"
	"Transaction-System/internal/domain/account"
	"Transaction-System/internal/domain/transaction"
	"errors"
	"testing"
)

// Mock para el repositorio de cuentas
type mockAccountRepository struct {
	accounts map[int]*account.Account
}

func (m *mockAccountRepository) Save(a *account.Account) error {
	m.accounts[a.ID] = a
	return nil
}

func (m *mockAccountRepository) FindByID(id int) (*account.Account, error) {
	if account, exists := m.accounts[id]; exists {
		return account, nil
	}
	return nil, errors.New("cuenta no encontrada")
}

// Mock para el repositorio de transacciones
type mockTransactionRepository struct{}

func (m *mockTransactionRepository) Save(t *transaction.Transaction) error {
	return nil
}

func TestProcessTransaction_Deposit(t *testing.T) {
	// Crear un mock del repositorio de cuentas
	accountRepo := &mockAccountRepository{
		accounts: map[int]*account.Account{
			1: {ID: 1, AccountNumber: "ACC123", Balance: 100.0},
		},
	}
	transactionRepo := &mockTransactionRepository{}
	service := application.NewTransactionService(accountRepo, transactionRepo)

	// Probar un depósito
	err := service.ProcessTransaction(1, 50.0, "deposit")
	if err != nil {
		t.Fatalf("Error al procesar el depósito: %v", err)
	}

	// Verificar que el balance se haya actualizado correctamente
	acc, _ := accountRepo.FindByID(1)
	if acc.Balance != 150.0 {
		t.Errorf("Balance incorrecto tras el depósito, esperado 150.0, obtenido %v", acc.Balance)
	}
}

func TestProcessTransaction_Withdraw(t *testing.T) {
	// Crear un mock del repositorio de cuentas
	accountRepo := &mockAccountRepository{
		accounts: map[int]*account.Account{
			1: {ID: 1, AccountNumber: "ACC123", Balance: 100.0},
		},
	}
	transactionRepo := &mockTransactionRepository{}
	service := application.NewTransactionService(accountRepo, transactionRepo)

	// Probar un retiro exitoso
	err := service.ProcessTransaction(1, 50.0, "withdrawal")
	if err != nil {
		t.Fatalf("Error al procesar el retiro: %v", err)
	}

	// Verificar que el balance se haya actualizado correctamente
	acc, _ := accountRepo.FindByID(1)
	if acc.Balance != 50.0 {
		t.Errorf("Balance incorrecto tras el retiro, esperado 50.0, obtenido %v", acc.Balance)
	}
}

func TestProcessTransaction_Withdraw_InsufficientFunds(t *testing.T) {
	// Crear un mock del repositorio de cuentas
	accountRepo := &mockAccountRepository{
		accounts: map[int]*account.Account{
			1: {ID: 1, AccountNumber: "ACC123", Balance: 100.0},
		},
	}
	transactionRepo := &mockTransactionRepository{}
	service := application.NewTransactionService(accountRepo, transactionRepo)

	// Probar un retiro con fondos insuficientes
	err := service.ProcessTransaction(1, 150.0, "withdrawal")
	if err == nil {
		t.Fatal("Se esperaba un error por fondos insuficientes, pero no se recibió ninguno")
	}

	// Verificar que el balance no haya cambiado
	acc, _ := accountRepo.FindByID(1)
	if acc.Balance != 100.0 {
		t.Errorf("El balance no debería haber cambiado, esperado 100.0, obtenido %v", acc.Balance)
	}
}
