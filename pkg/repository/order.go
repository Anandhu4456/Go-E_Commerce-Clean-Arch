package repository

import (
	"errors"
	"time"

	"github.com/Anandhu4456/go-Ecommerce/pkg/domain"
	"github.com/Anandhu4456/go-Ecommerce/pkg/utils/models"
	"gorm.io/gorm"
)

type orderRepository struct {
	DB *gorm.DB
}

// constructor function
func NewOrderRepository(DB *gorm.DB) *orderRepository {
	return &orderRepository{
		DB: DB,
	}
}

func (orr *orderRepository) GetOrders(id, page, limit int) ([]domain.Order, error) {
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 10
	}
	offset := (page - 1) * limit

	var getOrders []domain.Order

	err := orr.DB.Raw("SELECT * FROM orders WHERE id= ? limit ? offset ?", id, limit, offset).Scan(&getOrders).Error
	if err != nil {
		return []domain.Order{}, err
	}
	return getOrders, nil
}

func (orr *orderRepository) GetOrdersInRange(startDate, endDate time.Time) ([]domain.Order, error) {
	var getOrdersInTimeRange []domain.Order

	// to fetch orders with in a time range
	err := orr.DB.Raw("SELECT * FROM orders WHERE ordered_at BETWEEN ? AND ?", startDate, endDate).Scan(&getOrdersInTimeRange).Error
	if err != nil {
		return []domain.Order{}, err
	}
	return getOrdersInTimeRange, nil
}

func (orr *orderRepository) GetProductsQuantity() ([]domain.ProductReport, error) {

	var getProductQuantity []domain.ProductReport

	err := orr.DB.Raw("SELECT inventory_id,quantity FROM order_items").Scan(&getProductQuantity).Error
	if err != nil {
		return []domain.ProductReport{}, err
	}
	return getProductQuantity, nil
}

func (orr *orderRepository) CreditToUserWallet(amount float64, walletId int) error {

	if err := orr.DB.Exec("update wallets set amount=$1 where id=$2", amount, walletId).Error; err != nil {
		return err
	}

	return nil

}

func (orr *orderRepository) GetCart(userid int) (models.GetCart, error) {

	var cart models.GetCart
	err := orr.DB.Raw("SELECT inventories.product_name,cart_products.quantity,cart_products.total_price AS total FROM cart_products JOIN inventories ON cart_products.inventory_id=inventores.id WHERE user_id=?", userid).Scan(&cart).Error
	if err != nil {
		return models.GetCart{}, err
	}
	return cart, nil
}

func (orr *orderRepository) GetProductNameFromId(id int) (string, error) {

	var productName string

	err := orr.DB.Raw("SELECT product_name FROM inventories WHERE id = ?", id).Scan(&productName).Error
	if err != nil {
		return "", err
	}
	return productName, nil
}

func (orr *orderRepository) OrderItems(userid int, addressid int, paymentid int, total float64, coupon string) (int, error) {

	var id int

	query := `
	
	INSERT INTO orders
		(user_id,address_id,price,payment_method_id,total,coupon_used)
	VALUES
		(?,?,?,?,?)
	RETURNING id
	`
	err := orr.DB.Raw(query, userid, addressid, paymentid, total, coupon).Scan(&id).Error
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (orr *orderRepository) CreateNewWallet(userID int) (int, error) {

	var walletID int
	err := orr.DB.Exec("Insert into wallets(user_id,amount) values($1,$2)", userID, 0).Error
	if err != nil {
		return 0, err
	}

	if err := orr.DB.Raw("select id from wallets where user_id=$1", userID).Scan(&walletID).Error; err != nil {
		return 0, err
	}

	return walletID, nil
}

func (orr *orderRepository) AddOrderProducts(order_id int, cart []models.GetCart) error {
	query := `
	
	INSERT INTO 
		order_items
		(order_id,inventory_id,quantity,total_price)
	VALUES
		(?,?,?,?)

	`
	for _, cartVals := range cart {
		var invId int
		err := orr.DB.Raw("SELECT id FROM inventories WHERE product_name=?", cartVals.ProductName).Scan(&invId).Error
		if err != nil {
			return err
		}
		if err := orr.DB.Raw(query, order_id, invId, cartVals.Quantity, cartVals.Total).Error; err != nil {
			return err
		}

	}
	return nil
}

func (orr *orderRepository) FindWalletIdFromUserID(userId int) (int, error) {

	var count int
	err := orr.DB.Raw("select count(*) from wallets where user_id = ?", userId).Scan(&count).Error
	if err != nil {
		return 0, err
	}

	var walletID int
	if count > 0 {
		err := orr.DB.Raw("select id from wallets where user_id = ?", userId).Scan(&walletID).Error
		if err != nil {
			return 0, err
		}
	}

	return walletID, nil

}

func (orr *orderRepository) CancelOrder(orderid int) error {
	err := orr.DB.Exec("UPDATE orders SET order_status='CANCELED' WHERE id=?", orderid).Error
	if err != nil {
		return err
	}
	return nil
}

func (orr *orderRepository) EditOrderStatus(status string, id int) error {
	err := orr.DB.Exec("UPDATE orders SET order_status=? WHERE id=?", status, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (orr *orderRepository) MarkAsPaid(orderID int) error {
	if err := orr.DB.Exec("UPDATE orders SET payment_status='PAID' WHERE id=?", orderID).Error; err != nil {
		return err
	}
	return nil
}

func (orr *orderRepository) AdminOrders(status string) ([]domain.OrderDetails, error) {

	var orderDetails []domain.OrderDetails
	query := `
	SELECT orders.id AS id, users.name AS username, CONCAT('House Name:',addresses.house_name, ',', 'Street:', addresses.street, ',', 'City:', addresses.city, ',', 'State', addresses.state, ',', 'Phone:', addresses.phone) AS address, payment_methods.payment_name AS payment_method, orders.final_price As total FROM orders JOIN users ON users.id = orders.user_id JOIN payment_methods ON payment_methods.id = orders.payment_method_id JOIN addresses ON orders.address_id = addresses.id WHERE order_status = $1
	`
	err := orr.DB.Raw(query, status).Scan(&orderDetails).Error
	if err != nil {
		return []domain.OrderDetails{}, err
	}
	return orderDetails, nil
}

func (orr *orderRepository) CheckOrder(orderID string, userID int) error {

	var count int
	err := orr.DB.Raw("SELECT COUNT (*)FROM orders WHERE order_id=?", orderID).Scan(&count).Error
	if err != nil {
		return err
	}
	if count < 0 {
		return errors.New("no such order exist")
	}

	var checkUser int
	chUser := orr.DB.Raw("SELECT user_id FROM orders WHERE order_id=?", orderID).Scan(&checkUser).Error
	if err != nil {
		return chUser
	}
	if userID != checkUser {
		return errors.New("no order did by this user")
	}
	return nil
}

func (orr *orderRepository) GetOrderDetail(orderID string) (domain.Order, error) {
	var orderDetails domain.Order

	err := orr.DB.Raw("SELECT * FROM orders WHERE order_id=?", orderID).Scan(&orderDetails).Error
	if err != nil {
		return domain.Order{}, err
	}
	return orderDetails, nil
}

func (orr *orderRepository) FindAmountFromOrderID(orderID int) (float64, error) {
	var amount float64

	err := orr.DB.Raw("SELECT price FROM orders WHERE order_id=?", orderID).Scan(&amount).Error
	if err != nil {
		return 0, err
	}
	return amount, nil
}

func (orr *orderRepository) FindUserIdFromOrderID(orderID int) (int, error) {
	var userId int

	err := orr.DB.Raw("SELECT user_id FROM orders WHERE order_id=?", orderID).Scan(&userId).Error
	if err != nil {
		return 0, err
	}
	return userId, nil
}

func (orr *orderRepository) ReturnOrder(id int) error {
	err := orr.DB.Exec("UPDATE orders SET order_status='RETURNED' WHERE id=?", id).Error
	if err != nil {
		return err
	}
	return nil
}

func (orr *orderRepository) CheckOrderStatus(orderID int) (string, error) {
	var orderStatus string

	err := orr.DB.Raw("SELECT order_status FROM orders WHERE order_id=?", orderID).Scan(&orderStatus).Error
	if err != nil {
		return "", err
	}
	return orderStatus, nil
}

func (orr *orderRepository) CheckPaymentStatus(orderID int) (string, error) {
	var paymentStatus string
	err := orr.DB.Raw("SELECT payment_status FROM orders WHERE order_id=?", orderID).Scan(&paymentStatus).Error
	if err != nil {
		return "", err
	}
	return paymentStatus, nil
}
