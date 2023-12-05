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
	paymentHandler *handlers.PaymentHandler,
	wishlistHandler *handlers.WishlistHandler) {

	engine.POST("/signup", userHandler.SignUp)
	engine.POST("/login", userHandler.Login)
	engine.POST("/otplogin", otpHandler.SendOTP)
	engine.POST("/otpverify", otpHandler.VerifyOTP)

	// Auth middleware
	engine.Use(middleware.UserAuthMiddleware)
	{
		payment:=engine.Group("/payment")
		{
			payment.GET("/razorpay",paymentHandler.MakePamentRazorPay)
			payment.GET("/verify-status",paymentHandler.VerifyPayment)
		}
		
	}
}
