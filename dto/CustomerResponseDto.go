package dto

type CustomerResponse struct {
	Name             string `json:"name"`
	Contact          string `json:"contact"`
	DateOfBirth      string `json:"date_of_birth"`
	Status           string `json:"status"`
	CurrentAddress   string `json:"current_address,omitempty"`
	PermanentAddress string `json:"permanent_address,omitempty"`
}

func (c CustomerResponse) ToString() string {
	return "Customer[Name: " + c.Name +
		", Contact: " + c.Contact +
		", DateOfBirth: " + c.DateOfBirth + ", Status: " + c.Status + "]"
}
