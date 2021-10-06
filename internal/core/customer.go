package core

type Customer struct {
	Id string `json:"id" db:"customer_id"`
	Name string `json:"name"`
	City string `json:"city"`
	ZipCode string `json:"zipCode"`
	DateOfBirth string `json:"dateOfBirth" db:"date_of_birth"`
	Status bool `json:"status"`
}