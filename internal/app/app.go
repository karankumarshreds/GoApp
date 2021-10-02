package app

import (
	"github.com/gorilla/mux"
	"github.com/karankumarshreds/GoApp/internal/framework/db"
	"github.com/karankumarshreds/GoApp/internal/framework/http/handlers"
	"github.com/karankumarshreds/GoApp/internal/service"
	"log"
	"net/http"
)

func Start() {
	router := mux.NewRouter()

	// wiring application
	repo := db.NewCustomerRepositoryStub()
	cs := service.NewCustomerService(repo)
	h := handlers.NewCustomerHandlers(cs)

	// routes
	router.HandleFunc("/customers", h.GetAllCustomers)

	log.Fatal(http.ListenAndServe(":8000", router))
}
