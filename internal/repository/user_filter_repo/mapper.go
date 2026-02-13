package user_filter_repo

import (
	"fmt"
	"selection-service/internal/domain"

	"github.com/google/uuid"
)

func ToDomainUserFilter(e *GormUserFilter) *domain.UserFilter {
	sex := domain.SexEnumFromDBValue(e.Sex)

	return &domain.UserFilter{
		ID:           uuid.MustParse(e.ID),
		UserID:       uuid.MustParse(e.UserID),
		SearchTypeID: uuid.MustParse(e.SearchTypeID),
		Sex:          sex,
		UseTargetID:  uuid.MustParse(e.UseTargetID),
		AgeFrom:      e.AgeFrom,
		AgeTo:        e.AgeTo,
		HeightFrom:   e.HeightFrom,
		HeightTo:     e.HeightTo,
	}
}

func ToGormUserFilter(f *domain.UserFilter) *GormUserFilter {
	return &GormUserFilter{
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

func ToTagIDs(entities []GormUserFilterTag) ([]uuid.UUID, error) {
	tags := make([]uuid.UUID, len(entities))
	for i, ent := range entities {
		id, err := uuid.Parse(ent.TagID)
		if err != nil {
			return nil, domain.ErrInvalidTagID
		}
		tags[i] = id
	}
	return tags, nil
}

func ToDomainSearchType(e *GormSearchType) (*domain.SearchType, error) {
	id, err := uuid.Parse(e.ID)
	if err != nil {
		return nil, fmt.Errorf("invalid search type ID: %w", err)
	}
	return &domain.SearchType{
		ID:   id,
		Code: e.Code,
		Name: e.Name,
	}, nil
}

func ToGormSearchType(d *domain.SearchType) *GormSearchType {
	return &GormSearchType{
		ID:   d.ID.String(),
		Code: d.Code,
		Name: d.Name,
	}
}
