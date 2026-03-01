# Task API — Go REST API с PostgreSQL

![Go](https://img.shields.io/badge/Go-1.26-blue) ![PostgreSQL](https://img.shields.io/badge/PostgreSQL-15-blue) ![Docker](https://img.shields.io/badge/Docker-Compose-blue)

**Task API** — это REST API для управления списком задач, написанное на Go с использованием архитектуры **clean/hexagonal** и интеграцией с PostgreSQL.  

Проект полностью dockerized, с конфигурацией через environment variables, логгированием и поддержкой graceful shutdown.

---

## ⚙️ Особенности

- Чистая архитектура с разделением на слои:
  - **Handler (Transport layer)** — HTTP эндпоинты
  - **Service (Business logic)** — валидация, правила и обработка задач
  - **Storage (Persistence layer)** — InMemory или PostgreSQL
- Поддержка **context.Context** для отмены запросов и таймаутов
- Логгирование HTTP-запросов и сервисных вызовов
- Graceful shutdown сервера
- Docker + docker-compose для простого поднятия локальной среды
- Unit и integration тесты с table-driven подходом
- Легко расширяется: новые storage, авторизация, swagger

---

## 🧩 Технологии

- [Go 1.26](https://golang.org/)
- [chi router](https://github.com/go-chi/chi)
- [PostgreSQL 15](https://www.postgresql.org/)
- Docker / Docker Compose
- context.Context для graceful shutdown и таймаутов
- Unit / integration testing через `testing` пакет

---

## 🚀 Быстрый старт

### 1. Клонируем репозиторий

```bash
git clone https://github.com/northmaxota/task-api.git && cd task-api
```
### 2. Билдим через Docker Compose
```bash
docker compose up --build
```
API доступно через http://localhost:8080
PostgreSQL на localhost:5432 (taskuser/taskpass, база taskdb)
порты и верификация базы меняется в docker-compose.yml, как и смена DSN самой базы

### 3. Работа с task-api (примеры)
Создание новой задачи
```bash
curl -X POST localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"title":"Example Task"}'
```
Получение списка задач
```bash
curl localhost:8080/tasks
```

### 4. Тестирование
```bash
go test ./internal/task/ -v
```


