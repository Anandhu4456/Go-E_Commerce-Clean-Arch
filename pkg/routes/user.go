package routes

import (
	"github.com/Anandhu4456/go-Ecommerce/pkg/api/handlers"
	"github.com/Anandhu4456/go-Ecommerce/pkg/api/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(engine *gin.RouterGroup,
	userHandler *handlers.UserHandler,
	otpHandler *handlers.OtpHandler,
	inventoryHandler *handlers.InventoryHandler,
	cartHandler *handlers.CartHandler,
	orderHandler *handlers.OrderHandler,
	couponHandler *handlers.CouponHandler,
	paymentHandler *handlers.PaymentHandler,
	wishlistHandler *handlers.WishlistHandler) {

	engine.POST("/signup", userHandler.SignUp)
	engine.POST("/login", userHandler.Login)
	engine.POST("/otplogin", otpHandler.SendOTP)
	engine.POST("/otpverify", otpHandler.VerifyOTP)

	// Auth middleware
	engine.Use(middleware.UserAuthMiddleware)
	{
		payment := engine.Group("/payment")
		{
			payment.GET("/razorpay", paymentHandler.MakePaymentRazorPay)
			payment.GET("/verify-status", paymentHandler.VerifyPayment)
		}

		home := engine.Group("/home")
		{
			home.POST("/add-to-cart", cartHandler.AddtoCart)
			home.POST("/add-to-wishlist", wishlistHandler.AddtoWishlist)
		}

		profile := engine.Group("/profile")
		{
			profile.GET("/details", userHandler.GetUserDetails)
			profile.GET("/address", userHandler.GetAddresses)
			profile.POST("/address/add", userHandler.AddAddress)
			profile.PATCH("/edit", userHandler.EditUser)

			security := profile.Group("/security")
			{
				security.PATCH("/change-password", userHandler.ChangePassword)
			}
			wallet := profile.Group("/wallet")
			{
				wallet.GET("", userHandler.GetWallet)
			}
			orders := profile.Group("/orders")
			{
				orders.GET("", orderHandler.GetOrders)
				orders.POST("/cancel-order", orderHandler.CancelOrder)
				orders.POST("return-order", orderHandler.ReturnOrder)
			}
		}

		cart := engine.Group("/cart")
		{
			cart.GET("", userHandler.GetCart)
			cart.PATCH("/update-quantity-plus", userHandler.UpdateQuantityAdd)
			cart.PATCH("/update-quantity-minus", userHandler.UpdateQuantityLess)
			cart.DELETE("/remove", userHandler.RemoveFromCart)
		}
		wishlist := engine.Group("/wishlist")
		{
			wishlist.GET("", wishlistHandler.GetWishlist)
			wishlist.DELETE("/remove", wishlistHandler.RemoveFromWishlist)
		}
		checkout := engine.Group("/checkout")
		{
			checkout.GET("", cartHandler.CheckOut)
			checkout.GET("/download-invoice", orderHandler.DownloadInvoice)
			checkout.POST("/order", orderHandler.OrderItemsFromCart)
		}
		engine.GET("/coupons", couponHandler.Coupons)
	}
}
