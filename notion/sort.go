package notion

// These represent valid enum values for sort objects in the Notion API.
const (
	SortTimestampEnumCreatedTime    = "created_time"
	SortTimestampEnumLastEditedTime = "last_edited_time"

	SortDirectionEnumAscending  = "ascending"
	SortDirectionEnumDescending = "descending"
)

// SortTimestampEnum represents a valid timestamp sort type in the Notion API.
type SortTimestampEnum string

// SetValue sets the SortTimestampEnum to the given string
func (ste *SortTimestampEnum) SetValue(s string) {
	if ste != nil {
		*ste = SortTimestampEnum(s)
	}
}

// IsValidEnum returns true if the provide string represents a valid SortTimestampEnum in the Notion API.
func (ste *SortTimestampEnum) IsValidEnum() bool {
	if ste == nil {
		return false
	}
	return isValidEnum(string(*ste), SortTimestampEnumCreatedTime, SortTimestampEnumLastEditedTime)
}

// UnmarshalJSON returns an error if the string does not represent a valid SortTimestampEnum in the Notion API.
func (ste *SortTimestampEnum) UnmarshalJSON(b []byte) error {
	return unmarshalEnum(b, ste)
}

// SortDirectionEnum represents a valid timestamp sort type in the Notion API.
type SortDirectionEnum string

// SetValue sets the SortDirectionEnum to the given string
func (sde *SortDirectionEnum) SetValue(s string) {
	if sde != nil {
		*sde = SortDirectionEnum(s)
	}
}

// IsValidEnum returns true if the provide string represents a valid SortDirectionEnum in the Notion API.
func (sde *SortDirectionEnum) IsValidEnum() bool {
	return sde != nil && isValidEnum(string(*sde), SortDirectionEnumAscending, SortDirectionEnumDescending)
}

// UnmarshalJSON returns an error if the string does not represent a valid SortDirectionEnum in the Notion API.
func (sde *SortDirectionEnum) UnmarshalJSON(b []byte) error {
	return unmarshalEnum(b, sde)
}

// A Sort object represents a sort used to query a database in the Notion API.
type Sort struct {
	Property  *string            `json:"property,omitempty"`
	Timestamp *SortTimestampEnum `json:"timestamp,omitempty"`
	Direction *SortDirectionEnum `json:"direction,omitempty"`
}

// NewPropertySort returns a new Sort object with the provided property as name and the given sort direction
func NewPropertySort(s string, d SortDirectionEnum) (*Sort, error) {
	if !d.IsValidEnum() {
		return nil, NewInvalidEnumError("SortDirectionEnum", string(d))
	}
	return &Sort{Property: &s, Direction: &d}, nil
}

// NewTimestampSort returns a new Sort object with the provided timestamp and the given sort direction
func NewTimestampSort(t SortTimestampEnum, d SortDirectionEnum) (*Sort, error) {
	if !d.IsValidEnum() {
		return nil, NewInvalidEnumError("SortDirectionEnum", string(d))
	}
	if !t.IsValidEnum() {
		return nil, NewInvalidEnumError("SortTimestampEnum", string(t))
	}

	return &Sort{Timestamp: &t, Direction: &d}, nil
}
