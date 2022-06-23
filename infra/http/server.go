package http

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"tiagofv.com/transactions/core/adapters/controllers"
	"tiagofv.com/transactions/core/domain/use_cases"
	"tiagofv.com/transactions/infra/database"
	"time"
)

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
		WithHost(host+":"+port),
		WithDatabase(database.InitDB()),
		WithRepositories(&ctx),
	)
	base := controllers.BaseController{CreateTransactionUseCase: use_cases.NewCreateTransactionUseCase(srv.TransactionsRepository)}
	router.HandleFunc("/transactions", base.CreateTransaction).Methods("POST")
	//srv := &http.Server{
	//	Addr:         host + ":" + port,
	//	WriteTimeout: time.Second * 15,
	//	ReadTimeout:  time.Second * 15,
	//	IdleTimeout:  time.Second * 60,
	//	Handler:      router,
	//}

	go func() {
		if err = srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// intercept shutdown via SIGINT
	signal.Notify(c, os.Interrupt)
	<-c

	defer cancel()

	srv.Shutdown(ctx)

	log.Println("Shutting down gracefully!")
	os.Exit(0)
}
