package user_filter_service

import (
	"context"
	"selection-service/internal/domain"
	"selection-service/internal/repository"
	"selection-service/internal/service"

	"github.com/google/uuid"
)

type filterService struct {
	filterRepo repository.FilterRepo
}

func NewFilterService(filterRepo repository.FilterRepo) service.FilterService {
	return &filterService{
		filterRepo: filterRepo,
	}
}

func (s *filterService) UpdateFilter(ctx context.Context, filter domain.UserFilter) error {
	return nil
}
func (s *filterService) DeleteFilter(ctx context.Context, filterID uuid.UUID) error {
	return nil
}
func (s *filterService) GetFilterById(ctx context.Context, filterID uuid.UUID) (domain.UserFilter, error) {
	return domain.UserFilter{}, nil
}

func (s *filterService) SaveFilter(ctx context.Context, filter domain.UserFilter) error {
	// todo Попробуем найти существующий фильтр

	// todo Создаём новый

	// todo Обновляем существующий

	return nil
}
