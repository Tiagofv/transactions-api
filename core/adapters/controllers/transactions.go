package controllers

import (
	"encoding/json"
	"net/http"
	"tiagofv.com/transactions/core/domain/dto"
	"tiagofv.com/transactions/core/domain/use_cases"
)

type BaseController struct {
	CreateTransactionUseCase *use_cases.CreateTransactionUseCase
}

func NewBaseController(createTransactionUseCase *use_cases.CreateTransactionUseCase) *BaseController {
	return &BaseController{CreateTransactionUseCase: createTransactionUseCase}
}

func (b BaseController) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var data dto.TransactionDto
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response, err := b.CreateTransactionUseCase.Execute(data)

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
