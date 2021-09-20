package main

import (
	"github.com/gin-gonic/gin"

	"assignment2/config"
	"assignment2/controllers"
)

func main() {

	config.InitDB()

	r := gin.Default()

	r.GET("/orders", controllers.Read)
	r.POST("/orders", controllers.Create)
	r.PUT("/orders", controllers.Update)
	r.DELETE("/orders/:id", controllers.Delete)

	port := "3000"
	r.Run(":" + port)
}
