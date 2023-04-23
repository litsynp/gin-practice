package app

import "github.com/gin-gonic/gin"

func InitGinApp() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())

	InitRouter(r)

	return r
}
