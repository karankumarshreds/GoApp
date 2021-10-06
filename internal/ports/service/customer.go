package ports

import (
	"github.com/karankumarshreds/GoApp/internal/core"
	"github.com/karankumarshreds/GoApp/internal/dto"
	"github.com/karankumarshreds/GoApp/internal/pkg/custom_errors"
)

type CustomerPort interface {
	GetAllCustomers() ([]core.Customer, error)
	GetCustomer(id string) (*dto.CustomerResponse, *custom_errors.CustomError)
}

