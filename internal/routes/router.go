package routes

import (
	"controleEstoque/internal/products"

	"github.com/gin-gonic/gin"
)

func SetupRouter(handler *products.ProductHandler) *gin.Engine {
	r := gin.Default()

	productsGroup := r.Group("/products")
	{
		productsGroup.POST("", handler.Create)
		productsGroup.GET("", handler.GetAll)
		productsGroup.GET("/:id", handler.GetByID)
		productsGroup.PUT("/:id", handler.Update)
		productsGroup.DELETE("/:id", handler.Delete)
	}

	return r
}
