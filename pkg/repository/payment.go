package repository

import (
	"github.com/Anandhu4456/go-Ecommerce/pkg/repository/interfaces"
	"gorm.io/gorm"
)

type paymentRepository struct {
	DB *gorm.DB
}

// constructor function

func NewPaymentRepository(DB *gorm.DB) interfaces.PaymentRepository {
	return &paymentRepository{
		DB: DB,
	}
}

func (pr *paymentRepository) AddNewPaymentMethod(paymentMethod string) error {
	query := `INSERT INTO payment_methods(payment_method)VALUES(?)`

	if err := pr.DB.Exec(query, paymentMethod).Error; err != nil {
		return err
	}
	return nil
}

func (pr *paymentRepository)RemovePaymentMethod(paymentMethodId int) error{
	query:=`DELETE FROM payment_methods WHERE id=?`

	if err:=pr.DB.Exec(query,paymentMethodId).Error;err!=nil{
		return err
	}
	return nil
}