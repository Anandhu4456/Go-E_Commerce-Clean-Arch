package routes

import (
	"github.com/Anandhu4456/go-Ecommerce/pkg/api/handlers"
	"github.com/gin-gonic/gin"
)

// this function is for use this routes for all users without login

func InventoryRoutes(engine *gin.RouterGroup, inventoryHandler *handlers.InventoryHandler) {
	engine.GET("", inventoryHandler.ListProdutcs)
	engine.GET("/details", inventoryHandler.ShowIndividualProducts)
	engine.GET("/search", inventoryHandler.SearchProducts)
	engine.GET("/category", inventoryHandler.GetCategoryProducts)
}
