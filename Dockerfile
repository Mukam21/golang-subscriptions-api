FROM golang:1.24-alpine

WORKDIR /app

# Устанавливаем bash, git и curl (если нужен wait-for-it)
RUN apk add --no-cache bash git curl

# Копируем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходники
COPY . .

# Качаем скрипт ожидания БД
ADD https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh /wait-for-it.sh
RUN chmod +x /wait-for-it.sh

# Собираем Go-приложение
RUN go build -o main ./cmd/main.go

# Запуск через wait-for-it
CMD ["/wait-for-it.sh", "db:5432", "--", "bash", "-c", "./main"]
