package user_filter_repo

import (
	"context"
	"selection-service/internal/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FilterRepo struct {
	db *gorm.DB
}

func NewUserFilterRepo(db *gorm.DB) *FilterRepo {
	return &FilterRepo{
		db: db,
	}
}

func (r *FilterRepo) Create(ctx context.Context, filter domain.UserFilter) error {
	return nil
}

func (r *FilterRepo) Read(ctx context.Context, id uuid.UUID) (domain.UserFilter, error) {
	return domain.UserFilter{}, nil
}
func (r *FilterRepo) Update(ctx context.Context, filter domain.UserFilter) error {
	return nil
}
func (r *FilterRepo) Delete(ctx context.Context, id uuid.UUID) error {
	return nil
}
func (r *FilterRepo) FindByUserAndUseTarget(ctx context.Context, userID uuid.UUID, targetID uuid.UUID) (domain.UserFilter, error) {
	return domain.UserFilter{}, nil
}
