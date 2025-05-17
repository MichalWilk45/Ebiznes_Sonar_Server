package controllers

import (
	"net/http"
	"strconv"

	"e-commerce-app/config"
	"e-commerce-app/models"

	"github.com/labstack/echo/v4"
)

// Pobierz wszystkie produkty
func GetProducts(c echo.Context) error {
	var products []models.Product
	result := config.DB.Find(&products)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Nie udało się pobrać produktów"})
	}
	return c.JSON(http.StatusOK, products)
}

// Pobierz produkt po ID
func GetProduct(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Nieprawidłowe ID"})
	}

	var product models.Product
	result := config.DB.First(&product, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Produkt nie znaleziony"})
	}

	return c.JSON(http.StatusOK, product)
}

// Dodaj nowy produkt
func CreateProduct(c echo.Context) error {
	product := new(models.Product)
	if err := c.Bind(product); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Nieprawidłowe dane"})
	}

	result := config.DB.Create(&product)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Nie udało się utworzyć produktu"})
	}

	return c.JSON(http.StatusCreated, product)
}

// Zaktualizuj produkt
func UpdateProduct(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Nieprawidłowe ID"})
	}

	var product models.Product
	result := config.DB.First(&product, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Produkt nie znaleziony"})
	}

	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Nieprawidłowe dane"})
	}

	config.DB.Save(&product)
	return c.JSON(http.StatusOK, product)
}

// Usuń produkt
func DeleteProduct(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Nieprawidłowe ID"})
	}

	var product models.Product
	result := config.DB.First(&product, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Produkt nie znaleziony"})
	}

	config.DB.Delete(&product)
	return c.JSON(http.StatusOK, map[string]string{"message": "Produkt usunięty"})
}

// Użycie scope'ów
func GetProductsInStock(c echo.Context) error {
	var products []models.Product
	result := config.DB.Scopes(models.InStock).Find(&products)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Nie udało się pobrać produktów"})
	}
	return c.JSON(http.StatusOK, products)
}

func GetProductsByCategory(c echo.Context) error {
	categoryID, err := strconv.Atoi(c.Param("category_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Nieprawidłowe ID kategorii"})
	}

	var products []models.Product
	result := config.DB.Scopes(models.ByCategory(uint(categoryID))).Find(&products)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Nie udało się pobrać produktów"})
	}
	return c.JSON(http.StatusOK, products)
}
