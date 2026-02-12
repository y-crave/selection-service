package filter_controller

import (
	"context"
	"net/http"
	"selection-service/internal/domain"
	"selection-service/internal/service"

	"github.com/gorilla/mux"
)

type filterController struct {
	svc service.FilterService
}

func NewFilterController(svc service.FilterService) *filterController {
	return &filterController{svc: svc}
}

func (c *filterController) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/api/v1/filter", c.CreateFilter).Methods(http.MethodPost)

}

func (c *filterController) SaveFilter(w http.ResponseWriter, r *http.Request) {
	// ... парсинг

	// В сервисе: "сохрани или обнови"
	err := c.svc.SaveFilter(r.Context(), *filter)
}
