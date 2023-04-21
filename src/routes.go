package main

import (
	"database/sql"
	"gin-practice/src/users"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter(
	router *gin.Engine,
	db *sql.DB,
) {
	router.GET("/health", health)

	apiV1 := router.Group("/api/v1")
	users.AddUserRoutes(apiV1, db)

	return
}

func health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
