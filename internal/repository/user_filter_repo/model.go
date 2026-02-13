package user_filter_repo

import "time"

type GormUserFilter struct {
	ID           string    `gorm:"type:uuid;primaryKey;default:gen_random_uuid();autoIncrement:true"`
	UserID       string    `gorm:"type:uuid;not null"`
	SearchTypeID string    `gorm:"type:uuid;not null"`
	Sex          string    `gorm:"type:sex;not null"` // "male", "female"
	UseTargetID  string    `gorm:"type:uuid;not null"`
	AgeFrom      *int      `gorm:"type:int"`
	AgeTo        *int      `gorm:"type:int"`
	HeightFrom   *int      `gorm:"type:int"`
	HeightTo     *int      `gorm:"type:int"`
	CreatedAt    time.Time `gorm:"not null"`
	UpdatedAt    time.Time `gorm:"not null"`
}

type GormUserFilterTag struct {
	UserFilterID string `gorm:"type:uuid;primaryKey"`
	TagID        string `gorm:"type:uuid;primaryKey"`
}

type GormSearchType struct {
	ID   string `gorm:"type:uuid;primaryKey;default:gen_random_uuid();autoIncrement:true"`
	Code string `gorm:"type:varchar(50);not null"`
	Name string `gorm:"type:varchar(100);not null"`
}
