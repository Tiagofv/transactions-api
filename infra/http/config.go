package http

import (
	"context"
	"database/sql"
	"tiagofv.com/transactions/core/domain/repositories"
)

type Server struct {
	host                   string
	database               *sql.DB
	AccountsRepository     string
	TransactionsRepository repositories.TransactionsInterface
}

func New(options ...func(server *Server)) *Server {
	svr := &Server{}
	for _, o := range options {
		o(svr)
	}
	return svr
}

func WithHost(host string) func(server *Server) {
	return func(server *Server) {
		server.host = host
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
		server.AccountsRepository = "oi"
	}
}
