package domain

import (
	"github.com/Khanabeev/banking/dto"
	errors2 "github.com/Khanabeev/banking/errors"
	"github.com/Khanabeev/banking/logger"
	"github.com/jmoiron/sqlx"
	"strconv"
)

type TransactionRepositoryDb struct {
	dbClient *sqlx.DB
}

func (r TransactionRepositoryDb) Save(t Transaction) (*Transaction, *errors2.AppError) {

	tx, err := r.dbClient.Begin()
	if err != nil {
		logger.Error("Error during starting a new transaction for bank account: " + err.Error())
		return nil, errors2.UnexpectedError("Unexpected database error")
	}

	stmtInsert := `INSERT INTO transactions(account_id, amount, transaction_type, transaction_date)
					VALUES(?,?,?,?)`

	result, err := tx.Exec(stmtInsert, &t.AccountId, &t.Amount, &t.TransactionType, &t.TransactionDate)
	if err != nil {
		logger.Error("Error during inserting data to transaction table: " + err.Error())
		return nil, errors2.UnexpectedError("Unexpected database error")
	}

	var stmt string
	if t.TransactionType == dto.WITHDRAWAL {
		stmt = `UPDATE accounts SET amount = amount - ? WHERE account_id = ?`
	} else {
		stmt = `UPDATE accounts SET amount = amount + ? WHERE account_id = ?`
	}
	_, err = tx.Exec(stmt, t.Amount, t.AccountId)
	if err != nil {
		_ = tx.Rollback()

		logger.Error("Error while UPDATING account amount: " + err.Error())
		return nil, errors2.UnexpectedError("Unexpected error from database")
	}

	err = tx.Commit()
	if err != nil {
		logger.Error("Error while COMMITING new transaction: " + err.Error())
		return nil, errors2.UnexpectedError("Unexpected error from database")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error during getting last ID from transaction table: " + err.Error())
		return nil, errors2.UnexpectedError("Unexpected database error")
	}

	t.TransactionId = strconv.FormatInt(id, 10)

	return &t, nil
}

func NewTransactionRepositoryDb(dbClient *sqlx.DB) TransactionRepositoryDb {
	return TransactionRepositoryDb{
		dbClient: dbClient,
	}
}
