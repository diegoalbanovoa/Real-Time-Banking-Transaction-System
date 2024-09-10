package account

import (
	"fmt"
	"time"
)

// Account representa una cuenta bancaria en el dominio
type Account struct {
	ID            int
	AccountNumber string
	Balance       float64
	CreatedAt     time.Time
}

// NewAccount crea una nueva cuenta con un balance inicial
func NewAccount(accountNumber string, balance float64) *Account {
	return &Account{
		AccountNumber: accountNumber,
		Balance:       balance,
		CreatedAt:     time.Now(),
	}
}

// Withdraw realiza un retiro en la cuenta
func (a *Account) Withdraw(amount float64) error {
	if amount > a.Balance {
		return fmt.Errorf("fondos insuficientes")
	}
	a.Balance -= amount
	return nil
}

// Deposit realiza un depósito en la cuenta
func (a *Account) Deposit(amount float64) {
	a.Balance += amount
}
