package model

type Sex string

const(
	SexMale Sex = "male"
	SexFemale Sex = "female" //  filter.Sex = domain.SexFemale
	SexUnknown Sex = "unknown" 
)

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
