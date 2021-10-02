package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type Customer struct {
	Id string
	Name string
	City string
	ZipCode string
	DateOfBirth string
	Status bool
}

type CustomerRepositoryDb struct {
	db *sql.DB
}

func NewCustomerRepository (connString string) *CustomerRepositoryDb{
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatalf("Error while connecting to the database %v\n", err)
	}
	log.Println("Successfully connected to DB...")
	return &CustomerRepositoryDb{db}
}

func (d CustomerRepositoryDb) FindAll() ([]Customer, error){
	rows, err := d.db.Query(
		"SELECT customer_id, city, name, date_of_birth, zipcode, status FROM customers;",
	)
	if err != nil {
		log.Printf("Error while fetching customers %v", err)
		return nil, err
	}
	customers := []Customer{}

	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.City, &c.Name, &c.DateOfBirth, &c.ZipCode, &c.Status)
		if err != nil {
			return nil, err
		}
		customers = append(customers, c)
	}
	return customers, nil
}