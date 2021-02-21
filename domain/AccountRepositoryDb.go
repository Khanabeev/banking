package domain

import (
	errors2 "github.com/Khanabeev/banking/errors"
	"github.com/Khanabeev/banking/logger"
	"github.com/jmoiron/sqlx"
	"strconv"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (d AccountRepositoryDb) Save(a Account) (*Account, *errors2.AppError) {
	sqlInsert := `INSERT INTO accounts(customer_id, opening_date, account_type, amount, status)
				VALUES(?,?,?,?,?)`
	result, err := d.client.Exec(sqlInsert, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("Error while creating new account:" + err.Error())
		return nil, errors2.UnexpectedError("Unexpected error from database")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last insert id for new account:" + err.Error())
		return nil, errors2.UnexpectedError("Unexpected error from database")
	}
	a.AccountId = strconv.FormatInt(id, 10)

	return &a, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{dbClient}
}
