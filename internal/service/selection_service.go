package service

import (
	"context"

	"github.com/google/uuid"

	"selection-service/internal/domain"
	"selection-service/internal/repository/selection_repo"
)

// SelectionService handles queue and filter business logic.
type SelectionService interface {
	GetQueue(ctx context.Context, userID uuid.UUID) ([]*domain.Candidate, error)
	GetFilters(ctx context.Context, userID uuid.UUID) (*domain.SearchFilters, error)
	UpdateFilters(ctx context.Context, f *domain.SearchFilters) error
}

type selectionService struct {
	filtersRepo selection_repo.FiltersRepository
}

func NewSelectionService(filtersRepo selection_repo.FiltersRepository) SelectionService {
	return &selectionService{filtersRepo: filtersRepo}
}

// GetQueue returns the swipe-candidate queue for the given user.
// NOTE: real implementation will query a pre-computed candidate pool filtered
// by search_filters. For now it returns a stub queue so the mobile app can
// integrate before the full algorithm is ready.
func (s *selectionService) GetQueue(_ context.Context, _ uuid.UUID) ([]*domain.Candidate, error) {
	// Stub: return 3 synthetic candidates. Replace with DB/algorithm when ready.
	return stubCandidates(), nil
}

func (s *selectionService) GetFilters(ctx context.Context, userID uuid.UUID) (*domain.SearchFilters, error) {
	return s.filtersRepo.Get(ctx, userID)
}

func (s *selectionService) UpdateFilters(ctx context.Context, f *domain.SearchFilters) error {
	return s.filtersRepo.Upsert(ctx, f)
}

// stubCandidates produces a minimal synthetic queue matching the mobile DTO contract.
func stubCandidates() []*domain.Candidate {
	return []*domain.Candidate{
		{
			ID:         "candidate_stub_1",
			Name:       "Аня",
			Age:        24,
			City:       "Москва",
			DistanceKm: 1.5,
			IsVerified: true,
			UseTarget:  "Серьёзные отношения",
			PhotoURLs: []string{
				"https://picsum.photos/seed/sel-1-1/720/960",
				"https://picsum.photos/seed/sel-1-2/720/960",
			},
			Bio:          "Люблю путешествия и кофе.",
			InterestTags: []string{"Горы", "Кулинария", "Книги"},
			Layer2Items: []domain.KVPair{
				{Label: "Рост", Value: "168 см"},
				{Label: "Образование", Value: "Высшее"},
			},
			Layer3Items: []domain.KVPair{
				{Label: "Доход", Value: "Выше среднего"},
			},
			IsLayer3Unlocked: false,
		},
		{
			ID:         "candidate_stub_2",
			Name:       "Катя",
			Age:        26,
			City:       "Санкт-Петербург",
			DistanceKm: 3.0,
			IsVerified: false,
			UseTarget:  "Живое общение",
			PhotoURLs: []string{
				"https://picsum.photos/seed/sel-2-1/720/960",
			},
			Bio:          "Дизайнер по профессии.",
			InterestTags: []string{"Дизайн", "Кино", "Музыка"},
			Layer2Items:  []domain.KVPair{{Label: "Языки", Value: "RUS, ENG"}},
			Layer3Items:  []domain.KVPair{},
		},
	}
}
