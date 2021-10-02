package db

import (
	"github.com/karankumarshreds/GoApp/internal/core"
)

type CustomerRepositoryStub struct {
	customers []core.Customer
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []core.Customer{
		{ Id: "123", Name: "Karan", City: "Delhi", ZipCode: "00000", DateOfBirth: "09/09/09", Status: true},
		{ Id: "124", Name: "Rahul", City: "Delhi", ZipCode: "00001", DateOfBirth: "08/08/08", Status: true},
	}
	return CustomerRepositoryStub{customers}
}

func (c CustomerRepositoryStub) FindAll() ([]core.Customer, error) {
	return c.customers, nil
}

