package service

import (
	"github.com/karankumarshreds/GoApp/internal/ports/repository"
	"github.com/karankumarshreds/GoApp/internal/framework/db"
)

type CustomerService struct {
	// dependency injection of the driven adapter
	repo ports.CustomerRepositoryPort
}

// NewCustomerService instantiates a new customer service
func NewCustomerService( repo ports.CustomerRepositoryPort ) CustomerService {
	return CustomerService{repo}
}

func (cs CustomerService) GetAllCustomers() ([]db.Customer, error){
	return cs.repo.FindAll()
}