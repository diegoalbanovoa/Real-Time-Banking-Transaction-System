package main

import (
	_ "Transaction-System/internal/infrastructure/http-conection"
	http_conection "Transaction-System/internal/infrastructure/http-conection"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/trace" // Para habilitar el trace
	"time"

	_ "net/http/pprof" // Para habilitar pprof en tu servicio

	"Transaction-System/internal/application"
	_ "Transaction-System/internal/domain/account"
	_ "Transaction-System/internal/domain/transaction"
	"Transaction-System/internal/infrastructure/database"
	_ "github.com/go-sql-driver/mysql"
)

// Middleware para registrar las solicitudes HTTP entrantes
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Registrar la solicitud entrante
		log.Printf("Solicitud entrante: %s %s desde %s", r.Method, r.RequestURI, r.RemoteAddr)

		// Continuar con la siguiente función del middleware o el manejador final
		next.ServeHTTP(w, r)

		// Registrar el tiempo que tomó manejar la solicitud
		log.Printf("Solicitud completada en %v", time.Since(start))
	})
}

func main() {
	// Crear un archivo de trace
	traceFile, err := os.Create("trace.out")
	if err != nil {
		log.Fatalf("No se puede crear el archivo de trace: %v", err)
	}
	defer traceFile.Close()

	// Iniciar el trace
	if err := trace.Start(traceFile); err != nil {
		log.Fatalf("No se puede iniciar el trace: %v", err)
	}
	defer trace.Stop()

	// Configurar la conexión a la base de datos MySQL
	dsn := "bankuser:bankpassword@tcp(127.0.0.1:3306)/bankdb"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}
	defer db.Close()

	// Verificar la conexión a la base de datos
	if err := db.Ping(); err != nil {
		log.Fatalf("No se puede conectar a la base de datos: %v", err)
	}

	// Inicializar los repositorios
	accountRepo := database.NewAccountRepository(db)
	transactionRepo := database.NewTransactionRepository(db)

	// Inicializar el servicio de transacciones
	transactionService := application.NewTransactionService(accountRepo, transactionRepo)

	// Inicializar los controladores HTTP
	accountHandler := http_conection.NewAccountHandler(transactionService)

	// Crear un mux para las rutas
	mux := http.NewServeMux()

	// Definir las rutas HTTP
	mux.HandleFunc("/deposit", accountHandler.DepositHandler)
	mux.HandleFunc("/withdraw", accountHandler.WithdrawHandler)

	// Habilitar pprof en el puerto 6060
	go func() {
		log.Println("Iniciando el servidor de pprof en :6060")
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	// Envolver el mux con el middleware de logging
	loggingHandler := loggingMiddleware(mux)

	// Iniciar el servidor HTTP en el puerto especificado (por defecto 8080)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Iniciando el servidor HTTP en el puerto %s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), loggingHandler))
}
