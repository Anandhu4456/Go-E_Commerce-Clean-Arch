package routes

import (
	"github.com/Anandhu4456/go-Ecommerce/pkg/api/handlers"
	"github.com/Anandhu4456/go-Ecommerce/pkg/api/middleware"
	"github.com/gin-gonic/gin"
)

func AdminRoutes(engine *gin.RouterGroup, adminHandler *handlers.AdminHandler /*userHandler *handlers.UserHandler,*/, categoryHandler *handlers.CategoryHandler, inventoryHandler *handlers.InventoryHandler, orderHandler *handlers.OrderHandler, paymentHandler *handlers.PaymentHandler, offerHandler *handlers.OfferHandler, couponHandler *handlers.CouponHandler) {
	engine.POST("/adminlogin", adminHandler.LoginHandler)
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
		inventoryManagement.GET("", inventoryHandler.ListProdutcs)
		inventoryManagement.GET("/details", inventoryHandler.ShowIndividualProducts)
		inventoryManagement.POST("/add", inventoryHandler.AddInventory)
		inventoryManagement.POST("/add-image", inventoryHandler.AddImage)
		inventoryManagement.PATCH("/update", inventoryHandler.UpdateInventory)
		inventoryManagement.PATCH("/update-image", inventoryHandler.UpdateImage)
		inventoryManagement.DELETE("/delete-image", inventoryHandler.DeleteImage)
		inventoryManagement.DELETE("/delete", inventoryHandler.DeleteInventory)
	}

}
