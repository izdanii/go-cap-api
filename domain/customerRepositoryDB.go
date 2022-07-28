package domain

import (
	"capi/errs"
	"capi/logger"
	"database/sql"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type CustomerRepositoryDB struct {
	client *sqlx.DB
}

func NewCustomerRepositoryDB() CustomerRepositoryDB {
	connStr := "postgres://postgres:f41zd4n11@localhost/Go-API?sslmode=disable"
	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return CustomerRepositoryDB{db}
}

func (d CustomerRepositoryDB) FindByID(customerID string) (*Customer, *errs.AppErr) {
	query := "select * from customers where customer_id = $1"

	// row := d.client.QueryRow(query, customerID)

	var c Customer
	err := d.client.Get(&c, query, customerID)
	// err := row.Scan(&c.ID, &c.Name, &c.DateOfBirth, &c.City, &c.ZipCode, &c.Status)
	if err != nil {

		if err == sql.ErrNoRows {
			logger.Error("error customer data not found" + err.Error())
			return nil, errs.NewNotFoundError("Customer Not Found")
		} else {
			logger.Error("error scanning customer data" + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}

	}
	return &c, nil

}

func (d CustomerRepositoryDB) FindAll(customerStatus string) ([]Customer, *errs.AppErr) {
	var c []Customer
	if customerStatus == "" {
		query := "select * from customers"
		err := d.client.Select(&c, query)

		if err != nil {
			logger.Error("error query data to customer table" + err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
	} else {
		if customerStatus == "active" {
			customerStatus = "1"
		} else {
			customerStatus = "0"
		}

		query := "select * from customers where status = $1"
		err := d.client.Select(&c, query, customerStatus)

		if err != nil {
			logger.Error("error query data to customer table" + err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
	}

	return c, nil
}
