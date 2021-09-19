package notion

import (
	"encoding/json"
	"net/mail"
	"net/url"
	"strings"
	"time"

	"github.com/google/uuid"
)

const (
	noTimeDateLayout = "2006-01-02"

	IconTypeEnumFile  = "file"
	IconTypeEnumEmoji = "emoji"
)

// typed is an interface that returns the type of an object as returned by the Notion API.
type typed interface {
	getType() string
}

// expander is an interface that returns the fields of an object that should be expanded when marshaling.
type expander interface {
	typed
	fieldsToExpand() []string
}

// UUID4 represents the UUIDs used as "id" in the Notion API.
// It is not actually a UUID4, but that is the type used in the Notion API.
type UUID4 uuid.UUID

// String returns the UUID4 as a string
func (u *UUID4) String() string {
	if u == nil {
		return ""
	}
	return uuid.UUID(*u).String()
}

// UnmarshalJSON returns an error if the provide string is not a valid UUID
func (u *UUID4) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	if uu, err := uuid.Parse(s); err == nil {
		*u = UUID4(uu)
	} else {
		return err
	}

	return nil
}

// MarshalJSON marshals the UUID as a string
func (u *UUID4) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.String())
}

// An Object represents all the fields that are common to all objects in the Notion API.
type Object struct {
	ID     UUID4  `json:"id,omitempty"`
	Object string `json:"object,omitempty"`
}

// An Editable represents all objects in the Notion API that have a "created" and "last updated" time.
type Editable struct {
	CreatedTime    time.Time `json:"created_time,omitempty"`
	LastEditedTime time.Time `json:"last_edited_time,omitempty"`
}

// MarshalJSON marshals the date and time to one compatible with the Notion API.
func (e *Editable) MarshalJSON() ([]byte, error) {
	m := map[string]string{}
	if !e.CreatedTime.IsZero() {
		m["created_time"] = e.CreatedTime.Format(time.RFC3339)
	}
	if !e.LastEditedTime.IsZero() {
		m["last_edited_time"] = e.LastEditedTime.Format(time.RFC3339)
	}
	return json.Marshal(m)
}

// A Date object represents a date with a start and end date/time in the Notion API.
type Date struct {
	Start   time.Time `json:"start"`
	End     time.Time `json:"end,omitempty"`
	HasTime bool      `json:"-"`
}

// UnmarshalJSON unmarshals the date based on whether or not it has a time associated.
func (d *Date) UnmarshalJSON(b []byte) error {
	m := make(map[string]string)
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}

	start, end := m["start"], m["end"]
	err = d.parseTime(start, false)
	if err != nil {
		return err
	}

	return d.parseTime(end, true)
}

// MarshalJSON marshals the date and time to one compatible with the Notion API.
func (d *Date) MarshalJSON() ([]byte, error) {
	m := map[string]string{}
	layout := noTimeDateLayout
	if d.HasTime {
		layout = time.RFC3339
	}
	if !d.Start.IsZero() {
		m["start"] = d.Start.Format(layout)
	}
	if !d.End.IsZero() {
		m["end"] = d.End.Format(layout)
	}
	return json.Marshal(m)
}

func (d *Date) parseTime(s string, isEndTime bool) error {
	if s == "" {
		return nil
	}

	layout := noTimeDateLayout
	if strings.Contains(s, "T") {
		d.HasTime = true
		layout = time.RFC3339
	}

	t, err := time.Parse(layout, s)
	if err != nil {
		return err
	}

	if isEndTime {
		d.End = t
	} else {
		d.Start = t
	}
	return nil
}

// IconTypeEnum represents an icon type in the Notion API.
type IconTypeEnum string

// SetValue sets the IconTypeEnum to the given string
func (ite *IconTypeEnum) SetValue(s string) {
	if ite != nil {
		*ite = IconTypeEnum(s)
	}
}

// IsValidEnum returns true if the string represents a valid IconTypeEnum in the Notion API.
func (ite *IconTypeEnum) IsValidEnum() bool {
	return ite != nil && isValidEnum(string(*ite), IconTypeEnumFile, IconTypeEnumEmoji)
}

// UnmarshalJSON returns an error if the type is not a valid enum in the Notion API.
func (ite *IconTypeEnum) UnmarshalJSON(b []byte) error {
	return unmarshalEnum(b, ite)
}

// Icons represents the icon and cover icon for pages and databases in the Notion API
type Icons struct {
	Cover File `json:"cover,omitempty"`
	Icon  Icon `json:"icon,omitempty"`
}

// An Icon represents a page or database icon in the Notion API
type Icon struct {
	Type  IconTypeEnum `json:"type"`
	Emoji string       `json:"emoji,omitempty"`
	File  File         `json:"file,omitempty"`
}

type icon Icon

// GetType returns the type of the page property
func (i *icon) getType() string {
	if i == nil {
		return ""
	}
	return string(i.Type)
}

// FieldsToExpand implements the expander interface for Icon
func (i *icon) fieldsToExpand() []string {
	return []string{"emoji", "file"}
}

// UnmarshalJSON flattens the icon by its type
func (i *Icon) UnmarshalJSON(b []byte) error {
	ii := new(icon)
	if err := unmarshalJSONFlattenByType(b, ii); err != nil {
		return err
	}

	*i = Icon(*ii)
	return nil
}

// MarshalJSON marshals the Icon to be compatible with the Notion API.
func (i *Icon) MarshalJSON() ([]byte, error) {
	ii := icon(*i)
	return marshalJSONExpandByType(&ii)
}

type jsonURL url.URL

func (ju *jsonURL) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	if u, err := url.Parse(s); err == nil {
		*ju = jsonURL(*u)
	} else {
		return err
	}
	return nil
}

func (ju *jsonURL) MarshalJSON() ([]byte, error) {
	if ju == nil {
		return nil, nil
	}
	u := url.URL(*ju)
	return json.Marshal((&u).String())
}

type jsonEmail mail.Address

func (je *jsonEmail) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	if a, err := mail.ParseAddress(s); err == nil {
		*je = jsonEmail(*a)
	} else {
		return err
	}
	return nil
}

func (je *jsonEmail) MarshalJSON() ([]byte, error) {
	if je == nil {
		return nil, nil
	}
	a := mail.Address(*je)
	a.Name = ""
	return json.Marshal((&a).String())
}

func unmarshalJSONFlattenByType(bt []byte, b typed) error {
	if err := json.Unmarshal(bt, b); err != nil {
		return err
	}

	m := make(map[string]interface{})
	if err := json.Unmarshal(bt, &m); err != nil {
		return err
	}

	if bt, err := json.Marshal(m[b.getType()]); err == nil {
		if err := json.Unmarshal(bt, b); err != nil { //nolint:govet
			return err
		}
	} else {
		return err
	}

	return nil
}

func marshalJSONExpandByType(v expander) ([]byte, error) {
	var (
		b   []byte
		err error
	)
	if b, err = json.Marshal(v); err != nil {
		return nil, err
	}

	m := make(map[string]interface{})
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}

	typeMap := make(map[string]interface{})
	for _, field := range v.fieldsToExpand() {
		// Need to check the existence of a key because not all expanded fields will exist for all types.
		if v, ok := m[field]; ok {
			typeMap[field] = v
			delete(m, field)
		}
	}

	return json.Marshal(typeMap)
}
