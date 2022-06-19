package main

import (
	"github.com/gin-gonic/gin"
	"windolph.org/teve/db"
)

func main() {

	connection := db.Connect()
	db.CheckMigration(connection)

	r := gin.Default()
	r.GET("event", func(ctx *gin.Context) {
		ctx.Status(200)
	})
	r.Run()

}
