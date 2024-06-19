package main

import (
	"go_api/internal/config"
	"go_api/internal/controller"
	"go_api/internal/repository"
	"go_api/internal/service"

	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := config.SetupDatabaseConnection()

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	r.GET("/users", userController.GetUsers)
	r.POST("/users", userController.CreateUser)

	// Start the server and handle errors
	if err := r.Run(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
