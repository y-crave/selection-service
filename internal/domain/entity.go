package domain

import (
	"errors"
	"time"
	"database/sql/driver"
	"github.com/google/uuid"
)

type Sex string

const(
	SexMale Sex = "male"
	SexFemale Sex = "female" //  filter.Sex = domain.SexFemale
	SexUnknown Sex = "unknown" 
)

func (s Sex) Valid() bool {
	return s == SexMale || s == SexFemale || s == SexUnknown //  if !sex.Valid() { return error }
}

// // database/sql/driver.Scanner
// type Scanner interface {
//     Scan(src interface{}) error
// }
// type Valuer interface {
//     // Value возвращает значение для сохранения в базе данных.
//     // driver.Value может быть одним из следующих типов:
//     //
//     //   int64
//     //   float64
//     //   bool
//     //   []byte
//     //   string
//     //   time.Time
//     //   nil (для NULL)
//     Value() (Value, error)
// }

func (s *Sex) Scan (value interface{}) error {
	if value == nil {
		*s = SexUnknown
		return nil
	}
	str, ok := value.(string)
	if !ok {
		return errors.New("Invalid sex value")
	}
	*s = Sex(str)
	if !s.Valid(){
		return errors.New("Invalid sex enum value")
	}
	return nil
}

func (s Sex) Value() (driver.Value, error){
	if !s.Valid() { return nil, errors.New("Invalid")}
	return string(s), nil
}

type SearchType struct {
	ID   uuid.UUID `json:"id"`
	Code string    `json:"code"`
	Name string    `json:"name"` 
}

type UserFilter struct{
	ID           uuid.UUID `json:"id"`
	UserID       uuid.UUID `json:"user_id"`
	SearchTypeID uuid.UUID `json:"search_type_id"`
	Sex          Sex       `json:"sex"`
	UseTargetID  uuid.UUID `json:"use_target_id"`
	AgeFrom      *int      `json:"age_from,omitempty"`
	AgeTo        *int      `json:"age_to,omitempty"`
	HeightFrom   *int      `json:"height_from,omitempty"`
	HeightTo     *int      `json:"height_to,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	TagIDs       []uuid.UUID `json:"tag_ids"`
}

func (f *UserFilter) Validate() error {
	if !f.Sex.Valid() {
		return errors.New("invalid sex")
	}

	if f.AgeFrom != nil && f.AgeTo != nil {
		if *f.AgeFrom > *f.AgeTo {
			return errors.New("age_from must be <= age_to")
		}
		if *f.AgeFrom < 0 || *f.AgeTo > 150 {
			return errors.New("age must be between 0 and 150")
		}
	}

	if f.HeightFrom != nil && f.HeightTo != nil {
		if *f.HeightFrom > *f.HeightTo {
			return errors.New("height_from must be <= height_to")
		}
		if *f.HeightFrom < 50 || *f.HeightTo > 300 {
			return errors.New("height must be between 50 and 300 cm")
		}
	}

	return nil
}