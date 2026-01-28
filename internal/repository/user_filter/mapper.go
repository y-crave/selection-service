package user_filter

import (
	"fmt"
	"selection-service/internal/domain"

	"github.com/google/uuid"
)

func ToDomain(e *UserFilterEntity) (*domain.UserFilter, error) {
	sex, err := domain.SexEnumFromDBValue(e.Sex)
	if err != nil {
		return nil, err
	}
	id, err := uuid.Parse(e.ID)
	if err != nil {
		return nil, fmt.Errorf("invalid filter ID: %w", err)
	}
	userID, err := uuid.Parse(e.UserID)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID: %w", err)
	}
	searchTypeID, err := uuid.Parse(e.SearchTypeID)
	if err != nil {
		return nil, fmt.Errorf("invalid search type ID: %w", err)
	}
	useTargetID, err := uuid.Parse(e.UseTargetID)
	if err != nil {
		return nil, fmt.Errorf("invalid use target ID: %w", err)
	}

	return &domain.UserFilter{
		ID:           id,
		UserID:       userID,
		SearchTypeID: searchTypeID,
		Sex:          sex,
		UseTargetID:  useTargetID,
		AgeFrom:      e.AgeFrom,
		AgeTo:        e.AgeTo,
		HeightFrom:   e.HeightFrom,
		HeightTo:     e.HeightTo,
	}, nil
}

func ToEntity(f *domain.UserFilter) *UserFilterEntity {
	return &UserFilterEntity{
		ID:           f.ID.String(),
		UserID:       f.UserID.String(),
		SearchTypeID: f.SearchTypeID.String(),
		Sex:          f.Sex.ToDBValue(),
		UseTargetID:  f.UseTargetID.String(),
		AgeFrom:      f.AgeFrom,
		AgeTo:        f.AgeTo,
		HeightFrom:   f.HeightFrom,
		HeightTo:     f.HeightTo,
	}
}

func ToTagIDs(entities []UserFilterTagEntity) ([]uuid.UUID, error) {
	tags := make([]uuid.UUID, len(entities))
	for i, ent := range entities {
		id, err := uuid.Parse(ent.TagID)
		if err != nil {
			return nil, fmt.Errorf("invalid tag ID at index %d: %w", i, err)
		}
		tags[i] = id
	}
	return tags, nil
}
