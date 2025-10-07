package entities

import (
	"errors"
)

type Customer struct {
	ID               int64
	Name             string
	Email            string
	Phone            string
	Promocode        string
	OrderID          int64
	Order            []Order            `gorm:"foreignKey:CustomerID"`
	CustomerSegments []CustomerSegments `gorm:"many2many:customer_segments_link"`
}

func (c *Customer) Validate() error {
	if c.Email == "" || c.Name == "" || c.Phone == "" {
		return errors.New("customer name OR email are required to add a customer")
	}
	return nil
}

//список структур описывающие сущности в бд(и их сопуствующие) лучше вывести в отдельные файлы
//--entities
//----order.go
//----customer.go

type Order struct {
	ID         int64
	Status     string
	Price      int
	CustomerID int64
	Customer   *Customer `gorm:"foreignKey:CustomerID"`
}

type CustomerSegments struct {
	ID        int64
	Name      string
	Customers []Customer `gorm:"many2many:customer_segments_link"`
}

type CustomerSegmentsLink struct {
	CustomerID int64
	OrderID    int64
}

/*
Customer 1 → M Orders
Customer M ↔ M Segments
Order M → 1 Customer
*/
