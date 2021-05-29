package gotion

import (
	"context"
	"fmt"
	"net/http"

	"github.com/thedadams/gotion/notion"
)

// GetUser gets a user with the given id from the Notion API.
func (c *Client) GetUser(ctx context.Context, id string) (*notion.User, error) {
	u := &notion.User{}
	err := c.makeRequest(ctx, http.MethodGet, fmt.Sprintf("%s/v1/users/%s", apiBaseURL, id), nil, u)
	if err != nil {
		u = nil
	}

	return u, err
}

// GetUsers gets a user with the given id from the Notion API.
// If `maxResults < 0`, then this will get all users.
func (c *Client) GetUsers(ctx context.Context, cursor *string, maxResults int) ([]*notion.User, error) {
	var results notion.Users
	if err := c.getList(ctx, fmt.Sprintf("%s/v1/users", apiBaseURL), cursor, maxResults, &results); err != nil {
		return nil, err
	}
	return results, nil
}
