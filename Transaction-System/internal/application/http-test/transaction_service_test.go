package http_test

import (
	"Transaction-System/internal/application"
	"Transaction-System/internal/domain/account"
	"Transaction-System/internal/domain/transaction"
	"errors"
	"testing"
)

// Mock para el repositorio de cuentas
// Este mock almacena las cuentas en un mapa simulado en memoria, que actúa como base de datos temporal.
type mockAccountRepository struct {
	accounts map[int]*account.Account
}

// Método mock para guardar una cuenta
func (m *mockAccountRepository) Save(a *account.Account) error {
	m.accounts[a.ID] = a // Simula el almacenamiento de la cuenta en la "base de datos" (mapa en memoria)
	return nil
}

// Método mock para buscar una cuenta por ID
func (m *mockAccountRepository) FindByID(id int) (*account.Account, error) {
	if account, exists := m.accounts[id]; exists {
		return account, nil
	}
	return nil, errors.New("cuenta no encontrada") // Devuelve un error si la cuenta no existe
}

// Mock para el repositorio de transacciones
// Este mock simula la creación de transacciones sin interactuar con una base de datos real.
type mockTransactionRepository struct{}

// Método mock para guardar una transacción
func (m *mockTransactionRepository) Save(t *transaction.Transaction) error {
	// En este mock, no se hace nada, simplemente se simula que la transacción fue guardada correctamente.
	return nil
}

// Prueba para el procesamiento de depósitos
func TestProcessTransaction_Deposit(t *testing.T) {
	// Crear un mock del repositorio de cuentas con una cuenta inicial
	accountRepo := &mockAccountRepository{
		accounts: map[int]*account.Account{
			1: {ID: 1, AccountNumber: "ACC123", Balance: 100.0}, // Cuenta con un balance inicial de 100.0
		},
	}
	transactionRepo := &mockTransactionRepository{}                            // Mock del repositorio de transacciones
	service := application.NewTransactionService(accountRepo, transactionRepo) // Crear el servicio de transacciones

	// Probar un depósito de 50.0 a la cuenta
	err := service.ProcessTransaction(1, 50.0, "deposit")
	if err != nil {
		t.Fatalf("Error al procesar el depósito: %v", err)
	}

	// Verificar que el balance de la cuenta sea el correcto tras el depósito
	acc, _ := accountRepo.FindByID(1)
	if acc.Balance != 150.0 { // Balance esperado: 100.0 + 50.0 = 150.0
		t.Errorf("Balance incorrecto tras el depósito, esperado 150.0, obtenido %v", acc.Balance)
	}
}

// Prueba para el procesamiento de retiros
func TestProcessTransaction_Withdraw(t *testing.T) {
	// Crear un mock del repositorio de cuentas con una cuenta inicial
	accountRepo := &mockAccountRepository{
		accounts: map[int]*account.Account{
			1: {ID: 1, AccountNumber: "ACC123", Balance: 100.0}, // Balance inicial: 100.0
		},
	}
	transactionRepo := &mockTransactionRepository{}                            // Mock del repositorio de transacciones
	service := application.NewTransactionService(accountRepo, transactionRepo) // Crear el servicio de transacciones

	// Probar un retiro de 50.0 de la cuenta
	err := service.ProcessTransaction(1, 50.0, "withdrawal")
	if err != nil {
		t.Fatalf("Error al procesar el retiro: %v", err)
	}

	// Verificar que el balance de la cuenta sea el correcto tras el retiro
	acc, _ := accountRepo.FindByID(1)
	if acc.Balance != 50.0 { // Balance esperado: 100.0 - 50.0 = 50.0
		t.Errorf("Balance incorrecto tras el retiro, esperado 50.0, obtenido %v", acc.Balance)
	}
}

// Prueba para el retiro con fondos insuficientes
func TestProcessTransaction_Withdraw_InsufficientFunds(t *testing.T) {
	// Crear un mock del repositorio de cuentas con una cuenta inicial
	accountRepo := &mockAccountRepository{
		accounts: map[int]*account.Account{
			1: {ID: 1, AccountNumber: "ACC123", Balance: 100.0}, // Balance inicial: 100.0
		},
	}
	transactionRepo := &mockTransactionRepository{}                            // Mock del repositorio de transacciones
	service := application.NewTransactionService(accountRepo, transactionRepo) // Crear el servicio de transacciones

	// Probar un retiro de 150.0 (más de lo que hay en la cuenta)
	err := service.ProcessTransaction(1, 150.0, "withdrawal")
	if err == nil {
		// Se espera un error debido a fondos insuficientes
		t.Fatal("Se esperaba un error por fondos insuficientes, pero no se recibió ninguno")
	}

	// Verificar que el balance de la cuenta no haya cambiado
	acc, _ := accountRepo.FindByID(1)
	if acc.Balance != 100.0 { // Balance esperado: 100.0 (sin cambios)
		t.Errorf("El balance no debería haber cambiado, esperado 100.0, obtenido %v", acc.Balance)
	}
}
