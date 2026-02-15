package user_filter_service

import (
	"context"
	"selection-service/internal/contracts"
	"selection-service/internal/domain"

	"github.com/google/uuid"
)

type FilterService struct {
	filterRepo contracts.FilterRepo
}

func NewFilterService(filterRepo contracts.FilterRepo) *FilterService {
	return &FilterService{
		filterRepo: filterRepo,
	}
}

func (s *FilterService) UpdateFilter(ctx context.Context, filter domain.UserFilter) error {
	return nil
}
func (s *FilterService) DeleteFilter(ctx context.Context, filterID uuid.UUID) error {
	return nil
}
func (s *FilterService) GetFilterById(ctx context.Context, filterID uuid.UUID) (domain.UserFilter, error) {
	return domain.UserFilter{}, nil
}

func (s *FilterService) SaveFilter(ctx context.Context, filter domain.UserFilter) error {
	// todo Попробуем найти существующий фильтр

	// todo Создаём новый

	// todo Обновляем существующий

	return nil
}
