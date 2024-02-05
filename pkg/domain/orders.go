package domain

import (
	"time"

	"gorm.io/gorm"
)

// Order of user
type Order struct {
	gorm.Model
	ID              int           `json:"id" gorm:"primarykey;autoIncrement"`
	UserID          int           `json:"user_id" gorm:"not null"`
	User            User          `json:"-" gorm:"foreignkey:UserID"`
	AddressID       int           `json:"address_id" gorm:"not null"`
	Address         Address       `json:"-" gorm:"foreignkey:AddressID"`
	PaymentMethodID int           `json:"paymentmethod_id" gorm:"default:1"`
	PaymentMethod   PaymentMethod `json:"-" gorm:"foreignkey:PaymentMethodID"`
	PaymentID       string        `json:"payment_id"`
	Price           float64       `json:"price"`
	OrderedAt       time.Time     `json:"ordered_at"`
	OrderStatus     string        `json:"order_status" gorm:"order_status:4;default:'PENDING';check:order_status IN('PENDING','SHIPPED','DELIVERED','CANCELED','RETURNED')"`
	PaymentStatus   string        `json:"payment_status" gorm:"default:'PENDING'"`
}

// product details of the order

type OrderItem struct {
	ID          int       `json:"id" gorm:"primarykey;autoIncrement"`
	OrderID     int       `json:"order_id"`
	Order       Order     `json:"-" gorm:"foreignkey:OrderID"`
	InventoryID int       `json:"inventory_id"`
	Inventory   Inventory `json:"-" gorm:"foreignkey:InventoryID"`
	Quantity    int       `json:"quantity"`
	TotalPrice  float64   `json:"total_price"`
}

// Order details with order status

type AdminOrderResponse struct {
	Pending   []OrderDetails
	Shipped   []OrderDetails
	Delivered []OrderDetails
	Canceled  []OrderDetails
}

// Details of order

type OrderDetails struct {
	ID            int     `json:"id" gorm:"column:order_id"`
	Username      string  `json:"username"`
	Address       string  `json:"address"`
	PaymentMethod string  `json:"paymentmethod"`
	Total         float64 `json:"total"`
}

type PaymentMethod struct {
	ID           int    `json:"id" gorm:"primarykey"`
	PaymenMethod string `json:"paymentmethod" validate:"required" gorm:"unique"`
}

type SalesReport struct {
	Orders       []Order
	TotalRevenue float64
	TotalOrders  int
	BestSellers  []string
}

type ProductReport struct {
	InventoryID int
	Quantity    int
}
