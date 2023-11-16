package usecase

import (
	"errors"
	"fmt"
	"time"

	"github.com/Anandhu4456/go-Ecommerce/pkg/domain"
	"github.com/Anandhu4456/go-Ecommerce/pkg/helper"
	pdf "github.com/Anandhu4456/go-Ecommerce/pkg/helper/pdf"
	interfaces "github.com/Anandhu4456/go-Ecommerce/pkg/repository/interfaces"
	services "github.com/Anandhu4456/go-Ecommerce/pkg/usecase/interfaces"
	"github.com/Anandhu4456/go-Ecommerce/pkg/utils/models"
)

type orderUsecase struct {
	orderRepo   interfaces.OrderRepository
	userUsecase services.UserUsecase
	walletRepo  interfaces.WalletRepository
	couponRepo  interfaces.CouponRepository
}

func NewOrderUsecase(orderRepo interfaces.OrderRepository, userUsecase services.UserUsecase, walletRepo interfaces.WalletRepository, couponRepo interfaces.CouponRepository) *orderUsecase {
	return &orderUsecase{
		orderRepo:   orderRepo,
		userUsecase: userUsecase,
		walletRepo:  walletRepo,
		couponRepo:  couponRepo,
	}
}

func (orU *orderUsecase) GetOrders(id, page, limit int) ([]domain.Order, error) {
	orders, err := orU.orderRepo.GetOrders(id, page, limit)
	if err != nil {
		return []domain.Order{}, err
	}
	return orders, nil
}

func (orU *orderUsecase) OrderItemsFromCart(userid int, order models.Order, coupon string) (string, error) {

	cart, err := orU.userUsecase.GetCart(userid, 0, 0)
	if err != nil {
		return "", err
	}
	var total float64

	for _, items := range cart {
		total = total + items.DiscoundPrice
	}
	fmt.Println("total without coupon: ", total)

	if coupon != "" {
		valid, err := orU.couponRepo.ValidateCoupon(coupon)
		if err != nil || !valid {
			return "", err
		}

		// Find discount
		discount := orU.couponRepo.FindCouponDiscount(coupon)

		if discount > 0 {
			totalDiscount := total * float64(discount) / 100
			fmt.Println("Discount: ", discount, "Total Discount: ", totalDiscount, (discount / 100), int(total), int(discount/100))
			total = total - totalDiscount
		}
	}
	fmt.Println("Total amount: ", total)
	var invoiceItems []*pdf.InvoiceData

	for _, items := range cart {
		inventory, err := pdf.NewInvoiceData(items.ProductName, int(items.Quantity), (items.DiscoundPrice))
		if err != nil {
			panic(err)
		}
		invoiceItems = append(invoiceItems, inventory)
	}
	// Create single invoice
	invoice := pdf.CreateInvoice("Your Store", "www.your.store", invoiceItems)
	pdf.GenerateInvoicePdf(*invoice)
	fmt.Printf("The Total Invoice Amount Is : %f ", invoice.CalculateInvoiceTotalAmount())

	// Cash on Delivery

	if order.PaymentId == 1 {
		order_id, err := orU.orderRepo.OrderItems(userid, order, total)
		if err != nil {
			return "", err
		}
		if err := orU.orderRepo.AddOrderProducts(order_id, cart); err != nil {
			return "", err
		}

		cart_id, _ := orU.userUsecase.GetCartID(userid)
		if err := orU.userUsecase.ClearCart(cart_id); err != nil {
			return "", err
		}
	} else if order.PaymentId == 2 {
		// Razor pay
		order_id, err := orU.orderRepo.OrderItems(userid, order, total)
		if err != nil {
			return "", err
		}
		if err := orU.orderRepo.AddOrderProducts(order_id, cart); err != nil {
			return "", err
		}
		link := fmt.Sprintf("https://yoursstore.online/users/payment/razorpay?id=%d", order_id)
		return link, err

	}
	if order.PaymentId == 3 {
		// Payment Form Wallet
		order_id, err := orU.orderRepo.OrderItems(userid, order, total)
		if err != nil {
			return "", err
		}
		if err := orU.orderRepo.AddOrderProducts(order_id, cart); err != nil {
			return "", err
		}
		wallet_id, err := orU.walletRepo.FindWalletIdFromUserId(userid)
		if err != nil {
			return "", err
		}
		balance, err := orU.walletRepo.GetBalance(wallet_id)
		if err != nil {
			return "", err
		}
		if float64(balance) < total {
			return "insufficient balance on wallet", errors.New("insufficient balance")
		}
		newBalance, err := orU.walletRepo.PayFromWallet(userid, order_id, total)
		if err != nil {
			return "", err
		}
		orU.walletRepo.AddHistory(int(total*-1), wallet_id, "Order Placed")

		cart_id, _ := orU.userUsecase.GetCartID(userid)

		if err := orU.userUsecase.ClearCart(cart_id); err != nil {
			return "", err
		}
		return fmt.Sprintf("%f RS paid from wallet,New balance is : %f", total, newBalance), nil

	}
	return "", nil
}

func (orU *orderUsecase) EditOrderStatus(status string, id int) error {
	if err := orU.orderRepo.EditOrderStatus(status, id); err != nil {
		return err
	}
	return nil
}

func (orU *orderUsecase) MarkAsPaid(orderID int) error {
	if err := orU.orderRepo.MarkAsPaid(orderID); err != nil {
		return err
	}
	return nil
}

func (orU *orderUsecase) AdminOrders(page, limit int, status string) ([]domain.OrderDetails, error) {

	if status != "PENDING" && status != "SHIPPED" && status != "CANCELLED" && status != "RETURNED" && status != "DELIVERED" {
		return []domain.OrderDetails{}, errors.New("invalid order status")
	}
	orders, err := orU.orderRepo.AdminOrders(page, limit, status)
	if err != nil {
		return []domain.OrderDetails{}, err
	}
	return orders, nil
}

func (orU *orderUsecase) DailyOrders() (domain.SalesReport, error) {
	var SalesReport domain.SalesReport
	endDate := time.Now()
	startDate := time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 0, 0, 0, 0, time.UTC)
	SalesReport.Orders, _ = orU.orderRepo.GetOrdersInRange(startDate, endDate)
	SalesReport.TotalOrders = len(SalesReport.Orders)
	total := 0.0

	for _, items := range SalesReport.Orders {
		total += items.Price
	}
	SalesReport.TotalRevenue = total

	products, err := orU.orderRepo.GetProductsQuantity()
	if err != nil {
		return domain.SalesReport{}, err
	}
	bestSellerIds := helper.FindMostBroughtProduct(products)

	var bestSellers []string

	for _, items := range bestSellers {

		product, err := orU.orderRepo.GetProductNameFromId(items)
		if err != nil {
			return domain.SalesReport{}, err
		}
		bestSellers = append(bestSellers, product)
	}
	SalesReport.BestSellers = bestSellers

	return SalesReport, nil
}

func (orU *orderUsecase)WeeklyOrders() (domain.SalesReport, error){
	var SalesReport domain.SalesReport
	endDate := time.Now()
	startDate := endDate.Add(-time.Duration(endDate.Weekday()) * 24 * time.Hour)
	SalesReport.Orders, _ = orU.orderRepo.GetOrdersInRange(startDate, endDate)
	SalesReport.TotalOrders = len(SalesReport.Orders)
	total := 0.0

	for _, items := range SalesReport.Orders {
		total += items.Price
	}
	SalesReport.TotalRevenue = total

	products, err := orU.orderRepo.GetProductsQuantity()
	if err != nil {
		return domain.SalesReport{}, err
	}
	bestSellerIds := helper.FindMostBroughtProduct(products)

	var bestSellers []string

	for _, items := range bestSellers {

		product, err := orU.orderRepo.GetProductNameFromId(items)
		if err != nil {
			return domain.SalesReport{}, err
		}
		bestSellers = append(bestSellers, product)
	}
	SalesReport.BestSellers = bestSellers

	return SalesReport, nil
}

func (orU *orderUsecase)MonthlyOrders() (domain.SalesReport, error){
	var SalesReport domain.SalesReport
	endDate := time.Now()
	startDate := time.Date(endDate.Year(),endDate.Month(),1,0,0,0,0,time.UTC)
	SalesReport.Orders, _ = orU.orderRepo.GetOrdersInRange(startDate, endDate)
	SalesReport.TotalOrders = len(SalesReport.Orders)
	total := 0.0

	for _, items := range SalesReport.Orders {
		total += items.Price
	}
	SalesReport.TotalRevenue = total

	products, err := orU.orderRepo.GetProductsQuantity()
	if err != nil {
		return domain.SalesReport{}, err
	}
	bestSellerIds := helper.FindMostBroughtProduct(products)

	var bestSellers []string

	for _, items := range bestSellers {

		product, err := orU.orderRepo.GetProductNameFromId(items)
		if err != nil {
			return domain.SalesReport{}, err
		}
		bestSellers = append(bestSellers, product)
	}
	SalesReport.BestSellers = bestSellers

	return SalesReport, nil
}