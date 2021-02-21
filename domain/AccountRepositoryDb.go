package domain

import (
	"database/sql"
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

func (d AccountRepositoryDb) FindBy(accountId string) (*Account, *errors2.AppError) {

	stmt := `SELECT account_id, customer_id, opening_date, account_type, status, amount
				FROM accounts 
				WHERE account_id = ?`

	row := d.client.QueryRow(stmt, accountId)
	if row.Err() != nil {
		logger.Error("Error while getting account: " + row.Err().Error())
		return nil, errors2.UnexpectedError("Unexpected error from database")
	}
	var a Account

	switch err := row.Scan(
		&a.AccountId,
		&a.CustomerId,
		&a.OpeningDate,
		&a.AccountType,
		&a.Status,
		&a.Amount); err {
	case sql.ErrNoRows:
		return nil, errors2.NewNotFoundError("Bank account doesn't exist")
	case nil:
		return &a, nil
	default:
		logger.Error("Error while getting account: " + row.Err().Error())
		return nil, errors2.UnexpectedError("Unexpected error from database")
	}
}

func (d AccountRepositoryDb) UpdateAmount(id string, amount float64) *errors2.AppError {
	stmt := `UPDATE accounts 
				SET amount = ?
				WHERE account_id = ?`
	_, err := d.client.Exec(stmt, amount, id)
	if err != nil {
		logger.Error("Error while UPDATING account amount: " + err.Error())
		return errors2.UnexpectedError("Unexpected error from database")
	}

	return nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{dbClient}
}
