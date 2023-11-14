package usecase

import (
	"fmt"

	"github.com/Anandhu4456/go-Ecommerce/pkg/domain"
	pdf"github.com/Anandhu4456/go-Ecommerce/pkg/helper/pdf"
	interfaces"github.com/Anandhu4456/go-Ecommerce/pkg/repository/interfaces"
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

func (orU *orderUsecase) OrderItemsFromCart(userid int, order models.Order, coupon string) (string, error){

	cart,err:=orU.userUsecase.GetCart(userid,0,0)
	if err!=nil{
		return "",err
	}
	var total float64

	for _,items:=range cart{
		total= total+items.DiscoundPrice
	}
	fmt.Println("total without coupon: ",total)

	if coupon!=""{
		valid,err:=orU.couponRepo.ValidateCoupon(coupon)
		if err!=nil || !valid{
			return "",err
		}

		// Find discount 
		discount:=orU.couponRepo.FindCouponDiscount(coupon)
		
		if discount >0{
			totalDiscount:=total *float64(discount)/100
			fmt.Println("Discount: ",discount,"Total Discount: ",totalDiscount,(discount/100),int(total),int(discount/100))
			total= total-totalDiscount
		}
	}
	fmt.Println("Total amount: ",total)
	var invoiceItems []*pdf.InvoiceData

	for _,items:=range cart{
		inventory,err:=pdf.NewInvoiceData(items.ProductName,int(items.Quantity),(items.DiscoundPrice))
		if err!=nil{
			panic(err)
		}
		invoiceItems = append(invoiceItems, inventory)
	}
}
