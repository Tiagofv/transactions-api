package http

import (
	"context"
	"encoding/json"
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
// @description This is a sample transactions api.
// @termsOfService http://swagger.io/terms/

// @contact.name Tiago braga
// @contact.url http://www.swagger.io/support
// @contact.email tiagofvx@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func Run() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Could not read environment variables: %s", err.Error())
	}
	port := os.Getenv("SERVER_PORT")
	host := os.Getenv("HOST")

	router := mux.NewRouter()
	srv := New(
		WithHost(host),
		WithPort(port),
		WithDatabase(database.InitDB()),
		WithRepositories(&ctx),
	)
	base := controllers.BaseController{CreateTransactionUseCase: use_cases.NewCreateTransactionUseCase(srv.TransactionsRepository)}
	router.HandleFunc("/swagger/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8080/swagger/swagger.json")))
	router.HandleFunc("/transactions", base.CreateTransaction).Methods("GET")
	router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		// an example API handler
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	configServer := srv.Start(router)
	go func() {
		if err = configServer.ListenAndServe(); err != nil {
			log.Fatalf("Shutting down: %s", err.Error())
		}
	}()

	c := make(chan os.Signal, 1)
	// intercept shutdown via SIGINT
	signal.Notify(c, os.Interrupt)
	<-c

	defer cancel()

	configServer.Shutdown(ctx)

	log.Println("Shutting down gracefully!")
	os.Exit(0)
}
