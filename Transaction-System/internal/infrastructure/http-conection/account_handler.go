package http_conection

import (
	"Transaction-System/internal/application"
	"encoding/json"
	"net/http"
)

type AccountHandler struct {
	service *application.TransactionService
}

// NewAccountHandler crea un nuevo controlador de cuentas
func NewAccountHandler(service *application.TransactionService) *AccountHandler {
	return &AccountHandler{service: service}
}

// DepositHandler maneja las solicitudes de depósito
func (h *AccountHandler) DepositHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		AccountID int     `json:"account_id"`
		Amount    float64 `json:"amount"`
	}

	// Decodificar la solicitud JSON
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Solicitud inválida", http.StatusBadRequest)
		return
	}

	// Procesar la transacción de depósito
	err = h.service.ProcessTransaction(request.AccountID, request.Amount, "deposit")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Depósito exitoso"))
}

// WithdrawHandler maneja las solicitudes de retiro
func (h *AccountHandler) WithdrawHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		AccountID int     `json:"account_id"`
		Amount    float64 `json:"amount"`
	}

	// Decodificar la solicitud JSON
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Solicitud inválida", http.StatusBadRequest)
		return
	}

	// Procesar la transacción de retiro
	err = h.service.ProcessTransaction(request.AccountID, request.Amount, "withdrawal")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Retiro exitoso"))
}
