package http

import (
	handlers"github.com/Anandhu4456/go-Ecommerce/pkg/api/handlers"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerfiles "github.com/swaggo/files"
	"github.com/gin-gonic/gin"
)

// http server for the web application
type ServerHTTP struct{
	engine *gin.Engine
}

// Constructor function

func NewServerHTTP(categoryHandler *handlers.CategoryHandler,inventoryHandler *handlers.InventoryHandler,userHandler *handlers.UserHandler,adminHandler *handlers.AdminHandler,otpHandler *handlers.OtpHandler,cartHandler *handlers.CartHandler,orderHandler *handlers.OrderHandler,paymentHandler *handlers.PaymentHandler,wishlistHandler *handlers.WishlistHandler,offerHandler *handlers.OfferHandler,couponHandler *handlers.CouponHandler)*ServerHTTP{
	engine:=gin.New()
	engine.Use(gin.Logger())
	engine.LoadHTMLGlob("pkg/templates/*.html")
	engine.GET("/swagger/*any",ginSwagger.WrapHandler(swaggerfiles.Handler))
}