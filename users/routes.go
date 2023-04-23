package users

import (
	"github.com/gin-gonic/gin"
)

func AddUserRoutes(rg *gin.RouterGroup, userController *UserController) {
	users := rg.Group("/users")

	users.POST("/", userController.CreateUserAction)
	users.GET("/:id", userController.FindUserByIdAction)
	users.PUT("/:id", userController.UpdateUserAction)
	users.DELETE("/:id", userController.DeleteUserAction)
}
