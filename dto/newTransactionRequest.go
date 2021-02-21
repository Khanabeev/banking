package dto

import errors2 "github.com/Khanabeev/banking/errors"

const WITHDRAWAL = "withdrawal"
const DEPOSIT = "deposit"

type NewTransactionRequest struct {
	AccountId       string  `json:"account_id"`
	CustomerId		string	`json:"customer_id"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
	TransactionDate string  `json:"transaction_date"`
}

func (r NewTransactionRequest) Validate() *errors2.AppError {
	if r.TransactionType != WITHDRAWAL && r.TransactionType != DEPOSIT {
		return errors2.NewValidationError("Transaction type can be only withdrawal or deposit")
	}

	if r.Amount < 0 {
		return errors2.NewValidationError("Amount cannot be negative")
	}

	return nil
}
