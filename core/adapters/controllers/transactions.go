package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"tiagofv.com/transactions/core/domain/dto"
	"tiagofv.com/transactions/core/domain/use_cases"
)

// CreateTransaction godoc
// @Summary      Creates a transaction
// @Description  Creates a transaction
// @Tags         transaction
// @Accept       json
// @Produce      json
// @Param	data body dto.TransactionDto true "The data"
// @Success      200  {object}  presenters.TransactionPresenter
// @Failure      400  {object}  BaseError
// @Failure      404  {object}  BaseError
// @Failure      500  {object}  BaseError
// @Router       /transactions [post]
func (b BaseController) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var data dto.TransactionDto
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		b.WriteError(w, err, http.StatusBadRequest)
		return
	}

	if data.Amount <= 0 {
		b.WriteError(w, errors.New("The field amount must be bigger than zero."), http.StatusUnprocessableEntity)
		return
	}
	response, err := b.CreateTransactionUseCase.Execute(data)
	if errors.Is(err, use_cases.ErrInvalidOperation) || errors.Is(err, use_cases.ErrInvalidAccountId) {
		b.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}
	if err != nil {
		b.WriteError(w, err, http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (b BaseController) GetTransaction(w http.ResponseWriter, r *http.Request) {

}
