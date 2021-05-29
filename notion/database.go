package notion

import "encoding/json"

// These represent the valid enums for database properties int the Notion API
const (
	DatabasePropertyTypeEnumTitle          = "title"
	DatabasePropertyTypeEnumRichText       = "rich_text"
	DatabasePropertyTypeEnumNumber         = "number"
	DatabasePropertyTypeEnumSelect         = "select"
	DatabasePropertyTypeEnumMultiSelect    = "multi_select"
	DatabasePropertyTypeEnumDate           = "date"
	DatabasePropertyTypeEnumPeople         = "people"
	DatabasePropertyTypeEnumFile           = "file"
	DatabasePropertyTypeEnumCheckbox       = "checkbox"
	DatabasePropertyTypeEnumURL            = "url"
	DatabasePropertyTypeEnumEmail          = "email"
	DatabasePropertyTypeEnumPhoneNumber    = "phone_number"
	DatabasePropertyTypeEnumFormula        = "formula"
	DatabasePropertyTypeEnumRelation       = "relation"
	DatabasePropertyTypeEnumRollup         = "rollup"
	DatabasePropertyTypeEnumCreatedTime    = "created_time"
	DatabasePropertyTypeEnumLastEditedTime = "last_edited_time"
	DatabasePropertyTypeEnumLastEditedBy   = "last_edited_by"

	SelectColorEnumDefault = "default"
	SelectColorEnumGray    = "gray"
	SelectColorEnumBrown   = "brown"
	SelectColorEnumOrange  = "orange"
	SelectColorEnumYellow  = "yellow"
	SelectColorEnumGreen   = "green"
	SelectColorEnumBlue    = "blue"
	SelectColorEnumPurple  = "purple"
	SelectColorEnumPink    = "pink"
	SelectColorEnumRed     = "red"

	RollupFunctionEnumCountAll          = "count_all"
	RollupFunctionEnumCountValues       = "count_values"
	RollupFunctionEnumCountUniqueValues = "count_unique_values"
	RollupFunctionEnumCountEmpty        = "count_empty"
	RollupFunctionEnumCountNotEmpty     = "count_not_empty"
	RollupFunctionEnumPercentEmpty      = "percent_empty"
	RollupFunctionEnumPercentNotEmpty   = "percent_not_empty"
	RollupFunctionEnumSum               = "sum"
	RollupFunctionEnumAverage           = "average"
	RollupFunctionEnumMedian            = "median"
	RollupFunctionEnumMin               = "min"
	RollupFunctionEnumMax               = "max"
	RollupFunctionEnumRange             = "range"

	NumberConfigurationTypeEnumNumber           = "number"
	NumberConfigurationTypeEnumNumberWithCommas = "number_with_commas"
	NumberConfigurationTypeEnumPercent          = "percent"
	NumberConfigurationTypeEnumDollar           = "dollar"
	NumberConfigurationTypeEnumEuro             = "euro"
	NumberConfigurationTypeEnumPound            = "pound"
	NumberConfigurationTypeEnumYen              = "yen"
	NumberConfigurationTypeEnumRuble            = "ruble"
	NumberConfigurationTypeEnumRupee            = "rupee"
	NumberConfigurationTypeEnumWon              = "won"
	NumberConfigurationTypeEnumYuan             = "yuan"
)

// A DatabasePropertyTypeEnum represents a valid database property type in the Notion API.
type DatabasePropertyTypeEnum string

// SetValue sets the DatabasePropertyTypeEnum to the given string
func (dbte *DatabasePropertyTypeEnum) SetValue(s string) {
	if dbte != nil {
		*dbte = DatabasePropertyTypeEnum(s)
	}
}

// IsValidEnum returns true if the string represents a valid DatabasePropertyTypeEnum in the Notion API.
func (dbte *DatabasePropertyTypeEnum) IsValidEnum() bool {
	return dbte != nil && isValidEnum(string(*dbte), DatabasePropertyTypeEnumTitle,
		DatabasePropertyTypeEnumRichText,
		DatabasePropertyTypeEnumNumber,
		DatabasePropertyTypeEnumSelect,
		DatabasePropertyTypeEnumMultiSelect,
		DatabasePropertyTypeEnumDate,
		DatabasePropertyTypeEnumPeople,
		DatabasePropertyTypeEnumFile,
		DatabasePropertyTypeEnumCheckbox,
		DatabasePropertyTypeEnumURL,
		DatabasePropertyTypeEnumEmail,
		DatabasePropertyTypeEnumPhoneNumber,
		DatabasePropertyTypeEnumFormula,
		DatabasePropertyTypeEnumRelation,
		DatabasePropertyTypeEnumRollup,
		DatabasePropertyTypeEnumCreatedTime,
		DatabasePropertyTypeEnumLastEditedTime,
		DatabasePropertyTypeEnumLastEditedBy,
	)
}

// UnmarshalJSON verifies that the string is a valid DatabasePropertyTypeEnum for the Notion API.
func (dbte *DatabasePropertyTypeEnum) UnmarshalJSON(b []byte) error {
	return unmarshalEnum(b, dbte)
}

// SelectColorEnum represents the valid select colors in the Notion API
type SelectColorEnum string

// SetValue sets the SelectColorEnum to the given string
func (sce *SelectColorEnum) SetValue(s string) {
	if sce != nil {
		*sce = SelectColorEnum(s)
	}
}

// IsValidEnum returns true if the string represents a valid SelectColorEnum in the Notion API.
func (sce *SelectColorEnum) IsValidEnum() bool {
	return sce != nil && isValidEnum(string(*sce), SelectColorEnumDefault,
		SelectColorEnumGray,
		SelectColorEnumBrown,
		SelectColorEnumOrange,
		SelectColorEnumYellow,
		SelectColorEnumGreen,
		SelectColorEnumBlue,
		SelectColorEnumPurple,
		SelectColorEnumPink,
		SelectColorEnumRed,
	)
}

// UnmarshalJSON verifies that the string is a valid DatabasePropertyTypeEnum for the Notion API.
func (sce *SelectColorEnum) UnmarshalJSON(b []byte) error {
	return unmarshalEnum(b, sce)
}

// RollupFunctionEnumType represents a valid rollup function in a database in the Notion API.
type RollupFunctionEnumType string

// SetValue sets the RollupFunctionEnumType to the given string
func (rfet *RollupFunctionEnumType) SetValue(s string) {
	if rfet != nil {
		*rfet = RollupFunctionEnumType(s)
	}
}

// IsValidEnum returns true if the string is a valid RollupFunctionEnumType in the Notion API.
func (rfet *RollupFunctionEnumType) IsValidEnum() bool {
	return rfet != nil && isValidEnum(string(*rfet), RollupFunctionEnumCountAll,
		RollupFunctionEnumCountValues,
		RollupFunctionEnumCountUniqueValues,
		RollupFunctionEnumCountEmpty,
		RollupFunctionEnumCountNotEmpty,
		RollupFunctionEnumPercentEmpty,
		RollupFunctionEnumPercentNotEmpty,
		RollupFunctionEnumSum,
		RollupFunctionEnumAverage,
		RollupFunctionEnumMedian,
		RollupFunctionEnumMin,
		RollupFunctionEnumMax,
		RollupFunctionEnumRange,
	)
}

// UnmarshalJSON verifies that the string is a valid RollupFunctionEnumType for the Notion API.
func (rfet *RollupFunctionEnumType) UnmarshalJSON(b []byte) error {
	return unmarshalEnum(b, rfet)
}

// NumberConfigurationTypeEnum represents a number type in a database in the Notion API.
type NumberConfigurationTypeEnum string

// SetValue sets the NumberConfigurationTypeEnum to the given string
func (ncte *NumberConfigurationTypeEnum) SetValue(s string) {
	if ncte != nil {
		*ncte = NumberConfigurationTypeEnum(s)
	}
}

// IsValidEnum returns true if the string is a valid NumberConfigurationTypeEnum in the Notion API.
func (ncte *NumberConfigurationTypeEnum) IsValidEnum() bool {
	return ncte != nil && isValidEnum(string(*ncte), NumberConfigurationTypeEnumNumber,
		NumberConfigurationTypeEnumNumberWithCommas,
		NumberConfigurationTypeEnumPercent,
		NumberConfigurationTypeEnumDollar,
		NumberConfigurationTypeEnumEuro,
		NumberConfigurationTypeEnumPound,
		NumberConfigurationTypeEnumYen,
		NumberConfigurationTypeEnumRuble,
		NumberConfigurationTypeEnumRupee,
		NumberConfigurationTypeEnumWon,
		NumberConfigurationTypeEnumYuan,
	)
}

// UnmarshalJSON verifies that the string is a valid NumberConfigurationTypeEnum  in the Notion API
func (ncte *NumberConfigurationTypeEnum) UnmarshalJSON(b []byte) error {
	return unmarshalEnum(b, ncte)
}

// SelectOption represents the (multi)select options for database properties in the Notion API.
type SelectOption struct {
	Name  string          `json:"name"`
	ID    string          `json:"id"`
	Color SelectColorEnum `json:"color"`
}

// Relation represents a database relation in the Notion API.
type Relation struct {
	DatabaseID         UUID4   `json:"database_id"`
	SyncedPropertyName *string `json:"synced_property_name,omitempty"`
	SyncedPropertyID   *string `json:"synced_property_id,omitempty"`
}

// Rollup represents a rollup configuration in the Notion API.
type Rollup struct {
	RelationPropertyName string                 `json:"relation_property_name"`
	RelationPropertyID   string                 `json:"relation_property_id"`
	RollupPropertyName   string                 `json:"rollup_property_name"`
	RollupPropertyID     string                 `json:"rollup_property_id"`
	Function             RollupFunctionEnumType `json:"function"`
}

// A Database represents a database object in the Notion API
type Database struct {
	Object
	Editable
	Title      []RichText         `json:"title"`
	Properties DatabaseProperties `json:"properties"`
	Children   Blocks             `json:"-"`
}

// DatabaseProperties is a short-hand type for a slice of DatabaseProperty,
// used for (un)marshaling purposes
type DatabaseProperties []*DatabaseProperty

// UnmarshalJSON unmarshals the database properties from a "javascript" style to a more "go" style
func (dbps *DatabaseProperties) UnmarshalJSON(b []byte) error {
	var m map[string]*DatabaseProperty
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}

	*dbps = make([]*DatabaseProperty, 0, len(m))
	for name, prop := range m {
		prop.Name = name
		*dbps = append(*dbps, prop)
	}

	return nil
}

// MarshalJSON marshals the database properties into a form that is compatible with the Notion API.
func (dbps *DatabaseProperties) MarshalJSON() ([]byte, error) {
	if dbps == nil {
		return nil, nil
	}

	m := make(map[string]interface{})
	for _, prop := range *dbps {
		if prop != nil {
			m[prop.Name] = prop
		}
	}

	return json.Marshal(m)
}

// DatabaseProperty represents the properties of the items in a database
type DatabaseProperty struct {
	Name              string
	ID                string                       `json:"id"`
	Type              DatabasePropertyTypeEnum     `json:"type"`
	NumberFormat      *NumberConfigurationTypeEnum `json:"format,omitempty"`
	SelectOptions     []SelectOption               `json:"options,omitempty"`
	FormulaExpression *string                      `json:"expression,omitempty"`
	Relation          *Relation                    `json:"relation,omitempty"`
	Rollup            *Rollup                      `json:"rollup_configuration,omitempty"`
}

type databaseProperty DatabaseProperty

// GetName returns the name of the database property
func (dp *DatabaseProperty) GetName() string {
	if dp == nil {
		return ""
	}
	return dp.Name
}

// GetType returns the type of the database property
func (dp *databaseProperty) getType() string {
	if dp == nil {
		return ""
	}
	return string(dp.Type)
}

// FieldsToExpand implements the expander interface
func (dp *databaseProperty) fieldsToExpand() []string {
	return []string{"format", "options", "expression", "relation", "rollup_configuration"}
}

// UnmarshalJSON sets the JSON object based on the Type of the DatabaseProperty object.
func (dp *DatabaseProperty) UnmarshalJSON(bt []byte) error {
	d := new(databaseProperty)
	if err := unmarshalJSONFlattenByType(bt, d); err != nil {
		return err
	}

	*dp = DatabaseProperty(*d)
	return nil
}

// MarshalJSON returns a marshaled version of the Block that is compatible with the Notion API.
func (dp *DatabaseProperty) MarshalJSON() ([]byte, error) {
	if dp == nil {
		return nil, nil
	}
	d := databaseProperty(*dp)
	return marshalJSONExpandByType(&d)
}

// Databases represents a list of users from the Notion API.
type Databases []*Database

// Len returns the number of blocks in the the slice.
func (dbs *Databases) Len() int {
	if dbs == nil {
		return 0
	}
	return len(*dbs)
}

type databases Databases

// UnmarshalJSON appends the unmarshaled Databases to the slice.
func (dbs *Databases) UnmarshalJSON(b []byte) error {
	db := new(databases)
	if err := json.Unmarshal(b, db); err != nil || len(*db) == 0 {
		return err
	}

	if dbs == nil {
		*dbs = make([]*Database, 0, len(*db))
	}
	*dbs = append(*dbs, []*Database(*db)...)
	return nil
}
