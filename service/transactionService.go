package service

import (
	"github.com/Khanabeev/banking/domain"
	"github.com/Khanabeev/banking/dto"
	errors2 "github.com/Khanabeev/banking/errors"
	"time"
)

type TransactionService interface {
	CreateNewTransaction(request dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errors2.AppError)
}

type DefaultTransactionService struct {
	transactionRepo domain.TransactionRepository
	accountRepo     domain.AccountRepository
}

func (s DefaultTransactionService) CreateNewTransaction(r dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errors2.AppError) {

	err := r.Validate()
	if err != nil {
		return nil, err
	}

	account, appError := s.accountRepo.FindBy(r.AccountId, r.CustomerId)
	if appError != nil {
		return nil, appError
	}

	var newAmount float64

	if r.TransactionType == dto.WITHDRAWAL {
		// Check if enough money on account
		newAmount = account.Amount - r.Amount
		if newAmount < 0 {
			return nil, errors2.NewValidationError("Not enough money in account")
		}
	}

	if r.TransactionType == dto.DEPOSIT {
		newAmount = account.Amount + r.Amount
	}

	err = s.accountRepo.UpdateAmount(r.AccountId, newAmount)
	if err != nil {
		return nil, err
	}

	account.Amount = newAmount

	t := domain.Transaction{
		TransactionId:   "",
		AccountId:       r.AccountId,
		Amount:          r.Amount,
		TransactionType: r.TransactionType,
		TransactionDate: time.Now().Format("2006-01-02 15:04:05"),
	}

	newTransaction, err := s.transactionRepo.Save(t)

	if err != nil {
		return nil, err
	}

	result := newTransaction.TransactionToDtoResponse(account)

	return &result, nil
}

func NewTransactionService(repository domain.TransactionRepository, account domain.AccountRepository) DefaultTransactionService {
	return DefaultTransactionService{
		transactionRepo: repository,
		accountRepo:     account,
	}
}
