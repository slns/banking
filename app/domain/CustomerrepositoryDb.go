package domain

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/slns/banking/app/errs"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (d CustomerRepositoryDb) FindAll() ([]Customer, error) {

findAllSql := "SELECT customer_id, name, city, zipcode, date_of_birth, status from customers"

rows, err := d.client.Query(findAllSql)

if err != nil {
	log.Println("Error while quering customer table " + err.Error())
	return nil, err
}

customers := make([]Customer, 0)
for rows.Next() {
	var c Customer
	err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
	if err != nil {
		log.Println("Error while scaning customers " + err.Error())
		return nil, err
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
			log.Println("Error while scaning customer " + err.Error())
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