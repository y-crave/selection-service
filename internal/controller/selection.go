package controller

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"selection-service/internal/authidentity"
	"selection-service/internal/domain"
	"selection-service/internal/service"
)

// SelectionController handles /api/v1/selection/* endpoints.
type SelectionController struct {
	svc service.SelectionService
}

func NewSelectionController(svc service.SelectionService) *SelectionController {
	return &SelectionController{svc: svc}
}

// kvDTO mirrors mobile CandidateKvDto.
type kvDTO struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

// candidateDTO mirrors mobile CandidateDto (snake_case per @JsonKey annotations).
type candidateDTO struct {
	ID               string   `json:"id"`
	Name             string   `json:"name"`
	Age              int      `json:"age"`
	City             string   `json:"city"`
	DistanceKm       float64  `json:"distance_km"`
	IsVerified       bool     `json:"is_verified"`
	UseTarget        string   `json:"use_target"`
	PhotoURLs        []string `json:"photo_urls"`
	Bio              string   `json:"bio,omitempty"`
	InterestTags     []string `json:"interest_tags"`
	Layer2Items      []kvDTO  `json:"layer2_items"`
	Layer3Items      []kvDTO  `json:"layer3_items"`
	IsLayer3Unlocked bool     `json:"is_layer3_unlocked"`
}

// queueResponseDTO wraps the candidate list.
type queueResponseDTO struct {
	Candidates []candidateDTO `json:"candidates"`
}

// filtersDTO mirrors mobile SearchFiltersDto.
type filtersDTO struct {
	Genders       []string `json:"genders"`
	AgeMin        int      `json:"age_min"`
	AgeMax        int      `json:"age_max"`
	DistanceKm    int      `json:"distance_km"`
	HeightMin     *int     `json:"height_min,omitempty"`
	HeightMax     *int     `json:"height_max,omitempty"`
	Religions     []string `json:"religions"`
	WantsChildren *bool    `json:"wants_children,omitempty"`
	Values        []string `json:"values"`
	OnlyVerified  bool     `json:"only_verified"`
}

// GetQueue handles GET /api/v1/selection/queue
func (c *SelectionController) GetQueue(w http.ResponseWriter, r *http.Request) {
	identity, ok := authidentity.FromContext(r.Context())
	if !ok {
		writeJSON(w, http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
		return
	}

	candidates, err := c.svc.GetQueue(r.Context(), identity.UserID)
	if err != nil {
		slog.Error("GetQueue failed", "user_id", identity.UserID, "error", err)
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "internal server error"})
		return
	}

	dtos := make([]candidateDTO, 0, len(candidates))
	for _, cand := range candidates {
		dtos = append(dtos, toCandidateDTO(cand))
	}

	writeJSON(w, http.StatusOK, queueResponseDTO{Candidates: dtos})
}

// GetFilters handles GET /api/v1/selection/filters
func (c *SelectionController) GetFilters(w http.ResponseWriter, r *http.Request) {
	identity, ok := authidentity.FromContext(r.Context())
	if !ok {
		writeJSON(w, http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
		return
	}

	filters, err := c.svc.GetFilters(r.Context(), identity.UserID)
	if err != nil {
		slog.Error("GetFilters failed", "user_id", identity.UserID, "error", err)
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "internal server error"})
		return
	}

	writeJSON(w, http.StatusOK, toFiltersDTO(filters))
}

// PutFilters handles PUT /api/v1/selection/filters
func (c *SelectionController) PutFilters(w http.ResponseWriter, r *http.Request) {
	identity, ok := authidentity.FromContext(r.Context())
	if !ok {
		writeJSON(w, http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
		return
	}

	var dto filtersDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		return
	}

	f := fromFiltersDTO(identity.UserID.String(), dto)
	f.AppUserID = identity.UserID

	if err := c.svc.UpdateFilters(r.Context(), f); err != nil {
		slog.Error("UpdateFilters failed", "user_id", identity.UserID, "error", err)
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "internal server error"})
		return
	}

	writeJSON(w, http.StatusOK, toFiltersDTO(f))
}

// ── helpers ──────────────────────────────────────────────────────────────────

func toCandidateDTO(c *domain.Candidate) candidateDTO {
	layer2 := make([]kvDTO, 0, len(c.Layer2Items))
	for _, kv := range c.Layer2Items {
		layer2 = append(layer2, kvDTO{Label: kv.Label, Value: kv.Value})
	}
	layer3 := make([]kvDTO, 0, len(c.Layer3Items))
	for _, kv := range c.Layer3Items {
		layer3 = append(layer3, kvDTO{Label: kv.Label, Value: kv.Value})
	}
	tags := c.InterestTags
	if tags == nil {
		tags = []string{}
	}
	photos := c.PhotoURLs
	if photos == nil {
		photos = []string{}
	}
	return candidateDTO{
		ID:               c.ID,
		Name:             c.Name,
		Age:              c.Age,
		City:             c.City,
		DistanceKm:       c.DistanceKm,
		IsVerified:       c.IsVerified,
		UseTarget:        c.UseTarget,
		PhotoURLs:        photos,
		Bio:              c.Bio,
		InterestTags:     tags,
		Layer2Items:      layer2,
		Layer3Items:      layer3,
		IsLayer3Unlocked: c.IsLayer3Unlocked,
	}
}

func toFiltersDTO(f *domain.SearchFilters) filtersDTO {
	genders := f.Genders
	if genders == nil {
		genders = []string{"Все"}
	}
	religions := f.Religions
	if religions == nil {
		religions = []string{}
	}
	values := f.Values
	if values == nil {
		values = []string{}
	}
	return filtersDTO{
		Genders:       genders,
		AgeMin:        f.AgeMin,
		AgeMax:        f.AgeMax,
		DistanceKm:    f.DistanceKm,
		HeightMin:     f.HeightMin,
		HeightMax:     f.HeightMax,
		Religions:     religions,
		WantsChildren: f.WantsChildren,
		Values:        values,
		OnlyVerified:  f.OnlyVerified,
	}
}

func fromFiltersDTO(_ string, dto filtersDTO) *domain.SearchFilters {
	genders := dto.Genders
	if len(genders) == 0 {
		genders = []string{"Все"}
	}
	religions := dto.Religions
	if religions == nil {
		religions = []string{}
	}
	values := dto.Values
	if values == nil {
		values = []string{}
	}
	return &domain.SearchFilters{
		Genders:       genders,
		AgeMin:        dto.AgeMin,
		AgeMax:        dto.AgeMax,
		DistanceKm:    dto.DistanceKm,
		HeightMin:     dto.HeightMin,
		HeightMax:     dto.HeightMax,
		Religions:     religions,
		WantsChildren: dto.WantsChildren,
		Values:        values,
		OnlyVerified:  dto.OnlyVerified,
	}
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		slog.Error("writeJSON encode failed", "error", err)
	}
}
