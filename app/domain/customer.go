package domain

import "github.com/slns/banking/app/errs"

type Customer struct {
	Id string `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
	Zipcode string `json:"zipCode"`
	DateofBirth string `json:"date"`
	Status string `json:"status"`
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
	ById(string) (*Customer, *errs.AppError)
}