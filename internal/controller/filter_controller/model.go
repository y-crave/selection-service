package filter_controller

type JSONUserFilter struct {
	ID           string   `json:"id,omitempty"`
	UserID       string   `json:"user_id"`
	SearchTypeID string   `json:"search_type_id"`
	Sex          string   `json:"sex"`
	UseTargetID  string   `json:"use_target_id"`
	AgeFrom      *int     `json:"age_from,omitempty"`
	AgeTo        *int     `json:"age_to,omitempty"`
	HeightFrom   *int     `json:"height_from,omitempty"`
	HeightTo     *int     `json:"height_to,omitempty"`
	TagIDs       []string `json:"tag_ids,omitempty"`
}
