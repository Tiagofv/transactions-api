package controllers

import (
	"encoding/json"
	"net/http"
	"tiagofv.com/transactions/core/domain/use_cases"
)

type BaseController struct {
	CreateTransactionUseCase *use_cases.CreateTransactionUseCase
	CreateAccountUseCase     *use_cases.CreateAccountsUseCase
	GetAccountUseCase        *use_cases.GetAccountUseCase
}

func NewBaseController(createTransactionUseCase *use_cases.CreateTransactionUseCase, createAccountUseCase *use_cases.CreateAccountsUseCase, getAccountUseCase *use_cases.GetAccountUseCase) *BaseController {
	return &BaseController{CreateTransactionUseCase: createTransactionUseCase, CreateAccountUseCase: createAccountUseCase, GetAccountUseCase: getAccountUseCase}
}

type BaseError struct {
	Message string `json:"message"`
}

func (b BaseError) Error() string {
	return b.Message
}

func (b BaseController) WriteError(w http.ResponseWriter, err error, statusCode int) error {
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(BaseError{Message: err.Error()})
}
