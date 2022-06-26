package http

import (
	"context"
	"encoding/json"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"tiagofv.com/transactions/core/adapters/controllers"
	"tiagofv.com/transactions/core/domain/use_cases"
	_ "tiagofv.com/transactions/docs"
	"tiagofv.com/transactions/infra/database"
	"time"
)

// @title Swagger Transactions API
// @version 1.0
// @description This is the transactions API.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api
func Run() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Could not read environment variables: %s", err.Error())
	}
	port := os.Getenv("SERVER_PORT")
	host := os.Getenv("HOST")

	router := mux.NewRouter()
	backgroundCtx := context.Background()
	srv := New(
		WithHost(host),
		WithPort(port),
		WithDatabase(database.InitDB()),
		WithRepositories(&backgroundCtx),
	)
	base := controllers.NewBaseController(
		use_cases.NewCreateTransactionUseCase(srv.TransactionsRepository),
		use_cases.NewCreateAccountsUseCase(srv.AccountsRepository),
		use_cases.NewGetAccountUseCase(srv.AccountsRepository),
	)
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	router.HandleFunc("/api/transactions", base.CreateTransaction).Methods("POST")
	router.HandleFunc("/api/accounts", base.CreateAccount).Methods("POST")
	router.HandleFunc("/api/accounts/{id:[0-9]+}", base.GetAccount).Methods("GET")
	router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})
	loggedRouter := handlers.LoggingHandler(os.Stdout, router)
	configServer := srv.Start(loggedRouter)
	go func() {
		if err = configServer.ListenAndServe(); err != nil {
			log.Fatalf("Shutting down: %s", err.Error())
		}
	}()
	ctx, cancel := context.WithTimeout(backgroundCtx, time.Second*10)
	c := make(chan os.Signal, 1)
	// intercept shutdown via SIGINT
	signal.Notify(c, os.Interrupt)
	<-c

	defer cancel()

	configServer.Shutdown(ctx)

	log.Println("Shutting down gracefully!")
	os.Exit(0)
}
