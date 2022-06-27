package use_cases

import (
	"database/sql"
	"github.com/stretchr/testify/suite"
	"testing"
	"tiagofv.com/transactions/core/domain/dto"
	"tiagofv.com/transactions/infra/database/queries"
	"tiagofv.com/transactions/tests/mocks"
	"time"
)

type CreateTransactionUseCaseSuite struct {
	suite.Suite
}

func (c CreateTransactionUseCaseSuite) BaseMocks() (*mocks.TransactionsInterface, *mocks.OperationsInterface, *mocks.AccountsInterface) {
	return new(mocks.TransactionsInterface), new(mocks.OperationsInterface), new(mocks.AccountsInterface)
}

func (c CreateTransactionUseCaseSuite) TestItReturnErrorIfOperationInvalid() {
	tr, op, ac := c.BaseMocks()
	transactionDto := dto.TransactionDto{
		AccountID:       1,
		OperationTypeID: 1,
		Amount:          100,
	}
	op.On("GetOperation", transactionDto.AccountID).Return(queries.OperationType{}, sql.ErrNoRows)
	tr.On("CreateTransaction", transactionDto).Times(0)
	ac.On("FindAccount", transactionDto.AccountID).Times(0)
	handler := NewCreateTransactionUseCase(
		tr,
		op,
		ac)

	_, err := handler.Execute(transactionDto)
	c.NotNil(err)
}

func (c CreateTransactionUseCaseSuite) TestReturnErrorIfAccountNotFound() {
	tr, op, ac := c.BaseMocks()
	transactionDto := dto.TransactionDto{
		AccountID:       1,
		OperationTypeID: 1,
		Amount:          100,
	}
	op.On("GetOperation", transactionDto.AccountID).Return(queries.OperationType{}, nil)
	tr.On("CreateTransaction", transactionDto).Times(0)
	ac.On("FindAccount", transactionDto.AccountID).Times(1).Return(queries.Account{}, sql.ErrNoRows)
	handler := NewCreateTransactionUseCase(
		tr,
		op,
		ac)

	_, err := handler.Execute(transactionDto)
	c.NotNil(err)
}

func (c CreateTransactionUseCaseSuite) TestItCreatesTransaction() {
	tr, op, ac := c.BaseMocks()
	transactionDto := dto.TransactionDto{
		AccountID:       1,
		OperationTypeID: 1,
		Amount:          100,
	}
	op.On("GetOperation", transactionDto.AccountID).Return(queries.OperationType{Type: "NEGATIVE"}, nil)
	modifiedDto := transactionDto
	modifiedDto.Amount = -100
	tr.On("CreateTransaction", modifiedDto).Times(1).Return(queries.Transaction{
		ID:              1,
		AccountID:       1,
		OperationTypeID: 1,
		Amount:          -100,
		EventDate:       time.Now(),
	}, nil)
	ac.On("FindAccount", transactionDto.AccountID).Times(1).Return(queries.Account{}, nil)
	handler := NewCreateTransactionUseCase(
		tr,
		op,
		ac)

	transaction, err := handler.Execute(transactionDto)
	c.Nil(err)
	c.Equal(float32(-100), transaction.Amount)
}

func TestCreateTransactionUseCaseSuite(t *testing.T) {
	suite.Run(t, new(CreateTransactionUseCaseSuite))
}
