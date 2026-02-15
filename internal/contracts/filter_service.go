package contracts

import (
	"context"
	"selection-service/internal/domain"

	"github.com/google/uuid"
)

type FilterService interface {
	SaveFilter(ctx context.Context, filter domain.UserFilter) error

	GetFilterById(ctx context.Context, filterID uuid.UUID) (domain.UserFilter, error)

	DeleteFilter(ctx context.Context, filterID uuid.UUID) error

	UpdateFilter(ctx context.Context, filter domain.UserFilter) error
}
