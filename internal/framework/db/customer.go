package db

type Customer struct {
	Id string
	Name string
	City string
	ZipCode string
	DateOfBirth string
	Status bool
}

type CustomerRepositoryStub struct {
	customers []Customer
}

func (c CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return c.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{ Id: "123", Name: "Karan", City: "Delhi", ZipCode: "00000", DateOfBirth: "09/09/09", Status: true},
		{ Id: "124", Name: "Rahul", City: "Delhi", ZipCode: "00001", DateOfBirth: "08/08/08", Status: true},
	}
	return CustomerRepositoryStub{customers}
}