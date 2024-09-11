package database

import (
	"Transaction-System/internal/domain/account"
	"database/sql"
	"time"
)

// AccountRepository es una implementación de la interfaz account.Repository.
// Este repositorio se encarga de interactuar con la base de datos para las operaciones CRUD
// relacionadas con las cuentas bancarias.
type AccountRepository struct {
	db *sql.DB // Conexión a la base de datos SQL.
}

// Asegurar que AccountRepository implementa la interfaz account.Repository
// Esta línea asegura que AccountRepository satisface la interfaz account.Repository.
// Si AccountRepository no implementa todos los métodos de la interfaz, el compilador generará un error.
var _ account.Repository = &AccountRepository{}

// NewAccountRepository es un constructor que crea un nuevo repositorio de cuentas.
// Parámetros:
// - db: una instancia de *sql.DB que representa la conexión a la base de datos.
// Retorna:
// - Un puntero a AccountRepository que interactúa con la base de datos a través de db.
func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

// Save guarda una nueva cuenta en la base de datos.
// Parámetros:
// - a: un puntero a la estructura account.Account que contiene la información de la cuenta a guardar.
// Retorna:
// - error: retorna un error si la operación de guardado falla, de lo contrario, retorna nil.
func (r *AccountRepository) Save(a *account.Account) error {
	// La consulta INSERT inserta el número de cuenta, el balance y la fecha de creación en la tabla 'accounts'.
	// Utilizamos 'Exec' ya que no necesitamos recuperar ningún valor de la base de datos.
	_, err := r.db.Exec("INSERT INTO accounts (account_number, balance, created_at) VALUES (?, ?, ?)",
		a.AccountNumber, a.Balance, a.CreatedAt)

	// Si ocurre un error durante la ejecución de la consulta, se retorna el error.
	return err
}

// FindByID busca una cuenta en la base de datos por su ID único.
// Parámetros:
// - id: el ID de la cuenta que se desea buscar.
// Retorna:
// - *account.Account: un puntero a la estructura account.Account si la cuenta existe.
// - error: retorna un error si la cuenta no se encuentra o si ocurre algún problema durante la consulta.
func (r *AccountRepository) FindByID(id int) (*account.Account, error) {
	var a account.Account   // Estructura para almacenar los datos de la cuenta.
	var createdAtStr string // Variable para almacenar temporalmente la fecha de creación como string.

	// Realiza una consulta SELECT a la base de datos para obtener la cuenta con el ID proporcionado.
	// QueryRow se utiliza para ejecutar la consulta ya que esperamos un solo resultado (una sola fila).
	// Scan asigna los valores retornados por la consulta a las variables de destino.
	err := r.db.QueryRow("SELECT id, account_number, balance, created_at FROM accounts WHERE id = ?", id).
		Scan(&a.ID, &a.AccountNumber, &a.Balance, &createdAtStr)

	// Si ocurre algún error (como que no se encuentre la cuenta), se retorna el error.
	if err != nil {
		return nil, err
	}

	// Convertir el valor de la cadena createdAtStr en un valor de tipo time.Time.
	// El formato "2006-01-02 15:04:05" es el formato estándar que utiliza Go para analizar fechas.
	a.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdAtStr)
	if err != nil {
		return nil, err
	}

	// Retorna la cuenta encontrada.
	return &a, nil
}
