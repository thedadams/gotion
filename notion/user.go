package notion

import "encoding/json"

// These constants represent the valid enum types for users in the Notion API.
const (
	UserTypeEnumPerson = "person"
	UserTypeEnumBot    = "bot"
)

// UserTypeEnum represents a valid type for a user object in the Notion API.
type UserTypeEnum string

// SetValue sets the UserTypeEnum to the given string
func (ute *UserTypeEnum) SetValue(s string) {
	if ute != nil {
		*ute = UserTypeEnum(s)
	}
}

// IsValidEnum returns true if the string represents a valid user type in the Notion API.
func (ute *UserTypeEnum) IsValidEnum() bool {
	return ute != nil && isValidEnum(string(*ute), UserTypeEnumPerson, UserTypeEnumBot)
}

// UnmarshalJSON returns an error if the provided string is not a valid UserTypeEnum in the Notion API.
func (ute *UserTypeEnum) UnmarshalJSON(b []byte) error {
	return unmarshalEnum(b, ute)
}

// A User represents a user object in the Notion API.
type User struct {
	Object
	Type      UserTypeEnum `json:"type,omitempty"`
	Name      string       `json:"name,omitempty"`
	AvatarURL jsonURL      `json:"avatar_url,omitempty"`
	// Only set if the Type is "person"
	Email string `json:"email"`
}

type user User

// GetType implements the typed interface for User
func (u *user) getType() string {
	if u == nil {
		return ""
	}
	return string(u.Type)
}

// FieldsToExpand implements the expander interface for User
func (u *user) fieldsToExpand() []string {
	return []string{"email"}
}

// UnmarshalJSON parses the user object from the Notion API, flattening the "person" or "bot" attributes.
func (u *User) UnmarshalJSON(b []byte) error {
	uu := new(user)
	if err := unmarshalJSONFlattenByType(b, uu); err != nil {
		return err
	}

	*u = User(*uu)
	return nil
}

// MarshalJSON marshals the User object to be compatible with the Notion API.
func (u *User) MarshalJSON() ([]byte, error) {
	if u == nil {
		return nil, nil
	}
	uu := user(*u)
	return marshalJSONExpandByType(&uu)
}

// Users represents a list of users from the Notion API.
type Users []*User

// Len returns the number of users in the the slice.
func (us *Users) Len() int {
	if us == nil {
		return 0
	}
	return len(*us)
}

type users Users

// UnmarshalJSON appends the unmarshaled Users to the slice.
func (us *Users) UnmarshalJSON(b []byte) error {
	u := new(users)
	if err := json.Unmarshal(b, u); err != nil || len(*u) == 0 {
		return err
	}

	if us == nil {
		*us = make([]*User, 0, len(*u))
	}
	*us = append(*us, []*User(*u)...)
	return nil
}
