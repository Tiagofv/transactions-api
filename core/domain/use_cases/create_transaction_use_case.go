package use_cases

import (
	"tiagofv.com/transactions/core/domain/dto"
	"tiagofv.com/transactions/core/domain/repositories"
	"tiagofv.com/transactions/infra/database/queries"
)

type CreateTransactionUseCase struct {
	transactions repositories.TransactionsInterface
}

func NewCreateTransactionUseCase(transactions repositories.TransactionsInterface) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{transactions: transactions}
}

func (r CreateTransactionUseCase) Execute(dto dto.TransactionDto) (queries.Transaction, error) {
	transaction, err := r.transactions.CreateTransaction(dto)
	if err != nil {
		return queries.Transaction{}, err
	}
	return transaction, nil
}
