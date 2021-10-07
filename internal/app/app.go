package app

import (
	"database/sql"
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
	client, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatalf("Error while connecting to the database %v\n", err)
	}
	log.Println("Successfully connected to DB...")

	// wiring
	customerRepo := db.NewCustomerRepository(client)
	accountRepo  := db.NewAccountRepository(client)

	cs := service.NewCustomerService(customerRepo)
	as := service.NewAccountService(accountRepo)

	ch := handlers.NewCustomerHandlers(cs)
	ah := handlers.NewAccountHandlers(as)

	// routes
	router.HandleFunc("/customers", ch.GetAllCustomers).Methods("GET")
	router.HandleFunc("/customers/{id}", ch.GetSingleCustomer).Methods("GET")
	router.HandleFunc("/account", ah.CreateAccount).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))

}
