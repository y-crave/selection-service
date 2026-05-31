package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"selection-service/internal/config"
	"selection-service/internal/controller"
	"selection-service/internal/middleware"
	selectionrepo "selection-service/internal/repository/selection_repo"
	"selection-service/internal/service"
)

func main() {
	cfg := config.Load()

	db, err := sql.Open("postgres", cfg.PostgresDSN)
	if err != nil {
		log.Fatalf("sql.Open: %v", err)
	}
	defer db.Close()

	monitoringService := service.NewMonitoringService(db)
	monitoringController := controller.NewMonitoringController(monitoringService)

	filtersRepo := selectionrepo.NewFiltersRepository(db)
	selectionSvc := service.NewSelectionService(filtersRepo)
	selectionController := controller.NewSelectionController(selectionSvc)

	mainRouter := mux.NewRouter()

	appName := strings.ToLower(cfg.AppName)
	probeRouter := mainRouter.PathPrefix(fmt.Sprintf("/api/v1/%s", appName)).Subrouter()
	monitoringController.RegisterRoutes(probeRouter)

	mainRouter.HandleFunc("/healthz", monitoringController.LivenessProbe).Methods(http.MethodGet)
	mainRouter.HandleFunc("/ready", monitoringController.ReadinessProbe).Methods(http.MethodGet)

	selectionRouter := mainRouter.PathPrefix("/api/v1/selection").Subrouter()
	selectionRouter.Use(middleware.GatewayIdentity(cfg))
	selectionRouter.HandleFunc("/queue", selectionController.GetQueue).Methods(http.MethodGet)
	selectionRouter.HandleFunc("/filters", selectionController.GetFilters).Methods(http.MethodGet)
	selectionRouter.HandleFunc("/filters", selectionController.PutFilters).Methods(http.MethodPut)

	addr := fmt.Sprintf("%s:%d", cfg.AppHost, cfg.AppHttpPort)
	handler := config.LoggingMiddleware(mainRouter)

	config.PrintRoutes(mainRouter)
	log.Printf("Server started: %s", addr)
	log.Fatal(http.ListenAndServe(addr, handler))
}
