// Autor: [Tu nombre]
// Descripción: Este programa genera cuentas bancarias y transacciones aleatorias
//              y las inserta en una base de datos MySQL. Cada cuenta tiene un número
//              único y un balance inicial aleatorio, seguido de un conjunto de transacciones
//              (depósitos y retiros) que se generan aleatoriamente. Es útil para simular
//              escenarios de carga en un sistema bancario.
// Fecha: [Fecha actual]

package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"

	_ "github.com/go-sql-driver/mysql" // Importamos el driver MySQL
)

// Constantes de configuración
const (
	dsn             = "bankuser:bankpassword@tcp(127.0.0.1:3306)/bankdb" // DSN (Data Source Name) para la conexión a MySQL
	numAccounts     = 100                                                // Número de cuentas bancarias a generar
	numTransactions = 500                                                // Número de transacciones a generar por cada cuenta
)

func main() {
	// Conectar a la base de datos usando la DSN proporcionada
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		// Si hay un error al conectar, se imprime el error y se detiene el programa
		log.Fatal("Error al conectar a la base de datos:", err)
	}
	defer db.Close() // Cerrar la conexión a la base de datos al finalizar el programa

	// Verificar la conexión a la base de datos con un ping
	err = db.Ping()
	if err != nil {
		// Si no se puede hacer ping a la base de datos, se imprime el error y se detiene el programa
		log.Fatal("Error al hacer ping a la base de datos:", err)
	}

	// Sembrar la semilla del generador de números aleatorios con el tiempo actual
	// Esto garantiza que los datos generados sean diferentes cada vez que se ejecute el programa
	rand.Seed(time.Now().UnixNano())

	// Bucle para generar las cuentas bancarias
	for i := 0; i < numAccounts; i++ {
		// Crear un número de cuenta en formato "ACCXXXX" donde XXXX es un número incremental
		accountNumber := fmt.Sprintf("ACC%04d", i+1)

		// Generar un balance inicial aleatorio entre 0 y 10,000 para la cuenta
		initialBalance := rand.Float64() * 10000

		// Insertar la cuenta bancaria con su número y balance en la base de datos
		_, err := db.Exec("INSERT INTO accounts (account_number, balance) VALUES (?, ?)", accountNumber, initialBalance)
		if err != nil {
			// Si hay un error al insertar la cuenta, se imprime el error y se detiene el programa
			log.Fatalf("Error al insertar cuenta %s: %v", accountNumber, err)
		}

		// Obtener el ID de la cuenta recién creada para asociar transacciones a esa cuenta
		var accountID int
		err = db.QueryRow("SELECT id FROM accounts WHERE account_number = ?", accountNumber).Scan(&accountID)
		if err != nil {
			// Si no se puede obtener el ID de la cuenta, se imprime el error y se detiene el programa
			log.Fatalf("Error al obtener ID de la cuenta %s: %v", accountNumber, err)
		}

		// Bucle para generar transacciones para la cuenta
		for j := 0; j < numTransactions; j++ {
			// Generar un monto aleatorio para la transacción entre 0 y 1,000
			amount := rand.Float64() * 1000

			// Determinar aleatoriamente si la transacción es un depósito o un retiro
			transactionType := "deposit"
			if rand.Intn(2) == 0 {
				transactionType = "withdrawal"
			}

			// Insertar la transacción en la base de datos vinculada al ID de la cuenta
			_, err := db.Exec("INSERT INTO transactions (account_id, amount, transaction_type) VALUES (?, ?, ?)", accountID, amount, transactionType)
			if err != nil {
				// Si hay un error al insertar la transacción, se imprime el error y se detiene el programa
				log.Fatalf("Error al insertar transacción para la cuenta %d: %v", accountID, err)
			}
		}

		// Informar en el log que la cuenta fue creada exitosamente junto con sus transacciones
		log.Printf("Cuenta %s creada con éxito con %d transacciones.", accountNumber, numTransactions)
	}

	// Al final del proceso, imprimir que se ha completado la generación de datos
	log.Println("Proceso de generación de datos completado.")
}
