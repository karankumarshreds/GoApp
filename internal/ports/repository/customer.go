package ports

import "github.com/karankumarshreds/GoApp/internal/framework/db"

type CustomerRepositoryPort interface {
	FindAll() ([]db.Customer, error)
}
