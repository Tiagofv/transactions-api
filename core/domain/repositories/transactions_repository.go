package repositories

import (
	"context"
	"database/sql"
	"errors"
	"tiagofv.com/transactions/core/domain/dto"
	"tiagofv.com/transactions/infra/database/queries"
	"time"
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
	ctx, cancel := context.WithTimeout(*t.ctx, time.Second*30)
	defer cancel()
	db := queries.New(t.db)
	_, err := db.GetOperation(ctx, transactionDto.OperationTypeID)
	if errors.Is(err, sql.ErrNoRows) {
		return queries.Transaction{}, err
	}
	inserted, err := db.CreateTransaction(ctx, queries.CreateTransactionParams{
		AccountID:       transactionDto.AccountID,
		OperationTypeID: transactionDto.OperationTypeID,
		Amount:          transactionDto.Amount,
		EventDate:       time.Now(),
	})
	if err != nil {
		return queries.Transaction{}, err
	}
	return inserted, nil
}
