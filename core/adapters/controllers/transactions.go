package controllers

import (
	"encoding/json"
	"net/http"
	"tiagofv.com/transactions/core/domain/dto"
	"time"
)

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
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	data.EventDate = time.Now()
	if err != nil {
		b.WriteError(w, err, http.StatusBadRequest)
		return
	}
	response, err := b.CreateTransactionUseCase.Execute(data)
	if err != nil {
		b.WriteError(w, err, http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
