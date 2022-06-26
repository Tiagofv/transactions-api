package use_cases

import (
	"github.com/stretchr/testify/suite"
	"testing"
	"tiagofv.com/transactions/core/domain/dto"
	"tiagofv.com/transactions/infra/database/queries"
	"tiagofv.com/transactions/tests/mocks"
)

type CreateAccountTestSuite struct {
	suite.Suite
}

func (c CreateAccountTestSuite) TestIfDuplicatesReturnError() {
	randomCpf := "18965321912"

	accountsInterface := new(mocks.AccountsInterface)
	accountsInterface.On("FindAccountByDocument", randomCpf).Return(queries.Account{
		Document: "123",
	}, nil)
	handler := NewCreateAccountsUseCase(accountsInterface)
	_, err := handler.Execute(dto.AccountDto{
		Document: randomCpf,
	})
	c.NotNil(err, "Error should not be nil!")

}

func TestCreateAccountTestSuite(t *testing.T) {
	suite.Run(t, new(CreateAccountTestSuite))
}
