package controllers

import (
	"net/http"
	"strconv"

	"e-commerce-app/config"
	"e-commerce-app/models"

	"github.com/labstack/echo/v4"
)

// Pobierz koszyk użytkownika
func GetCart(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Nieprawidłowe ID użytkownika"})
	}

	var cart models.Cart
	result := config.DB.Where("user_id = ?", userID).Scopes(models.WithCartItems).First(&cart)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Koszyk nie znaleziony"})
	}

	return c.JSON(http.StatusOK, cart)
}

// Dodaj produkt do koszyka
func AddToCart(c echo.Context) error {
	cartItem := new(models.CartItem)
	if err := c.Bind(cartItem); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Nieprawidłowe dane"})
	}

	// Sprawdź, czy produkt istnieje
	var product models.Product
	if result := config.DB.First(&product, cartItem.ProductID); result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Produkt nie znaleziony"})
	}

	// Sprawdź, czy koszyk istnieje
	var cart models.Cart
	if result := config.DB.First(&cart, cartItem.CartID); result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Koszyk nie znaleziony"})
	}

	// Sprawdź, czy produkt już jest w koszyku
	var existingItem models.CartItem
	result := config.DB.Where("cart_id = ? AND product_id = ?", cartItem.CartID, cartItem.ProductID).First(&existingItem)
	if result.Error == nil {
		// Produkt już istnieje w koszyku, zaktualizuj ilość
		existingItem.Quantity += cartItem.Quantity
		config.DB.Save(&existingItem)
		return c.JSON(http.StatusOK, existingItem)
	}

	// Dodaj nowy produkt do koszyka
	result = config.DB.Create(&cartItem)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Nie udało się dodać produktu do koszyka"})
	}

	return c.JSON(http.StatusCreated, cartItem)
}

// Usuń produkt z koszyka
func RemoveFromCart(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Nieprawidłowe ID"})
	}

	var cartItem models.CartItem
	if result := config.DB.First(&cartItem, id); result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Element koszyka nie znaleziony"})
	}

	config.DB.Delete(&cartItem)
	return c.JSON(http.StatusOK, map[string]string{"message": "Produkt usunięty z koszyka"})
}
