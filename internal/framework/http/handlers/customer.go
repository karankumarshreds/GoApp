package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/karankumarshreds/GoApp/internal/pkg/custom_errors"
	"github.com/karankumarshreds/GoApp/internal/pkg/logger"
	"github.com/karankumarshreds/GoApp/internal/ports/service"
	"net/http"
	"strconv"
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
		logger.ErrorLog("Unable to get all customers", err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func (ch *CustomerHandlers) GetSingleCustomer(w http.ResponseWriter, r*http.Request) {
	vars := mux.Vars(r)
	customerId := vars["id"]
	_, e := strconv.Atoi(customerId)
	if customerId == "" || e != nil{
		logger.ErrorLog("Invalid customer ID", e)
		custom_errors.NewBadRequestError("Invalid customer ID").ErrorJsonResponse(w)
		return
	}
 	customer, err := ch.service.GetCustomer(customerId)
	if err != nil {
		logger.ErrorLog("Unable to get single customer" + err.Message, nil)
		err.ErrorJsonResponse(w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}