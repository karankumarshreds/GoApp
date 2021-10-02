package service

import (
	"github.com/karankumarshreds/GoApp/internal/core"
	"github.com/karankumarshreds/GoApp/internal/ports/repository"
)

type CustomerService struct {
	// dependency injection of the driven adapter
	repo ports.CustomerRepositoryPort
}

// NewCustomerService instantiates a new customer service
func NewCustomerService( repo ports.CustomerRepositoryPort ) CustomerService {
	return CustomerService{repo}
}

func (cs CustomerService) GetAllCustomers() ([]core.Customer, error){
	return cs.repo.FindAll()
}

func (cs CustomerService) GetCustomer(id string) (*core.Customer, error) {
	return cs.repo.FindById(id)
}