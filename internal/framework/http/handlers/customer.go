package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/karankumarshreds/GoApp/internal/ports/service"
	"log"
	"net/http"
)

type CustomerHandlers struct {
	service ports.CustomerPort
}

func NewCustomerHandlers(service ports.CustomerPort) CustomerHandlers{
	return CustomerHandlers{ service }
}

func (ch *CustomerHandlers) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := ch.service.GetAllCustomers()
	if err != nil {
		log.Println("ERROR: Unable to get all customers", err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func (ch *CustomerHandlers) GetSingleCustomer(w http.ResponseWriter, r*http.Request) {
	vars := mux.Vars(r)
	customerId := vars["id"]
	if customerId == "" {
		log.Println("Invalid customer ID")
	}
 	customer, err := ch.service.GetCustomer(customerId)
	if err != nil {
		log.Println("ERROR: Unable to get single customer", err)
		err.ErrorJsonResponse(w, err.Message)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}