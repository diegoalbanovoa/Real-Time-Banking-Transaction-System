package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dsn             = "bankuser:bankpassword@tcp(127.0.0.1:3306)/bankdb"
	numAccounts     = 100 // Número de cuentas a generar
	numTransactions = 500 // Número de transacciones a generar por cuenta
)

func main() {
	// Conectar a la base de datos
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error al conectar a la base de datos:", err)
	}
	defer db.Close()

	// Verificar la conexión
	err = db.Ping()
	if err != nil {
		log.Fatal("Error al hacer ping a la base de datos:", err)
	}

	// Semilla para los datos aleatorios
	rand.Seed(time.Now().UnixNano())

	// Generar cuentas bancarias
	for i := 0; i < numAccounts; i++ {
		accountNumber := fmt.Sprintf("ACC%04d", i+1)
		initialBalance := rand.Float64() * 10000 // Balance inicial aleatorio entre 0 y 10,000

		// Insertar cuenta en la base de datos
		_, err := db.Exec("INSERT INTO accounts (account_number, balance) VALUES (?, ?)", accountNumber, initialBalance)
		if err != nil {
			log.Fatalf("Error al insertar cuenta %s: %v", accountNumber, err)
		}

		// Obtener el ID de la cuenta recién creada
		var accountID int
		err = db.QueryRow("SELECT id FROM accounts WHERE account_number = ?", accountNumber).Scan(&accountID)
		if err != nil {
			log.Fatalf("Error al obtener ID de la cuenta %s: %v", accountNumber, err)
		}

		// Generar transacciones para cada cuenta
		for j := 0; j < numTransactions; j++ {
			amount := rand.Float64() * 1000 // Monto aleatorio entre 0 y 1,000
			transactionType := "deposit"
			if rand.Intn(2) == 0 {
				transactionType = "withdrawal"
			}

			// Insertar transacción en la base de datos
			_, err := db.Exec("INSERT INTO transactions (account_id, amount, transaction_type) VALUES (?, ?, ?)", accountID, amount, transactionType)
			if err != nil {
				log.Fatalf("Error al insertar transacción para la cuenta %d: %v", accountID, err)
			}
		}

		log.Printf("Cuenta %s creada con éxito con %d transacciones.", accountNumber, numTransactions)
	}

	log.Println("Proceso de generación de datos completado.")
}
