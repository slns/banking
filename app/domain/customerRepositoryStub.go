package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer {
		{"1001", "Robert", "New York", "110011", "2000-01-01", "1"},
		{"1002", "Sérgio", "New York", "110011", "2000-01-01", "1"},
	}
	return CustomerRepositoryStub{ customers }
}