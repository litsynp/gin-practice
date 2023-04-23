package main

import (
	"gin-practice/app"
	"gin-practice/db"
	"gin-practice/internal"
)

func main() {
	db.GetDb()
	r := app.InitGinApp()

	r.Run(":" + internal.PORT)
}
