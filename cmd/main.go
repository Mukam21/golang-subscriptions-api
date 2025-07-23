package main

import (
	"log"

	"golang-subscriptions-api/internal/config"
	"golang-subscriptions-api/internal/database"
	"golang-subscriptions-api/internal/handler"
	"golang-subscriptions-api/internal/repository"
	"golang-subscriptions-api/internal/router"
	"golang-subscriptions-api/internal/service"
)

func main() {
	// Загружаем конфигурацию из env или дефолтов
	cfg := config.LoadConfig()

	// Инициализируем подключение к базе данных
	db, err := database.InitDB(cfg)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}

	// Выполняем миграции
	if err := database.RunMigrations(db); err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}

	repo := repository.NewRepository(db)
	svc := service.NewService(repo)
	h := handler.NewHandler(svc)

	// Настраиваем маршрутизатор Gin с хендлером
	r := router.SetupRouter(h)

	// Запускаем сервер на порту из конфигурации
	log.Printf("Starting server on port %s...", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
