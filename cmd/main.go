package main

import (
	"TChat/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode("debug")
	router := controllers.InitRouter()
	if err := router.Run(":8080"); err != nil {
		return
	}
}
