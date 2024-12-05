package models

// User model
type User struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `json:"name" binding:"required"`
	Age  int    `json:"age" binding:"gte=0,lte=130"`
}
