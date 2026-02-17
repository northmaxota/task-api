# Task API (Go)

Minimalistic REST API for managing tasks.  
Built as a learning project with focus on clean architecture, dependency injection, decorators, and graceful shutdown.

---

## ✨ Features

- Create tasks
- List tasks
- In-memory storage
- Service layer with business rules
- Logging decorator
- Graceful shutdown
- Clean architecture (transport → service → storage)

---

## 🏗 Architecture

The project follows layered architecture:



### Key principles:

- **Transport layer** knows only about HTTP
- **Service layer** contains business logic
- **Storage layer** is responsible only for persistence
- Layers depend only downward
- Business logic is isolated from HTTP

---

## 🔌 Dependency Injection

Dependencies are assembled in `main.go`:

## Example of using in entry point of program

```go
storage := task.NewInMemoryStorage()
baseService := task.NewService(storage)
service := task.NewLoggingService(baseService)
handler := httpServer.NewHandler(service)
router = httpServer.NewRouter(handler)
server.Handler = router
```

---

## Feel free to reuse code from internal package by

**go get github.com/northmaxota/task-api**
