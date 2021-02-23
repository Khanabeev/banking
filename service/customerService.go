package service

import (
	"github.com/Khanabeev/banking/domain"
	"github.com/Khanabeev/banking/dto"
	errors2 "github.com/Khanabeev/banking/errors"
)
//go:generate mockgen -destination=../mocks/service/mockCustomerService.go -package=service github.com/Khanabeev/banking/service CustomerService
type CustomerService interface {
	GetAllCustomers(string) ([]dto.CustomerResponse, *errors2.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *errors2.AppError)
}

type DefaultCustomerService struct {
	repository domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers(status string) ([]dto.CustomerResponse, *errors2.AppError) {

	switch status {
	case "active":
		status = "1"
	case "inactive":
		status = "0"
	default:
		status = ""
	}

	c, err := s.repository.FindAll(status)
	if err != nil {
		return nil, err
	}

	var arr []dto.CustomerResponse

	for _, cr := range c {
		res := cr.ToDto()
		arr = append(arr, res)
	}

	return arr, nil
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errors2.AppError) {
	c, err := s.repository.ById(id)
	if err != nil {
		return nil, err
	}

	response := c.ToDto()

	return &response, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
