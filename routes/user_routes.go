// user_routes.go
package routes

import (
	handler "api-v2/handler"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	// r.GET("/users", handler.GetUsers)
	r.POST("/users_login", handler.ValidUser)
	r.POST("/users", handler.CreateUser)
	// r.PUT("/users/:id", handler.UpdateUser)
	// r.DELETE("/users/:id", handler.DeleteUser)
	// ...
}
