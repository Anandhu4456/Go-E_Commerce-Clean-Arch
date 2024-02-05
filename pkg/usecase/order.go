package usecase

import (
	"errors"
	"fmt"
	"time"

	"github.com/Anandhu4456/go-Ecommerce/pkg/domain"
	"github.com/Anandhu4456/go-Ecommerce/pkg/helper"
	interfaces "github.com/Anandhu4456/go-Ecommerce/pkg/repository/interfaces"
	services "github.com/Anandhu4456/go-Ecommerce/pkg/usecase/interfaces"
	"github.com/Anandhu4456/go-Ecommerce/pkg/utils/models"
)

type orderUsecase struct {
	orderRepo   interfaces.OrderRepository
	userUsecase services.UserUsecase
	// walletRepo  interfaces.WalletRepository
	couponRepo interfaces.CouponRepository
}

func NewOrderUsecase(orderRepo interfaces.OrderRepository, userUsecase services.UserUsecase /*walletRepo interfaces.WalletRepository*/, couponRepo interfaces.CouponRepository) services.OrderUsecase {
	return &orderUsecase{
		orderRepo:   orderRepo,
		userUsecase: userUsecase,
		// walletRepo:  walletRepo,
		couponRepo: couponRepo,
	}
}

func (orU *orderUsecase) GetOrders(id, page, limit int) ([]domain.Order, error) {
	orders, err := orU.orderRepo.GetOrders(id, page, limit)
	if err != nil {
		return []domain.Order{}, err
	}
	return orders, nil
}

func (orU *orderUsecase) OrderItemsFromCart(userid int, addressid int, paymentid int, couponid int) error {
	cart, err := orU.userUsecase.GetCart(userid)
	if err != nil {
		return err
	}

	var total float64
	for _, v := range cart.Values {
		total = total + v.DiscountPrice
	}
	// Find discount if any
	coupon, err := orU.couponRepo.FindCouponDetails(couponid)
	if err != nil {
		return err
	}
	totalDiscount := (total * float64(coupon.DiscountRate)) / 100
	total = total - totalDiscount

	order_id, err := orU.orderRepo.OrderItems(userid, addressid, paymentid, total, coupon.Coupon)
	if err != nil {
		return err
	}
	if err := orU.orderRepo.AddOrderProducts(order_id, cart.Values); err != nil {
		return err
	}
	for _, v := range cart.Values {
		if err := orU.userUsecase.RemoveFromCart(cart.Id, v.Id); err != nil {
			return err
		}
	}
	return nil

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

func (orU *orderUsecase) AdminOrders() (domain.AdminOrderResponse, error) {
	var response domain.AdminOrderResponse

	pending, err := orU.orderRepo.AdminOrders("PENDING")
	if err != nil {
		return domain.AdminOrderResponse{}, err
	}
	shipped, err := orU.orderRepo.AdminOrders("SHIPPED")
	if err != nil {
		return domain.AdminOrderResponse{}, err
	}
	delivered, err := orU.orderRepo.AdminOrders("DELIVERED")
	if err != nil {
		return domain.AdminOrderResponse{}, err
	}
	returned, err := orU.orderRepo.AdminOrders("RETURNED")
	if err != nil {
		return domain.AdminOrderResponse{}, err
	}
	canceled, err := orU.orderRepo.AdminOrders("CANCELED")
	if err != nil {
		return domain.AdminOrderResponse{}, err
	}

	response.Pending = pending
	response.Shipped = shipped
	response.Delivered = delivered
	response.Returned = returned
	response.Canceled = canceled

	return response, nil

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

	for _, items := range bestSellerIds {

		product, err := orU.orderRepo.GetProductNameFromId(items)
		if err != nil {
			return domain.SalesReport{}, err
		}
		bestSellers = append(bestSellers, product)
	}
	SalesReport.BestSellers = bestSellers

	return SalesReport, nil
}

func (orU *orderUsecase) WeeklyOrders() (domain.SalesReport, error) {
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

	for _, items := range bestSellerIds {

		product, err := orU.orderRepo.GetProductNameFromId(items)
		if err != nil {
			return domain.SalesReport{}, err
		}
		bestSellers = append(bestSellers, product)
	}
	SalesReport.BestSellers = bestSellers

	return SalesReport, nil
}

func (orU *orderUsecase) MonthlyOrders() (domain.SalesReport, error) {
	var SalesReport domain.SalesReport
	endDate := time.Now()
	startDate := time.Date(endDate.Year(), endDate.Month(), 1, 0, 0, 0, 0, time.UTC)
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

	for _, items := range bestSellerIds {

		product, err := orU.orderRepo.GetProductNameFromId(items)
		if err != nil {
			return domain.SalesReport{}, err
		}
		bestSellers = append(bestSellers, product)
	}
	SalesReport.BestSellers = bestSellers

	return SalesReport, nil
}
func (orU *orderUsecase) AnnualOrders() (domain.SalesReport, error) {
	var SalesReport domain.SalesReport
	endDate := time.Now()
	startDate := time.Date(endDate.Year(), 1, 1, 0, 0, 0, 0, time.UTC)
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

	for _, items := range bestSellerIds {

		product, err := orU.orderRepo.GetProductNameFromId(items)
		if err != nil {
			return domain.SalesReport{}, err
		}
		bestSellers = append(bestSellers, product)
	}
	SalesReport.BestSellers = bestSellers

	return SalesReport, nil
}

func (orU *orderUsecase) CustomDateOrders(dates models.CustomDates) (domain.SalesReport, error) {
	var SalesReport domain.SalesReport
	endDate := dates.EndDate
	startDate := dates.StartingDate
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

	for _, items := range bestSellerIds {

		product, err := orU.orderRepo.GetProductNameFromId(items)
		if err != nil {
			return domain.SalesReport{}, err
		}
		bestSellers = append(bestSellers, product)
	}
	SalesReport.BestSellers = bestSellers

	return SalesReport, nil
}

func (orU *orderUsecase) ReturnOrder(id int) error {

	status, err := orU.orderRepo.CheckOrderStatus(id)
	if err != nil {
		return err
	}
	if status == "RETURNED" {
		return errors.New("item already returned")
	}

	if status != "DELIVERED" {
		return errors.New("user is trying to return an item which is still not delivered")
	}

	// make order is return order
	if err := orU.orderRepo.ReturnOrder(id); err != nil {
		return err
	}
	// find amount to return to user

	amount, err := orU.orderRepo.FindAmountFromOrderID(id)
	fmt.Println(amount)
	if err != nil {
		return err
	}
	// find the user

	userId, err := orU.orderRepo.FindUserIdFromOrderID(id)
	fmt.Println(userId)
	if err != nil {
		return err
	}
	// find if the user having a wallet
	walletId, err := orU.orderRepo.FindWalletIdFromUserID(userId)
	fmt.Println(walletId)
	if err != nil {
		return err
	}
	// if no wallet,create new wallet for user

	if walletId == 0 {
		walletId, err = orU.orderRepo.CreateNewWallet(userId)
		if err != nil {
			return err
		}
	}
	// credit the amount into user wallet
	if err := orU.orderRepo.CreditToUserWallet(amount, walletId); err != nil {
		return err
	}
	return nil
}

func (orU *orderUsecase) CancelOrder(id, orderid int) error {

	status, err := orU.orderRepo.CheckOrderStatus(orderid)
	if err != nil {
		return err
	}
	if status == "CANCELLED" {
		return errors.New("item already cancelled")
	}
	if status == "DELIVERED" {
		return errors.New("item delivered")
	}

	if status == "PENDING" || status == "SHIPPED" {
		if err := orU.orderRepo.CancelOrder(orderid); err != nil {
			return err
		}
	}
	// check if already payed

	paymentStatus, err := orU.orderRepo.CheckPaymentStatus(orderid)
	if err != nil {
		return err
	}
	if paymentStatus != "PAID" {
		return nil
	}

	// find amount
	amount, err := orU.orderRepo.FindAmountFromOrderID(id)
	fmt.Println(amount)
	if err != nil {
		return err
	}

	return nil
}
