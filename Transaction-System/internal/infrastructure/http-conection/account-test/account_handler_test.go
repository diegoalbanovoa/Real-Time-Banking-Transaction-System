package account_test

import (
	"Transaction-System/internal/application"
	"Transaction-System/internal/domain/account"
	"Transaction-System/internal/domain/transaction"
	"Transaction-System/internal/infrastructure/http-conection"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

// mockAccountRepository es una implementación simulada (mock) del repositorio de cuentas.
// Este mock actúa como una base de datos en memoria utilizando un mapa (map[int]*account.Account)
// para almacenar las cuentas. Se utiliza para realizar pruebas unitarias sin necesidad de interactuar
// con una base de datos real.
type mockAccountRepository struct {
	accounts map[int]*account.Account // Mapa que simula la base de datos de cuentas, donde la clave es el ID.
}

// Save simula el guardado de una cuenta en el repositorio.
// El método agrega la cuenta al mapa de cuentas usando su ID como clave.
func (m *mockAccountRepository) Save(a *account.Account) error {
	// Simula el guardado de la cuenta en el repositorio (almacena en el mapa).
	m.accounts[a.ID] = a
	return nil
}

func (m *mockAccountRepository) FindByID(id int) (*account.Account, error) {
	if account, exists := m.accounts[id]; exists {
		return account, nil
	}
	return nil, errors.New("cuenta no encontrada")
}

// Mock del repositorio de transacciones
type mockTransactionRepository struct{}

func (m *mockTransactionRepository) Save(t *transaction.Transaction) error {
	return nil
}

// Prueba del endpoint /deposit
func TestDepositHandler(t *testing.T) {
	// Crear mocks de los repositorios
	accountRepo := &mockAccountRepository{
		accounts: map[int]*account.Account{
			100: {ID: 100, AccountNumber: "ACC0100", Balance: 5000.0},
		},
	}
	transactionRepo := &mockTransactionRepository{}

	// Crear el servicio de transacciones usando los mocks
	service := application.NewTransactionService(accountRepo, transactionRepo)

	// Crear el handler usando el servicio mockeado
	handler := http_conection.NewAccountHandler(service)

	// Crear una solicitud de prueba para un depósito
	body := map[string]interface{}{
		"account_id": 100,
		"amount":     100.0,
	}
	bodyBytes, _ := json.Marshal(body)
	req, err := http.NewRequest("POST", "/deposit", bytes.NewBuffer(bodyBytes))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Crear un ResponseRecorder para capturar la respuesta
	rr := httptest.NewRecorder()

	// Llamar al handler
	handler.DepositHandler(rr, req)

	// Verificar que la respuesta sea HTTP 200 OK
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Código de estado incorrecto: obtenido %v, esperado %v", status, http.StatusOK)
	}

	// Verificar que el cuerpo de la respuesta sea el esperado
	expected := "Depósito exitoso"
	if rr.Body.String() != expected {
		t.Errorf("Respuesta incorrecta: obtenida %v, esperada %v", rr.Body.String(), expected)
	}
}

// Prueba del endpoint /withdraw
func TestWithdrawHandler(t *testing.T) {
	// Crear mocks de los repositorios
	accountRepo := &mockAccountRepository{
		accounts: map[int]*account.Account{
			100: {ID: 100, AccountNumber: "ACC0100", Balance: 5000.0},
		},
	}
	transactionRepo := &mockTransactionRepository{}

	// Crear el servicio de transacciones usando los mocks
	service := application.NewTransactionService(accountRepo, transactionRepo)

	// Crear el handler usando el servicio mockeado
	handler := http_conection.NewAccountHandler(service)

	// Crear una solicitud de prueba para un retiro
	body := map[string]interface{}{
		"account_id": 100,
		"amount":     200.0,
	}
	bodyBytes, _ := json.Marshal(body)
	req, err := http.NewRequest("POST", "/withdraw", bytes.NewBuffer(bodyBytes))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Crear un ResponseRecorder para capturar la respuesta
	rr := httptest.NewRecorder()

	// Llamar al handler
	handler.WithdrawHandler(rr, req)

	// Verificar que la respuesta sea HTTP 200 OK
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Código de estado incorrecto: obtenido %v, esperado %v", status, http.StatusOK)
	}

	// Verificar que el cuerpo de la respuesta sea el esperado
	expected := "Retiro exitoso"
	if rr.Body.String() != expected {
		t.Errorf("Respuesta incorrecta: obtenida %v, esperada %v", rr.Body.String(), expected)
	}
}

// Prueba del endpoint /withdraw con fondos insuficientes
func TestWithdrawHandler_InsufficientFunds(t *testing.T) {
	// Crear mocks de los repositorios
	accountRepo := &mockAccountRepository{
		accounts: map[int]*account.Account{
			100: {ID: 100, AccountNumber: "ACC0100", Balance: 100.0},
		},
	}
	transactionRepo := &mockTransactionRepository{}

	// Crear el servicio de transacciones usando los mocks
	service := application.NewTransactionService(accountRepo, transactionRepo)

	// Crear el handler usando el servicio mockeado
	handler := http_conection.NewAccountHandler(service)

	// Crear una solicitud de prueba para un retiro que excede el balance
	body := map[string]interface{}{
		"account_id": 100,
		"amount":     200.0,
	}
	bodyBytes, _ := json.Marshal(body)
	req, err := http.NewRequest("POST", "/withdraw", bytes.NewBuffer(bodyBytes))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Crear un ResponseRecorder para capturar la respuesta
	rr := httptest.NewRecorder()

	// Llamar al handler
	handler.WithdrawHandler(rr, req)

	// Verificar que la respuesta sea HTTP 500 Internal Server Error (por fondos insuficientes)
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("Código de estado incorrecto: obtenido %v, esperado %v", status, http.StatusInternalServerError)
	}

	// Verificar que el cuerpo de la respuesta contenga un mensaje de error
	expected := "fondos insuficientes"
	if !bytes.Contains(rr.Body.Bytes(), []byte(expected)) {
		t.Errorf("Respuesta incorrecta: obtenida %v, esperada %v", rr.Body.String(), expected)
	}
}
