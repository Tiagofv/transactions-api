package repositories

import (
	"context"
	"database/sql"
	"tiagofv.com/transactions/core/domain/dto"
	"tiagofv.com/transactions/infra/database/queries"
	"time"
)

type AccountsInterface interface {
	CreateAccount(data dto.AccountDto) (queries.Account, error)
	FindAccount(id int64) (queries.Account, error)
	FindAccountByDocument(document string) (queries.Account, error)
}

type AccountsRepository struct {
	db  *sql.DB
	ctx *context.Context
}

func NewAccountsRepository(db *sql.DB, ctx *context.Context) *AccountsRepository {
	return &AccountsRepository{db: db, ctx: ctx}
}

func (r AccountsRepository) CreateAccount(data dto.AccountDto) (queries.Account, error) {
	db := queries.New(r.db)

	ctx, cancel := context.WithTimeout(*r.ctx, time.Second*30)
	defer cancel()
	account, err := db.CreateAccount(ctx, data.Document)
	if err != nil {
		return account, err
	}
	return account, nil
}

func (r AccountsRepository) FindAccount(id int64) (queries.Account, error) {
	db := queries.New(r.db)
	ctx, cancel := context.WithTimeout(*r.ctx, time.Second*30)
	defer cancel()
	account, err := db.GetAccount(ctx, id)
	if err != nil {
		return account, err
	}
	return account, err
}

func (r AccountsRepository) FindAccountByDocument(document string) (queries.Account, error) {
	ctx, cancel := context.WithTimeout(*r.ctx, time.Second*30)
	defer cancel()
	db := queries.New(r.db)
	account, err := db.GetAccountByDocument(ctx, document)
	if err != nil {
		return account, err
	}
	return account, nil
}
