package database

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
	"time"
)

func InitDB() *sql.DB {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	DSN := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_DB"),
	)
	db, err := sql.Open("postgres", DSN)
	if err != nil {
		log.Fatalf("Error on database connection: %s", err.Error())
	}
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatalf("Error on database ping: %s", err.Error())
	}
	return db
}
