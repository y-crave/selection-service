package filter_controller

import (
	"net/http"
	"selection-service/internal/contracts"

	"github.com/gorilla/mux"
)

type FilterController struct {
	svc contracts.FilterService
}

func NewFilterController(svc contracts.FilterService) *FilterController {
	return &FilterController{svc: svc}
}

func (c *FilterController) RegisterRoutes(router *mux.Router) {
	// todo handler
	/*	router.HandleFunc("/api/v1/filter", c.CreateFilter).Methods(http.MethodPost)*/
}

func (c *FilterController) SaveFilter(w http.ResponseWriter, r *http.Request) {
	//todo В сервисе: "сохрани или обнови"
}
