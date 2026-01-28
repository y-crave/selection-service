package domain

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Sex string

const (
	SexMale    Sex = "male"
	SexFemale  Sex = "female" //  filter.Sex = domain.SexFemale
	SexUnknown Sex = "unknown"
)

func (s Sex) Valid() bool {
	return s == SexMale || s == SexFemale || s == SexUnknown
}

type SearchType struct {
	ID   uuid.UUID
	Code string
	Name string
}

type UserFilter struct {
	ID           uuid.UUID
	UserID       uuid.UUID
	SearchTypeID uuid.UUID
	Sex          Sex
	UseTargetID  uuid.UUID
	AgeFrom      *int
	AgeTo        *int
	HeightFrom   *int
	HeightTo     *int
	CreatedAt    time.Time
	UpdatedAt    time.Time
	TagIDs       []uuid.UUID
}

func (f *UserFilter) Validate() error {
	if err := f.validSex(); err != nil {
		return err
	}
	if err := f.validAge(); err != nil {
		return err
	}
	if err := f.validHeight(); err != nil {
		return err
	}
	return nil
}

func (f *UserFilter) validSex() error {
	if !f.Sex.Valid() {
		return errors.New("invalid sex")
	}
	return nil
}

const (
	minAge = 0
	maxAge = 150
)

func (f *UserFilter) validAge() error {
	if f.AgeFrom != nil {
		if *f.AgeFrom < minAge || *f.AgeFrom > maxAge {
			return fmt.Errorf("age_from must be between %d and %d years old", minAge, maxAge)
		}
	}
	if f.AgeTo != nil {
		if *f.AgeTo < minAge || *f.AgeTo > maxAge {
			return fmt.Errorf("age_to must be between %d and %d years old", minAge, maxAge)
		}
	}
	if f.AgeFrom != nil && f.AgeTo != nil {
		if *f.AgeFrom > *f.AgeTo {
			return errors.New("age_from must be <= age_to")
		}
	}
	return nil
}

const (
	minHeight = 50
	maxHeight = 300
)

func (f *UserFilter) validHeight() error {
	if f.HeightFrom != nil {
		if *f.HeightFrom < minHeight || *f.HeightFrom > maxHeight {
			return fmt.Errorf("height_from must be between %d and %d cm", minHeight, maxHeight)
		}
	}
	if f.HeightTo != nil {
		if *f.HeightTo < minHeight || *f.HeightTo > maxHeight {
			return fmt.Errorf("height_to must be between %d and %d cm", minHeight, maxHeight)
		}
	}
	if f.HeightFrom != nil && f.HeightTo != nil {
		if *f.HeightFrom > *f.HeightTo {
			return errors.New("height_from must be <= height_to")
		}
	}

	return nil
}
