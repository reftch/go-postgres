package services

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/reftch/go-postgres/models" // Replace with your actual package path
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(dsn string) (*UserService, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}

	// Auto migrate the schema
	if err := db.AutoMigrate(&models.User{}); err != nil {
		return nil, fmt.Errorf("failed to auto migrate schema: %w", err)
	}

	return &UserService{db: db}, nil
}

func (s *UserService) GetUsers(c echo.Context) error {
	var users []models.User
	result := s.db.Find(&users)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": result.Error.Error()})
	}
	return c.JSON(http.StatusOK, users)
}

func (s *UserService) CreateUser(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	result := s.db.Create(user)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": result.Error.Error()})
	}
	return c.JSON(http.StatusCreated, user)
}

func (s *UserService) GetUserByID(c echo.Context) error {
	id := c.Param("id")
	var user models.User
	result := s.db.First(&user, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}
	return c.JSON(http.StatusOK, user)
}

func (s *UserService) UpdateUser(c echo.Context) error {
	id := c.Param("id")
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	result := s.db.Model(&models.User{}).Where("id = ?", id).Updates(user)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": result.Error.Error()})
	}
	return c.JSON(http.StatusOK, user)
}

func (s *UserService) DeleteUser(c echo.Context) error {
	id := c.Param("id")
	result := s.db.Delete(&models.User{}, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": result.Error.Error()})
	}
	return c.JSON(http.StatusNoContent, nil)
}
