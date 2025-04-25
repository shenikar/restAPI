package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	_ "github.com/user-management/docs"
	"github.com/user-management/internal/config"
	"github.com/user-management/internal/handlers"
	"github.com/user-management/internal/middleware"
	"github.com/user-management/internal/repository"
)

// @title User Management API
// @version 1.0
// @description API для управления пользователями и заданиями
// @host localhost:8080
// @BasePath /
func main() {
	// Загрузка конфигурации
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Подключение к базе данных
	db, err := sql.Open("postgres",
		"host="+cfg.DBHost+
			" port="+cfg.DBPort+
			" user="+cfg.DBUser+
			" password="+cfg.DBPassword+
			" dbname="+cfg.DBName+
			" sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Инициализация репозитория
	userRepo := repository.NewUserRepository(db)

	// Инициализация обработчиков
	userHandler := handlers.NewUserHandler(userRepo)
	authHandler := handlers.NewAuthHandler(cfg.JWTSecret, userRepo)

	// Создание маршрутизатора
	r := mux.NewRouter()

	// Добавляем Swagger UI
	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	))

	// Публичные маршруты
	r.HandleFunc("/auth/login", authHandler.Login).Methods("POST")

	// Защищенные маршруты
	protected := r.PathPrefix("").Subrouter()
	protected.Use(middleware.AuthMiddleware(cfg.JWTSecret))

	protected.HandleFunc("/users/{id}/status", userHandler.GetUserStatus).Methods("GET")
	protected.HandleFunc("/users/leaderboard", userHandler.GetLeaderboard).Methods("GET")
	protected.HandleFunc("/users/{id}/task/complete", userHandler.CompleteTask).Methods("POST")
	protected.HandleFunc("/users/{id}/referrer", userHandler.SetReferrer).Methods("POST")

	// Запуск сервера
	log.Printf("Server starting on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
