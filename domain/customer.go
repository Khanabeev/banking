package domain

import errors2 "github.com/Khanabeev/banking/errors"

type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string
	Status      string
}

type CustomerRepository interface {
	FindAll(string) ([]Customer, *errors2.AppError)
	ById(string) (*Customer, *errors2.AppError)
}

