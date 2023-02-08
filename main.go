package main

import (
	routes "api-v2/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.UserRoutes(r)
	routes.ProductRoutes(r)

	r.Run(":8080")
}
