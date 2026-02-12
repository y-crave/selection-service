package service

import (
	"context"
	"selection-service/internal/domain"

	"github.com/google/uuid"
)

type FilterService interface {
	// SaveFilter сохраняет или обновляет фильтр.
	// Если фильтр для (user_id, use_target_id) уже существует — обновляет.
	// Иначе — создаёт новый.
	SaveFilter(ctx context.Context, filter domain.UserFilter) error

	// GetFilterById получает фильтр по ID.
	GetFilterById(ctx context.Context, filterID uuid.UUID) (domain.UserFilter, error)

	// DeleteFilter удаляет фильтр по ID.
	DeleteFilter(ctx context.Context, filterID uuid.UUID) error

	// UpdateFilter обновляет фильтр (альтернатива Save, если нужно явное обновление по ID).
	UpdateFilter(ctx context.Context, filter domain.UserFilter) error
}
