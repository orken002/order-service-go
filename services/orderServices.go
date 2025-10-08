package services

import (
	"OrderProject/entities"
	"OrderProject/repository"
)

type OrderService interface {
	AddOrder(order entities.Order)
	UpdateOrder(order entities.Order)
	DeleteOrder(id int)
	GetAllOrders() []entities.Order
	GetByIdOrder(id int) entities.Order
}

type orderService struct {
	orderRepository repository.OrderRepository
}

func NewOrderServices(orderRepository repository.OrderRepository) OrderService {
	return &orderService{orderRepository: orderRepository}
}

func (orderService *orderService) AddOrder(order entities.Order) {
	//названия каждой репозиторйи можешь просто сделать
	//Create, GetByID, GetMany, Update, Delete и тд
	//как должно быть: orderService.orderRepository.Create(order)
	//нет смысла orderRepository.CraeteOrder писать, так как ясно что orderRepo создаст только order и ничего более
	orderService.orderRepository.AddOrder(order)

}

func (orderService *orderService) UpdateOrder(order entities.Order) {
	orderService.orderRepository.UpdateOrder(order)
}

func (orderService *orderService) DeleteOrder(id int) {
	orderService.orderRepository.DeleteOrder(id)
}

func (orderService *orderService) GetAllOrders() []entities.Order {
	return orderService.orderRepository.GetAllOrders()
}

func (orderService *orderService) GetByIdOrder(id int) entities.Order {
	return orderService.orderRepository.GetByIdOrder(id)
}
