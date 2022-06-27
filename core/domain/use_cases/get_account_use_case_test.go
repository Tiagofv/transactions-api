package use_cases

import (
	"database/sql"
	"github.com/stretchr/testify/suite"
	"testing"
	"tiagofv.com/transactions/infra/database/queries"
	"tiagofv.com/transactions/tests/mocks"
)

type GetAccountUseCaseSuite struct {
	suite.Suite
	id int64
}

func (g GetAccountUseCaseSuite) SetupTest() {
	g.id = 123
}

func (g GetAccountUseCaseSuite) TestAccountNotFoundReturnsError() {
	accountsRepo := new(mocks.AccountsInterface)
	accountsRepo.On("FindAccount", g.id).Return(queries.Account{}, sql.ErrNoRows)
	handler := GetAccountUseCase{accountsRepo: accountsRepo}
	_, err := handler.Execute(g.id)
	g.NotNil(err, "Error Should not be nil")
	g.ErrorIs(err, sql.ErrNoRows)
}

func (g GetAccountUseCaseSuite) TestItFindsAccounts() {
	accountsRepo := new(mocks.AccountsInterface)
	accountsRepo.On("FindAccount", g.id).Return(queries.Account{
		ID:       g.id,
		Document: "12345678911",
	}, nil)
	handler := GetAccountUseCase{accountsRepo: accountsRepo}
	account, err := handler.Execute(g.id)
	g.Nil(err)
	g.Equal(g.id, account.ID)
}

func TestGetAccountUseCaseSuite(t *testing.T) {
	suite.Run(t, new(GetAccountUseCaseSuite))
}
