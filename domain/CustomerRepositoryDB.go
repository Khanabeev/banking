package domain

import (
	"database/sql"
	errors2 "github.com/Khanabeev/banking/errors"
	"github.com/Khanabeev/banking/logger"
	"github.com/jmoiron/sqlx"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDB struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDB) FindAll(status string) ([]Customer, *errors2.AppError) {
	customers := make([]Customer, 0)
	findAllSql := `SELECT 
					   customer_id, 
					   name, 
					   city, 
					   zipcode, 
					   date_of_birth, 
					   status 
					FROM customers`

	if status != "" {
		findAllSql += ` WHERE status = ` + status
	}
	err := d.client.Select(&customers,findAllSql)

	if err != nil {
		logger.Error("Error while querying customer table " + err.Error())
		return nil, errors2.NewNotFoundError("Unexpected database error")
	}

	return customers, nil
}

func (d CustomerRepositoryDB) ById(id string) (*Customer, *errors2.AppError) {
	customerSql := `SELECT customer_id, name, city, zipcode, date_of_birth, status from customers WHERE customer_id = ?`

	var c Customer
	err := d.client.Get(&c, customerSql, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors2.NewNotFoundError("Customer not found")
		} else {
			logger.Error("Error while scanning customer " + err.Error())
			return nil, errors2.UnexpectedError("unexpected database error")
		}
	}
	return &c, nil
}

func NewCustomerRepositoryDB() CustomerRepositoryDB {
	client, err := sqlx.Open("mysql", "root:root@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(2)

	return CustomerRepositoryDB{
		client,
	}
}
