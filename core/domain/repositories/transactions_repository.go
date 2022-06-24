package repositories

import (
	"context"
	"database/sql"
	"tiagofv.com/transactions/core/domain/dto"
	"tiagofv.com/transactions/infra/database/queries"
)

type TransactionsInterface interface {
	CreateTransaction(transaction dto.TransactionDto) (queries.Transaction, error)
}

type TransactionsRepository struct {
	db  *sql.DB
	ctx *context.Context
}

func NewTransactionsRepository(db *sql.DB, ctx *context.Context) *TransactionsRepository {
	return &TransactionsRepository{
		db:  db,
		ctx: ctx,
	}
}

func (t TransactionsRepository) CreateTransaction(transactionDto dto.TransactionDto) (queries.Transaction, error) {
	db := queries.New(t.db)
	inserted, err := db.CreateTransaction(*t.ctx, queries.CreateTransactionParams{
		AccountID:       transactionDto.AccountID,
		OperationTypeID: transactionDto.OperationTypeID,
		Amount:          transactionDto.Amount,
		EventDate:       transactionDto.EventDate,
	})
	if err != nil {
		return queries.Transaction{}, err
	}
	return inserted, nil
}
