package gotion

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/sethgrid/pester"
	"github.com/thedadams/gotion/notion"
	"golang.org/x/time/rate"
)

const apiBaseURL = "https://api.notion.com"

type list interface {
	Len() int
}

type paginated interface {
	setPage(*string, int) ([]byte, error)
	getCursor() *string
	getMaxResults() int
}

// A Result is a response from the Notion API when getting more than one thing back (i.e. listing)
type Result struct {
	NextCursor *string     `json:"next_cursor"`
	HasMore    bool        `json:"has_more"`
	Results    interface{} `json:"results"`
}

// Client is a client used to make calls to the Notion API.
type Client struct {
	settings    *notion.Settings
	httpClient  *pester.Client
	rateLimiter *rate.Limiter
}

// NewClient creates a new gotion client to use with the API.
// By default:
// - the client will use the most recent accepted Notion API version.
// - Timeout is 30 seconds
// - Backoff strategy is set to pester.ExponentialJitterBackoff
// - Rate limiter is set to 3 requests per second
// - MaxRetries is set to 8
// - the client will retry on 429 errors.
func NewClient(apiKey string, options ...Option) *Client {
	c := &Client{settings: &notion.Settings{APIKey: apiKey}}
	WithPesterClient(pester.New())(c)
	WithBackoffStrategy(pester.ExponentialJitterBackoff)(c)
	WithTimeout(defaultTimeout)(c)
	WithMaxRetries(defaultMaxRetries)(c)
	WithRetryOnHTTP429()(c)
	WithRateLimiter(defaultRateLimiter)(c)
	for _, o := range options {
		o(c)
	}
	return c
}

func addQueryParams(url string, cursor *string, pageSize int) string {
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 100
	}

	if cursor == nil {
		return fmt.Sprintf("%s?page_size=%d", url, pageSize)
	}
	return fmt.Sprintf("%s?start_cursor=%s&page_size=%d", url, *cursor, pageSize)
}

func (c *Client) makeRequest(ctx context.Context, method, url string, body io.Reader, respObject interface{}) error {
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return err
	}

	c.settings.ToHeaders(req)

	if err := c.rateLimiter.Wait(ctx); err != nil {
		return err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return parseError(resp.Status, method, url, respBody)
	}

	if respObject != nil {
		return json.Unmarshal(respBody, &respObject)
	}
	return nil
}

func (c *Client) createObject(ctx context.Context, url string, body map[string]interface{}, respObject interface{}) error {
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return err
	}

	return c.makeRequest(ctx, http.MethodPost, url, bytes.NewBuffer(bodyBytes), respObject)
}

func (c *Client) updateObject(ctx context.Context, url string, body, respObject interface{}) error {
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return err
	}

	return c.makeRequest(ctx, http.MethodPatch, url, bytes.NewBuffer(bodyBytes), respObject)
}

func (c *Client) queryForList(ctx context.Context, url string, body paginated, results list) error {
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return err
	}

	hasMore := true
	maxResults := body.getMaxResults()
	r := Result{Results: results}

	for hasMore && (maxResults < 0 || results.Len() < maxResults) {
		err = c.makeRequest(ctx, http.MethodPost, url, bytes.NewReader(bodyBytes), &r)
		if err != nil {
			return err
		}

		hasMore = r.HasMore
		maxResults -= results.Len()
		if bodyBytes, err = body.setPage(r.NextCursor, maxResults); err != nil {
			return err
		}
	}

	return nil
}

func (c *Client) getList(ctx context.Context, url string, cursor *string, maxResults int, results list) error {
	hasMore := true
	r := Result{Results: results}

	for hasMore && (maxResults < 0 || results.Len() < maxResults) {
		err := c.makeRequest(ctx, http.MethodGet, addQueryParams(url, cursor, maxResults), nil, &r)
		if err != nil {
			return err
		}

		hasMore = r.HasMore
		maxResults -= results.Len()
		cursor = r.NextCursor
	}

	return nil
}

func parseError(status, method, url string, body []byte) error {
	apiError := notion.APIError{}
	if err := json.Unmarshal(body, &apiError); err != nil {
		return fmt.Errorf("error parsing error json: %w", err)
	}
	return fmt.Errorf("%s request to %s with status %s: %w", method, url, status, apiError)
}
