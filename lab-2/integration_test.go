package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"go-rest-api/controllers"
	"go-rest-api/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestPostgresIntegration(t *testing.T) {
	ctx := context.Background()

	// Піднімаємо реальний контейнер з PostgreSQL
	postgresContainer, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("postgres:15-alpine"),
		postgres.WithDatabase("testdb"),
		postgres.WithUsername("user"),
		postgres.WithPassword("password"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second),
		),
	)
	if err != nil {
		t.Fatalf("Не вдалося запустити контейнер: %s", err)
	}

	// Гарантуємо, що контейнер видалиться після завершення тесту
	defer func() {
		if err := postgresContainer.Terminate(ctx); err != nil {
			t.Fatalf("Не вдалося зупинити контейнер: %s", err)
		}
	}()

	// Отримуємо рядок підключення до піднятої бази
	connStr, err := postgresContainer.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		t.Fatalf("Не вдалося отримати рядок підключення: %s", err)
	}

	// Підключаємо наш додаток до тестової бази
	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		t.Fatalf("Не вдалося підключитися до БД: %s", err)
	}
	db.AutoMigrate(&models.Product{})
	models.DB = db

	// 4. Виконуємо інтеграційний тест
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/products", controllers.FindProducts)

	req, _ := http.NewRequest("GET", "/products", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	// Перевіряємо, що запит до реальної БД пройшов успішно (повертається порожній масив, бо БД нова)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "[]", w.Body.String())
}
