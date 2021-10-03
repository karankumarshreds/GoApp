package ports

import (
	"github.com/karankumarshreds/GoApp/internal/core"
	"github.com/karankumarshreds/GoApp/internal/pkg/custom_errors"
)

type CustomerPort interface {
	GetAllCustomers() ([]core.Customer, error)
	GetCustomer(id string) (*core.Customer, *custom_errors.CustomError)
}

