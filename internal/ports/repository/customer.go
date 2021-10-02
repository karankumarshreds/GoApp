package ports

import "github.com/karankumarshreds/GoApp/internal/core"

type CustomerRepositoryPort interface {
	FindAll() ([]core.Customer, error)
	// returning a pointer because we might want to return nil in case the user is not found
	// we can only return nil with pointers, it would not work for normal struct type
	FindById(id string) (*core.Customer, error)
}
