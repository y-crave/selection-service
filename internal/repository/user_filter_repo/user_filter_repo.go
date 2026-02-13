package user_filter_repo

import (
	"context"
	"selection-service/internal/domain"
	"selection-service/internal/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type filterRepo struct {
	db *gorm.DB
}

func NewUserFilterRepo(db *gorm.DB) repository.FilterRepo {
	return &filterRepo{
		db: db,
	}
}

func (r *filterRepo) Create(ctx context.Context, filter domain.UserFilter) error {
	return nil
}

func (r *filterRepo) Read(ctx context.Context, id uuid.UUID) (domain.UserFilter, error) {
	return domain.UserFilter{}, nil
}
func (r *filterRepo) Update(ctx context.Context, filter domain.UserFilter) error {
	return nil
}
func (r *filterRepo) Delete(ctx context.Context, id uuid.UUID) error {
	return nil
}
func (r *filterRepo) FindByUserAndUseTarget(ctx context.Context, userID uuid.UUID, targetID uuid.UUID) (domain.UserFilter, error) {
	return domain.UserFilter{}, nil
}
