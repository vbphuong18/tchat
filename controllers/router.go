package controllers

import (
	"TChat/repositories"
	"TChat/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	dsn := "root:dt@tcp(127.0.0.1:3306)/tchat?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err) // thoát luôn rồi không cần phải return nữa
	}
	messageRepository := repositories.NewMessageRepository(db)
	messageService := services.NewMessageService(messageRepository)
	messageHandler := NewMessageHandler(messageService)

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	v := validator.New()
	userHandler := NewUserHandler(userService, v)

	authenRepository := repositories.NewAuthenRepository(db)
	authenService := services.NewAuthenService(authenRepository)
	authenHandler := NewAuthenHandler(authenService, v)

	apiGroup := r.Group("/api")
	messageGroup := apiGroup.Group("/message")
	{
		messageGroup.POST("/create", messageHandler.CreateMessage)
		messageGroup.GET("/list", messageHandler.ListMessage)
		messageGroup.DELETE("/delete", messageHandler.DeleteMessage)
	}
	userGroup := apiGroup.Group("/user")
	{
		userGroup.POST("/create", userHandler.CreateUser)
		userGroup.GET("/list", userHandler.ListUser)
		userGroup.GET("/search", userHandler.SearchUser)
		userGroup.DELETE("/delete", userHandler.DeleteUser)
	}
	authenGroup := apiGroup.Group("/authen")
	{
		authenGroup.POST("/login", authenHandler.Login)
		authenGroup.POST("/register", userHandler.CreateUser)
	}
	return r
}
