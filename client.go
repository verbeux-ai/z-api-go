package z_api

import (
	"net/http"
)

type Client struct {
	token      string
	secret     string
	instance   string
	baseUrl    string
	httpClient *http.Client
}

// Option is a function that configures a client
type Option func(*Client)

// NewClient creates a new client with the provided options
func NewClient(opts ...Option) *Client {
	c := &Client{
		baseUrl:    "https://api.z-api.io",
		httpClient: http.DefaultClient,
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

// WithToken sets the token of the client
func WithToken(token string) Option {
	return func(c *Client) {
		c.token = token
	}
}

// WithSecret sets the token of the client
func WithSecret(secret string) Option {
	return func(c *Client) {
		c.secret = secret
	}
}

// WithBaseUrl sets the base URL of the client
func WithBaseUrl(baseUrl string) Option {
	return func(c *Client) {
		c.baseUrl = baseUrl
	}
}

// WithInstance sets the instance id of the client
func WithInstance(instance string) Option {
	return func(c *Client) {
		c.instance = instance
	}
}

// WithHttpClient sets the http client of the client
func WithHttpClient(httpClient *http.Client) Option {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}
