package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/karankumarshreds/GoApp/internal/core"
	"github.com/karankumarshreds/GoApp/internal/pkg/custom_errors"
	"github.com/karankumarshreds/GoApp/internal/pkg/logger"
	_ "github.com/lib/pq"
)

type AccountRepositoryDb struct {
	db *sqlx.DB
}

func NewAccountRepository(dbClient *sqlx.DB) *AccountRepositoryDb {
	return &AccountRepositoryDb{dbClient}
}

func (d AccountRepositoryDb) Save(a core.Account) (*core.Account, *custom_errors.CustomError) {
	q := `
		INSERT INTO accounts(customer_id, opening_date, account_type, amount, status)
		VALUES  ($1, $2, $3, $4, $5) RETURNING account_id;
	`
	id := 0
	err := d.db.QueryRow(q,a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status).Scan(&id)
	if err != nil {
		logger.ErrorLog("Unable to create account", err)
		return nil, custom_errors.NewInternalServerError("Unable to create account at the moment")
	}
	a.AccountId = fmt.Sprintf("%v", id)
	return &a, nil
}

