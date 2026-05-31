package controller_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"

	"selection-service/internal/authidentity"
	"selection-service/internal/controller"
	"selection-service/internal/domain"
)

// ── fake service ──────────────────────────────────────────────────────────────

type fakeSelectionService struct {
	filters      *domain.SearchFilters
	savedFilters *domain.SearchFilters
}

func (f *fakeSelectionService) GetQueue(_ context.Context, _ uuid.UUID) ([]*domain.Candidate, error) {
	return []*domain.Candidate{
		{
			ID:           "c1",
			Name:         "Test",
			Age:          25,
			City:         "Москва",
			DistanceKm:   2.0,
			InterestTags: []string{},
			Layer2Items:  []domain.KVPair{},
			Layer3Items:  []domain.KVPair{},
		},
	}, nil
}

func (f *fakeSelectionService) GetFilters(_ context.Context, userID uuid.UUID) (*domain.SearchFilters, error) {
	if f.filters != nil {
		return f.filters, nil
	}
	return domain.DefaultSearchFilters(userID), nil
}

func (f *fakeSelectionService) UpdateFilters(_ context.Context, filters *domain.SearchFilters) error {
	f.savedFilters = filters
	return nil
}

// ── helpers ───────────────────────────────────────────────────────────────────

func withIdentity(r *http.Request, userID uuid.UUID) *http.Request {
	ctx := authidentity.WithContext(r.Context(), authidentity.Identity{UserID: userID})
	return r.WithContext(ctx)
}

func newTestController() (*controller.SelectionController, *fakeSelectionService) {
	svc := &fakeSelectionService{}
	c := controller.NewSelectionController(svc)
	return c, svc
}

// ── tests ─────────────────────────────────────────────────────────────────────

func TestGetQueue_OK(t *testing.T) {
	ctrl, _ := newTestController()
	userID := uuid.New()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/selection/queue", nil)
	req = withIdentity(req, userID)
	w := httptest.NewRecorder()

	ctrl.GetQueue(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	var body map[string]any
	if err := json.Unmarshal(w.Body.Bytes(), &body); err != nil {
		t.Fatalf("invalid json: %v", err)
	}
	candidates, ok := body["candidates"].([]any)
	if !ok || len(candidates) == 0 {
		t.Fatal("expected non-empty candidates array")
	}
}

func TestGetQueue_Unauthorized(t *testing.T) {
	ctrl, _ := newTestController()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/selection/queue", nil)
	w := httptest.NewRecorder()

	ctrl.GetQueue(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Fatalf("expected 401, got %d", w.Code)
	}
}

func TestGetFilters_DefaultForNewUser(t *testing.T) {
	ctrl, _ := newTestController()
	userID := uuid.New()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/selection/filters", nil)
	req = withIdentity(req, userID)
	w := httptest.NewRecorder()

	ctrl.GetFilters(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	var dto map[string]any
	if err := json.Unmarshal(w.Body.Bytes(), &dto); err != nil {
		t.Fatalf("invalid json: %v", err)
	}
	if dto["age_min"] != float64(18) {
		t.Fatalf("expected age_min=18, got %v", dto["age_min"])
	}
	if dto["age_max"] != float64(65) {
		t.Fatalf("expected age_max=65, got %v", dto["age_max"])
	}
}

func TestPutFilters_SavesAndReturns(t *testing.T) {
	ctrl, svc := newTestController()
	userID := uuid.New()

	payload := map[string]any{
		"genders":       []string{"Женский"},
		"age_min":       20,
		"age_max":       35,
		"distance_km":   30,
		"religions":     []string{},
		"values":        []string{},
		"only_verified": true,
	}
	body, _ := json.Marshal(payload)

	req := httptest.NewRequest(http.MethodPut, "/api/v1/selection/filters", bytes.NewReader(body))
	req = withIdentity(req, userID)
	w := httptest.NewRecorder()

	ctrl.PutFilters(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", w.Code, w.Body.String())
	}
	if svc.savedFilters == nil {
		t.Fatal("expected filters to be saved")
	}
	if svc.savedFilters.AgeMin != 20 {
		t.Fatalf("expected age_min=20, got %d", svc.savedFilters.AgeMin)
	}
	if !svc.savedFilters.OnlyVerified {
		t.Fatal("expected only_verified=true")
	}
}

func TestPutFilters_InvalidBody(t *testing.T) {
	ctrl, _ := newTestController()
	userID := uuid.New()

	req := httptest.NewRequest(http.MethodPut, "/api/v1/selection/filters",
		bytes.NewBufferString("not-json"))
	req = withIdentity(req, userID)
	w := httptest.NewRecorder()

	ctrl.PutFilters(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}
