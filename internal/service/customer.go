package service

import (
	"github.com/karankumarshreds/GoApp/internal/core"
	"github.com/karankumarshreds/GoApp/internal/dto"
	"github.com/karankumarshreds/GoApp/internal/pkg/custom_errors"
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

func (cs CustomerService) GetCustomer(id string) (*dto.CustomerResponse, *custom_errors.CustomError) {
	c, err := cs.repo.FindById(id)
	if err != nil {
		return nil, err
	}
	response := dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		ZipCode:     c.ZipCode,
		DateOfBirth: c.DateOfBirth,
		Status:      false,
	}
	return &response, nil
}