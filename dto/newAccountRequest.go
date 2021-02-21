package dto

import (
	errors2 "github.com/Khanabeev/banking/errors"
	"strings"
)

type NewAccountRequest struct {
	CustomerId string `json:"customer_id"`
	AccountType string `json:"account_type"`
	Amount float64 `json:"amount"`
}

func (r NewAccountRequest) Validate() *errors2.AppError {
	if r.Amount < 5000 {
		return errors2.NewValidationError("To open new account you need to deposit at least 5000 ")
	}

	if strings.ToLower(r.AccountType) != "saving" &&  strings.ToLower(r.AccountType) != "checking" {
		return errors2.NewValidationError("Account type should be checking or saving")
	}

	return nil
}


