package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"tiagofv.com/transactions/core/domain/dto"
	"tiagofv.com/transactions/core/domain/use_cases"
)

// CreateAccount godoc
// @Summary      Creates an account
// @Description  Creates an account
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param	data body dto.AccountDto true "The data"
// @Success      200  {object}  queries.Account
// @Failure      400  {object}  BaseError
// @Failure      422  {object}  BaseError
// @Failure      500  {object}  BaseError
// @Router       /accounts [post]
func (b BaseController) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var data dto.AccountDto

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&data)
	if err != nil {
		b.WriteError(w, err, http.StatusBadRequest)
		return
	}

	if data.Document == "" || len(data.Document) != 11 || !IsCPF(data.Document) {
		b.WriteError(w, errors.New("The field document must be a valid CPF."), http.StatusUnprocessableEntity)
		return
	}

	account, err := b.CreateAccountUseCase.Execute(data)
	if errors.Is(err, use_cases.ErrAccountExists) {
		b.WriteError(w, err, http.StatusUnprocessableEntity)
		return
	}
	if err != nil {
		b.WriteError(w, err, http.StatusInternalServerError)
		return
	}

	encoder := json.NewEncoder(w)
	err = encoder.Encode(account)
	if err != nil {
		b.WriteError(w, err, http.StatusInternalServerError)
		return
	}
}

// GetAccount
// @Summary      Get account by id
// @Description  Get account by id
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param	id path int true "The id"
// @Success      200  {object}  queries.Account
// @Failure      400  {object}  BaseError
// @Failure      404  {object}  BaseError
// @Failure      500  {object}  BaseError
// @Router       /accounts/{id} [get]
func (b BaseController) GetAccount(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		b.WriteError(w, err, http.StatusBadRequest)
		return
	}
	account, err := b.GetAccountUseCase.Execute(int64(id))

	if errors.Is(err, sql.ErrNoRows) {
		b.WriteError(w, errors.New("Account not found"), http.StatusNotFound)
		return
	}
	if err != nil {
		b.WriteError(w, err, http.StatusInternalServerError)
		return
	}
	encoder := json.NewEncoder(w)
	err = encoder.Encode(account)
	if err != nil {
		b.WriteError(w, err, http.StatusInternalServerError)
		return
	}
}
