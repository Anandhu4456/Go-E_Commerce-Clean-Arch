package interfaces

import (
	"github.com/Anandhu4456/go-Ecommerce/pkg/domain"
	"github.com/Anandhu4456/go-Ecommerce/pkg/utils/models"
	"time"
)

type OrderRepository interface{
	GetOrders(id,page,limit int)([]domain.Order,error)
	GetProductsQuantity()([]domain.ProductReport,error)
	GetOrdersInRange(startDate,endDate time.Time)([]domain.Order,error)
	GetProductNameFromId(id int)(string,error)
	GetCart(userid int)(models.GetCart,error)

	OrderItems(userid int,addressid int,paymentid int,total float64,coupon string) (int,error)
	AddOrderProducts(order_id int, cart []models.GetCart) error
	CancelOrder(orderid int) error
	EditOrderStatus(status string, id int) error
	MarkAsPaid(orderID int) error
	AdminOrders(status string) ([]domain.OrderDetails, error)

	CheckOrder(orderID string, userID int) error
	GetOrderDetail(orderID string) (domain.Order, error)
	FindUserIdFromOrderID(orderID int) (int, error)
	FindWalletIdFromUserID(userId int) (int, error)
	CreateNewWallet(userID int) (int, error)
	CreditToUserWallet(amount float64, walletID int) error
	FindAmountFromOrderID(orderID int) (float64, error)
	ReturnOrder(id int) error
	CheckOrderStatus(orderID int) (string, error)
	CheckPaymentStatus(orderID int) (string, error)
}