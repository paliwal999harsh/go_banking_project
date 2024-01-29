package domain

import (
	"banking/dto"
	"banking/errs"
)

type Address struct {
	LineOne    string `json:"line_one" bson:"line_one"`
	LineTwo    string `json:"line_two,omitempty" bson:"line_two"`
	City       string `json:"city" bson:"city"`
	State      string `json:"state" bson:"state"`
	Country    string `json:"country" bson:"country"`
	PostalCode string `json:"postal_code" bson:"postal_code"`
}
type Customer struct {
	Id               string  `json:"customer_id" bson:"_id"`
	FirstName        string  `json:"first_name" bson:"first_name"`
	LastName         string  `json:"last_name" bson:"last_name"`
	Contact          string  `json:"contact" bson:"contact"`
	DateOfBirth      string  `json:"date_of_birth" bson:"date_of_birth"`
	Status           bool    `json:"status" bson:"status"`
	CurrentAddress   Address `json:"current_address" bson:"current_address"`
	PermanentAddress Address `json:"permanent_address" bson:"permanent_address"`
}

func (c Customer) ToCustomerResponseDto() dto.CustomerResponse {
	return dto.CustomerResponse{
		Name:             c.FirstName + " " + c.LastName,
		Contact:          c.Contact,
		DateOfBirth:      c.DateOfBirth,
		Status:           c.statusAsText(),
		CurrentAddress:   c.addressAsText("current"),
		PermanentAddress: c.addressAsText("permanent"),
	}
}

func (c Customer) statusAsText() string {
	status := "active"
	if !c.Status {
		status = "inactive"
	}
	return status
}

func (c Customer) addressAsText(addressType string) string {
	if addressType == "current" {
		return c.CurrentAddress.LineOne + ", " + c.CurrentAddress.LineTwo + ", " + c.CurrentAddress.City +
			", " + c.CurrentAddress.State + ", " + c.CurrentAddress.Country + " - " + c.CurrentAddress.PostalCode
	} else if addressType == "permanent" {
		return c.PermanentAddress.LineOne + ", " + c.PermanentAddress.LineTwo + ", " + c.PermanentAddress.City +
			", " + c.PermanentAddress.State + ", " + c.PermanentAddress.Country + " - " + c.PermanentAddress.PostalCode
	}
	return ""
}

func (a Address) ToString() string {
	return "Address[LineOne: " + a.LineOne +
		", LineTwo: " + a.LineTwo + ", City: " + a.City +
		", State: " + a.State + ", Country: " + a.Country +
		", PostalCode: " + a.PostalCode + "]"
}

func (c Customer) ToString() string {
	return "Customer[Id: " + c.Id + ", FirstName: " + c.FirstName +
		", LastName: " + c.LastName + ", Contact: " + c.Contact +
		", DateOfBirth: " + c.DateOfBirth + ", Status: " + c.statusAsText() +
		", CurrentAddress: " + c.CurrentAddress.ToString() +
		", PremanentAddress: " + c.PermanentAddress.ToString() + "]"
}

type CustomerRepository interface {
	FindAll(status bool) ([]Customer, *errs.AppError)
	FindById(string) (*Customer, *errs.AppError)
}
