package domain

import (
	errors2 "github.com/Khanabeev/banking/errors"
	"github.com/Khanabeev/banking/logger"
	"github.com/jmoiron/sqlx"
	"strconv"
)

type TransactionRepositoryDb struct {
	dbClient *sqlx.DB
}

func (r TransactionRepositoryDb) Save(t Transaction) (*Transaction, *errors2.AppError) {

	stmtInsert := `INSERT INTO transactions(account_id, amount, transaction_type, transaction_date)
					VALUES(?,?,?,?)`

	result, err := r.dbClient.Exec(stmtInsert, &t.AccountId, &t.Amount, &t.TransactionType, &t.TransactionDate)
	if err != nil {
		logger.Error("Error during inserting data to transaction table: " + err.Error())
		return nil, errors2.UnexpectedError("Unexpected database error")
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
