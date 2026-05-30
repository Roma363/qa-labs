package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"go-rest-api/controllers"
	"go-rest-api/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/stretchr/testify/assert"
)

// Ініціалізація БД для перевірки контрактів
func setupContractTestDB() {
	db, _ := gorm.Open("sqlite3", ":memory:")
	db.AutoMigrate(&models.Product{})
	models.DB = db
}

// Процес 1: Контракт на отримання списку продуктів
func TestContract_GetProducts(t *testing.T) {
	gin.SetMode(gin.TestMode)
	setupContractTestDB()
	r := gin.Default()
	r.GET("/products", controllers.FindProducts)

	req, _ := http.NewRequest("GET", "/products", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Перевіряємо контракт HTTP-рівня
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
}

// Процес 2: Контракт на створення продукту
func TestContract_CreateProduct(t *testing.T) {
	gin.SetMode(gin.TestMode)
	setupContractTestDB()
	r := gin.Default()
	r.POST("/products", controllers.CreateProduct)

	payload := []byte(`{"name":"Contract Test Item","price":1500}`)
	req, _ := http.NewRequest("POST", "/products", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Перевіряємо контракт HTTP-рівня
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))

	// Перевіряємо контракт JSON-структури
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	// Жорстка перевірка наявності ключів у відповіді (The Contract)
	assert.Contains(t, response, "ID", "Контракт порушено: відсутнє поле ID")
	assert.Contains(t, response, "name", "Контракт порушено: відсутнє поле Name")
	assert.Contains(t, response, "price", "Контракт порушено: відсутнє поле Price")
}
