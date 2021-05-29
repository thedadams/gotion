package notion

import "encoding/json"

// These constants represent the valid type enums for a block type
const (
	BlockTypeEnumParagraph        = "paragraph"
	BlockTypeEnumHeading1         = "heading_1"
	BlockTypeEnumHeading2         = "heading_2"
	BlockTypeEnumHeading3         = "heading_3"
	BlockTypeEnumBulletedListItem = "bulleted_list_item"
	BlockTypeEnumNumberedListItem = "numbered_list_item"
	BlockTypeEnumToDo             = "to_do"
	BlockTypeEnumToggle           = "toggle"
	BlockTypeEnumChildPage        = "child_page"
	BlockTypeEnumUnsupported      = "unsupported"
)

// A BlockTypeEnum represents a type of a block in the Notion API.
type BlockTypeEnum string

// SetValue sets the BlockTypeEnum to the given string
func (bte *BlockTypeEnum) SetValue(s string) {
	if bte != nil {
		*bte = BlockTypeEnum(s)
	}
}

// IsValidEnum returns true if the string represents a valid type of a block object in the Notion API.
func (bte *BlockTypeEnum) IsValidEnum() bool {
	return bte != nil && isValidEnum(string(*bte), BlockTypeEnumParagraph,
		BlockTypeEnumHeading1,
		BlockTypeEnumHeading2,
		BlockTypeEnumHeading3,
		BlockTypeEnumBulletedListItem,
		BlockTypeEnumNumberedListItem,
		BlockTypeEnumToDo,
		BlockTypeEnumToggle,
		BlockTypeEnumChildPage,
		BlockTypeEnumUnsupported,
	)
}

// UnmarshalJSON returns an error if the provided string is not a valid BlockTypeEnum in the Notion API.
func (bte *BlockTypeEnum) UnmarshalJSON(b []byte) error {
	return unmarshalEnum(b, bte)
}

// Block represents the common fields in a "block" in Notion (i.e. text, headings, list, etc)
type Block struct {
	Object
	Editable
	Type        BlockTypeEnum `json:"type"`
	HasChildren bool          `json:"has_children"`
	Text        []*RichText   `json:"text,omitempty"`
	Children    []*Block      `json:"children,omitempty"`
	// Only valid for To Do blocks
	Checked *bool `json:"checked,omitempty"`
	// Only valid for Child Page blocks
	Title *string `json:"title,omitempty"`
}

type block Block

// GetType implements the typed interface
func (b *block) getType() string {
	if b == nil {
		return ""
	}
	return string(b.Type)
}

// FieldsToExpand implements the expander interface
func (b *block) fieldsToExpand() []string {
	return []string{"text", "children", "checked", "title"}
}

// UnmarshalJSON sets the JSON object based on the Type of the Block object.
// Text and Children are set based on the Type of the Block.
func (b *Block) UnmarshalJSON(bt []byte) error {
	bb := new(block)
	if err := unmarshalJSONFlattenByType(bt, bb); err != nil {
		return err
	}

	*b = Block(*bb)
	return nil
}

// MarshalJSON returns a marshaled version of the Block that is compatible with the Notion API.
func (b *Block) MarshalJSON() ([]byte, error) {
	if b == nil {
		return nil, nil
	}
	bb := block(*b)
	return marshalJSONExpandByType(&bb)
}

// IsChecked returns true if the Block is a To Do block and it is checked,
// and returns false otherwise
func (b *Block) IsChecked() bool {
	return b != nil && b.Checked != nil && *b.Checked
}

// GetTitle returns the title of the Block if the Block is a child_page Block,
// and returns the empty string otherwise
func (b *Block) GetTitle() string {
	if b == nil || b.Title == nil {
		return ""
	}
	return *b.Title
}

// Blocks represents a list of users from the Notion API.
type Blocks []*Block

// Len returns the number of blocks in the the slice.
func (bs *Blocks) Len() int {
	if bs == nil {
		return 0
	}
	return len(*bs)
}

type blocks Blocks

// UnmarshalJSON appends the unmarshaled Blocks to the slice.
func (bs *Blocks) UnmarshalJSON(bt []byte) error {
	b := new(blocks)
	if err := json.Unmarshal(bt, b); err != nil || len(*b) == 0 {
		return err
	}

	if bs == nil {
		*bs = make([]*Block, 0, len(*b))
	}
	*bs = append(*bs, []*Block(*b)...)
	return nil
}
