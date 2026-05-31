package domain

import "github.com/google/uuid"

// SearchFilters stores per-user discovery filter preferences.
type SearchFilters struct {
	AppUserID     uuid.UUID
	Genders       []string
	AgeMin        int
	AgeMax        int
	DistanceKm    int
	HeightMin     *int
	HeightMax     *int
	Religions     []string
	WantsChildren *bool
	Values        []string
	OnlyVerified  bool
}

// DefaultSearchFilters returns the default filter set for a new user.
func DefaultSearchFilters(userID uuid.UUID) *SearchFilters {
	return &SearchFilters{
		AppUserID:  userID,
		Genders:    []string{"Все"},
		AgeMin:     18,
		AgeMax:     65,
		DistanceKm: 50,
		Religions:  []string{},
		Values:     []string{},
	}
}

// Candidate is a swipe-queue profile stub.
// The real implementation will join with user-service profile data via internal gRPC.
// For now we return enough for the mobile DTO contract.
type Candidate struct {
	ID               string
	Name             string
	Age              int
	City             string
	DistanceKm       float64
	IsVerified       bool
	UseTarget        string
	PhotoURLs        []string
	Bio              string
	InterestTags     []string
	Layer2Items      []KVPair
	Layer3Items      []KVPair
	IsLayer3Unlocked bool
}

// KVPair is a labelled key-value item displayed in profile layers.
type KVPair struct {
	Label string
	Value string
}
