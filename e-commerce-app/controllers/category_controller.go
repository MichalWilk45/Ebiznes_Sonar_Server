package controllers

import (
	"net/http"
	"strconv"

	"e-commerce-app/config"
	"e-commerce-app/models"

	"github.com/labstack/echo/v4"
)

// Pobierz wszystkie kategorie
func GetCategories(c echo.Context) error {
	var categories []models.Category
	result := config.DB.Find(&categories)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Nie udało się pobrać kategorii"})
	}
	return c.JSON(http.StatusOK, categories)
}

// Pobierz kategorię po ID
func GetCategory(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Nieprawidłowe ID"})
	}

	var category models.Category
	result := config.DB.First(&category, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Kategoria nie znaleziona"})
	}

	return c.JSON(http.StatusOK, category)
}

// Dodaj nową kategorię
func CreateCategory(c echo.Context) error {
	category := new(models.Category)
	if err := c.Bind(category); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Nieprawidłowe dane"})
	}

	result := config.DB.Create(&category)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Nie udało się utworzyć kategorii"})
	}

	return c.JSON(http.StatusCreated, category)
}

// Pobierz kategorię wraz z produktami
func GetCategoryWithProducts(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Nieprawidłowe ID"})
	}

	var category models.Category
	result := config.DB.Scopes(models.WithProducts).First(&category, id)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Kategoria nie znaleziona"})
	}

	return c.JSON(http.StatusOK, category)
}
