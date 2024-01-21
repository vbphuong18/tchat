package controllers

import (
	"TChat/pkg/middleware"
	"TChat/pkg/repositories"
	"TChat/pkg/services"
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
	v := validator.New()
	authenRepository := repositories.NewAuthenRepository(db)
	authenService := services.NewAuthenService(authenRepository)
	authenHandler := NewAuthenHandler(authenService, v)

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userHandler := NewUserHandler(userService, v)

	friendRepository := repositories.NewFriendRepository(db)
	friendService := services.NewFriendService(friendRepository)
	friendHandler := NewFriendHandler(friendService, userService, v)

	messageRepository := repositories.NewMessageRepository(db)
	messageService := services.NewMessageService(messageRepository)
	messageHandler := NewMessageHandler(messageService, userService, v)

	groupRepository := repositories.NewGroupRepository(db)
	groupService := services.NewGroupService(groupRepository, userRepository)
	groupHandler := NewGroupHandler(groupService, userService, v)

	apiGroup := r.Group("/api")
	messageGroup := apiGroup.Group("/message")
	{
		messageGroup.Use(middleware.HTTPAuthentication)
		messageGroup.POST("/create", messageHandler.CreateMessage)
		messageGroup.GET("/list", messageHandler.ListMessage)
		messageGroup.DELETE("/delete", messageHandler.DeleteMessage)
	}
	userGroup := apiGroup.Group("/user")
	{
		userGroup.Use(middleware.HTTPAuthentication)
		userGroup.POST("/create", userHandler.CreateUser)
		userGroup.GET("/list", userHandler.ListUser)
		userGroup.GET("/search", userHandler.SearchUser)
		userGroup.GET("/get-user-by-user-id", userHandler.GetUserByUserID)
		userGroup.DELETE("/delete", userHandler.DeleteUser)
	}
	authenGroup := apiGroup.Group("/authen")
	{
		authenGroup.POST("/login", authenHandler.Login)
		authenGroup.POST("/register", userHandler.CreateUser)
	}
	friendGroup := apiGroup.Group("/friend")
	{
		friendGroup.Use(middleware.HTTPAuthentication)
		friendGroup.POST("/add", friendHandler.CreateFriend)
		friendGroup.GET("/list", friendHandler.ListFriend)
		friendGroup.DELETE("/delete", friendHandler.DeleteFriend)
	}
	groupGroup := apiGroup.Group("/group")
	{
		groupGroup.POST("/create", groupHandler.CreateGroup)
	}
	return r
}
