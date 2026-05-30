## Запуск тестів

### 1. Модульні тести (Unit Tests)

```bash
go test -v ./controllers

```

### 2. Вузькі інтеграційні тести


```bash
go test -v integration_test.go

```

*Примітка: Перший запуск може тривати довше через завантаження Docker-образу (postgres:15-alpine).*

### 3. Контрактні тести


```bash
go test -v contract_test.go

```

### Запуск усіх тестів одночасно:

```bash
go test -v ./...

```

---

## Збір метрик покриття коду (Coverage)


1. Очистіть кеш тестів (рекомендовано для запобігання конфліктів):

```bash
go clean -testcache

```

2. Зберіть статистику для папки з контролерами:

```bash
go test -coverprofile coverage.out ./controllers

```

3. Перегляньте результати по кожній функції прямо у терміналі:

```bash
go tool cover -func coverage.out

```

4. Відкрийте детальний візуальний звіт у браузері:

```bash
go tool cover -html coverage.out

```