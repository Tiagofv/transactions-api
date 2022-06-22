package http

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func Run() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Could not read environment variables: %s", err.Error())
	}
	port := os.Getenv("SERVER_PORT")
	host := os.Getenv("HOST")

	router := mux.NewRouter()

	srv := &http.Server{
		Addr:         host + ":" + port,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}

	go func() {
		if err = srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// intercept shutdown via SIGINT
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	srv.Shutdown(ctx)

	log.Println("Shutting down gracefully!")
	os.Exit(0)
}