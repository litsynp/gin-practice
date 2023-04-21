package main

import (
	"database/sql"
	"gin-practice/src/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	db := utils.InitDb()
	r := InitGinApp(db.DB)

	r.Run(":" + utils.PORT)
}

func InitGinApp(db *sql.DB) *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())

	InitRouter(r, db)

	return r
}
