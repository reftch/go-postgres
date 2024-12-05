package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/reftch/go-postgres/services"
)

var userService *services.UserService

func main() {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	var err error
	userService, err := services.NewUserService(dsn)
	if err != nil {
		log.Fatalf("Failed to create user service: %v", err)
	}

	e := echo.New()

	// Define routes
	e.GET("/users", userService.GetUsers)
	e.POST("/users", userService.CreateUser)
	e.GET("/users/:id", userService.GetUserByID)
	e.PUT("/users/:id", userService.UpdateUser)
	e.DELETE("/users/:id", userService.DeleteUser)

	// Start the server
	log.Println("Server is running on :8080")
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
