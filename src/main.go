package main

import (
	"gin-practice/src/db"
	"gin-practice/src/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	db.GetDb()
	r := InitGinApp()

	r.Run(":" + utils.PORT)
}

func InitGinApp() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())

	InitRouter(r)

	return r
}
