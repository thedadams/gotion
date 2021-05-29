package gotion

import (
	"time"

	"github.com/sethgrid/pester"
	"github.com/thedadams/gotion/notion"
	"golang.org/x/time/rate"
)

const (
	defaultTimeout    = 30 * time.Second
	defaultMaxRetries = 8
)

var defaultRateLimiter = rate.NewLimiter(rate.Every(time.Second), 3)

// An Option is a way of customizing the gotion client
type Option func(*Client)

// WithPesterClient uses the given pester client for the gotion client
func WithPesterClient(pesterClient *pester.Client) Option {
	return func(c *Client) {
		if c != nil {
			c.httpClient = pesterClient
		}
	}
}

// WithSettings uses the given settings with the gotion client.
func WithSettings(s *notion.Settings) Option {
	return func(c *Client) {
		if c != nil {
			c.settings = s
		}
	}
}

// WithToken uses the given auth token with the gotion client.
func WithToken(t string) Option {
	return func(c *Client) {
		if c != nil {
			c.settings.APIKey = t
		}
	}
}

// WithUserAgent uses the given user agent string with the gotion client.
func WithUserAgent(u string) Option {
	return func(c *Client) {
		if c != nil {
			c.settings.UserAgent = u
		}
	}
}

// WithAPIVersion uses the given Notion API version with the gotion client.
// It checks that the string provided is an accepted API version.
// If the string does not represent an accepted API version, then this is a no-op.
func WithAPIVersion(v string) Option {
	return func(c *Client) {
		for _, version := range notion.Versions {
			if v == version {
				if c != nil {
					c.settings.Version = v
				}
				break
			}
		}
	}
}

// WithRateLimiter users the given rate limiter with the gotion client.
func WithRateLimiter(r *rate.Limiter) Option {
	return func(c *Client) {
		if c != nil {
			c.rateLimiter = r
		}
	}
}

// WithBackoffStrategy applies the given backoff strategy to the gotion client
func WithBackoffStrategy(b pester.BackoffStrategy) Option {
	return func(c *Client) {
		if c != nil {
			c.httpClient.Backoff = b
		}
	}
}

// WithRetryOnHTTP429 sets the gotion client to retry on 429 errors.
// The rate limiter should help with this as well.
func WithRetryOnHTTP429() Option {
	return func(c *Client) {
		if c != nil {
			c.httpClient.RetryOnHTTP429 = true
		}
	}
}

// WithoutRetryOnHTTP429 sets the gotion client to NOT retry on 429 errors.
// The rate limiter should help with this as well.
func WithoutRetryOnHTTP429() Option {
	return func(c *Client) {
		if c != nil {
			c.httpClient.RetryOnHTTP429 = false
		}
	}
}

// WithTimeout sets the timeout parameter for the gotion client.
func WithTimeout(t time.Duration) Option {
	return func(c *Client) {
		if c != nil {
			c.httpClient.Timeout = t
		}
	}
}

// WithMaxRetries sets the maximum number of retries for the gotion client.
func WithMaxRetries(n int) Option {
	return func(c *Client) {
		if c != nil {
			c.httpClient.MaxRetries = n
		}
	}
}
