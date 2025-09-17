package entities

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
