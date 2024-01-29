package service

import (
	"banking/domain"
	"banking/dto"
	"banking/errs"
	"banking/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

//go:generate mockgen -destination=../mocks/service/mockCustomerService.go -package=service ./banking/service CustomerService
type CustomerService interface {
	GetAllCustomer(string) ([]dto.CustomerResponse, *errs.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppError)
	NewAccount(domain.Customer) (**mongo.InsertOneResult, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepositoryDb
}

func (s DefaultCustomerService) GetAllCustomer(status string) ([]dto.CustomerResponse, *errs.AppError) {

	var customers []domain.Customer
	var err *errs.AppError
	if status == "active" {
		customers, err = s.repo.FindAllByStatus(true)
	} else if status == "inactive" {
		customers, err = s.repo.FindAllByStatus(false)
	} else if status == "" {
		customers, err = s.repo.FindAll()
	}
	if err != nil {
		return nil, err
	}
	response := make([]dto.CustomerResponse, 0)
	for _, c := range customers {
		response = append(response, c.ToCustomerResponseDto())
	}
	return response, err
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}
	response := c.ToCustomerResponseDto()
	return &response, nil
}

func (s DefaultCustomerService) NewAccount(c domain.Customer) (**mongo.InsertOneResult, *errs.AppError) {
	response, err := s.repo.Save(c)
	if err != nil {
		logger.Error("NewAccount() | Error while saving the customer record | " + err.Message)
		return nil, err
	}
	return response, nil
}
func NewCustomerService(repository domain.CustomerRepositoryDb) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
