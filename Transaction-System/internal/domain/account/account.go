package account

import (
	"fmt"  // Paquete para formatear y manejar errores
	"time" // Paquete para manejar fechas y horas
)

// Account representa una cuenta bancaria en el dominio del sistema.
// Contiene un número de cuenta, un balance actual, una identificación única y
// la fecha de creación de la cuenta.
type Account struct {
	ID            int       // Identificador único de la cuenta
	AccountNumber string    // Número de cuenta único
	Balance       float64   // Balance actual de la cuenta
	CreatedAt     time.Time // Fecha de creación de la cuenta
}

// NewAccount es un constructor que crea una nueva instancia de una cuenta bancaria.
// Recibe el número de cuenta y el balance inicial como parámetros.
func NewAccount(accountNumber string, balance float64) *Account {
	return &Account{
		AccountNumber: accountNumber, // Asigna el número de cuenta
		Balance:       balance,       // Asigna el balance inicial
		CreatedAt:     time.Now(),    // Establece la fecha de creación como la fecha y hora actual
	}
}

// Withdraw realiza un retiro de la cuenta bancaria.
// Si el monto del retiro es mayor que el balance actual, devuelve un error de fondos insuficientes.
func (a *Account) Withdraw(amount float64) error {
	// Verificar si hay suficientes fondos
	if amount > a.Balance {
		return fmt.Errorf("fondos insuficientes") // Devuelve un error si no hay fondos suficientes
	}
	a.Balance -= amount // Disminuye el balance con el monto retirado
	return nil          // No se devuelve ningún error si el retiro es exitoso
}

// Deposit realiza un depósito en la cuenta bancaria.
// Simplemente aumenta el balance de la cuenta con el monto especificado.
func (a *Account) Deposit(amount float64) {
	a.Balance += amount // Aumenta el balance con el monto depositado
}
