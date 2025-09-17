package services

import (
	"OrderProject/dtos"
	"OrderProject/entities"
	"OrderProject/mappers"
	"OrderProject/repository"
	"github.com/google/uuid"
	"log"
)

type CustomerService interface {
	AddCustomer(customer entities.Customer)
	UpdateCustomer(customer entities.Customer)
	DeleteCustomer(id int)
	GetAllCustomers() []dtos.CustomerDTO
	GetByIdCustomer(id int) dtos.CustomerDTO
	GetByNameCustomer(name string) []dtos.CustomerDTO
}

type customerService struct {
	customerRepository repository.CustomerRepository
}

func NewCustomerService(customerRepository repository.CustomerRepository) CustomerService {
	return &customerService{customerRepository: customerRepository}
}

func (customerService *customerService) AddCustomer(customer entities.Customer) {
	if customer.Name == "" || customer.Email == "" || customer.Phone == "" {
		log.Fatal("Customer name and email are required to add a customer")
	} else {
		customer.Promocode = uuid.New().String()
		customerService.customerRepository.AddCustomer(customer)
	}
}

func (customerService *customerService) UpdateCustomer(customer entities.Customer) {
	if customer.Name == "" || customer.Email == "" || customer.Phone == "" {
		log.Fatal("Customer name and email are required to update a customer")
	} else {
		var currentCustomer entities.Customer = customerService.customerRepository.GetByIdCustomer(int(customer.ID))
		customer.Promocode = currentCustomer.Promocode
		customerService.customerRepository.UpdateCustomer(customer)
	}
}

func (customerService *customerService) DeleteCustomer(id int) {
	customerService.customerRepository.DeleteCustomer(id)
}

func (customerService *customerService) GetAllCustomers() []dtos.CustomerDTO {
	return mappers.MapToDTOList(customerService.customerRepository.GetAllCustomers())
}

func (customerService *customerService) GetByIdCustomer(id int) dtos.CustomerDTO {
	return mappers.MapToDTO(customerService.customerRepository.GetByIdCustomer(id))
}

func (customerService *customerService) GetByNameCustomer(name string) []dtos.CustomerDTO {
	return mappers.MapToDTOList(customerService.customerRepository.GetByNameCustomer(name))
}
