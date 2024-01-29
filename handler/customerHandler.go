package handler

import (
	"banking/domain"
	"banking/logger"
	"banking/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type CustomerHandler struct {
	service service.CustomerService
}

func (ch *CustomerHandler) GetAllCustomers(w http.ResponseWriter, r *http.Request) {

	status := r.URL.Query().Get("status")

	customers, err := ch.service.GetAllCustomer(status)

	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customers)
	}
}

func (ch *CustomerHandler) GetCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]

	logger.Debug("GetCustomer() | Customer ID is " + id)

	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customer)
	}
}

func (ch *CustomerHandler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var request domain.Customer
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		logger.Debug("CreateCustomer() | Bad Request")
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		if logger.IsDebugEnabled() {
			logger.Debug("CreateCustomer() | Customer Request Body | " + request.ToString())
		}
		account, appError := ch.service.NewAccount(request)
		if appError != nil {
			writeResponse(w, appError.Code, appError.AsMessage())
		} else {
			writeResponse(w, http.StatusCreated, account)
		}
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}

func NewCustomerHandler(customerService service.DefaultCustomerService) CustomerHandler {
	return CustomerHandler{service: customerService}
}
