package main

import (
	"e-commerce-app/config"
	"e-commerce-app/models"
	"e-commerce-app/routes"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main(){
	// Inicjalizacja Echo
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Połączenie z bazą danych
	config.ConnectDatabase()

	// Auto-migracja dla modeli
	db := config.DB
	err := db.AutoMigrate(&models.Product{}, &models.Category{}, &models.User{}, &models.Cart{}, &models.CartItem{})
	if err != nil {
    	log.Fatalf("AutoMigrate failed: %v", err)
	}


	// Konfiguracja tras
	routes.SetupRoutes(e)

	// Start serwera
	e.Logger.Fatal(e.Start(":8080"))
}