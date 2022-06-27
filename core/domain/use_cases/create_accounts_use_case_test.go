package use_cases

import (
	"database/sql"
	"github.com/stretchr/testify/suite"
	"testing"
	"tiagofv.com/transactions/core/domain/dto"
	"tiagofv.com/transactions/infra/database/queries"
	"tiagofv.com/transactions/tests/mocks"
)

type CreateAccountTestSuite struct {
	suite.Suite
	randomCpf string
}

func (c CreateAccountTestSuite) SetupTest() {
	c.randomCpf = "18965321912"
}

func (c CreateAccountTestSuite) TestIfDuplicatesReturnError() {

	accountsInterface := new(mocks.AccountsInterface)
	accountsInterface.On("FindAccountByDocument", c.randomCpf).Return(queries.Account{
		Document: "123",
	}, nil)
	handler := NewCreateAccountsUseCase(accountsInterface)
	_, err := handler.Execute(dto.AccountDto{
		Document: c.randomCpf,
	})
	c.NotNil(err, "Error should not be nil!")
}

func (c CreateAccountTestSuite) TestItDoesntReturnErrorIfErrNoRows() {
	accountsInterface := new(mocks.AccountsInterface)
	accountsInterface.On("FindAccountByDocument", c.randomCpf).Return(queries.Account{}, sql.ErrNoRows)
	accountsInterface.On("CreateAccount", dto.AccountDto{Document: c.randomCpf}).Return(queries.Account{Document: c.randomCpf}, nil)
	handler := NewCreateAccountsUseCase(accountsInterface)
	_, err := handler.Execute(dto.AccountDto{
		Document: c.randomCpf,
	})
	c.Nil(err, "Error should be nil!")
}

func (c CreateAccountTestSuite) TestItCreatesAccount() {
	accountsInterface := new(mocks.AccountsInterface)
	accountsInterface.On("FindAccountByDocument", c.randomCpf).Return(queries.Account{}, nil)
	accountsInterface.On("CreateAccount", dto.AccountDto{Document: c.randomCpf}).Return(queries.Account{Document: c.randomCpf}, nil)
	handler := NewCreateAccountsUseCase(accountsInterface)
	account, err := handler.Execute(dto.AccountDto{
		Document: c.randomCpf,
	})
	c.Nil(err, "Error should be nil")
	c.Equal(c.randomCpf, account.Document, "Cpf needs to be equal, found different!")
}

func TestCreateAccountTestSuite(t *testing.T) {
	suite.Run(t, new(CreateAccountTestSuite))
}
