package db

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/karankumarshreds/GoApp/internal/core"
	"github.com/karankumarshreds/GoApp/internal/pkg/custom_errors"
	"github.com/karankumarshreds/GoApp/internal/pkg/logger"
	_ "github.com/lib/pq"
	"log"
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
		//"SELECT customer_id, city, name, date_of_birth, zipcode, status FROM customers;",
		"SELECT * FROM customers;",
	)
	if err != nil {
		log.Printf("Error while fetching customers %v", err)
		return nil, err
	}
	customers := make([]core.Customer, 0)
	err = sqlx.StructScan(rows, &customers)
	if err != nil {
		logger.ErrorLog("Unable to scan all the users", err)
		return nil, err
	}
	return customers, nil

	//for rows.Next() {
	//	var c core.Customer
	//	err := rows.Scan(&c.City, &c.Name, &c.DateOfBirth, &c.ZipCode, &c.Status, &c.Id)
	//	if err != nil {
	//		return nil, err
	//	}
	//	customers = append(customers, c)
	//}
}

func (d CustomerRepositoryDb) FindById(id string) (*core.Customer, *custom_errors.CustomError) {
	customerId, _ := strconv.Atoi(id)
	q := "SELECT customer_id, city, name, date_of_birth, zipcode, status FROM customers WHERE customer_id=$1"
	row:= d.db.QueryRow( q, customerId)
	c := core.Customer{}
	err := row.Scan(&c.Id, &c.City, &c.Name, &c.DateOfBirth, &c.ZipCode, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("ERROR: No customer found")
			return nil, custom_errors.NewNotFoundError("customer not found")
		} else {
			log.Printf("ERROR: Unable to scan customer %v", err)
			return nil, custom_errors.NewInternalServerError("unable to query users")
		}
	}
	return &c, nil
}
























