package usecase

import (
	"errors"
	"fmt"
	"strconv"

	interfaces "github.com/Anandhu4456/go-Ecommerce/pkg/repository/interfaces"
	"github.com/Anandhu4456/go-Ecommerce/pkg/utils/models"
	"github.com/razorpay/razorpay-go"
)

type paymentUsecase struct {
	paymentRepo interfaces.PaymentRepository
	userRepo    interfaces.UserRepository
}

// Constructor function
func NewPaymentUsecase(paymentRepo interfaces.PaymentRepository, userRepo interfaces.UserRepository) *paymentUsecase {
	return &paymentUsecase{
		paymentRepo: paymentRepo,
		userRepo:    userRepo,
	}
}

func (payU *paymentUsecase) AddNewPaymentMethod(paymentMethod string) error {
	if paymentMethod == "" {
		return errors.New("enter payment method")
	}
	if err := payU.paymentRepo.AddNewPaymentMethod(paymentMethod); err != nil {
		return err
	}
	return nil
}

func (payU *paymentUsecase) RemovePaymentMethod(paymentMethodID int) error {
	if paymentMethodID == 0 {
		return errors.New("enter method id")
	}
	if err := payU.paymentRepo.RemovePaymentMethod(paymentMethodID); err != nil {
		return err
	}
	return nil
}

func (payU *paymentUsecase) GetPaymentMethods() ([]models.PaymentMethod, error) {
	paymentMethods, err := payU.paymentRepo.GetPaymentMethods()
	if err != nil {
		return []models.PaymentMethod{}, err
	}
	return paymentMethods, nil
}

func (payU *paymentUsecase) MakePaymentRazorPay(orderID string, userID int) (models.OrderPaymentDetails, error) {
	var orderDetails models.OrderPaymentDetails

	// Get order id
	orderId, err := strconv.Atoi(orderID)
	if err != nil {
		return models.OrderPaymentDetails{}, err
	}
	orderDetails.OrderID = orderId
	orderDetails.UserID = userID

	// Get username
	username, err := payU.paymentRepo.FindUsername(userID)
	if err != nil {
		return models.OrderPaymentDetails{}, err
	}
	orderDetails.Username = username

	// Get total
	total, err := payU.paymentRepo.FindPrice(orderId)
	if err != nil {
		return models.OrderPaymentDetails{}, err
	}
	orderDetails.FinalPrice = total

	// need to add key and secret
	client := razorpay.NewClient("key", "secret")

	data := map[string]interface{}{
		"amount":   int(orderDetails.FinalPrice) * 100,
		"currency": "INR",
		"receipt":  "some receipt id",
	}

	fmt.Println("razorpay::91", orderDetails, data)

	body, err := client.Order.Create(data, nil)
	if err != nil {
		fmt.Println(err)
		return models.OrderPaymentDetails{}, err
	}
	razorpayOrderId := body["id"].(int)
	orderDetails.RazorID = razorpayOrderId

	fmt.Println("razorpay::100", orderDetails)

	return orderDetails, nil
}

func (payU *paymentUsecase) VerifyPayment(paymentID string, razorID string, orderID string) error {

	if err := payU.paymentRepo.UpdatePaymentDetails(orderID, paymentID, razorID); err != nil {
		return err
	}

	// Clear cart
	// orderIdInt, err := strconv.Atoi(orderID)
	// if err != nil {
	// 	return err
	// }

	// userId, err := payU.userRepo.FindUserIDByOrderID(orderIdInt)
	// if err != nil {
	// 	return err
	// }
	// cartId, err := payU.userRepo.GetCartID(userId)
	// if err != nil {
	// 	return err
	// }
	// if err := payU.userRepo.ClearCart(cartId); err != nil {
	// 	return err
	// }
	return nil
}
