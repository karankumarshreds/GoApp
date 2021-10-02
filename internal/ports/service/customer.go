package ports

import "github.com/karankumarshreds/GoApp/internal/framework/db"

type CustomerPort interface {
	GetAllCustomers() ([]db.Customer, error)
}

