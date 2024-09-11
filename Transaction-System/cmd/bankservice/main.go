// Autor: Diego Fernando Alba Novoa
// Descripción: Este programa implementa un sistema de transacciones bancarias en tiempo real utilizando Go.
//              Proporciona APIs para depositar y retirar dinero de cuentas bancarias. Utiliza MySQL como base de datos
//              para almacenar cuentas y transacciones. El programa también incluye monitoreo de rendimiento con pprof
//              y generación de trace para análisis detallado del sistema.
// Fecha: 10/09/2024

package main

import (
	_ "Transaction-System/internal/infrastructure/http-conection"              // Importación implícita para inicializar la conexión HTTP
	http_conection "Transaction-System/internal/infrastructure/http-conection" // Alias explícito para manejar HTTP
	"database/sql"                                                             // Paquete para trabajar con bases de datos SQL
	"fmt"                                                                      // Paquete para formatear cadenas
	"log"                                                                      // Paquete para loguear mensajes de información o errores
	"net/http"                                                                 // Paquete para manejar solicitudes HTTP
	"os"                                                                       // Paquete para interactuar con el sistema operativo, en este caso para crear archivos
	"runtime/trace"                                                            // Paquete para habilitar tracing y análisis de rendimiento
	"time"                                                                     // Paquete para trabajar con fechas y tiempos

	_ "net/http/pprof" // Paquete para habilitar el perfilado de pprof en el servidor

	"Transaction-System/internal/application"             // Módulo de aplicación para manejar la lógica de negocio
	_ "Transaction-System/internal/domain/account"        // Módulo de dominio para gestionar cuentas
	_ "Transaction-System/internal/domain/transaction"    // Módulo de dominio para gestionar transacciones
	"Transaction-System/internal/infrastructure/database" // Módulo de infraestructura para interactuar con la base de datos
	_ "github.com/go-sql-driver/mysql"                    // Driver MySQL para Go
)

// Middleware para registrar las solicitudes HTTP entrantes
// Este middleware registra el método, URI y la IP de origen de cada solicitud, así como el tiempo de ejecución.
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now() // Captura el tiempo de inicio de la solicitud

		// Registra la solicitud entrante
		log.Printf("Solicitud entrante: %s %s desde %s", r.Method, r.RequestURI, r.RemoteAddr)

		// Pasa la solicitud al siguiente manejador en la cadena
		next.ServeHTTP(w, r)

		// Calcula y registra el tiempo que tomó procesar la solicitud
		log.Printf("Solicitud completada en %v", time.Since(start))
	})
}

func main() {
	// Crear un archivo de trace que almacenará el rastro de ejecución del sistema
	traceFile, err := os.Create("trace.out")
	if err != nil {
		// Si no se puede crear el archivo de trace, se registra un error y se detiene el programa
		log.Fatalf("No se puede crear el archivo de trace: %v", err)
	}
	defer traceFile.Close() // Asegurarse de cerrar el archivo de trace al finalizar el programa

	// Iniciar el trace para comenzar a capturar eventos de ejecución
	if err := trace.Start(traceFile); err != nil {
		log.Fatalf("No se puede iniciar el trace: %v", err)
	}
	defer trace.Stop() // Detener el trace cuando el programa finalice

	// Configurar la conexión a la base de datos MySQL usando el DSN (Data Source Name)
	// El DSN incluye las credenciales y la dirección del servidor MySQL
	dsn := "bankuser:bankpassword@tcp(127.0.0.1:3306)/bankdb"
	db, err := sql.Open("mysql", dsn) // Abre una conexión a la base de datos
	if err != nil {
		// Si hay un error al conectar a la base de datos, se registra y se detiene el programa
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}
	defer db.Close() // Cerrar la conexión a la base de datos al finalizar el programa

	// Verificar la conexión a la base de datos con un "ping"
	if err := db.Ping(); err != nil {
		// Si no se puede establecer una conexión estable, se registra el error y se detiene el programa
		log.Fatalf("No se puede conectar a la base de datos: %v", err)
	}

	// Inicializar los repositorios de cuentas y transacciones, que interactúan con la base de datos
	accountRepo := database.NewAccountRepository(db)
	transactionRepo := database.NewTransactionRepository(db)

	// Crear el servicio de transacciones, que contiene la lógica para manejar las transacciones de cuentas
	transactionService := application.NewTransactionService(accountRepo, transactionRepo)

	// Crear los controladores HTTP para manejar las solicitudes de depósito y retiro
	accountHandler := http_conection.NewAccountHandler(transactionService)

	// Crear un nuevo "mux" que se encargará de enrutar las solicitudes HTTP
	mux := http.NewServeMux()

	// Definir las rutas HTTP y asociarlas con los manejadores correspondientes
	// La ruta "/deposit" manejará las solicitudes POST para depósitos en cuentas
	mux.HandleFunc("/deposit", accountHandler.DepositHandler)
	// La ruta "/withdraw" manejará las solicitudes POST para retiros de cuentas
	mux.HandleFunc("/withdraw", accountHandler.WithdrawHandler)

	// Habilitar pprof en un puerto separado (6060) para permitir el monitoreo de rendimiento
	go func() {
		log.Println("Iniciando el servidor de pprof en :6060")
		// Iniciar el servidor pprof
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	// Aplicar el middleware de logging al mux, para que todas las solicitudes pasen por el logger
	loggingHandler := loggingMiddleware(mux)

	// Determinar el puerto en el que el servidor HTTP principal escuchará solicitudes
	// Si no se define un puerto en las variables de entorno, usar el puerto por defecto (8080)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Iniciar el servidor HTTP en el puerto especificado
	log.Printf("Iniciando el servidor HTTP en el puerto %s\n", port)
	// El servidor usará el manejador con logging aplicado
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), loggingHandler))
}
