package http

import (
	"context"
	"database/sql"
	"net/http"
	"tiagofv.com/transactions/core/domain/repositories"
	"time"
)

type Server struct {
	host                   string
	port                   string
	database               *sql.DB
	AccountsRepository     repositories.AccountsInterface
	TransactionsRepository repositories.TransactionsInterface
	OperationsRepository   repositories.OperationsInterface
}

func New(options ...func(server *Server)) *Server {
	svr := &Server{}
	for _, o := range options {
		o(svr)
	}
	return svr
}

func (s Server) Start(router http.Handler) *http.Server {
	return &http.Server{
		Addr:         ":" + s.port,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		Handler:      router,
	}
}

func WithHost(host string) func(server *Server) {
	return func(server *Server) {
		server.host = host
	}
}

func WithPort(port string) func(server *Server) {
	return func(server *Server) {
		server.port = port
	}
}

func WithDatabase(db *sql.DB) func(server *Server) {
	return func(server *Server) {
		server.database = db
	}
}

func WithRepositories(ctx *context.Context) func(server *Server) {
	return func(server *Server) {
		server.TransactionsRepository = repositories.NewTransactionsRepository(server.database, ctx)
		server.AccountsRepository = repositories.NewAccountsRepository(server.database, ctx)
		server.OperationsRepository = repositories.NewOperationsRepository(server.database, ctx)
	}
}
