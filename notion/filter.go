package notion

import (
	"encoding/json"
	"time"
)

// These constants represent valid enum values for filter objects in the Notion API.
const (
	FilterConditionTypeEnumText        = "text"
	FilterConditionTypeEnumNumber      = "number"
	FilterConditionTypeEnumCheckbox    = "checkbox"
	FilterConditionTypeEnumSelect      = "select"
	FilterConditionTypeEnumMultiSelect = "multi_select"
	FilterConditionTypeEnumDate        = "date"
	FilterConditionTypeEnumPeople      = "people"
	FilterConditionTypeEnumFiles       = "files"
	FilterConditionTypeEnumRelation    = "relation"
	FilterConditionTypeEnumFormula     = "formula"
	FilterConditionTypeEnumObject      = "object"

	FilterTextConditionEnumEquals         = "equals"
	FilterTextConditionEnumDoesNotEqual   = "does_not_equal"
	FilterTextConditionEnumContains       = "contains"
	FilterTextConditionEnumDoesNotContain = "does_not_contain"
	FilterTextConditionEnumStartsWith     = "starts_with"
	FilterTextConditionEnumEndWith        = "ends_with"
	FilterTextConditionEnumIsEmpty        = "is_empty"
	FilterTextConditionEnumIsNotEmpty     = "is_not_empty"

	FilterNumberConditionEnumEquals               = "equals"
	FilterNumberConditionEnumDoesNotEqual         = "does_not_equal"
	FilterNumberConditionEnumGreaterThan          = "greater_than"
	FilterNumberConditionEnumLessThan             = "less_than"
	FilterNumberConditionEnumGreaterThanOrEqualTo = "greater_than_or_equal_to"
	FilterNumberConditionEnumLessThanOrEqualTo    = "less_than_or_equal_to"
	FilterNumberConditionEnumIsEmpty              = "is_empty"
	FilterNumberConditionEnumIsNotEmpty           = "is_not_empty"

	FilterCheckboxConditionEnumEquals       = "equals"
	FilterCheckboxConditionEnumDoesNotEqual = "does_not_equal"

	FilterSelectConditionEnumEquals       = "equals"
	FilterSelectConditionEnumDoesNotEqual = "does_not_equal"
	FilterSelectConditionEnumIsEmpty      = "is_empty"
	FilterSelectConditionEnumIsNotEmpty   = "is_not_empty"

	FilterMultiSelectConditionEnumContains       = "contains"
	FilterMultiSelectConditionEnumDoesNotContain = "does_not_contain"
	FilterMultiSelectConditionEnumIsEmpty        = "is_empty"
	FilterMultiSelectConditionEnumIsNotEmpty     = "is_not_empty"

	FilterDateConditionEnumEquals     = "equals"
	FilterDateConditionEnumBefore     = "before"
	FilterDateConditionEnumAfter      = "after"
	FilterDateConditionEnumOnOrBefore = "on_or_before"
	FilterDateConditionEnumOnOrAfter  = "on_or_after"
	FilterDateConditionEnumIsEmpty    = "is_empty"
	FilterDateConditionEnumIsNotEmpty = "is_not_empty"
	FilterDateConditionEnumPastWeek   = "past_week"
	FilterDateConditionEnumPastMonth  = "past_month"
	FilterDateConditionEnumPastYear   = "past_year"
	FilterDateConditionEnumNextWeek   = "next_week"
	FilterDateConditionEnumNextMonth  = "next_month"
	FilterDateConditionEnumNextYear   = "next_year"

	FilterPeopleConditionEnumContains       = "contains"
	FilterPeopleConditionEnumDoesNotContain = "does_not_contain"
	FilterPeopleConditionEnumIsEmpty        = "is_empty"
	FilterPeopleConditionEnumIsNotEmpty     = "is_not_empty"

	FilterFilesConditionEnumIsEmpty    = "is_empty"
	FilterFilesConditionEnumIsNotEmpty = "is_not_empty"

	FilterRelationConditionEnumContains       = "contains"
	FilterRelationConditionEnumDoesNotContain = "does_not_contain"
	FilterRelationConditionEnumIsEmpty        = "is_empty"
	FilterRelationConditionEnumIsNotEmpty     = "is_not_empty"

	FilterFormulaConditionEnumText     = "text"
	FilterFormulaConditionEnumCheckbox = "checkbox"
	FilterFormulaConditionEnumNumber   = "number"
	FilterFormulaConditionEnumDate     = "date"

	FilterObjectConditionEnumPage     = "page"
	FilterObjectConditionEnumDatabase = "database"
)

// A FilterConditionTypeEnum represents the valid filter types in the Notion API.
type FilterConditionTypeEnum string

// SetValue sets the FilterConditionTypeEnum to the given string
func (fcte *FilterConditionTypeEnum) SetValue(s string) {
	if fcte != nil {
		*fcte = FilterConditionTypeEnum(s)
	}
}

// IsValidEnum returns true if the string represents a valid FilterConditionTypeEnum in the Notion API.
func (fcte *FilterConditionTypeEnum) IsValidEnum() bool {
	return fcte != nil && isValidEnum(string(*fcte), FilterConditionTypeEnumText,
		FilterConditionTypeEnumNumber,
		FilterConditionTypeEnumCheckbox,
		FilterConditionTypeEnumSelect,
		FilterConditionTypeEnumMultiSelect,
		FilterConditionTypeEnumDate,
		FilterConditionTypeEnumPeople,
		FilterConditionTypeEnumFiles,
		FilterConditionTypeEnumRelation,
		FilterConditionTypeEnumFormula,
	)
}

// UnmarshalJSON returns an error if the string is not a valid FilterConditionTypeEnum in the Notion API.
func (fcte *FilterConditionTypeEnum) UnmarshalJSON(b []byte) error {
	return unmarshalEnum(b, fcte)
}

// A FilterTextConditionEnum represents a valid text filter in the Notion API.
type FilterTextConditionEnum string

// SetValue sets the FilterTextConditionEnum to the given string
func (ftce *FilterTextConditionEnum) SetValue(s string) {
	if ftce != nil {
		*ftce = FilterTextConditionEnum(s)
	}
}

// IsValidEnum returns true if the string represents a valid FilterTextConditionEnum in the Notion API.
func (ftce *FilterTextConditionEnum) IsValidEnum() bool {
	return ftce != nil && isValidEnum(string(*ftce), FilterTextConditionEnumEquals,
		FilterTextConditionEnumDoesNotEqual,
		FilterTextConditionEnumContains,
		FilterTextConditionEnumDoesNotContain,
		FilterTextConditionEnumStartsWith,
		FilterTextConditionEnumEndWith,
		FilterTextConditionEnumIsEmpty,
		FilterTextConditionEnumIsNotEmpty,
	)
}

// UnmarshalJSON returns an error if the string is not a valid FilterTextConditionEnum in the Notion API.
func (ftce *FilterTextConditionEnum) UnmarshalJSON(b []byte) error {
	return unmarshalEnum(b, ftce)
}

// A FilterNumberConditionEnum represents a valid text filter in the Notion API.
type FilterNumberConditionEnum string

// SetValue sets the FilterNumberConditionEnum to the given string
func (fnce *FilterNumberConditionEnum) SetValue(s string) {
	if fnce != nil {
		*fnce = FilterNumberConditionEnum(s)
	}
}

// IsValidEnum returns true if the string represents a valid FilterNumberConditionEnum in the Notion API.
func (fnce *FilterNumberConditionEnum) IsValidEnum() bool {
	return fnce != nil && isValidEnum(string(*fnce), FilterNumberConditionEnumEquals,
		FilterNumberConditionEnumDoesNotEqual,
		FilterNumberConditionEnumGreaterThan,
		FilterNumberConditionEnumLessThan,
		FilterNumberConditionEnumGreaterThanOrEqualTo,
		FilterNumberConditionEnumLessThanOrEqualTo,
		FilterNumberConditionEnumIsEmpty,
		FilterNumberConditionEnumIsNotEmpty,
	)
}

// UnmarshalJSON returns an error if the string is not a valid FilterNumberConditionEnum in the Notion API.
func (fnce *FilterNumberConditionEnum) UnmarshalJSON(b []byte) error {
	return unmarshalEnum(b, fnce)
}

// A FilterCheckboxConditionEnum represents a valid text filter in the Notion API.
type FilterCheckboxConditionEnum string

// SetValue sets the FilterCheckboxConditionEnum to the given string
func (fcce *FilterCheckboxConditionEnum) SetValue(s string) {
	if fcce != nil {
		*fcce = FilterCheckboxConditionEnum(s)
	}
}

// IsValidEnum returns true if the string represents a valid FilterCheckboxConditionEnum in the Notion API.
func (fcce *FilterCheckboxConditionEnum) IsValidEnum() bool {
	return fcce != nil && isValidEnum(string(*fcce), FilterCheckboxConditionEnumEquals, FilterCheckboxConditionEnumDoesNotEqual)
}

// UnmarshalJSON returns an error if the string is not a valid FilterCheckboxConditionEnum in the Notion API.
func (fcce *FilterCheckboxConditionEnum) UnmarshalJSON(b []byte) error {
	return unmarshalEnum(b, fcce)
}

// A FilterSelectConditionEnum represents a valid text filter in the Notion API.
type FilterSelectConditionEnum string

// SetValue sets the FilterSelectConditionEnum to the given string
func (fsce *FilterSelectConditionEnum) SetValue(s string) {
	if fsce != nil {
		*fsce = FilterSelectConditionEnum(s)
	}
}

// IsValidEnum returns true if the string represents a valid FilterSelectConditionEnum in the Notion API.
func (fsce *FilterSelectConditionEnum) IsValidEnum() bool {
	return fsce != nil && isValidEnum(string(*fsce), FilterSelectConditionEnumEquals,
		FilterSelectConditionEnumDoesNotEqual,
		FilterSelectConditionEnumIsEmpty,
		FilterSelectConditionEnumIsNotEmpty,
	)
}

// UnmarshalJSON returns an error if the string is not a valid FilterSelectConditionEnum in the Notion API.
func (fsce *FilterSelectConditionEnum) UnmarshalJSON(b []byte) error {
	return unmarshalEnum(b, fsce)
}

// A FilterMultiSelectConditionEnum represents a valid text filter in the Notion API.
type FilterMultiSelectConditionEnum string

// SetValue sets the FilterMultiSelectConditionEnum to the given string
func (fmce *FilterMultiSelectConditionEnum) SetValue(s string) {
	if fmce != nil {
		*fmce = FilterMultiSelectConditionEnum(s)
	}
}

// IsValidEnum returns true if the string represents a valid FilterMultiSelectConditionEnum in the Notion API.
func (fmce *FilterMultiSelectConditionEnum) IsValidEnum() bool {
	return fmce != nil && isValidEnum(string(*fmce), FilterMultiSelectConditionEnumContains,
		FilterMultiSelectConditionEnumDoesNotContain,
		FilterMultiSelectConditionEnumIsEmpty,
		FilterMultiSelectConditionEnumIsNotEmpty,
	)
}

// UnmarshalJSON returns an error if the string is not a valid FilterMultiSelectConditionEnum in the Notion API.
func (fmce *FilterMultiSelectConditionEnum) UnmarshalJSON(b []byte) error {
	return unmarshalEnum(b, fmce)
}

// A FilterDateConditionEnum represents a valid text filter in the Notion API.
type FilterDateConditionEnum string

// SetValue sets the FilterDateConditionEnum to the given string
func (fdce *FilterDateConditionEnum) SetValue(s string) {
	if fdce != nil {
		*fdce = FilterDateConditionEnum(s)
	}
}

// IsValidEnum returns true if the string represents a valid FilterDateConditionEnum in the Notion API.
func (fdce *FilterDateConditionEnum) IsValidEnum() bool {
	return fdce != nil && isValidEnum(string(*fdce), FilterDateConditionEnumEquals,
		FilterDateConditionEnumBefore,
		FilterDateConditionEnumAfter,
		FilterDateConditionEnumOnOrBefore,
		FilterDateConditionEnumOnOrAfter,
		FilterDateConditionEnumIsEmpty,
		FilterDateConditionEnumIsNotEmpty,
		FilterDateConditionEnumPastWeek,
		FilterDateConditionEnumPastMonth,
		FilterDateConditionEnumPastYear,
		FilterDateConditionEnumNextWeek,
		FilterDateConditionEnumNextMonth,
		FilterDateConditionEnumNextYear,
	)
}

// UnmarshalJSON returns an error if the string is not a valid FilterDateConditionEnum in the Notion API.
func (fdce *FilterDateConditionEnum) UnmarshalJSON(b []byte) error {
	return unmarshalEnum(b, fdce)
}

// A FilterPeopleConditionEnum represents a valid text filter in the Notion API.
type FilterPeopleConditionEnum string

// SetValue sets the FilterPeopleConditionEnum to the given string
func (fpce *FilterPeopleConditionEnum) SetValue(s string) {
	if fpce != nil {
		*fpce = FilterPeopleConditionEnum(s)
	}
}

// IsValidEnum returns true if the string represents a valid FilterPeopleConditionEnum in the Notion API.
func (fpce *FilterPeopleConditionEnum) IsValidEnum() bool {
	return fpce != nil && isValidEnum(string(*fpce), FilterPeopleConditionEnumContains,
		FilterPeopleConditionEnumDoesNotContain,
		FilterPeopleConditionEnumIsEmpty,
		FilterPeopleConditionEnumIsNotEmpty,
	)
}

// UnmarshalJSON returns an error if the string is not a valid FilterPeopleConditionEnum in the Notion API.
func (fpce *FilterPeopleConditionEnum) UnmarshalJSON(b []byte) error {
	return unmarshalEnum(b, fpce)
}

// A FilterFilesConditionEnum represents a valid text filter in the Notion API.
type FilterFilesConditionEnum string

// SetValue sets the FilterFilesConditionEnum to the given string
func (ffce *FilterFilesConditionEnum) SetValue(s string) {
	if ffce != nil {
		*ffce = FilterFilesConditionEnum(s)
	}
}

// IsValidEnum returns true if the string represents a valid FilterFilesConditionEnum in the Notion API.
func (ffce *FilterFilesConditionEnum) IsValidEnum() bool {
	return ffce != nil && isValidEnum(string(*ffce), FilterFormulaConditionEnumText,
		FilterFormulaConditionEnumCheckbox,
		FilterFormulaConditionEnumNumber,
		FilterFormulaConditionEnumDate,
	)
}

// UnmarshalJSON returns an error if the string is not a valid FilterFilesConditionEnum in the Notion API.
func (ffce *FilterFilesConditionEnum) UnmarshalJSON(b []byte) error {
	return unmarshalEnum(b, ffce)
}

// A FilterRelationConditionEnum represents a valid text filter in the Notion API.
type FilterRelationConditionEnum string

// SetValue sets the FilterRelationConditionEnum to the given string
func (frce *FilterRelationConditionEnum) SetValue(s string) {
	if frce != nil {
		*frce = FilterRelationConditionEnum(s)
	}
}

// IsValidEnum returns true if the string represents a valid FilterRelationConditionEnum in the Notion API.
func (frce *FilterRelationConditionEnum) IsValidEnum() bool {
	return frce != nil && isValidEnum(string(*frce), FilterRelationConditionEnumContains,
		FilterRelationConditionEnumDoesNotContain,
		FilterRelationConditionEnumIsEmpty,
		FilterRelationConditionEnumIsNotEmpty,
	)
}

// UnmarshalJSON returns an error if the string is not a valid FilterRelationConditionEnum in the Notion API.
func (frce *FilterRelationConditionEnum) UnmarshalJSON(b []byte) error {
	return unmarshalEnum(b, frce)
}

// A FilterFormulaConditionEnum represents a valid text filter in the Notion API.
type FilterFormulaConditionEnum string

// SetValue sets the FilterFormulaConditionEnum to the given string
func (ffce *FilterFormulaConditionEnum) SetValue(s string) {
	if ffce != nil {
		*ffce = FilterFormulaConditionEnum(s)
	}
}

// IsValidEnum returns true if the string represents a valid FilterFormulaConditionEnum in the Notion API.
func (ffce *FilterFormulaConditionEnum) IsValidEnum() bool {
	return ffce != nil && isValidEnum(string(*ffce), FilterFilesConditionEnumIsEmpty,
		FilterFilesConditionEnumIsNotEmpty,
	)
}

// UnmarshalJSON returns an error if the string is not a valid FilterFormulaConditionEnum in the Notion API.
func (ffce *FilterFormulaConditionEnum) UnmarshalJSON(b []byte) error {
	return unmarshalEnum(b, ffce)
}

// A FilterObjectConditionEnum represents a valid text filter in the Notion API.
type FilterObjectConditionEnum string

// SetValue sets the FilterObjectConditionEnum to the given string
func (ofce *FilterObjectConditionEnum) SetValue(s string) {
	if ofce != nil {
		*ofce = FilterObjectConditionEnum(s)
	}
}

// IsValidEnum returns true if the string represents a valid FilterObjectConditionEnum in the Notion API.
func (ofce *FilterObjectConditionEnum) IsValidEnum() bool {
	return ofce != nil && isValidEnum(string(*ofce), FilterObjectConditionEnumPage,
		FilterObjectConditionEnumDatabase,
	)
}

// UnmarshalJSON returns an error if the string is not a valid FilterObjectConditionEnum in the Notion API.
func (ofce *FilterObjectConditionEnum) UnmarshalJSON(b []byte) error {
	return unmarshalEnum(b, ofce)
}

// The EmptyFilter is one where a property can be filter on whether or not it is empty.
type EmptyFilter struct {
	IsEmpty    bool `json:"is_empty,omitempty"`
	IsNotEmpty bool `json:"is_not_empty,omitempty"`
}

func (e *EmptyFilter) getValue() interface{} {
	return e != nil && e.IsEmpty && e.IsNotEmpty
}

// A Filter represents a filter object with which to query a database in the Notion API.
type Filter struct {
	// All that is required for the "property" field is the name or id of the property,
	// but allowing to pass the whole property is more flexible. Just requires custom (un)marshaling.
	Property    DatabaseProperty
	Type        FilterConditionTypeEnum `json:"type"`
	Text        *TextFilter             `json:"text,omitempty"`
	Number      *NumberFilter           `json:"number,omitempty"`
	Checkbox    *CheckboxFilter         `json:"checkbox,omitempty"`
	Select      *SelectFilter           `json:"select,omitempty"`
	MultiSelect *MultiSelectFilter      `json:"multi_select,omitempty"`
	Date        *DateFilter             `json:"date,omitempty"`
	People      *PeopleFilter           `json:"people,omitempty"`
	Files       *FilesFilter            `json:"files,omitempty"`
	Relation    *RelationFilter         `json:"relation,omitempty"`
	Formula     *FormulaFilter          `json:"formula,omitempty"`
}

type filter Filter

// GetType returns the Type of a text filter.
func (f *filter) getType() string {
	if f == nil {
		return ""
	}
	return string(f.Type)
}

// FieldsToExpand implements the expander interface for filter.
func (f *filter) fieldsToExpand() []string {
	return []string{}
}

// MarshalJSON prepares the filter to be compatible with the Notion API.
func (f *Filter) MarshalJSON() ([]byte, error) {
	if f == nil {
		return nil, nil
	}
	ff := filter(*f)
	return marshalJSONExpandByType(&ff)
}

// A TextFilter represents a filter object with which to query a database in the Notion API.
type TextFilter struct {
	EmptyFilter
	Type  FilterTextConditionEnum
	Value *string
}

// GetType returns the Type of a text filter.
func (tf *TextFilter) getType() string {
	if tf == nil {
		return ""
	}
	return string(tf.Type)
}

func (tf *TextFilter) getValue() interface{} {
	if tf == nil {
		return nil
	}

	if e, ok := tf.EmptyFilter.getValue().(bool); ok && e {
		return true
	}

	return tf.Value
}

// MarshalJSON prepares the filter to be compatible with the Notion API.
func (tf *TextFilter) MarshalJSON() ([]byte, error) {
	return marshalJSONFilter(tf)
}

// A NumberFilter represents a filter object with which to query a database in the Notion API.
type NumberFilter struct {
	EmptyFilter
	Type  FilterNumberConditionEnum
	Value *float64
}

// GetType returns the Type of a number filter.
func (nf *NumberFilter) getType() string {
	if nf == nil {
		return ""
	}
	return string(nf.Type)
}

func (nf *NumberFilter) getValue() interface{} {
	if nf == nil {
		return nil
	}

	if e, ok := nf.EmptyFilter.getValue().(bool); ok && e {
		return true
	}

	return nf.Value
}

// MarshalJSON prepares the filter to be compatible with the Notion API.
func (nf *NumberFilter) MarshalJSON() ([]byte, error) {
	return marshalJSONFilter(nf)
}

// A CheckboxFilter represents a filter object with which to query a database in the Notion API.
type CheckboxFilter struct {
	EmptyFilter
	Type  FilterCheckboxConditionEnum
	Value *bool
}

// GetType returns the Type of a number filter.
func (cf *CheckboxFilter) getType() string {
	if cf == nil {
		return ""
	}
	return string(cf.Type)
}

func (cf *CheckboxFilter) getValue() interface{} {
	if cf == nil {
		return nil
	}

	if e, ok := cf.EmptyFilter.getValue().(bool); ok && e {
		return true
	}

	return cf.Value
}

// MarshalJSON prepares the filter to be compatible with the Notion API.
func (cf *CheckboxFilter) MarshalJSON() ([]byte, error) {
	return marshalJSONFilter(cf)
}

// A SelectFilter represents a filter object with which to query a database in the Notion API.
type SelectFilter struct {
	EmptyFilter
	Type  FilterSelectConditionEnum
	Value *string
}

// GetType returns the Type of a number filter.
func (sf *SelectFilter) getType() string {
	if sf == nil {
		return ""
	}
	return string(sf.Type)
}

func (sf *SelectFilter) getValue() interface{} {
	if sf == nil {
		return nil
	}

	if e, ok := sf.EmptyFilter.getValue().(bool); ok && e {
		return true
	}

	return sf.Value
}

// MarshalJSON prepares the filter to be compatible with the Notion API.
func (sf *SelectFilter) MarshalJSON() ([]byte, error) {
	return marshalJSONFilter(sf)
}

// A MultiSelectFilter represents a filter object with which to query a database in the Notion API.
type MultiSelectFilter struct {
	EmptyFilter
	Type  FilterMultiSelectConditionEnum
	Value *string
}

// GetType returns the Type of a number filter.
func (mf *MultiSelectFilter) getType() string {
	if mf == nil {
		return ""
	}
	return string(mf.Type)
}

func (mf *MultiSelectFilter) getValue() interface{} {
	if mf == nil {
		return nil
	}

	if e, ok := mf.EmptyFilter.getValue().(bool); ok && e {
		return true
	}

	return mf.Value
}

// MarshalJSON prepares the filter to be compatible with the Notion API.
func (mf *MultiSelectFilter) MarshalJSON() ([]byte, error) {
	return marshalJSONFilter(mf)
}

// A DateFilter represents a filter object with which to query a database in the Notion API.
type DateFilter struct {
	EmptyFilter
	Type  FilterDateConditionEnum
	Value time.Time
}

// GetType returns the Type of a number filter.
func (df *DateFilter) getType() string {
	if df == nil {
		return ""
	}
	return string(df.Type)
}

func (df *DateFilter) getValue() interface{} {
	if df == nil {
		return nil
	}

	if e, ok := df.EmptyFilter.getValue().(bool); ok && e {
		return true
	}

	return df.Value
}

// MarshalJSON prepares the filter to be compatible with the Notion API.
func (df *DateFilter) MarshalJSON() ([]byte, error) {
	return marshalJSONFilter(df)
}

// A PeopleFilter represents a filter object with which to query a database in the Notion API.
type PeopleFilter struct {
	EmptyFilter
	Type  FilterPeopleConditionEnum
	Value *UUID4
}

// GetType returns the Type of a number filter.
func (pf *PeopleFilter) getType() string {
	if pf == nil {
		return ""
	}
	return string(pf.Type)
}

func (pf *PeopleFilter) getValue() interface{} {
	if pf == nil {
		return nil
	}

	if e, ok := pf.EmptyFilter.getValue().(bool); ok && e {
		return true
	}

	return pf.Value
}

// MarshalJSON prepares the filter to be compatible with the Notion API.
func (pf *PeopleFilter) MarshalJSON() ([]byte, error) {
	return marshalJSONFilter(pf)
}

// A FilesFilter represents a filter object with which to query a database in the Notion API.
type FilesFilter struct {
	EmptyFilter
	Type FilterFilesConditionEnum
}

// GetType returns the Type of a number filter.
func (ff *FilesFilter) getType() string {
	if ff == nil {
		return ""
	}
	return string(ff.Type)
}

// MarshalJSON prepares the filter to be compatible with the Notion API.
func (ff *FilesFilter) MarshalJSON() ([]byte, error) {
	return marshalJSONFilter(ff)
}

// A RelationFilter represents a filter object with which to query a database in the Notion API.
type RelationFilter struct {
	EmptyFilter
	Type  FilterRelationConditionEnum
	Value *UUID4
}

// GetType returns the Type of a number filter.
func (rf *RelationFilter) getType() string {
	if rf == nil {
		return ""
	}
	return string(rf.Type)
}

func (rf *RelationFilter) getValue() interface{} {
	if rf == nil {
		return nil
	}

	if e, ok := rf.EmptyFilter.getValue().(bool); ok && e {
		return true
	}
	return rf.Value
}

// MarshalJSON prepares the filter to be compatible with the Notion API.
func (rf *RelationFilter) MarshalJSON() ([]byte, error) {
	return marshalJSONFilter(rf)
}

// A ObjectFilter represents a filter object with which to query a database in the Notion API.
type ObjectFilter struct {
	Type  FilterObjectConditionEnum
	Value *string
}

// GetType returns the Type of an object filter.
func (of *ObjectFilter) getType() string {
	if of == nil {
		return ""
	}
	return string(of.Type)
}

func (of *ObjectFilter) getValue() interface{} {
	if of == nil {
		return nil
	}

	return of.Value
}

// MarshalJSON prepares the filter to be compatible with the Notion API.
func (of *ObjectFilter) MarshalJSON() ([]byte, error) {
	return marshalJSONFilter(of)
}

// A FormulaFilter represents a filter object with which to query a database in the Notion API.
type FormulaFilter struct {
	Type     FilterFormulaConditionEnum `json:"-"`
	Text     *TextFilter                `json:"text,omitempty"`
	Checkbox *CheckboxFilter            `json:"checkbox,omitempty"`
	Number   *NumberFilter              `json:"number,omitempty"`
	Date     *DateFilter                `json:"date,omitempty"`
}

// A CompoundFilter represents a compound filter object with which to query a database in the Notion API.
// TODO: nested compound filters are not yet implemented.
type CompoundFilter struct {
	Or  []*Filter `json:"or,omitempty"`
	And []*Filter `json:"and,omitempty"`
}

type valued interface {
	typed
	getValue() interface{}
}

func marshalJSONFilter(v valued) ([]byte, error) {
	return json.Marshal(map[string]interface{}{v.getType(): v.getValue()})
}
