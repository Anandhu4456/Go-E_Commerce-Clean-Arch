package interfaces

import (
	"github.com/Anandhu4456/go-Ecommerce/pkg/utils/models"
)

type PaymentRepository interface {
	AddNewPaymentMethod(paymentMethod string) error
	RemovePaymentMethod(paymentMethodId int) error
	GetPaymentMethods() ([]models.PaymentMethod, error)
	FindUsername(user_id int) (string, error)
	FindPrice(order_id int) (float64, error)
	UpdatePaymentDetails(orderId, paymentId, razorId string) error
}
