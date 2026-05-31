package selection_repo

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/lib/pq"

	"selection-service/internal/domain"
)

// FiltersRepository persists search filters per user.
type FiltersRepository interface {
	Get(ctx context.Context, userID uuid.UUID) (*domain.SearchFilters, error)
	Upsert(ctx context.Context, f *domain.SearchFilters) error
}

type filtersRepository struct {
	db *sql.DB
}

func NewFiltersRepository(db *sql.DB) FiltersRepository {
	return &filtersRepository{db: db}
}

const selectFiltersSQL = `
SELECT app_user_id, genders, age_min, age_max, distance_km,
       height_min, height_max, religions, wants_children, values, only_verified
FROM search_filters
WHERE app_user_id = $1
`

func (r *filtersRepository) Get(ctx context.Context, userID uuid.UUID) (*domain.SearchFilters, error) {
	row := r.db.QueryRowContext(ctx, selectFiltersSQL, userID)
	f := &domain.SearchFilters{}
	var heightMin, heightMax sql.NullInt16
	var wantsChildren sql.NullBool

	err := row.Scan(
		&f.AppUserID,
		pq.Array(&f.Genders),
		&f.AgeMin, &f.AgeMax,
		&f.DistanceKm,
		&heightMin, &heightMax,
		pq.Array(&f.Religions),
		&wantsChildren,
		pq.Array(&f.Values),
		&f.OnlyVerified,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// return default filters for new users
			return domain.DefaultSearchFilters(userID), nil
		}
		return nil, err
	}
	if heightMin.Valid {
		v := int(heightMin.Int16)
		f.HeightMin = &v
	}
	if heightMax.Valid {
		v := int(heightMax.Int16)
		f.HeightMax = &v
	}
	if wantsChildren.Valid {
		f.WantsChildren = &wantsChildren.Bool
	}
	return f, nil
}

const upsertFiltersSQL = `
INSERT INTO search_filters (
    app_user_id, genders, age_min, age_max, distance_km,
    height_min, height_max, religions, wants_children, values, only_verified, updated_at
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, NOW())
ON CONFLICT (app_user_id) DO UPDATE SET
    genders        = EXCLUDED.genders,
    age_min        = EXCLUDED.age_min,
    age_max        = EXCLUDED.age_max,
    distance_km    = EXCLUDED.distance_km,
    height_min     = EXCLUDED.height_min,
    height_max     = EXCLUDED.height_max,
    religions      = EXCLUDED.religions,
    wants_children = EXCLUDED.wants_children,
    values         = EXCLUDED.values,
    only_verified  = EXCLUDED.only_verified,
    updated_at     = NOW()
`

func (r *filtersRepository) Upsert(ctx context.Context, f *domain.SearchFilters) error {
	var heightMin, heightMax sql.NullInt16
	if f.HeightMin != nil {
		heightMin = sql.NullInt16{Int16: int16(*f.HeightMin), Valid: true} //nolint:gosec
	}
	if f.HeightMax != nil {
		heightMax = sql.NullInt16{Int16: int16(*f.HeightMax), Valid: true} //nolint:gosec
	}
	var wantsChildren sql.NullBool
	if f.WantsChildren != nil {
		wantsChildren = sql.NullBool{Bool: *f.WantsChildren, Valid: true}
	}
	_, err := r.db.ExecContext(ctx, upsertFiltersSQL,
		f.AppUserID,
		pq.Array(f.Genders),
		f.AgeMin, f.AgeMax,
		f.DistanceKm,
		heightMin, heightMax,
		pq.Array(f.Religions),
		wantsChildren,
		pq.Array(f.Values),
		f.OnlyVerified,
	)
	return err
}
