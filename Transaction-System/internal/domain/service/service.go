package application

import (
	"Transaction-System/internal/domain/account"     // Importa el dominio de cuentas
	"Transaction-System/internal/domain/transaction" // Importa el dominio de transacciones
	"fmt"                                            // Paquete para manejar errores y formatear mensajes
)

// TransactionService es el servicio que gestiona las transacciones bancarias (depósitos y retiros)
type TransactionService struct {
	accountRepo     account.Repository     // Repositorio de cuentas para acceder y modificar cuentas
	transactionRepo transaction.Repository // Repositorio de transacciones para almacenar transacciones
}

// NewTransactionService es un constructor que crea una nueva instancia del servicio de transacciones.
// Recibe como parámetros los repositorios necesarios para cuentas y transacciones.
func NewTransactionService(aRepo account.Repository, tRepo transaction.Repository) *TransactionService {
	return &TransactionService{
		accountRepo:     aRepo,
		transactionRepo: tRepo,
	}
}

// ProcessTransaction procesa una transacción de depósito o retiro en la cuenta especificada.
// Recibe el ID de la cuenta, el monto de la transacción y el tipo de transacción ("deposit" o "withdrawal").
func (s *TransactionService) ProcessTransaction(accountID int, amount float64, transactionType string) error {
	// Buscar la cuenta por su ID utilizando el repositorio de cuentas
	acc, err := s.accountRepo.FindByID(accountID)
	if err != nil {
		// Si no se encuentra la cuenta, devolver el error
		return err
	}

	// Determinar el tipo de transacción (depósito o retiro)
	switch transactionType {
	case "deposit":
		// Si es un depósito, sumar el monto al balance de la cuenta
		acc.Deposit(amount)
	case "withdrawal":
		// Si es un retiro, intentar restar el monto del balance
		// Si no hay suficientes fondos, acc.Withdraw devolverá un error
		if err := acc.Withdraw(amount); err != nil {
			return err
		}
	default:
		// Si el tipo de transacción no es válido, devolver un error
		return fmt.Errorf("tipo de transacción no válido")
	}

	// Crear una nueva transacción y guardarla en el repositorio de transacciones
	tr := transaction.New(accountID, amount, transactionType)
	return s.transactionRepo.Save(tr)
}
