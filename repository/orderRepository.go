package repository

import (
	"OrderProject/entities"
	"gorm.io/gorm"
)

type OrderRepository interface {
	AddOrder(order entities.Order)
	UpdateOrder(order entities.Order)
	DeleteOrder(id int)
	GetAllOrders() []entities.Order
	GetByIdOrder(id int) entities.Order
}

type orderRepository struct {
	gormDB *gorm.DB
}

func NewOrderRepository(gormDB *gorm.DB) OrderRepository {
	return &orderRepository{gormDB: gormDB}
}

func (orderRepository *orderRepository) AddOrder(order entities.Order) {
	orderRepository.gormDB.Create(&order)
}

func (orderRepository *orderRepository) UpdateOrder(order entities.Order) {
	orderRepository.gormDB.Save(&order)
}

func (orderRepository *orderRepository) DeleteOrder(id int) {
	orderRepository.gormDB.Delete(&entities.Order{}, id)
}

func (orderRepository *orderRepository) GetAllOrders() []entities.Order {
	var orders []entities.Order
	orderRepository.gormDB.Find(&orders)
	return orders
}

func (orderRepository *orderRepository) GetByIdOrder(id int) entities.Order {
	var order entities.Order
	orderRepository.gormDB.First(&order, id)
	return order
}
