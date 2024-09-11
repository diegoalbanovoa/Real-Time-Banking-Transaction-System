package http_conection

import (
	"Transaction-System/internal/application"
	"encoding/json"
	"net/http"
)

// AccountHandler maneja las solicitudes HTTP relacionadas con las cuentas bancarias.
// Utiliza el servicio TransactionService para procesar las transacciones de depósito y retiro.
type AccountHandler struct {
	service *application.TransactionService // Servicio de transacciones que procesa depósitos y retiros
}

// NewAccountHandler crea un nuevo controlador de cuentas (AccountHandler).
// Parámetros:
// - service: una instancia de TransactionService que procesa las transacciones.
// Retorna:
// - Un puntero a AccountHandler, que se utiliza para manejar las solicitudes HTTP relacionadas con las cuentas.
func NewAccountHandler(service *application.TransactionService) *AccountHandler {
	return &AccountHandler{service: service}
}

// DepositHandler maneja las solicitudes de depósito realizadas a través de HTTP.
// Procesa una transacción de depósito para la cuenta especificada en la solicitud.
// Parámetros:
// - w: el escritor de respuesta HTTP (http.ResponseWriter) que se utiliza para enviar la respuesta al cliente.
// - r: la solicitud HTTP entrante (http.Request), que contiene los datos del depósito.
func (h *AccountHandler) DepositHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		AccountID int     `json:"account_id"` // ID de la cuenta en la que se realizará el depósito
		Amount    float64 `json:"amount"`     // Monto del depósito
	}

	// Decodificar la solicitud JSON en la estructura request
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		// Si la solicitud no puede ser decodificada (por ejemplo, está mal formateada), devolver un error 400
		http.Error(w, "Solicitud inválida", http.StatusBadRequest)
		return
	}

	// Procesar la transacción de depósito utilizando el servicio
	err = h.service.ProcessTransaction(request.AccountID, request.Amount, "deposit")
	if err != nil {
		// Si ocurre un error al procesar la transacción, devolver un error 500
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Si la transacción es exitosa, devolver un código de estado 200 y un mensaje de éxito
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Depósito exitoso"))
}

// WithdrawHandler maneja las solicitudes de retiro realizadas a través de HTTP.
// Procesa una transacción de retiro para la cuenta especificada en la solicitud.
// Parámetros:
// - w: el escritor de respuesta HTTP (http.ResponseWriter) que se utiliza para enviar la respuesta al cliente.
// - r: la solicitud HTTP entrante (http.Request), que contiene los datos del retiro.
func (h *AccountHandler) WithdrawHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		AccountID int     `json:"account_id"` // ID de la cuenta de la que se retirarán los fondos
		Amount    float64 `json:"amount"`     // Monto del retiro
	}

	// Decodificar la solicitud JSON en la estructura request
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		// Si la solicitud no puede ser decodificada, devolver un error 400
		http.Error(w, "Solicitud inválida", http.StatusBadRequest)
		return
	}

	// Procesar la transacción de retiro utilizando el servicio
	err = h.service.ProcessTransaction(request.AccountID, request.Amount, "withdrawal")
	if err != nil {
		// Si ocurre un error al procesar la transacción, devolver un error 500
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Si la transacción es exitosa, devolver un código de estado 200 y un mensaje de éxito
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Retiro exitoso"))
}
