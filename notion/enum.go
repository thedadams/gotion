package notion

import (
	"encoding/json"
	"fmt"
)

// StringEnum interface are the methods needed to unmarshal string enums.
type StringEnum interface {
	SetValue(v string)
	IsValidEnum() bool
}

// InvalidEnumError represents an invalid enum value for a given enum
type InvalidEnumError struct {
	Enum, InvalidValue string
}

// Error turns an InvalidEnumError into an error
func (i *InvalidEnumError) Error() string {
	return fmt.Sprintf("%s is not a valid value for an enum %s", i.InvalidValue, i.Enum)
}

// NewInvalidEnumError returns an error value representing an invalid value for the given enum.
func NewInvalidEnumError(enum, invalidValue string) *InvalidEnumError {
	return &InvalidEnumError{
		Enum:         enum,
		InvalidValue: invalidValue,
	}
}

func isValidEnum(s string, validValues ...string) bool {
	for _, v := range validValues {
		if v == s {
			return true
		}
	}

	return false
}

func unmarshalEnum(b []byte, e StringEnum) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	e.SetValue(s)
	if !e.IsValidEnum() {
		return fmt.Errorf("%s is not a valid %T", s, e)
	}

	return nil
}
