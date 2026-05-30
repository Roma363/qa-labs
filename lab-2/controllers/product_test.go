package controllers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"go-rest-api/models" // Переконайся, що шлях відповідає твоєму go.mod

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	// Імпортуємо GORM та драйвер SQLite для створення тимчасової БД
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// setupTestDB створює тимчасову базу даних у пам'яті спеціально для тестів
func setupTestDB() {
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		panic("Failed to connect to test database!")
	}
	db.AutoMigrate(&models.Product{})
	models.DB = db // Підміняємо глобальну змінну з твого коду
}

func TestCreateProduct(t *testing.T) {
	// 1. Arrange (Підготовка) - Вимоги AAA паттерну
	gin.SetMode(gin.TestMode)
	setupTestDB()

	// Наша таблиця тестових сценаріїв
	tests := []struct {
		name           string // Назва тесту
		payload        string // Вхідний JSON
		expectedStatus int    // Очікуваний HTTP статус
	}{
		{
			name:           "Успішне створення продукту",
			payload:        `{"name": "MacBook Pro", "price": 2500}`,
			expectedStatus: http.StatusCreated,
		},
		{
			name:           "Помилка валідації (невірний тип price)",
			payload:        `{"name": "MacBook Pro", "price": "expensive"}`, // price має бути числом
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Помилка валідації (зламаний JSON)",
			payload:        `{name: "MacBook",}`, // Відсутні лапки
			expectedStatus: http.StatusBadRequest,
		},
	}

	// 2. Act (Дія) та 3. Assert (Перевірка)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Налаштовуємо роутер для кожного тесту
			r := gin.Default()
			r.POST("/products", CreateProduct)

			// Створюємо HTTP-запит з даними з нашої таблиці
			req, _ := http.NewRequest("POST", "/products", bytes.NewBuffer([]byte(tt.payload)))
			req.Header.Set("Content-Type", "application/json")

			// Recorder запише відповідь нашого контролера
			w := httptest.NewRecorder()

			// Виконуємо запит
			r.ServeHTTP(w, req)

			// Перевіряємо, чи збігається реальний статус-код з очікуваним
			assert.Equal(t, tt.expectedStatus, w.Code)
		})
	}
}
