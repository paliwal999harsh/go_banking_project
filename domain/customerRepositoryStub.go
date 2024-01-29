package domain

import "banking/errs"

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll(status bool) ([]Customer, *errs.AppError) {
	return s.customers, nil
}

func (s CustomerRepositoryStub) FindById(id string) (Customer, *errs.AppError) {
	return s.customers[0], nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{
			Id:          "C1001",
			FirstName:   "Harsh",
			LastName:    "Paliwal",
			Contact:     "+91-XXX-XXX-XXXX",
			DateOfBirth: "2000-01-01",
			Status:      true,
			CurrentAddress: Address{
				LineOne:    "House No 1",
				LineTwo:    "Lane No 1",
				City:       "Mumbai",
				State:      "Maharashtra",
				Country:    "India",
				PostalCode: "123456",
			},
			PermanentAddress: Address{
				LineOne:    "House No 1",
				LineTwo:    "Lane No 1",
				City:       "Mumbai",
				State:      "Maharashtra",
				Country:    "India",
				PostalCode: "123456",
			},
		},
		{
			Id:          "C1002",
			FirstName:   "Sachin",
			LastName:    "Kapkoti",
			Contact:     "+91-XXX-XXX-XXXX",
			DateOfBirth: "2000-01-02",
			Status:      true,
			CurrentAddress: Address{
				LineOne:    "House No 2",
				LineTwo:    "Lane No 1",
				City:       "Mumbai",
				State:      "Maharashtra",
				Country:    "India",
				PostalCode: "123456",
			},
			PermanentAddress: Address{
				LineOne:    "House No 2	",
				LineTwo:    "Lane No 1",
				City:       "Mumbai",
				State:      "Maharashtra",
				Country:    "India",
				PostalCode: "123456",
			},
		},
	}
	return CustomerRepositoryStub{customers}
}
