package domain

import (
	"github.com/Khanabeev/banking/dto"
	errors2 "github.com/Khanabeev/banking/errors"
)

type Account struct {
	AccountId               string
	CustomerId              string
	OpeningDate             string
	AccountType             string
	Amount                  float64
	Status                  string
}

type AccountRepository interface {
	Save(account Account) (*Account, *errors2.AppError)
	FindBy(id string) (*Account, *errors2.AppError)
	UpdateAmount(id string, amount float64) *errors2.AppError
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{
		AccountId: a.AccountId,
	}
}