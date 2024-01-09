package main

import (
	"TChat/pkg/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode("debug")
	router := controllers.InitRouter()
	if err := router.Run(":8080"); err != nil {
		return
	}
}
