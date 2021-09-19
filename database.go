package gotion

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/thedadams/gotion/notion"
)

// DBQuery represents the parameters needed to query a database in the Notion API.
type DBQuery struct {
	Filter     *notion.Filter `json:"filter,omitempty"`
	Sort       *notion.Sort   `json:"sort,omitempty"`
	Cursor     *string        `json:"cursor,omitempty"`
	MaxResults *int           `json:"page_size,omitempty"`
}

func (db *DBQuery) setPage(cursor *string, maxResults int) ([]byte, error) {
	oldCursor, oldPageSize := db.Cursor, db.MaxResults
	defer func() {
		db.Cursor = oldCursor
		db.MaxResults = oldPageSize
	}()

	db.Cursor, db.MaxResults = cursor, &maxResults
	return json.Marshal(db)
}

func (db *DBQuery) getCursor() *string {
	return db.Cursor
}

func (db *DBQuery) getMaxResults() int {
	return *db.MaxResults
}

// QueryDatabase will  query the database with the given id in the Notion API.
func (c *Client) QueryDatabase(ctx context.Context, id string, query *DBQuery) ([]*notion.Page, error) {
	var results notion.Pages
	if err := c.queryForList(ctx, fmt.Sprintf("%s/v1/databases/%s/query", apiBaseURL, id), query, &results); err != nil {
		return nil, err
	}
	return results, nil
}

// CreateDatabase will send a request to create the given database in the Notion API.
// All that is needed in the notion.Database object are the Parent and Properties.
// Optionally, a Title can be set.
// No IDs need to be given.
// On success, the notion.Database will be the complete page from the Notion API.
// On error, the notion.Page will not be changed.
func (c *Client) CreateDatabase(ctx context.Context, db *notion.Database) error {
	body := map[string]interface{}{
		"parent":     &db.Parent,
		"properties": &db.Properties,
	}
	if len(db.Title) != 0 {
		body["title"] = &db.Title
	}

	return c.createObject(ctx, fmt.Sprintf("%s/v1/databases", apiBaseURL), body, db)
}

// UpdateDatabase updates the database in the Notion API.
// On success, the database is the complete database from the Notion API.
// On error, the database is not updated.
func (c *Client) UpdateDatabase(ctx context.Context, db *notion.Database) error {
	body := map[string]interface{}{
		"title":      db.Title,
		"properties": db.Properties,
	}

	return c.updateObject(ctx, fmt.Sprintf("%s/v1/databases/%s", apiBaseURL, db.ID), body, db)
}

// GetDatabase gets a database with the given id from the Notion API.
func (c *Client) GetDatabase(ctx context.Context, id string) (*notion.Database, error) {
	db := &notion.Database{}
	err := c.makeRequest(ctx, http.MethodGet, fmt.Sprintf("%s/v1/databases/%s", apiBaseURL, id), nil, db)
	if err != nil {
		db = nil
	}
	return db, err
}

// GetDatabaseAndChildren gets a database with the given id from the Notion API,
// as well as the children of the database. If `maxResults < 0`, then this gets all children.
func (c *Client) GetDatabaseAndChildren(ctx context.Context, id string, maxResults int) (*notion.Database, error) {
	db, err := c.GetDatabase(ctx, id)
	if err != nil {
		return nil, err
	}

	db.Children, err = c.GetBlockChildren(ctx, id, nil, maxResults)
	return db, err
}

// GetDatabases gets a number of databases with from the Notion API.
// If `maxResults < 0`, then all databases are retrieved.
func (c *Client) GetDatabases(ctx context.Context, cursor *string, maxResults int) (*notion.Databases, error) {
	var results *notion.Databases
	if err := c.getList(ctx, fmt.Sprintf("%s/v1/databases", apiBaseURL), cursor, maxResults, results); err != nil {
		return nil, err
	}

	return results, nil
}
