package dto

type CustomerResponse struct {
	Id string `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
	ZipCode string `json:"zipCode"`
	DateOfBirth string `json:"dateOfBirth"`
	Status bool `json:"status"`
}