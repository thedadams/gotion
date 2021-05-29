package gotion

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/thedadams/gotion/notion"
)

// SearchQuery represents the body needed to search the Notion API.
// Currently, one can only filter based on object: page or database.
// Currently, one can only sort based on last_edited_time.
type SearchQuery struct {
	Query      string         `json:"query"`
	Filter     *notion.Filter `json:"filter,omitempty"`
	Sort       *notion.Sort   `json:"sort,omitempty"`
	Cursor     *string        `json:"cursor,omitempty"`
	MaxResults *int           `json:"page_size,omitempty"`
}

// SearchResults represent the returned results from searching the Notion API.
type SearchResults struct {
	Pages     []*notion.Page
	Databases []*notion.Database
}

// Len returns the total number of results, pages plus databases
func (sr *SearchResults) Len() int {
	return len(sr.Pages) + len(sr.Databases)
}

// UnmarshalJSON unmarshals the SearchResults into Pages and Databases.
func (sr *SearchResults) UnmarshalJSON(b []byte) error {
	m := make([]map[string]interface{}, 0)
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}

	var err error
	for _, r := range m {
		if objectType, ok := r["object"]; ok {
			var bt []byte
			if bt, err = json.Marshal(r); err != nil {
				return err
			}
			switch objectType {
			case notion.FilterObjectConditionEnumPage:
				page := &notion.Page{}
				if err := json.Unmarshal(bt, page); err != nil {
					return err
				}
				sr.Pages = append(sr.Pages, page)
			case notion.FilterObjectConditionEnumDatabase:
				db := &notion.Database{}
				if err := json.Unmarshal(bt, db); err != nil {
					return err
				}
				sr.Databases = append(sr.Databases, db)
			}
		}
	}

	return nil
}

func (sq *SearchQuery) setPage(cursor *string, maxResults int) ([]byte, error) {
	oldCursor, oldPageSize := sq.Cursor, sq.MaxResults
	defer func() {
		sq.Cursor = oldCursor
		sq.MaxResults = oldPageSize
	}()

	sq.Cursor, sq.MaxResults = cursor, &maxResults
	return json.Marshal(sq)
}

func (sq *SearchQuery) getCursor() *string {
	return sq.Cursor
}

func (sq *SearchQuery) getMaxResults() int {
	return *sq.MaxResults
}

// Search allows searching the Notion API. Currently, this is limited by Notion.
// query -- search the titles of objects (pages and databases)
// sort -- can only sort by last_edited_time
// filter -- can only filter by object type: pages or databases.
func (c *Client) Search(ctx context.Context, query *SearchQuery) (*SearchResults, error) {
	results := SearchResults{}
	if err := c.queryForList(ctx, fmt.Sprintf("%s/v1/search", apiBaseURL), query, &results); err != nil {
		return nil, err
	}
	return &results, nil
}
