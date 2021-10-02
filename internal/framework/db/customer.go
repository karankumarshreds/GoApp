package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"github.com/karankumarshreds/GoApp/internal/core"
	"strconv"
)

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

func (d CustomerRepositoryDb) FindAll() ([]core.Customer, error){
	rows, err := d.db.Query(
		"SELECT customer_id, city, name, date_of_birth, zipcode, status FROM customers;",
	)
	if err != nil {
		log.Printf("Error while fetching customers %v", err)
		return nil, err
	}
	customers := []core.Customer{}

	for rows.Next() {
		var c core.Customer
		err := rows.Scan(&c.Id, &c.City, &c.Name, &c.DateOfBirth, &c.ZipCode, &c.Status)
		if err != nil {
			return nil, err
		}
		customers = append(customers, c)
	}
	return customers, nil
}

func (d CustomerRepositoryDb) FindById(id string) (*core.Customer, error) {
	customerId, _ := strconv.Atoi(id)
	row:= d.db.QueryRow(
		"" +
			"SELECT customer_id, city, name, date_of_birth, zipcode, status FROM customers WHERE customer_id=$1",
			customerId,
	)
	c := core.Customer{}
	err := row.Scan(&c.Id, &c.City, &c.Name, &c.DateOfBirth, &c.ZipCode, &c.Status)
	if err != nil {
		log.Printf("ERROR: Unable to scan customer %v", err)
		return nil, err
	}
	return &c, nil
}
























