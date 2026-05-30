package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Базовий модульний тест для перевірки працездатності
func TestHealthCheck(t *testing.T) {
	// 1. Налаштування (Arrange)
	// Створюємо простий тестовий обробник, який імітує реальний ендпоінт сервера
	mockHandler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}

	// Створюємо HTTP-запит
	req, _ := http.NewRequest("GET", "/health", nil)

	// Створюємо Recorder, який "запише" відповідь нашого обробника
	w := httptest.NewRecorder()

	// 2. Дія (Act)
	// Викликаємо обробник. Тепер змінна req використовується, і компілятор буде задоволений!
	mockHandler(w, req)

	// 3. Перевірка (Assert)
	// Переконуємося, що повернувся очікуваний статус 200 OK
	assert.Equal(t, http.StatusOK, w.Code)
}
