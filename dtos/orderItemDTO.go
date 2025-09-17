package dtos

import "OrderProject/entities"

type CustomerDTO struct {
	ID               int64
	Name             string
	Email            string
	Phone            string
	OrderID          int64
	Order            []entities.Order            `gorm:"foreignKey:CustomerID"`
	CustomerSegments []entities.CustomerSegments `gorm:"many2many:customer_segments_link"`
}
