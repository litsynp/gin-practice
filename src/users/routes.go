package users

import (
	"github.com/gin-gonic/gin"
)

func AddUserRoutes(rg *gin.RouterGroup) {
	users := rg.Group("/users")

	users.POST("/", CreateUserAction)
	users.GET("/:id", FindUserByIdAction)
	users.PUT("/:id", UpdateUserAction)
	users.DELETE("/:id", DeleteUserAction)
}
