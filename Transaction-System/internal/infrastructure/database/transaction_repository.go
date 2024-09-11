package database

import (
	"Transaction-System/internal/domain/transaction"
	"database/sql"
)

// TransactionRepository es un repositorio para interactuar con las transacciones en la base de datos.
// Implementa las operaciones CRUD (en este caso, sólo guardado) para las transacciones en la base de datos.
type TransactionRepository struct {
	db *sql.DB // Conexión a la base de datos SQL
}

// Aseguramos que TransactionRepository implementa la interfaz transaction.Repository.
// Esta línea asegura que TransactionRepository cumple con la interfaz transaction.Repository,
// lo que significa que si no implementa todos los métodos requeridos por la interfaz,
// el compilador generará un error.
var _ transaction.Repository = &TransactionRepository{}

// NewTransactionRepository crea una nueva instancia de TransactionRepository.
// Parámetros:
// - db: una instancia de *sql.DB que representa la conexión a la base de datos.
// Retorna:
// - Un puntero a TransactionRepository que puede usarse para interactuar con la base de datos.
func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

// Save guarda una transacción en la base de datos.
// Este método realiza una consulta SQL de tipo INSERT para almacenar los detalles de una transacción bancaria.
// Parámetros:
// - t: un puntero a una estructura transaction.Transaction que contiene los detalles de la transacción a guardar.
// Retorna:
// - error: retorna un error si la operación de guardado falla, de lo contrario retorna nil.
func (r *TransactionRepository) Save(t *transaction.Transaction) error {
	// La consulta INSERT inserta los detalles de la transacción en la tabla 'transactions'.
	// Los valores de account_id, amount, transaction_type y created_at se insertan en la tabla.
	_, err := r.db.Exec("INSERT INTO transactions (account_id, amount, transaction_type, created_at) VALUES (?, ?, ?, ?)",
		t.AccountID, t.Amount, t.TransactionType, t.CreatedAt)

	// Si ocurre algún error durante la inserción, lo retornamos para que pueda ser manejado por la lógica de la aplicación.
	if err != nil {
		return err
	}

	// Si la transacción se guarda correctamente, no hay errores que devolver.
	return nil
}
