package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Products    []Product
}

// Scope'y dla zapytań GORM
func WithProducts(db *gorm.DB) *gorm.DB {
	return db.Preload("Products")
}
