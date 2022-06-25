package dto

import (
	"time"
)

type TransactionDto struct {
	AccountID       int64     `json:"account_id,omitempty"`
	OperationTypeID int64     `json:"operation_type_id,omitempty"`
	Amount          float32   `json:"amount,omitempty"`
	EventDate       time.Time `json:"event_date"`
}
