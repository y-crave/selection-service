package main

import (
	"fmt"
	"log"
	"net/http"
	"selection-service/internal/config"
	"selection-service/internal/controller/filter_controller"
	"selection-service/internal/controller/monitoring_controller"
	"selection-service/internal/repository/user_filter_repo"
	"selection-service/internal/service/monitoring_service"
	"selection-service/internal/service/user_filter_service"

	"github.com/gorilla/mux"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// 1. Загружаем конфигурацию
	cfg := config.Load()

	// 2. Инициализируем GORM
	db, err := gorm.Open(postgres.Open(cfg.PostgresDSN), &gorm.Config{
		// Опционально: включить логирование SQL (только для dev!)
		// Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}

	monitoringService := monitoring_service.NewMonitoringService(db)
	monitoringController := monitoring_controller.NewMonitoringController(monitoringService)

	mainRouter := mux.NewRouter()

	apiPrefix := fmt.Sprintf("/api/v1/%s/", cfg.AppName)
	router := mainRouter.PathPrefix(apiPrefix).Subrouter()
	monitoringController.RegisterRoutes(router)

	addr := fmt.Sprintf("%s:%d", cfg.AppHost, cfg.AppHttpPort)
	handler := config.LoggingMiddleware(router)

	config.PrintRoutes(router)
	log.Printf("Server started: %s", addr)
	log.Fatal(http.ListenAndServe(addr, handler))

	// 3. Создаём репозиторий
	filterRepo := user_filter_repo.NewUserFilterRepo(db)

	// 4. Создаём сервис
	filterService := user_filter_service.NewFilterService(filterRepo)

	// 5. Дальше — инициализация HTTP-сервера, роутер и т.д.
	// Например:
	// router := mux.NewRouter()
	// filterCtrl := filter_controller.NewFilterController(filterService, cfg)
	// filterCtrl.RegisterRoutes(router)
	// log.Fatal(http.ListenAndServe(":8080", router))
}
