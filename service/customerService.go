package service

import (
	"github.com/Khanabeev/banking/domain"
	errors2 "github.com/Khanabeev/banking/errors"
)

type CustomerService interface {
	GetAllCustomers(string) ([]domain.Customer, *errors2.AppError)
	GetCustomer(string) (*domain.Customer, *errors2.AppError)
}

type DefaultCustomerService struct {
	repository domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers(status string) ([]domain.Customer, *errors2.AppError) {

	switch status {
	case "active":
		status = "1"
		break
	case "inactive":
		status = "0"
		break
	default:
		status = ""
	}

	return s.repository.FindAll(status)
}

func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *errors2.AppError) {
	return s.repository.ById(id)
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
