package app

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gin-practice/db"
	"gin-practice/users"
)

func InitRouter(router *gin.Engine) {
	router.GET("/health", health)

	apiV1 := router.Group("/api/v1")

	userController := users.NewUserController(
		users.NewUserService(
			users.NewUserRepository(db.GetDb()),
		),
	)

	users.AddUserRoutes(apiV1, userController)

	return
}

func health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
