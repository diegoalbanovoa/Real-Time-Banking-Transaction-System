package transaction

import "time"

// Transaction representa una transacción bancaria en el sistema.
// Cada transacción contiene información sobre el ID de la cuenta, el monto,
// el tipo de transacción (por ejemplo, depósito o retiro) y la fecha de creación.
type Transaction struct {
	ID              int       // Identificador único de la transacción (probablemente asignado por la base de datos)
	AccountID       int       // ID de la cuenta a la que se aplica la transacción
	Amount          float64   // Monto de la transacción (puede ser positivo para depósitos, negativo para retiros)
	TransactionType string    // Tipo de transacción: puede ser "deposit" o "withdrawal"
	CreatedAt       time.Time // Marca de tiempo que indica cuándo fue creada la transacción
}
