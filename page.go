package gotion

import (
	"context"
	"fmt"
	"net/http"

	"github.com/thedadams/gotion/notion"
)

// GetPage gets a page with the given id from the Notion API.
// To get the contents of a page, use `GetPageWithChildren` with the page id.
func (c *Client) GetPage(ctx context.Context, id string) (*notion.Page, error) {
	page := &notion.Page{}
	err := c.makeRequest(ctx, http.MethodGet, fmt.Sprintf("%s/v1/pages/%s", apiBaseURL, id), nil, page)
	if err != nil {
		page = nil
	}

	return page, err
}

// GetPageAndChildren gets a page with the given id from the Notion API,
// as well as the children of the page. If `maxResults < 0`, then this gets all children.
func (c *Client) GetPageAndChildren(ctx context.Context, id string, maxResults int) (*notion.Page, error) {
	page, err := c.GetPage(ctx, id)
	if err != nil {
		return nil, err
	}

	page.Children, err = c.GetBlockChildren(ctx, id, nil, maxResults)
	return page, err
}

// CreatePage will send a request to create the given page in the Notion API.
// All that is needed in the notion.Page object are the Parent, Properties, and Children.
// No IDs need to be given.
// On success, the notion.Page returned will be the complete page from the Notion API.
// On error, the notion.Page returned is the original one.
func (c *Client) CreatePage(ctx context.Context, page *notion.Page) (*notion.Page, error) {
	body := map[string]interface{}{
		"parent":     &page.Parent,
		"properties": &page.Properties,
	}

	return page, c.createObject(ctx, fmt.Sprintf("%s/v1/pages", apiBaseURL), body, page)
}

// UpdatePageProperties updates the page properties in the Notion API.
// All that is needed in the notion.Page are the page ID and Properties.
// A caller doesn't need to provide all properties, but only the updates ones. However, providing all properties
// works fine as well.
// On success, the notion.Page will be the complete page from the Notion API.
// On error, the notion.Page will be changed.
func (c *Client) UpdatePageProperties(ctx context.Context, page *notion.Page) error {
	body := map[string]interface{}{"properties": page.Properties}

	return c.updateObject(ctx, fmt.Sprintf("%s/v1/pages/%s", apiBaseURL, page.ID), body, page)
}
