package use_cases

import (
	"database/sql"
	"errors"
	"tiagofv.com/transactions/core/domain/dto"
	"tiagofv.com/transactions/core/domain/repositories"
	"tiagofv.com/transactions/infra/database/queries"
)

type CreateAccountsUseCase struct {
	accountsRepo repositories.AccountsInterface
}

var ErrAccountExists = errors.New("account number with this document is already associated")

func NewCreateAccountsUseCase(accountsRepo repositories.AccountsInterface) *CreateAccountsUseCase {
	return &CreateAccountsUseCase{
		accountsRepo: accountsRepo,
	}
}

func (c CreateAccountsUseCase) Execute(dto dto.AccountDto) (queries.Account, error) {
	account, err := c.accountsRepo.FindAccountByDocument(dto.Document)
	if (!errors.Is(err, sql.ErrNoRows) && err != nil) || account.Document != "" {
		return account, ErrAccountExists
	}
	return c.accountsRepo.CreateAccount(dto)
}
