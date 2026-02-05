package domain

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type SexEnum int

const (
	SexMale SexEnum = iota
	SexFemale
	NotSelected
)

var sexName = map[SexEnum]string{
	SexMale:     "Мужчина",
	SexFemale:   "Женщина",
	NotSelected: "Не выбрано",
}

func (s SexEnum) String() string {
	return sexName[s]
}

func (s SexEnum) Valid() bool {
	return s == SexMale || s == SexFemale || s == NotSelected
}

func (s SexEnum) ToDBValue() string {
	switch s {
	case SexMale:
		return "male"
	case SexFemale:
		return "female"
	default:
		return "unknown"
	}
}

func SexEnumFromDBValue(value string) (SexEnum, error) {
	switch value {
	case "male":
		return SexMale, nil
	case "female":
		return SexFemale, nil
	case "unknown":
		return NotSelected, nil
	default:
		return NotSelected, errors.New("invalid sex value in database")
	}
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
	Sex          SexEnum
	UseTargetID  uuid.UUID
	AgeFrom      *int // Использование *int для nullable полей → правильно, так как в PostgreSQL NULL ≠ 0
	AgeTo        *int
	HeightFrom   *int
	HeightTo     *int
	TagIDs       []uuid.UUID
}

func (f *UserFilter) Validate(minAge, maxAge, minHeight, maxHeight int) error {
	if err := f.validSex(); err != nil {
		return err
	}
	if err := f.validAge(minAge, maxAge); err != nil {
		return err
	}
	if err := f.validHeight(minHeight, maxHeight); err != nil {
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

func (f *UserFilter) validAge(minAge, maxAge int) error {
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

func (f *UserFilter) validHeight(minHeight, maxHeight int) error {
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
