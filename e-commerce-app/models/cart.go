package models

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserID    uint       `json:"user_id"`
	User      User       `gorm:"foreignKey:UserID"`
	CartItems []CartItem `json:"cart_items"`
}

// Scope'y dla zapytań GORM
func WithCartItems(db *gorm.DB) *gorm.DB {
	return db.Preload("CartItems").Preload("CartItems.Product")
}
