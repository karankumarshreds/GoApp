package dto

type AccountCreateDto struct{
	CustomerId string `json:customerId`
	AccountType string `json:accountType`
	Amount float64 `json:amount`
}


