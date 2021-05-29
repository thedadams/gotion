package notion

import (
	"encoding/json"
	"net/url"
)

// These constants represent the enums in the Notion API for rich_text objects
const (
	RichTextTypeEnumText     = "text"
	RichTextTypeEnumMention  = "mention"
	RichTextTypeEnumEquation = "equation"

	AnnotationColorEnumDefault          = "default"
	AnnotationColorEnumGray             = "gray"
	AnnotationColorEnumBrown            = "brown"
	AnnotationColorEnumOrange           = "orange"
	AnnotationColorEnumYellow           = "yellow"
	AnnotationColorEnumGreen            = "green"
	AnnotationColorEnumBlue             = "blue"
	AnnotationColorEnumPurple           = "purple"
	AnnotationColorEnumPink             = "pink"
	AnnotationColorEnumRed              = "red"
	AnnotationColorEnumGrayBackground   = "gray_background"
	AnnotationColorEnumBrownBackground  = "brown_background"
	AnnotationColorEnumOrangeBackground = "orange_background"
	AnnotationColorEnumYellowBackground = "yellow_background"
	AnnotationColorEnumGreenBackground  = "green_background"
	AnnotationColorEnumBlueBackground   = "blue_background"
	AnnotationColorEnumPurpleBackground = "purple_background"
	AnnotationColorEnumPinkBackground   = "pink_background"
	AnnotationColorEnumRedBackground    = "red_background"

	MentionTypeEnumUser     = "user"
	MentionTypeEnumPage     = "page"
	MentionTypeEnumDatabase = "database"
	MentionTypeEnumData     = "date"
)

// RichTextTypeEnum represents the different types of a rich_text object in the Notion API.
type RichTextTypeEnum string

// SetValue sets the RichTextTypeEnum to the given string
func (rtte *RichTextTypeEnum) SetValue(s string) {
	if rtte != nil {
		*rtte = RichTextTypeEnum(s)
	}
}

// IsValidEnum returns true if the string represents a valid rich_text type in the Notion API.
func (rtte *RichTextTypeEnum) IsValidEnum() bool {
	return rtte != nil && isValidEnum(string(*rtte), RichTextTypeEnumText, RichTextTypeEnumMention, RichTextTypeEnumEquation)
}

// UnmarshalJSON returns an error if the provided string is not a valid RichTextTypeEnum
func (rtte *RichTextTypeEnum) UnmarshalJSON(b []byte) error {
	return unmarshalEnum(b, rtte)
}

// RichText represents the common fields in a rich text object in Notion.
type RichText struct {
	Type        RichTextTypeEnum `json:"type"`
	HRef        string           `json:"href,omitempty"`
	Annotations Annotations      `json:"annotations"`
	PlainText   string           `json:"plain_text"`
	Text        *Text            `json:"text,omitempty"`
	Mention     *Mention         `json:"mention,omitempty"`
	Equation    *Equation        `json:"equation,omitempty"`
}

// AnnotationColorEnum represents the colors of a rich_text object in the Notion API.
type AnnotationColorEnum string

// SetValue sets the AnnotationColorEnum to the given string
func (ace *AnnotationColorEnum) SetValue(s string) {
	if ace != nil {
		*ace = AnnotationColorEnum(s)
	}
}

// IsValidEnum returns true if the provided string represents a valid color in the Notion API.
func (ace *AnnotationColorEnum) IsValidEnum() bool {
	return ace != nil && isValidEnum(string(*ace), AnnotationColorEnumDefault,
		AnnotationColorEnumGray,
		AnnotationColorEnumBrown,
		AnnotationColorEnumOrange,
		AnnotationColorEnumYellow,
		AnnotationColorEnumGreen,
		AnnotationColorEnumBlue,
		AnnotationColorEnumPurple,
		AnnotationColorEnumPink,
		AnnotationColorEnumRed,
		AnnotationColorEnumGrayBackground,
		AnnotationColorEnumBrownBackground,
		AnnotationColorEnumOrangeBackground,
		AnnotationColorEnumYellowBackground,
		AnnotationColorEnumGreenBackground,
		AnnotationColorEnumBlueBackground,
		AnnotationColorEnumPurpleBackground,
		AnnotationColorEnumPinkBackground,
		AnnotationColorEnumRedBackground,
	)
}

// UnmarshalJSON returns an error if the provided string is not a valid AnnotationColorEnum
func (ace *AnnotationColorEnum) UnmarshalJSON(b []byte) error {
	return unmarshalEnum(b, ace)
}

// Annotations are the annotations on a rich_text object in the Notion API.
type Annotations struct {
	Bold          bool                `json:"bold"`
	Italic        bool                `json:"italic"`
	Strikethrough bool                `json:"strikethrough"`
	Underline     bool                `json:"underline"`
	Code          bool                `json:"code"`
	Color         AnnotationColorEnum `json:"color"`
}

// Text type represents a text rich_text object in the Notion API.
type Text struct {
	Content string `json:"content"`
	Link    *Link  `json:"link,omitempty"`
}

// GetURL returns a pointer to the URL of the Link of the Text object, if it exists,
// and returns nil otherwise.
func (t *Text) GetURL() *url.URL {
	if t == nil || t.Link == nil {
		return nil
	}

	u := url.URL(t.Link.URL)
	return &u
}

// Link type represents a link in a rich_text object in the Notion API.
type Link struct {
	URL jsonURL `json:"url"`
}

// MarshalJSON marshals a Link object, always setting the `type` parameter to "url"
func (l *Link) MarshalJSON() ([]byte, error) {
	if l == nil {
		return nil, nil
	}
	return json.Marshal(map[string]interface{}{
		"type": "url",
		"url":  l.URL,
	})
}

// MentionTypeEnum represents the valid mention types in the Notion API.
type MentionTypeEnum string

// SetValue sets the MentionTypeEnum to the given string
func (mte *MentionTypeEnum) SetValue(s string) {
	if mte != nil {
		*mte = MentionTypeEnum(s)
	}
}

// IsValidEnum returns true if the provided string represents a valid mention type in the Notion API.
func (mte *MentionTypeEnum) IsValidEnum() bool {
	return mte != nil && isValidEnum(string(*mte), MentionTypeEnumUser, MentionTypeEnumPage, MentionTypeEnumDatabase, MentionTypeEnumData)
}

// UnmarshalJSON returns an error if the provided string is not a valid MentionTypeEnum
func (mte *MentionTypeEnum) UnmarshalJSON(b []byte) error {
	return unmarshalEnum(b, mte)
}

// A Mention represents a mention object in a rich_text object in the Notion API.
type Mention struct {
	Type MentionTypeEnum `json:"type"`
	User *User           `json:"user,omitempty"`
	Ref  *UUID4          `json:"id,omitempty"`
	Date *Date           `json:"date,omitempty"`
}

type mention Mention

// GetType implements the typed interface for Mention
func (m *mention) getType() string {
	if m == nil {
		return ""
	}
	return string(m.Type)
}

// FieldsToExpand implements the expander interface for Mention
func (m *mention) fieldsToExpand() []string {
	return []string{"user", "id", "date"}
}

// UnmarshalJSON parses the Mention object, flattening the page or database reference into the Ref parameter
func (m *Mention) UnmarshalJSON(b []byte) error {
	mm := new(mention)
	if err := unmarshalJSONFlattenByType(b, mm); err != nil {
		return err
	}

	*m = Mention(*mm)
	return nil
}

// MarshalJSON marshals the Mention object to be compatible with the Notion API.
func (m *Mention) MarshalJSON() ([]byte, error) {
	mm := mention(*m)
	return marshalJSONExpandByType(&mm)
}

// An Equation represents an equation in a rich_text object in the Notion API.
type Equation struct {
	Expression string `json:"expression"`
}
