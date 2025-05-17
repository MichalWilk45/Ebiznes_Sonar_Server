package routes

import (
	"e-commerce-app/controllers"

	"github.com/labstack/echo/v4"
)

const productByID = "/products/:id"


func SetupRoutes(e *echo.Echo) {
	// Trasy dla produktów
	e.GET("/products", controllers.GetProducts)
	e.GET(productByID, controllers.GetProduct)
	e.POST("/products", controllers.CreateProduct)
	e.PUT(productByID, controllers.UpdateProduct)
	e.DELETE(productByID, controllers.DeleteProduct)
	e.GET("/products/in-stock", controllers.GetProductsInStock)
	e.GET("/products/category/:category_id", controllers.GetProductsByCategory)

	// Trasy dla kategorii
	e.GET("/categories", controllers.GetCategories)
	e.GET("/categories/:id", controllers.GetCategory)
	e.POST("/categories", controllers.CreateCategory)
	e.GET("/categories/:id/products", controllers.GetCategoryWithProducts)

	// Trasy dla koszyka
	e.GET("/carts/:user_id", controllers.GetCart)
	e.POST("/cart-items", controllers.AddToCart)
	e.DELETE("/cart-items/:id", controllers.RemoveFromCart)
}
