package filter_controller

import (
	"net/http"
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
	// todo handler
	/*	router.HandleFunc("/api/v1/filter", c.CreateFilter).Methods(http.MethodPost)*/
}

func (c *filterController) SaveFilter(w http.ResponseWriter, r *http.Request) {
	//todo В сервисе: "сохрани или обнови"
}
