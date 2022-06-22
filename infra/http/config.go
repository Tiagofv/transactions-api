package http

import "context"

type Server struct {
	host                   string
	database               string
	AccountsRepository     string
	TransactionsRepository string
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

func WithDatabase(db string) func(server *Server) {
	return func(server *Server) {
		server.database = db
	}
}

func WithRepositories(ctx *context.Context) func(server *Server) {
	return func(server *Server) {
		server.TransactionsRepository = "oi"
		server.AccountsRepository = "oi"
	}
}
