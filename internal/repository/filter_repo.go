package repository

import (
	"context"
	"github.com/google/uuid"
	"selection-service/internal/domain"
)

// FilterRepo управляет фильтрами в хранилище.
type FilterRepo interface {
	// Create создаёт новый фильтр. ID должен быть уникальным.
	Create(ctx context.Context, filter domain.UserFilter) error

	// Read получает фильтр по ID.
	Read(ctx context.Context, id uuid.UUID) (domain.UserFilter, error)

	// Update обновляет существующий фильтр (по ID).
	Update(ctx context.Context, filter domain.UserFilter) error

	// Delete удаляет фильтр по ID.
	Delete(ctx context.Context, id uuid.UUID) error

	// FindByUserAndUseTarget ищет фильтр по паре (user_id, use_target_id).
	// Возвращает ErrNotFound, если не найден.
	FindByUserAndUseTarget(ctx context.Context, userID, useTargetID uuid.UUID) (domain.UserFilter, error)
}
