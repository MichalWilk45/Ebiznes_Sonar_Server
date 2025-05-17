package controllers

import (
	"net/http"
	"strconv"

	"e-commerce-app/config"
	"e-commerce-app/models"

	"github.com/labstack/echo/v4"
)

// Stałe z komunikatami błędów i sukcesów
const (
	ErrFetchProducts        = "Nie udało się pobrać produktów"
	ErrInvalidID            = "Nieprawidłowe ID"
	ErrProductNotFound      = "Produkt nie znaleziony"
	ErrInvalidData          = "Nieprawidłowe dane"
	ErrCreateProduct        = "Nie udało się utworzyć produktu"
	MsgProductDeleted       = "Produkt usunięty"
	ErrInvalidCategoryID    = "Nieprawidłowe ID kategorii"
)

func GetProducts(c echo.Context) error {
	var products []models.Product
	result := config.DB.Find(&products)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": ErrFetchProducts})
	}
	return c.JSON(http.StatusOK, products)
}

func GetProduct(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": ErrInvalidID})
	}

	var product models.Product
	result := config.DB.First(&product, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": ErrProductNotFound})
	}

	return c.JSON(http.StatusOK, product)
}

func CreateProduct(c echo.Context) error {
	product := new(models.Product)
	if err := c.Bind(product); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": ErrInvalidData})
	}

	result := config.DB.Create(&product)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": ErrCreateProduct})
	}

	return c.JSON(http.StatusCreated, product)
}

func UpdateProduct(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": ErrInvalidID})
	}

	var product models.Product
	result := config.DB.First(&product, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": ErrProductNotFound})
	}

	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": ErrInvalidData})
	}

	config.DB.Save(&product)
	return c.JSON(http.StatusOK, product)
}

func DeleteProduct(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": ErrInvalidID})
	}

	var product models.Product
	result := config.DB.First(&product, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": ErrProductNotFound})
	}

	config.DB.Delete(&product)
	return c.JSON(http.StatusOK, map[string]string{"message": MsgProductDeleted})
}

func GetProductsInStock(c echo.Context) error {
	var products []models.Product
	result := config.DB.Scopes(models.InStock).Find(&products)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": ErrFetchProducts})
	}
	return c.JSON(http.StatusOK, products)
}

func GetProductsByCategory(c echo.Context) error {
	categoryID, err := strconv.Atoi(c.Param("category_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": ErrInvalidCategoryID})
	}

	var products []models.Product
	result := config.DB.Scopes(models.ByCategory(uint(categoryID))).Find(&products)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": ErrFetchProducts})
	}
	return c.JSON(http.StatusOK, products)
}
