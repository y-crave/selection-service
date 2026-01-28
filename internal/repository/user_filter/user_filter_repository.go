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