package interfaces

import (
	"github.com/Anandhu4456/go-Ecommerce/pkg/domain"
	"github.com/Anandhu4456/go-Ecommerce/pkg/utils/models"
)

type OrderUsecase interface {
	GetOrders(id, page, limit int) ([]domain.Order, error)
	OrderItemsFromCart(userid int, order models.Order, coupon string) (string, error)
	CancelOrder(id, orderid int) error
	EditOrderStatus(status string, id int) error
	MarkAsPaid(orderID int) error
	AdminOrders(page, limit int, status string) ([]domain.OrderDetails, error)
	DailyOrders() (domain.SalesReport, error)
	WeeklyOrders() (domain.SalesReport, error)
	MonthlyOrders() (domain.SalesReport, error)
	AnnualOrders() (domain.SalesReport, error)
	CustomDateOrders(dates models.CustomDates) (domain.SalesReport, error)
	ReturnOrder(id int) error
}
