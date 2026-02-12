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

func (s *filterService) UpdateFilter(ctx context.Context, filter domain.UserFilter) error {}
func (s *filterService) DeleteFilter(ctx context.Context, filterID uuid.UUID) error       {}
func (s *filterService) GetFilterById(ctx context.Context, filterID uuid.UUID) (domain.UserFilter, error) {
}

func (s *filterService) SaveFilter(ctx context.Context, filter domain.UserFilter) error {
	/*    // Попробуем найти существующий фильтр
	      existing, err := s.filterRepo.FindByUserAndUseTarget(ctx, filter.UserID, filter.UseTargetID)
	      if err != nil && !errors.Is(err, domain.ErrNotFound) {
	          return err
	      }

	      if errors.Is(err, domain.ErrNotFound) {
	          // Создаём новый
	          return s.filterRepo.Create(ctx, filter)
	      } else {
	          // Обновляем существующий
	          filter.ID = existing.ID // сохраняем тот же ID!
	          return s.filterRepo.Update(ctx, filter)
	      }*/
	return nil
}
