package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/karankumarshreds/GoApp/internal/ports/service"
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