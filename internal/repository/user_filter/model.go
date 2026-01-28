package user_filter

import "time"

type UserFilterEntity struct {
	ID           string    `db:"id"` // UUID как строка для sqlx
	UserID       string    `db:"user_id"`
	SearchTypeID string    `db:"search_type_id"`
	Sex          string    `db:"sex"` // "male"/"female"
	UseTargetID  string    `db:"use_target_id"`
	AgeFrom      *int      `db:"age_from"` // *int позволяет отличить «не задано» от «0 лет»
	AgeTo        *int      `db:"age_to"`
	HeightFrom   *int      `db:"height_from"`
	HeightTo     *int      `db:"height_to"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

type UserFilterTagEntity struct {
	FilterID string `db:"filter_id"`
	TagID    string `db:"tag_id"`
}

type SearchTypeEntity struct {
	ID   string `db:"id"`
	Code string `db:"code"`
	Name string `db:"name"`
}
