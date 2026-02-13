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
	cfg := config.Load()

	db, err := gorm.Open(postgres.Open(cfg.PostgresDSN), &gorm.Config{
		// todo логирование sql
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

	filterRepo := user_filter_repo.NewUserFilterRepo(db)

	filterService := user_filter_service.NewFilterService(filterRepo)

	// todo Дальше — инициализация HTTP-сервера, роутер и т.д.

	filterCtrl := filter_controller.NewFilterController(filterService)

}
