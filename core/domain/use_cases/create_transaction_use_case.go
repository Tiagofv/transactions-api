package use_cases

import (
	"database/sql"
	"errors"
	"tiagofv.com/transactions/core/domain/dto"
	"tiagofv.com/transactions/core/domain/repositories"
	"tiagofv.com/transactions/core/presenters"
)

type CreateTransactionUseCase struct {
	transactions repositories.TransactionsInterface
	operations   repositories.OperationsInterface
	accounts     repositories.AccountsInterface
}

func NewCreateTransactionUseCase(transactions repositories.TransactionsInterface, operations repositories.OperationsInterface, accounts repositories.AccountsInterface) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{transactions: transactions, operations: operations, accounts: accounts}
}

var ErrInvalidOperation = errors.New("Invalid operation ID.")
var ErrInvalidAccountId = errors.New("Invalid account ID.")

func (r CreateTransactionUseCase) Execute(dto dto.TransactionDto) (presenters.TransactionPresenter, error) {
	operation, err := r.operations.GetOperation(dto.OperationTypeID)
	if errors.Is(err, sql.ErrNoRows) {
		return presenters.TransactionPresenter{}, ErrInvalidOperation
	}

	_, err = r.accounts.FindAccount(dto.AccountID)
	if errors.Is(err, sql.ErrNoRows) {
		return presenters.TransactionPresenter{}, ErrInvalidAccountId
	}
	if operation.Type == "NEGATIVE" {
		dto.Amount = dto.Amount * -1
	}
	transaction, err := r.transactions.CreateTransaction(dto)
	if err != nil {
		return presenters.TransactionPresenter{}, err
	}
	return presenters.TransactionPresenter{
		AccountID:       transaction.AccountID,
		OperationTypeID: transaction.OperationTypeID,
		Amount:          transaction.Amount,
		EventDate:       transaction.EventDate,
	}, nil
}
