package filter_controller

import (
	"github.com/google/uuid"
	"selection-service/internal/domain"
)

func (j *JSONUserFilter) JSONToDomain() (*domain.UserFilter, error) {
	id := uuid.Nil
	if j.ID != "" {
		var err error
		id, err = uuid.Parse(j.ID)
		if err != nil {
			return nil, domain.ErrInvalidFilterID
		}
	}

	userID, err := uuid.Parse(j.UserID)
	if err != nil {
		return nil, domain.ErrInvalidUserID
	}

	searchTypeID, err := uuid.Parse(j.SearchTypeID)
	if err != nil {
		return nil, domain.ErrInvalidSearchTypeID
	}

	useTargetID, err := uuid.Parse(j.UseTargetID)
	if err != nil {
		return nil, domain.ErrInvalidUseTargetID
	}

	if !domain.IsValidSexValue(j.Sex) {
		return nil, domain.ErrInvalidSexValue
	}
	sex := domain.SexEnumFromDBValue(j.Sex)

	tagIDs := make([]uuid.UUID, len(j.TagIDs))
	for i, s := range j.TagIDs {
		id, err := uuid.Parse(s)
		if err != nil {
			return nil, domain.ErrInvalidTagID
		}
		tagIDs[i] = id
	}
	return &domain.UserFilter{
		ID:           id,
		UserID:       userID,
		SearchTypeID: searchTypeID,
		Sex:          sex,
		UseTargetID:  useTargetID,
		AgeFrom:      j.AgeFrom,
		AgeTo:        j.AgeTo,
		HeightFrom:   j.HeightFrom,
		HeightTo:     j.HeightTo,
		TagIDs:       tagIDs,
	}, nil
}

func FromDomain(f domain.UserFilter) JSONUserFilter {
	tagIDs := make([]string, len(f.TagIDs))
	for i, id := range f.TagIDs {
		tagIDs[i] = id.String()
	}

	return JSONUserFilter{
		ID:           f.ID.String(),
		UserID:       f.UserID.String(),
		SearchTypeID: f.SearchTypeID.String(),
		Sex:          f.Sex.ToDBValue(),
		UseTargetID:  f.UseTargetID.String(),
		AgeFrom:      f.AgeFrom,
		AgeTo:        f.AgeTo,
		HeightFrom:   f.HeightFrom,
		HeightTo:     f.HeightTo,
		TagIDs:       tagIDs,
	}
}
