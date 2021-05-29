package gotion

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/thedadams/gotion/notion"
)

// GetBlockChildren gets a the children of the block with the given id from the Notion API.
// If `maxResults < 0`, then this will get all the children of the block.
func (c *Client) GetBlockChildren(ctx context.Context, id string, cursor *string, maxResults int) ([]*notion.Block, error) {
	var results notion.Blocks
	if err := c.getList(ctx, fmt.Sprintf("%s/v1/blocks/%s/children", apiBaseURL, id), cursor, maxResults, &results); err != nil {
		return nil, err
	}

	return results, nil
}

// AppendBlockChildren adds the given blocks as children of the block with the ID provided in the Notion API.
// All that is needed in the notion.Block are the page ID and the blocks that need to be added.
// Providing the current children of the block will result in duplicate blocks.
// On success, the notion.Block returned will be the complete block from the Notion API.
// On error, the notion.Block returned is the original one.
func (c *Client) AppendBlockChildren(ctx context.Context, block *notion.Block) (*notion.Block, error) {
	bodyBytes, err := json.Marshal(map[string]interface{}{"children": block.Children})
	if err != nil {
		return block, err
	}

	err = c.makeRequest(ctx, http.MethodPatch, fmt.Sprintf("%s/v1/blocks/%s/children", apiBaseURL, block.ID), bytes.NewReader(bodyBytes), nil)
	if err != nil {
		return block, err
	}

	block.Children, err = c.GetBlockChildren(ctx, block.ID.String(), nil, -1)
	return block, err
}
