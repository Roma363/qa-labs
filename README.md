# Go REST API Project

## Про проект

Цей проект базується на boilerplate-коді, клонованому з репозиторію [gin-boilerplate](https://github.com/Massad/gin-boilerplate.git).

**Важливо:** Оригінальний проект не мав тестів. На нього було накладено комплексну систему тестування, включаючи:
- **Модульні тести (Unit Tests)** - тестування окремих компонентів
- **Інтеграційні тести (Integration Tests)** - тестування взаємодії компонентів з базою даних
- **Контрактні тести (Contract Tests)** - тестування контрактів між сервісами

## Стек технологій

- **Go 1.25.0** - мова програмування
- **Gin Web Framework** - фреймворк для створення REST API
- **GORM** - ORM для роботи з базою даних
- **PostgreSQL** - база даних
- **Testify** - фреймворк для написання тестів
- **Testcontainers** - использование Docker контейнерів для тестування

## Структура проекту

```
project/
├── main.go                   # Точка входу додатку
├── go.mod / go.sum          # Файли залежностей
├── config/                   # Конфігурація додатку
├── controllers/              # Контролери (обробка запитів)
├── middlewares/              # Middleware функції
├── models/                   # Моделі даних і робота з БД
├── routes/                   # Маршрути API
├── utils/                    # Утиліти (helper функції)
├── *_test.go                 # Тести
└── coverage/                 # Звіти покриття коду
```

## Встановлення та запуск

### Встановлення залежностей

```bash
go mod download
```

### Запуск додатку

```bash
go run main.go
```

### Запуск тестів

Детальну інформацію про запуск тестів дивіться у файлі [TESTS_README.md](TESTS_README.md).

#### Модульні тести
```bash
go test -v ./controllers
```

#### Інтеграційні тести
```bash
go test -v integration_test.go
```

#### Контрактні тести
```bash
go test -v contract_test.go
```

#### Все тести одночасно
```bash
go test -v ./...
```

### Покриття коду

```bash
go test -coverprofile coverage.out ./controllers
go tool cover -html=coverage.out
```

## Основні кінцеві точки (Endpoints)

API надає функціональність для управління товарами (Products):
- **GET** `/products` - отримання списку товарів
- **GET** `/products/:id` - отримання товару за ID
- **POST** `/products` - створення нового товару
- **PUT** `/products/:id` - оновлення товару
- **DELETE** `/products/:id` - видалення товару

## Ліцензія

Проект поширюється під ліцензією, яка зазначена у файлі [LICENSE](LICENSE).

## Автор

Базований на [gin-boilerplate](https://github.com/Massad/gin-boilerplate.git)

---

**Примітка:** Перший запуск інтеграційних тестів може тривати довше через завантаження Docker-образу (postgres:15-alpine).

