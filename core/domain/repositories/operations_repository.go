package repositories

import (
	"context"
	"database/sql"
	"tiagofv.com/transactions/infra/database/queries"
	"time"
)

type OperationsInterface interface {
	GetOperation(id int64) (queries.OperationType, error)
}

type OperationsRepository struct {
	db  *sql.DB
	ctx *context.Context
}

func NewOperationsRepository(db *sql.DB, ctx *context.Context) *OperationsRepository {
	return &OperationsRepository{db: db, ctx: ctx}
}

func (o OperationsRepository) GetOperation(id int64) (queries.OperationType, error) {
	db := queries.New(o.db)
	ctx, cancel := context.WithTimeout(*o.ctx, time.Second*30)
	defer cancel()
	return db.GetOperation(ctx, id)
}
