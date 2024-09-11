package application

import (
	"Transaction-System/internal/domain/account"     // Importación del dominio de cuentas
	"Transaction-System/internal/domain/transaction" // Importación del dominio de transacciones
	"fmt"                                            // Paquete para formatear errores
)

// TransactionService es el servicio encargado de procesar transacciones
// como depósitos y retiros. Este servicio utiliza repositorios para interactuar con
// la capa de persistencia (base de datos).
type TransactionService struct {
	accountRepo     account.Repository     // Repositorio de cuentas, utilizado para acceder a las cuentas
	transactionRepo transaction.Repository // Repositorio de transacciones, utilizado para guardar transacciones
}

// NewTransactionService crea una instancia del servicio de transacciones
// Recibe dos repositorios: uno para cuentas y otro para transacciones.
func NewTransactionService(aRepo account.Repository, tRepo transaction.Repository) *TransactionService {
	return &TransactionService{
		accountRepo:     aRepo,
		transactionRepo: tRepo,
	}
}

// ProcessTransaction procesa una transacción de depósito o retiro para una cuenta dada
// Parametros:
//   - accountID: ID de la cuenta a la que se aplicará la transacción
//   - amount: Monto de la transacción
//   - transactionType: Tipo de transacción ("deposit" o "withdrawal")
//
// Devuelve un error si la transacción no puede ser procesada.
func (s *TransactionService) ProcessTransaction(accountID int, amount float64, transactionType string) error {
	// Obtener la cuenta por su ID
	acc, err := s.accountRepo.FindByID(accountID)
	if err != nil {
		// Si la cuenta no se encuentra, devolver un error
		return err
	}

	// Procesar la transacción dependiendo del tipo (depósito o retiro)
	switch transactionType {
	case "deposit":
		// Si es un depósito, aumentar el balance de la cuenta
		acc.Deposit(amount)
	case "withdrawal":
		// Si es un retiro, intentar disminuir el balance de la cuenta
		// Si los fondos son insuficientes, devolver un error
		if err := acc.Withdraw(amount); err != nil {
			return err
		}
	default:
		// Si el tipo de transacción no es válido, devolver un error
		return fmt.Errorf("tipo de transacción no válido")
	}

	// Crear una nueva transacción y guardarla en la base de datos
	tr := transaction.New(accountID, amount, transactionType)
	return s.transactionRepo.Save(tr)
}
