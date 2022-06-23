package dto

import (
	"time"
)

type TransactionDto struct {
	AccountID       int64
	OperationTypeID int64
	Amount          float32
	EventDate       time.Time
}
