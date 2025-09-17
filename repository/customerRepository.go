package repository

import (
	"OrderProject/entities"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	AddCustomer(customer entities.Customer)
	UpdateCustomer(customer entities.Customer)
	DeleteCustomer(id int)
	GetAllCustomers() []entities.Customer
	GetByIdCustomer(id int) entities.Customer
	GetByNameCustomer(name string) []entities.Customer
}

type customerRepository struct {
	gormDB *gorm.DB
}

func NewCustomerRepository(gormDB *gorm.DB) CustomerRepository {
	return &customerRepository{gormDB: gormDB}
}

func (customRepository *customerRepository) AddCustomer(customer entities.Customer) {
	customRepository.gormDB.Create(&customer)
}

func (customRepository *customerRepository) UpdateCustomer(customer entities.Customer) {
	customRepository.gormDB.Save(&customer)
}

func (customRepository *customerRepository) DeleteCustomer(id int) {
	customRepository.gormDB.Delete(&entities.Customer{}, id)
}

func (customRepository *customerRepository) GetAllCustomers() []entities.Customer {
	var customers []entities.Customer
	customRepository.gormDB.Preload("Order").Preload("CustomerSegments").Find(&customers)
	return customers
}

func (customRepository *customerRepository) GetByIdCustomer(id int) entities.Customer {
	var customer entities.Customer
	customRepository.gormDB.First(&customer, "id = ?", id)
	return customer
}

func (customRepository *customerRepository) GetByNameCustomer(name string) []entities.Customer {
	var customers []entities.Customer
	customRepository.gormDB.Where("name = ?", name).Find(&customers)
	return customers
}
