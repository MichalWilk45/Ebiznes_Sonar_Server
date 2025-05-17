package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	StockCount  int     `json:"stock_count"`
	CategoryID  uint    `json:"category_id"`
	Category    Category
	CartItems   []CartItem
}

// Scope'y dla zapytań GORM
func InStock(db *gorm.DB) *gorm.DB {
	return db.Where("stock_count > 0")
}

func PriceBelow(price float64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("price < ?", price)
	}
}

func ByCategory(categoryID uint) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("category_id = ?", categoryID)
	}
}
