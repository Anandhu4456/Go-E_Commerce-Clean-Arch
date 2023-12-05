package routes

import (
	"github.com/Anandhu4456/go-Ecommerce/pkg/api/handlers"
	"github.com/Anandhu4456/go-Ecommerce/pkg/api/middleware"
	"github.com/gin-gonic/gin"
)

func AdminRoutes(engine *gin.RouterGroup,
	adminHandler *handlers.AdminHandler,
	// userHandler *handlers.UserHandler,
	categoryHandler *handlers.CategoryHandler,
	inventoryHandler *handlers.InventoryHandler,
	orderHandler *handlers.OrderHandler,
	paymentHandler *handlers.PaymentHandler,
	offerHandler *handlers.OfferHandler,
	couponHandler *handlers.CouponHandler) {
	engine.POST("/adminlogin", adminHandler.LoginHandler)
	{
		engine.Use(middleware.AdminAuthMiddleware)

		userManagement := engine.Group("/users")
		{
			userManagement.POST("/block", adminHandler.BlockUser)
			userManagement.POST("/unblock", adminHandler.UnblockUser)
			userManagement.GET("/getusers", adminHandler.GetUsers)
		}

		categoryManagement := engine.Group("/category")
		{
			categoryManagement.GET("/categories", categoryHandler.Categories)
			categoryManagement.POST("/add", categoryHandler.AddCategory)
			categoryManagement.PATCH("/update", categoryHandler.UpdateCategory)
			categoryManagement.DELETE("/delete", categoryHandler.DeleteCategory)
		}

		inventoryManagement := engine.Group("/inventories")
		{
			// inventoryManagement.GET("", inventoryHandler.ListProdutcs)
			// inventoryManagement.GET("/details", inventoryHandler.ShowIndividualProducts)
			inventoryManagement.POST("/add", inventoryHandler.AddInventory)
			inventoryManagement.POST("/add-image", inventoryHandler.AddImage)
			inventoryManagement.PATCH("/update", inventoryHandler.UpdateInventory)
			inventoryManagement.PATCH("/update-image", inventoryHandler.UpdateImage)
			inventoryManagement.DELETE("/delete-image", inventoryHandler.DeleteImage)
			inventoryManagement.DELETE("/delete", inventoryHandler.DeleteInventory)
		}

		orderManagement := engine.Group("/orders")
		{
			orderManagement.GET("", orderHandler.AdminOrders)
			orderManagement.PATCH("/edit/status", orderHandler.EditOrderStatus)
			orderManagement.PATCH("/edit/mark-as-paid", orderHandler.MarkAsPaid)
		}

		paymentManangement := engine.Group("/paymentmethods")
		{
			paymentManangement.GET("/", paymentHandler.GetPaymentMethods)
			paymentManangement.POST("/add", paymentHandler.AddNewPaymentMethod)
			paymentManangement.DELETE("/delete", paymentHandler.RemovePaymentMethod)
		}

		offerManagement := engine.Group("/offers")
		{
			offerManagement.GET("", offerHandler.Offers)
			offerManagement.POST("/create", offerHandler.AddOffer)
			offerManagement.POST("/expire", offerHandler.ExpireValidity)
		}

		couponManagement := engine.Group("/coupons")
		{
			couponManagement.GET("", couponHandler.Coupons)
			couponManagement.POST("/create", couponHandler.CreateNewCoupon)
			couponManagement.POST("/expire", couponHandler.MakeCouponInvalid)
		}
		salesManagement := engine.Group("/sales")
		{
			salesManagement.GET("/daily", orderHandler.AdminSalesDailyReport)
			salesManagement.GET("/weekly", orderHandler.AdminSalesWeeklyReports)
			salesManagement.GET("/monthly", orderHandler.AdminSalesMonthlyReport)
			salesManagement.GET("/annual", orderHandler.AdminSalesAnnualReport)
			salesManagement.POST("/custom", orderHandler.AdminSaleCustomReport)
		}
		productsManagement := engine.Group("/products")
		{
			productsManagement.GET("", inventoryHandler.AdminListProdutcs)
			productsManagement.GET("/details", inventoryHandler.ShowIndividualProducts)
			productsManagement.GET("/search", inventoryHandler.SearchProducts)
			productsManagement.GET("/category", inventoryHandler.GetCategoryProducts)
		}
	}
}
