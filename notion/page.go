package notion

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// These represent the enum types for pages in the Notion API.
const (
	ParentTypeEnumWorkspace = "workspace"
	ParentTypeEnumPage      = "page_id"
	ParentTypeEnumDatabase  = "database_id"

	FormulaTypeEnumString  = "string"
	FormulaTypeEnumNumber  = "number"
	FormulaTypeEnumBoolean = "boolean"
	FormulaTypeEnumDate    = "date"

	RollupValueTypeEnumNumber = "number"
	RollupValueTypeEnumDate   = "date"
	RollupValueTypeEnumArray  = "array"

	FileTypeEnumFile     = "file"
	FileTypeEnumExternal = "external"
)

// ParentTypeEnum represents the parent type in the Notion API.
type ParentTypeEnum string

// SetValue sets the ParentTypeEnum to the given string
func (pte *ParentTypeEnum) SetValue(s string) {
	if pte != nil {
		*pte = ParentTypeEnum(s)
	}
}

// IsValidEnum returns true if the string represents a valid ParentTypeEnum in the Notion API.
func (pte *ParentTypeEnum) IsValidEnum() bool {
	return pte != nil && isValidEnum(string(*pte), ParentTypeEnumWorkspace, ParentTypeEnumPage, ParentTypeEnumDatabase)
}

// UnmarshalJSON returns an error if the type is not a valid enum in the Notion API.
func (pte *ParentTypeEnum) UnmarshalJSON(b []byte) error {
	return unmarshalEnum(b, pte)
}

// FormulaTypeEnum represents the parent type in the Notion API.
type FormulaTypeEnum string

// SetValue sets the FormulaTypeEnum to the given string
func (fte *FormulaTypeEnum) SetValue(s string) {
	if fte != nil {
		*fte = FormulaTypeEnum(s)
	}
}

// IsValidEnum returns true if the string represents a valid FormulaTypeEnum in the Notion API.
func (fte *FormulaTypeEnum) IsValidEnum() bool {
	return fte != nil && isValidEnum(string(*fte), FormulaTypeEnumString, FormulaTypeEnumNumber, FormulaTypeEnumBoolean, FormulaTypeEnumDate)
}

// UnmarshalJSON returns an error if the type is not a valid enum in the Notion API.
func (fte *FormulaTypeEnum) UnmarshalJSON(b []byte) error {
	return unmarshalEnum(b, fte)
}

// RollupValueTypeEnum represents a valid rollup value type for a page in a database in the Notion API.
type RollupValueTypeEnum string

// SetValue sets the RollupValueTypeEnum to the given string
func (rvte *RollupValueTypeEnum) SetValue(s string) {
	if rvte != nil {
		*rvte = RollupValueTypeEnum(s)
	}
}

// IsValidEnum returns true if the string represents a valid RollupValueTypeEnum in the Notion API.
func (rvte *RollupValueTypeEnum) IsValidEnum() bool {
	return rvte != nil && isValidEnum(string(*rvte), RollupValueTypeEnumNumber, RollupValueTypeEnumDate, RollupValueTypeEnumArray)
}

// UnmarshalJSON returns an error if the type is not a valid enum in the Notion API.
func (rvte *RollupValueTypeEnum) UnmarshalJSON(b []byte) error {
	return unmarshalEnum(b, rvte)
}

// FileTypeEnum represents a file type in the Notion API.
type FileTypeEnum string

// SetValue sets the FileTypeEnum to the given string
func (pte *FileTypeEnum) SetValue(s string) {
	if pte != nil {
		*pte = FileTypeEnum(s)
	}
}

// IsValidEnum returns true if the string represents a valid FileTypeEnum in the Notion API.
func (pte *FileTypeEnum) IsValidEnum() bool {
	return pte != nil && isValidEnum(string(*pte), FileTypeEnumFile, FileTypeEnumExternal)
}

// UnmarshalJSON returns an error if the type is not a valid enum in the Notion API.
func (pte *FileTypeEnum) UnmarshalJSON(b []byte) error {
	return unmarshalEnum(b, pte)
}

// Parent represents the parent object of a page in the Notion API.
type Parent struct {
	Type ParentTypeEnum `json:"type"`
	ID   string
}

// UnmarshalJSON sets the ID field based on the parent type
func (p *Parent) UnmarshalJSON(b []byte) error {
	m := make(map[string]interface{})
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}
	if t, ok := m["type"].(string); ok {
		p.Type = ParentTypeEnum(t)
	} else {
		return fmt.Errorf("the type parameter of parent is not a string")
	}

	if id, ok := m[string(p.Type)].(string); ok {
		p.ID = id
	}

	return nil
}

// MarshalJSON sets the correct type id based on the type.
func (p *Parent) MarshalJSON() ([]byte, error) {
	if p == nil {
		return nil, nil
	}

	m := map[string]interface{}{"type": p.Type}
	if string(p.Type) != ParentTypeEnumWorkspace {
		m[string(p.Type)] = p.ID
	}
	return json.Marshal(m)
}

// Formula represents a formula entry in a page in the Notion API.
type Formula struct {
	Type    FormulaTypeEnum `json:"type"`
	String  *string         `json:"string,omitempty"`
	Number  *float64        `json:"number,omitempty"`
	Boolean *bool           `json:"boolean,omitempty"`
	Date    *Date           `json:"date,omitempty"`
}

type formula Formula

// GetType returns the type of the Formula
func (f *formula) getType() string {
	if f == nil {
		return ""
	}
	return string(f.Type)
}

// FieldsToExpand implements the expander interface for marshaling
func (f *formula) fieldsToExpand() []string {
	return []string{"string", "number", "boolean", "date"}
}

// MarshalJSON expands the formula by its type to be compatible with the Notion API.
func (f *Formula) MarshalJSON() ([]byte, error) {
	ff := formula(*f)
	return marshalJSONExpandByType(&ff)
}

// A RelationID represents a single relation in a page in a database in the Notion API.
type RelationID UUID4

type relationID RelationID

// GetType is a hack to use the unmarshaler functions. It has not meaning in the context of a RelationID.
func (rid *relationID) getType() string {
	return "id"
}

// UnmarshalJSON flattens a relation reference for a page in a database in the Notion API.
func (rid *RelationID) UnmarshalJSON(b []byte) error {
	r := new(relationID)
	if err := unmarshalJSONFlattenByType(b, r); err != nil {
		return err
	}

	*rid = RelationID(*r)
	return nil
}

// MarshalJSON expands a RelationID to be compatible with the Notion API.
func (rid *RelationID) MarshalJSON() ([]byte, error) {
	if rid == nil {
		return nil, nil
	}
	return json.Marshal(map[string]string{"id": uuid.UUID(*rid).String()})
}

// Relations represent a set of relations for a page in a database in the Notion API.
type Relations []*RelationID

// RollupValue represents the value of a rollup property of a page in a database in the Notion API.
type RollupValue struct {
	Type   RollupValueTypeEnum `json:"type"`
	Number *float64            `json:"number,omitempty"`
	Date   *Date               `json:"date,omitempty"`
	Array  []*PageProperty     `json:"array,omitempty"`
}

// A File represents a file property of a page in a database in the Notion API.
type File struct {
	Name       string       `json:"name"`
	Type       FileTypeEnum `json:"type"`
	URL        *jsonURL     `json:"url"`
	ExpiryTime time.Time    `json:"expiry_time,omitempty"`
}

type file File

// GetType returns the type of the page property
func (f *file) getType() string {
	if f == nil {
		return ""
	}
	return string(f.Type)
}

// FieldsToExpand implements the expander interface for File
func (f *file) fieldsToExpand() []string {
	return []string{"url", "expiry_time"}
}

// UnmarshalJSON flattens the page property by its type
func (f *File) UnmarshalJSON(b []byte) error {
	ff := new(file)
	if err := unmarshalJSONFlattenByType(b, ff); err != nil {
		return err
	}

	*f = File(*ff)
	return nil
}

// MarshalJSON marshals the File to be compatible with the Notion API.
func (f *File) MarshalJSON() ([]byte, error) {
	ff := file(*f)
	return marshalJSONExpandByType(&ff)
}

// PageProperty represents a property of a page in a database in the Notion API.
type PageProperty struct {
	Editable
	Name         string                   `json:"-"`
	ID           string                   `json:"id,omitempty"`
	Type         DatabasePropertyTypeEnum `json:"type,omitempty"`
	Title        []RichText               `json:"title,omitempty"`
	RichText     []RichText               `json:"rich_text,omitempty"`
	Number       *float64                 `json:"number,omitempty"`
	Select       *SelectOption            `json:"select,omitempty"`
	MultiSelect  []SelectOption           `json:"multi_select,omitempty"`
	Date         *Date                    `json:"date,omitempty"`
	Formula      *Formula                 `json:"formula,omitempty"`
	Relations    Relations                `json:"relation,omitempty"`
	People       []*User                  `json:"people,omitempty"`
	Files        []*File                  `json:"files,omitempty"`
	Checkbox     *bool                    `json:"checkbox,omitempty"`
	URL          *jsonURL                 `json:"url,omitempty"`
	Email        *jsonEmail               `json:"email,omitempty"`
	PhoneNumber  *string                  `json:"phone_number,omitempty"`
	CreatedBy    *User                    `json:"created_by,omitempty"`
	LastEditedBy *User                    `json:"last_edited_by,omitempty"`
}

type pageProperty PageProperty

// GetType returns the type of the page property
func (pp *pageProperty) getType() string {
	if pp == nil {
		return ""
	}
	return string(pp.Type)
}

// FieldsToExpand implements the expander interface for PageProperty
func (pp *pageProperty) fieldsToExpand() []string {
	return []string{"title", "rich_text", "number", "select", "multi_select", "date", "formula", "relation",
		"people", "files", "checkbox", "url", "email", "phone_number", "created_by", "last_edited_by", "created_time", "last_edited_time"}
}

// UnmarshalJSON flattens the page property by its type
func (pp *PageProperty) UnmarshalJSON(b []byte) error {
	ppp := new(pageProperty)
	if err := unmarshalJSONFlattenByType(b, ppp); err != nil {
		return err
	}

	*pp = PageProperty(*ppp)
	return nil
}

// MarshalJSON marshals the page property to be compatible with the Notion API.
func (pp *PageProperty) MarshalJSON() ([]byte, error) {
	ppp := pageProperty(*pp)
	return marshalJSONExpandByType(&ppp)
}

// PageProperties represents the properties of a page in a database in the Notion API.
type PageProperties []*PageProperty

// UnmarshalJSON unmarshals the page properties from a "javascript" style to a more "go" style
func (pps *PageProperties) UnmarshalJSON(b []byte) error {
	var m map[string]*pageProperty
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}

	*pps = make([]*PageProperty, 0, len(m))
	for name, prop := range m {
		prop.Name = name
		p := PageProperty(*prop)
		*pps = append(*pps, &p)
	}

	return nil
}

// MarshalJSON marshals the database properties into a form that is compatible with the Notion API.
func (pps *PageProperties) MarshalJSON() ([]byte, error) {
	if pps == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	for _, prop := range *pps {
		if prop != nil {
			m[prop.Name] = prop
		}
	}

	return json.Marshal(m)
}

// A Page represents a page object in the Notion API.
type Page struct {
	Object
	Editable
	Icons
	Archived   bool           `json:"archived"`
	Parent     Parent         `json:"parent"`
	Properties PageProperties `json:"properties"`
	Children   Blocks         `json:"children,omitempty"`
	URL        *jsonURL       `json:"url"`
}

// Pages is a slice of pages from the Notion API.
type Pages []*Page

// Len returns the number of pages in the slice.
func (pgs *Pages) Len() int {
	if pgs == nil {
		return 0
	}
	return len(*pgs)
}

type pages Pages

// UnmarshalJSON appends the unmarshaled Users to the slice.
func (pgs *Pages) UnmarshalJSON(b []byte) error {
	p := new(pages)
	if err := json.Unmarshal(b, p); err != nil || len(*p) == 0 {
		return err
	}

	if pgs == nil {
		*pgs = make([]*Page, 0, len(*p))
	}
	*pgs = append(*pgs, []*Page(*p)...)
	return nil
}
