package domain

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/slns/banking/app/errs"
	"github.com/slns/banking/app/logger"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	var rows *sql.Rows
	var err error
if status == "" {
	findAllSql := "SELECT customer_id, name, city, zipcode, date_of_birth, status from customers"
	rows, err = d.client.Query(findAllSql)
} else {
	findAllSql := "SELECT customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
	rows, err = d.client.Query(findAllSql, status)
}

if err != nil {
	logger.Error("Error while quering customer table " + err.Error())
	return nil, errs.NewUnexpectedError("Unexpected database error")
}

customers := make([]Customer, 0)
for rows.Next() {
	var c Customer
	err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
	if err != nil {
		logger.Error("Error while scaning customers " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	customers = append(customers, c)
}
return customers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	customerSql := "SELECT customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	row := d.client.QueryRow(customerSql, id)
	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customers not found")
		} else {
			logger.Error("Error while scaning customer " + err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}	
	}
	return &c, nil
}

func NewcustomerrepositoryDb() CustomerRepositoryDb  {
	client, err := sql.Open("mysql", "root:root@tcp(localhost:8889)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDb{client}
}