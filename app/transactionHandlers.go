package app

import (
	"encoding/json"
	"github.com/Khanabeev/banking/dto"
	"github.com/Khanabeev/banking/service"
	"github.com/gorilla/mux"
	"net/http"
)

type TransactionHandler struct {
	service service.TransactionService
}

func (h TransactionHandler) NewTransaction(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	accountId := vars["account_id"]
	customerId := vars["customer_id"]

	var request dto.NewTransactionRequest

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.AccountId = accountId
		request.CustomerId = customerId

		response, appError := h.service.CreateNewTransaction(request)
		if appError != nil {
			writeResponse(w, appError.Code, appError.AsMessage())
		} else {
			writeResponse(w, http.StatusOK, response)
		}
	}
}
