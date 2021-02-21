package domain

import (
	"github.com/Khanabeev/banking/dto"
	errors2 "github.com/Khanabeev/banking/errors"
)

type Transaction struct {
	TransactionId   string
	AccountId       string
	Amount          float64
	TransactionType string
	TransactionDate string
}

func (t Transaction) TransactionToDtoResponse(a *Account) dto.NewTransactionResponse {
	return dto.NewTransactionResponse{
		TransactionId:   t.TransactionId,
		AccountId:       a.AccountId,
		NewBalance:      a.Amount,
		TransactionType: t.TransactionType,
		TransactionDate: t.TransactionDate,
	}
}

type TransactionRepository interface {
	Save(transaction Transaction) (*Transaction, *errors2.AppError)
}
