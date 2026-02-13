package monitoring_controller

import (
	"context"
	"log"
	"net/http"
	"selection-service/internal/service/monitoring_service"
	"time"

	"github.com/gorilla/mux"
)

type MonitoringController struct {
	service monitoring_service.MonitoringService
}

func NewMonitoringController(service monitoring_service.MonitoringService) *MonitoringController {
	return &MonitoringController{service: service}
}

func (c *MonitoringController) LivenessProbe(w http.ResponseWriter, r *http.Request) {
	log.Println("GET /healthz")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("alive"))
}

func (c *MonitoringController) ReadinessProbe(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	if err := c.service.CheckDB(ctx); err != nil {
		http.Error(w, "DB not ready", http.StatusServiceUnavailable)
		return
	}

	// TODO: добавить проверку других зависимостей: Redis, Kafka и т.д.

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ready"))
}

func (c *MonitoringController) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/healthz", c.LivenessProbe).Methods(http.MethodGet)
	router.HandleFunc("/ready", c.ReadinessProbe).Methods(http.MethodGet)
}
