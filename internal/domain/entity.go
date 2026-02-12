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

var sexToString = map[SexEnum]string{
	SexMale:     "male",
	SexFemale:   "female",
	NotSelected: "unknown",
}

var stringToSex = map[string]SexEnum{
	"male":    SexMale,
	"female":  SexFemale,
	"unknown": NotSelected,
}

func (s SexEnum) ToDBValue() string {
	if v, ok := sexToString[s]; ok {
		return v
	}
	return "unknown"
}

func (s SexEnum) String() string {
	return s.ToDBValue()
}

func (s SexEnum) Valid() bool {
	_, ok := sexToString[s]
	return ok
}

func SexEnumFromDBValue(value string) SexEnum {
	if s, ok := stringToSex[value]; ok {
		return s
	}
	return NotSelected
}

func IsValidSexValue(s string) bool {
	_, ok := stringToSex[s]
	return ok
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
	if err := f.validAge(minAge, maxAge); err != nil {
		return err
	}
	if err := f.validHeight(minHeight, maxHeight); err != nil {
		return err
	}
	return nil
}

func (f *UserFilter) validAge(minAge, maxAge int) error {
	if f.AgeFrom != nil {
		if *f.AgeFrom < minAge || *f.AgeFrom > maxAge {
			return ErrAgeOutOfRange
		}
	}
	if f.AgeTo != nil {
		if *f.AgeTo < minAge || *f.AgeTo > maxAge {
			return ErrAgeOutOfRange
		}
	}
	if f.AgeFrom != nil && f.AgeTo != nil {
		if *f.AgeFrom > *f.AgeTo {
			return ErrInvalidAgeRange
		}
	}
	return nil
}

func (f *UserFilter) validHeight(minHeight, maxHeight int) error {
	if f.HeightFrom != nil {
		if *f.HeightFrom < minHeight || *f.HeightFrom > maxHeight {
			return ErrHeightOutOfRange
		}
	}
	if f.HeightTo != nil {
		if *f.HeightTo < minHeight || *f.HeightTo > maxHeight {
			return ErrHeightOutOfRange
		}
	}
	if f.HeightFrom != nil && f.HeightTo != nil {
		if *f.HeightFrom > *f.HeightTo {
			return ErrInvalidHeightRange
		}
	}

	return nil
}
