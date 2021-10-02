package ports

import "github.com/karankumarshreds/GoApp/internal/core"

type CustomerPort interface {
	GetAllCustomers() ([]core.Customer, error)
	GetCustomer(id string) (*core.Customer, error)
}

