package dtos

import "OrderProject/entities"

// название файла orderItemDTO, что подразумевает что тут данные о заказах
// но сюда ты добавил данные клиента, что сильно путает и является плохой практикой
// так же названия файлов в ГО желательно писать через snake case (order_item_dto.go)
type CustomerDTO struct {
	ID               int64
	Name             string
	Email            string
	Phone            string
	OrderID          int64
	Order            []entities.Order            `gorm:"foreignKey:CustomerID"`
	CustomerSegments []entities.CustomerSegments `gorm:"many2many:customer_segments_link"`
}
