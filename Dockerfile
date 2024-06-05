# Базовый образ.
FROM golang:1.21

# Задаем рабочую директорию в контейнере.
WORKDIR /app

# Копируем go mod и sum файлы.
COPY go.mod go.sum ./

# Загружаем зависимости.
RUN go mod download

# Копируем исходный код в рабочую директорию контейнера.
COPY . .

# Собираем приложение.
RUN go build -o main ./cmd/main.go

# Запускаем приложение.
CMD ["./main"]