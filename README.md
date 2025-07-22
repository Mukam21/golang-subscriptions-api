# Golang Subscriptions API

REST-сервис для управления онлайн-подписками пользователей.

---

## Описание проекта

Сервис позволяет выполнять CRUD операции с записями о подписках пользователей, а также подсчитывать суммарную стоимость подписок за заданный период с фильтрацией по пользователю и названию сервиса.

---

## Функционал

- Создание, чтение, обновление, удаление и просмотр списка подписок
- Подсчет суммарной стоимости подписок за указанный период с фильтрацией по `user_id` и `service_name`
- Хранение данных в PostgreSQL
- Логирование основных операций и ошибок
- Конфигурация через `.env` файл
- Автоматически генерируемая Swagger-документация
- Запуск приложения и БД через Docker Compose

---

## Технологии

- Go 1.24 (Alpine)
- Gin (HTTP-фреймворк)
- GORM (ORM для работы с PostgreSQL)
- PostgreSQL
- Swagger (swaggo/swag)
- Docker, Docker Compose
- Viper (конфигурация)
- Zap / стандартный логгер (логирование)

---

## Быстрый старт

### Клонировать репозиторий

```bash
git clone https://github.com/yourusername/golang-subscriptions-api.git

cd golang-subscriptions-api
```

### Настроить .env файл

Создайте файл .env в корне проекта на основе .env.example и укажите параметры подключения к БД

        DB_HOST=postgres-subscriptions
        DB_PORT=5432
        DB_USER=postgres
        DB_PASSWORD=yourpassword
        DB_NAME=subscriptions
        SERVER_PORT=8080

## Запуск через Docker Compose

        docker-compose up --build

Сервис будет доступен по адресу: http://localhost:8080

Swagger-документация: http://localhost:8080/swagger/index.html

API Endpoints

    POST	/subscriptions	Создать новую подписку
    GET	/subscriptions	Получить список всех подписок
    GET	/subscriptions/:id	Получить подписку по ID
    PUT	/subscriptions/:id	Обновить подписку по ID
    DELETE	/subscriptions/:id	Удалить подписку по ID
    GET	/subscriptions/total	Получить сумму подписок за период с фильтрами

## Пример запроса на создание подписки

        {
            "service_name": "Yandex Plus",
            "price": 400,
            "user_id": "60601fee-2bf1-4721-ae6f-7636e79a0cba",
            "start_date": "07-2025",
            "end_date": "12-2025"   // опционально
        }

## Миграции базы данных

    Миграции находятся в папке migrations/. При запуске контейнера PostgreSQL база и таблицы создаются автоматически.

    Если нужно применить миграци  и вручную:

        migrate -path ./migrations -database "postgresql://user:pass@host:port/dbname?sslmode=disable" up


Логирование
-   Логи выводятся в консоль с указанием времени, уровня и места в коде

-   Логируются все основные операции и ошибки

## Docker

## Dockerfile
-   Используется образ golang:1.24-alpine

-   Устанавливается bash, git и curl

-   Копируются исходники, скачиваются зависимости и билдится бинарник

-   Используется скрипт wait-for-it.sh для ожидания запуска БД перед стартом приложения

## docker-compose.yml
-   Описывает два сервиса: app (Go-приложение) и postgres-subscriptions (PostgreSQL)

-   Используется volume для хранения данных БД

-   Проброшены порты 8080 (приложение) и 5432 (БД)

## Swagger

Для генерации и просмотра документации:

    -   Команда для генерации:

        swag init

    -   Просмотр документации:

        http://localhost:8080/swagger/index.html
        
## Как внести изменения и запустить локально

1. Внести изменения в код

2. Собрать приложение:

        go build -o main ./cmd/main.go

3. Запустить локально (если есть доступ к базе):

        ./main


## Контакты:

    Если есть вопросы или предложения, пишите на почту: usmanowmukam2003@gmail.com