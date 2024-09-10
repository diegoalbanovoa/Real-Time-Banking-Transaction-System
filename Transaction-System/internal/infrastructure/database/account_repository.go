package database

import (
	"Transaction-System/internal/domain/account"
	"database/sql"
	"time"
)

type AccountRepository struct {
	db *sql.DB
}

// Asegurar que AccountRepository implementa account.Repository
var _ account.Repository = &AccountRepository{}

// NewAccountRepository crea un nuevo repositorio de cuentas
func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

func (r *AccountRepository) Save(a *account.Account) error {
	_, err := r.db.Exec("INSERT INTO accounts (account_number, balance, created_at) VALUES (?, ?, ?)",
		a.AccountNumber, a.Balance, a.CreatedAt)
	return err
}

func (r *AccountRepository) FindByID(id int) (*account.Account, error) {
	var a account.Account
	var createdAtStr string // Cambia el tipo para que primero obtengas el valor como string

	// Realiza la consulta
	err := r.db.QueryRow("SELECT id, account_number, balance, created_at FROM accounts WHERE id = ?", id).
		Scan(&a.ID, &a.AccountNumber, &a.Balance, &createdAtStr) // Escanea el valor creado como string
	if err != nil {
		return nil, err
	}

	// Convertir la cadena de fecha en time.Time
	a.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdAtStr)
	if err != nil {
		return nil, err
	}

	return &a, nil
}
