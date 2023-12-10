package http

import (
	handlers "github.com/Anandhu4456/go-Ecommerce/pkg/api/handlers"
	"github.com/Anandhu4456/go-Ecommerce/pkg/routes"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	routes.UserRoutes(engine.Group("/users"),userHandler,otpHandler,inventoryHandler,cartHandler,orderHandler,couponHandler,paymentHandler,wishlistHandler)
	routes.AdminRoutes(engine.Group("/admin"),adminHandler,categoryHandler,inventoryHandler,orderHandler,paymentHandler,offerHandler,couponHandler)
	routes.InventoryRoutes(engine.Group("/products"),inventoryHandler)

	return &ServerHTTP{
		engine: engine,
	}
}

func (sh *ServerHTTP)Start(){
	sh.engine.Run(":8080")
}