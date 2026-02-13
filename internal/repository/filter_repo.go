package repository

import (
	"context"
	"selection-service/internal/domain"

	"github.com/google/uuid"
)

type FilterRepo interface {
	Create(ctx context.Context, filter domain.UserFilter) error

	Read(ctx context.Context, id uuid.UUID) (domain.UserFilter, error)

	Update(ctx context.Context, filter domain.UserFilter) error

	Delete(ctx context.Context, id uuid.UUID) error

	FindByUserAndUseTarget(ctx context.Context, userID, useTargetID uuid.UUID) (domain.UserFilter, error)
}
