<<<<<<< HEAD
/*package user_filter

import (
	"github.com/jmoiron/sqlx"
	"selection-service/internal/domain"
)

type Repository struct {
	db *sqlx.DB
}

func (r *Repository) GetByID(filterID string) (*domain.UserFilter, error) {
	// 1. Загружаем основную сущность
	var ent UserFilterEntity
	err := r.db.Get(&ent, "SELECT * FROM user_filters WHERE id = $1", filterID)
	if err != nil {
		return nil, err // оберни в domain.ErrNotFound при необходимости
	}

	// 2. Загружаем теги
	var tagEnts []UserFilterTagEntity
	err = r.db.Select(&tagEnts, "SELECT tag_id FROM user_filter_tags WHERE filter_id = $1", filterID)
	if err != nil {
		return nil, err
	}

	// 3. Маппинг в домен
	filter := ToDomain(&ent)
	filter.Tags = ToTagIDs(tagEnts) // предположим, Tags = []string в домене

	return filter, nil
}
*/
=======
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
	/*	gormFilter := ToGormUserFilter(&filter)
		err := r.db.WithContext(ctx).Create(&gormFilter).Error
		if err != nil {
			return fmt.Errorf("create user filter: %w", err)
		}

		if len(filter.TagIDs) >0 {
			links:= make([]GormUserFilter, len(filter.TagIDs))
			for i, tagID := range filter.TagIDs {

			}

		}*/
	return nil
}

func (r *filterRepo) Read(ctx context.Context, id uuid.UUID) (domain.UserFilter, error) {
	/*	var gormFilter GormUserFilter
		if err := r.db.WithContext(ctx).First(&gormFilter, id).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return domain.UserFilter{}, domain.ErrInternal
			}
		}*/
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
>>>>>>> 2098faa (add interfaces for repo&service;add empty methods;add controller)
