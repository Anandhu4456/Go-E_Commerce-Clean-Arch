package repository

import (
	"github.com/Anandhu4456/go-Ecommerce/pkg/repository/interfaces"
	"gorm.io/gorm"
)

type orderRepository struct {
	DB *gorm.DB
}

// constructor function
func NewOrderRepository(DB *gorm.DB)interfaces.OrderRepository{
	return &orderRepository{
		DB:DB,
	}
}

