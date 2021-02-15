package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer {
		{"1001", "John", "Moscow", "23123","1978-02-06", "1"},
		{"1002", "Jane", "Krasnodar", "23123","2000-12-23", "2"},
		{"1003", "Rob", "Minsk", "23123","1999-11-19", "3"},
	}

	return CustomerRepositoryStub{
		customers: customers,
	}
}
