package use_cases

import (
	"tiagofv.com/transactions/core/domain/repositories"
	"tiagofv.com/transactions/core/presenters"
)

type GetAccountUseCase struct {
	accountsRepo repositories.AccountsInterface
}

func NewGetAccountUseCase(accountsRepo repositories.AccountsInterface) *GetAccountUseCase {
	return &GetAccountUseCase{accountsRepo: accountsRepo}
}

func (g GetAccountUseCase) Execute(id int64) (presenters.AccountPresenter, error) {
	account, err := g.accountsRepo.FindAccount(id)
	return presenters.AccountPresenter{
		ID:       account.ID,
		Document: account.Document,
	}, err
}
