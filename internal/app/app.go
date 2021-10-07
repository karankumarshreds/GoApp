package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/karankumarshreds/GoApp/internal/framework/db"
	"github.com/karankumarshreds/GoApp/internal/framework/http/handlers"
	"github.com/karankumarshreds/GoApp/internal/service"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"time"
)

func Start() {

	router := mux.NewRouter()
	client := dbClient()

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

func dbClient() *sqlx.DB {
	connString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", "postgres", "password", "banking")
	client, err := sqlx.Open("postgres", connString)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}
