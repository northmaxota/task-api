# Stage 1: Build
FROM golang:1.26-alpine AS builder

WORKDIR /app

# Установим зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь код
COPY . .

# Собираем бинарник
RUN go build -o task-api ./cmd/api/main.go

# Stage 2: Run
FROM alpine:3.18

WORKDIR /app

# Копируем бинарник из builder
COPY --from=builder /app/task-api .

# Приложение будет слушать порт 8080
EXPOSE 8080

# Запуск
CMD ["./task-api"]
