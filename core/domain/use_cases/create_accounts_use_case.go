package use_cases

import (
	"database/sql"
	"errors"
	"tiagofv.com/transactions/core/domain/dto"
	"tiagofv.com/transactions/core/domain/repositories"
	"tiagofv.com/transactions/core/presenters"
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

func (c CreateAccountsUseCase) Execute(dto dto.AccountDto) (presenters.AccountPresenter, error) {
	account, err := c.accountsRepo.FindAccountByDocument(dto.Document)
	if (!errors.Is(err, sql.ErrNoRows) && err != nil) || account.Document != "" {
		return presenters.AccountPresenter{}, ErrAccountExists
	}
	createAccount, err := c.accountsRepo.CreateAccount(dto)
	return presenters.AccountPresenter{
		ID:       createAccount.ID,
		Document: createAccount.Document,
	}, err
}
