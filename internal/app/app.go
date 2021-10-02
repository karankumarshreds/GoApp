package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/karankumarshreds/GoApp/internal/framework/db"
	"github.com/karankumarshreds/GoApp/internal/framework/http/handlers"
	"github.com/karankumarshreds/GoApp/internal/service"
	"log"
	"net/http"
)

func Start() {
	router := mux.NewRouter()

	// wiring
	connString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", "postgres", "password", "banking")
	repo := db.NewCustomerRepository(connString)
	cs := service.NewCustomerService(repo)
	h := handlers.NewCustomerHandlers(cs)

	// routes
	router.HandleFunc("/customers", h.GetAllCustomers)

	log.Fatal(http.ListenAndServe(":8000", router))
}
