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
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		log.Println("ERROR: Unable to get single customer", err)
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	json.NewEncoder(w).Encode(customer)
}