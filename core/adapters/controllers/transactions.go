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

type BaseError struct {
	Message string
}

func (b BaseError) Error() string {
	return b.Message
}

func NewBaseController(createTransactionUseCase *use_cases.CreateTransactionUseCase) *BaseController {
	return &BaseController{CreateTransactionUseCase: createTransactionUseCase}
}

// CreateTransaction godoc
// @Summary      Creates a transaction
// @Description  Creates a transaction
// @Tags         transaction
// @Accept       json
// @Produce      json
// @Param	data body dto.TransactionDto true "The data"
// @Success      200  {object}  queries.Transaction
// @Failure      400  {object}  BaseError
// @Failure      404  {object}  BaseError
// @Failure      500  {object}  BaseError
// @Router       /transactions [post]
func (b BaseController) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var data dto.TransactionDto
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		_ = json.NewEncoder(w).Encode(BaseError{Message: err.Error()})
		return
	}
	response, err := b.CreateTransactionUseCase.Execute(data)

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
