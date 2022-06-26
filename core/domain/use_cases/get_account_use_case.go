package use_cases

import (
	"tiagofv.com/transactions/core/domain/repositories"
	"tiagofv.com/transactions/infra/database/queries"
)

type GetAccountUseCase struct {
	accountsRepo repositories.AccountsInterface
}

func NewGetAccountUseCase(accountsRepo repositories.AccountsInterface) *GetAccountUseCase {
	return &GetAccountUseCase{accountsRepo: accountsRepo}
}

func (g GetAccountUseCase) Execute(id int64) (queries.Account, error) {
	return g.accountsRepo.FindAccount(id)
}
