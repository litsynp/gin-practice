package main

import (
	"gin-practice/src/users"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter(router *gin.Engine) {
	router.GET("/health", health)

	apiV1 := router.Group("/api/v1")
	users.AddUserRoutes(apiV1)

	return
}

func health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
