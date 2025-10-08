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

// хорошая практика, когда методы в репозиторий так же возвращают ощибки
// представь кейс, когда ты отправил пост запрос на сохранение, все ок, эндпойнт вернул 200. Но запись в бд не сохраняется
// поэтому надо хендлить ошибки и возвращать их в том числе
func (customRepository *customerRepository) AddCustomerWithErrorHandler(customer entities.Customer, err error) {
	err = customRepository.gormDB.
		Create(&customer).
		Error

	return
}

func (customRepository *customerRepository) AddCustomer(customer entities.Customer) {
	customRepository.gormDB.Create(&customer)
}

// кстати, метод save не работает без where :)
func (customRepository *customerRepository) UpdateCustomer(customer entities.Customer) {
	customRepository.gormDB.Save(&customer)
}

func (customRepository *customerRepository) DeleteCustomer(id int) {
	customRepository.gormDB.Delete(&entities.Customer{}, id)
}

// используй табуляцию, это тебе не джава что бы все в ряд писать
func (customRepository *customerRepository) GetAllCustomers() []entities.Customer {
	var customers []entities.Customer
	customRepository.gormDB.
		Preload("Order").
		Preload("CustomerSegments").
		Find(&customers)
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
