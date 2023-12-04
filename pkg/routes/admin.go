package routes

import (
	"github.com/Anandhu4456/go-Ecommerce/pkg/api/handlers"
	"github.com/Anandhu4456/go-Ecommerce/pkg/api/middleware"
	"github.com/gin-gonic/gin"
)

func AdminRoutes(engine *gin.RouterGroup,adminHandler *handlers.AdminHandler /*userHandler *handlers.UserHandler,*/,categoryHandler *handlers.CategoryHandler,inventoryHandler *handlers.InventoryHandler,orderHandler *handlers.OrderHandler,paymentHandler *handlers.PaymentHandler,offerHandler *handlers.OfferHandler,couponHandler *handlers.CouponHandler ){
	engine.POST("/adminlogin",adminHandler.LoginHandler)
	engine.Use(middleware.AdminAuthMiddleware)

	userManagement:=engine.Group("/users")
	{
		userManagement.POST("/block",adminHandler.BlockUser)
		userManagement.POST("/unblock",adminHandler.UnblockUser)
		userManagement.GET("/getusers",adminHandler.GetUsers)
	}
}