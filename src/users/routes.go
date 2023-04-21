package users

import (
	"database/sql"
	"github.com/gin-gonic/gin"
)

func AddUserRoutes(rg *gin.RouterGroup, db *sql.DB) {
	users := rg.Group("/users")

	users.POST("/", CreateUserAction(db))
	users.GET("/:id", FindUserByIdAction(db))
	users.PUT("/:id", UpdateUserAction(db))
	users.DELETE("/:id", DeleteUserAction(db))
}
