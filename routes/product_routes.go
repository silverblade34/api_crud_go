// product_routes.go
package routes

import (
	handler "api-v2/handler"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.Engine) {
	r.GET("/products", handler.GetProducts)
	r.GET("/products/:id", handler.GetProduct)
	r.POST("/products", handler.CreateProduct)
	r.PUT("/products/:id", handler.UpdateProduct)
	r.DELETE("/products/:id", handler.DeleteProduct)
	// ...
}
