// Copyright (c) 2025-2025 All rights reserved.
//
// The original source code is licensed under the Apache License 2.0.
//
// You may review the terms of both licenses in the LICENSE file.

package gitness

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/imroc/req/v3"
)

const (
	defaultBaseURL = "https://gitness.com/"
	apiVersionPath = "api/v1"
	userAgent      = "go-gitness"
)

// Client represents a Gitness API client
type Client struct {
	client  *req.Client
	baseURL string
	token   string

	// Services
	Admin          *AdminService
	Audit          *AuditService
	Auth           *AuthService
	Checks         *ChecksService
	CiCache        *CiCacheService
	Connectors     *ConnectorsService
	Gitspaces      *GitspacesService
	InfraProviders *InfraProvidersService
	Pipelines      *PipelinesService
	Principals     *PrincipalsService
	Plugins        *PluginsService
	PullRequests   *PullRequestsService
	Repositories   *RepositoriesService
	Resource       *ResourceService
	Secrets        *SecretsService
	Spaces         *SpacesService
	System         *SystemService
	Templates      *TemplatesService
	Upload         *UploadService
	Users          *UsersService
	Webhooks       *WebhooksService
}

// ClientOptionFunc defines option functions for configuring the client
type ClientOptionFunc func(*Client) error

// NewClient creates a new Gitness API client
func NewClient(token string, options ...ClientOptionFunc) (*Client, error) {
	baseURL := defaultBaseURL

	// Create req client with default configuration
	reqClient := req.C().
		SetUserAgent(userAgent).
		SetTimeout(10 * time.Second).
		SetCommonBearerAuthToken(token).
		SetCommonContentType("application/json")

	c := &Client{
		client:  reqClient,
		baseURL: baseURL,
		token:   token,
	}

	// Apply options
	for _, option := range options {
		if err := option(c); err != nil {
			return nil, err
		}
	}

	// Set the base URL with API version
	apiURL := c.baseURL + apiVersionPath
	c.client.SetBaseURL(apiURL)

	// Initialize services
	c.Admin = &AdminService{client: c}
	c.Audit = &AuditService{client: c}
	c.Auth = &AuthService{client: c}
	c.Checks = &ChecksService{client: c}
	c.CiCache = &CiCacheService{client: c}
	c.Connectors = &ConnectorsService{client: c}
	c.Gitspaces = &GitspacesService{client: c}
	c.InfraProviders = &InfraProvidersService{client: c}
	c.Pipelines = &PipelinesService{client: c}
	c.Principals = &PrincipalsService{client: c}
	c.Plugins = &PluginsService{client: c}
	c.PullRequests = &PullRequestsService{client: c}
	c.Repositories = &RepositoriesService{client: c}
	c.Resource = &ResourceService{client: c}
	c.Secrets = &SecretsService{client: c}
	c.Spaces = &SpacesService{client: c}
	c.System = &SystemService{client: c}
	c.Templates = &TemplatesService{client: c}
	c.Upload = &UploadService{client: c}
	c.Users = &UsersService{client: c}
	c.Webhooks = &WebhooksService{client: c}

	return c, nil
}

// WithBaseURL sets a custom base URL for the client
func WithBaseURL(baseURL string) ClientOptionFunc {
	return func(c *Client) error {
		u, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.baseURL = u.String()
		return nil
	}
}

// WithHTTPClient sets a custom HTTP client
func WithHTTPClient(httpClient *http.Client) ClientOptionFunc {
	return func(c *Client) error {
		// For req/v3, we can set transport via the underlying client
		// This is a workaround since req/v3 doesn't expose SetHTTPClient directly
		return nil // Skip setting HTTP client for now
	}
}

// WithTimeout sets a custom timeout for HTTP requests
func WithTimeout(timeout time.Duration) ClientOptionFunc {
	return func(c *Client) error {
		c.client.SetTimeout(timeout)
		return nil
	}
}

// WithDebug enables debug logging for HTTP requests
func WithDebug() ClientOptionFunc {
	return func(c *Client) error {
		c.client.EnableDebugLog()
		return nil
	}
}

// WithRetry enables retry mechanism with default configuration
func WithRetry(retryCount int) ClientOptionFunc {
	return func(c *Client) error {
		if retryCount > 0 {
			c.client.SetCommonRetryCount(retryCount)
		}
		return nil
	}
}

// Response wraps an HTTP response from req/v3 with pagination information
type Response struct {
	*req.Response

	// Pagination info from headers
	Page       *int `json:"page,omitempty"`
	PerPage    *int `json:"per_page,omitempty"`
	NextPage   *int `json:"next_page,omitempty"`
	Total      *int `json:"total,omitempty"`
	TotalPages *int `json:"total_pages,omitempty"`
}

// ErrorResponse represents an error response from the API
type ErrorResponse struct {
	Response *req.Response `json:"-"`
	Message  string        `json:"message"`
	Details  string        `json:"details,omitempty"`
}

func (e *ErrorResponse) Error() string {
	if e.Response != nil {
		return fmt.Sprintf("%v %v: %d %s",
			e.Response.Request.Method, e.Response.Request.URL,
			e.Response.StatusCode, e.Message)
	}
	return e.Message
}

// Get performs a GET request
func (c *Client) Get(ctx context.Context, path string, result any) (*Response, error) {
	fullURL := c.buildFullURL(path)
	resp, err := c.client.R().
		SetContext(ctx).
		SetSuccessResult(result).
		Get(fullURL)

	if err != nil {
		return nil, err
	}

	if err := c.checkResponse(resp); err != nil {
		return &Response{Response: resp}, err
	}

	// Parse pagination headers
	response := &Response{Response: resp}
	c.parsePaginationHeaders(response)

	return response, nil
}

// Post performs a POST request
func (c *Client) Post(ctx context.Context, path string, body any, result any) (*Response, error) {
	fullURL := c.buildFullURL(path)
	req := c.client.R().SetContext(ctx)

	if body != nil {
		req.SetBodyJsonMarshal(body)
	}

	if result != nil {
		req.SetSuccessResult(result)
	}

	resp, err := req.Post(fullURL)
	if err != nil {
		return nil, err
	}

	if err := c.checkResponse(resp); err != nil {
		return &Response{Response: resp}, err
	}

	return &Response{Response: resp}, nil
}

// Put performs a PUT request
func (c *Client) Put(ctx context.Context, path string, body any, result any) (*Response, error) {
	fullURL := c.buildFullURL(path)
	req := c.client.R().SetContext(ctx)

	if body != nil {
		req.SetBodyJsonMarshal(body)
	}

	if result != nil {
		req.SetSuccessResult(result)
	}

	resp, err := req.Put(fullURL)
	if err != nil {
		return nil, err
	}

	if err := c.checkResponse(resp); err != nil {
		return &Response{Response: resp}, err
	}

	return &Response{Response: resp}, nil
}

// Patch performs a PATCH request
func (c *Client) Patch(ctx context.Context, path string, body any, result any) (*Response, error) {
	fullURL := c.buildFullURL(path)
	req := c.client.R().SetContext(ctx)

	if body != nil {
		req.SetBodyJsonMarshal(body)
	}

	if result != nil {
		req.SetSuccessResult(result)
	}

	resp, err := req.Patch(fullURL)
	if err != nil {
		return nil, err
	}

	if err := c.checkResponse(resp); err != nil {
		return &Response{Response: resp}, err
	}

	return &Response{Response: resp}, nil
}

// Delete performs a DELETE request
func (c *Client) Delete(ctx context.Context, path string, body any) (*Response, error) {
	fullURL := c.buildFullURL(path)
	req := c.client.R().SetContext(ctx)

	if body != nil {
		req.SetBodyJsonMarshal(body)
	}

	resp, err := req.Delete(fullURL)
	if err != nil {
		return nil, err
	}

	if err := c.checkResponse(resp); err != nil {
		return &Response{Response: resp}, err
	}

	return &Response{Response: resp}, nil
}

// DeleteWithResponse performs a DELETE request and returns the response body
func (c *Client) DeleteWithResponse(ctx context.Context, path string, body any, result any) (*Response, error) {
	fullURL := c.buildFullURL(path)
	req := c.client.R().SetContext(ctx)

	if body != nil {
		req.SetBodyJsonMarshal(body)
	}

	if result != nil {
		req.SetSuccessResult(result)
	}

	resp, err := req.Delete(fullURL)
	if err != nil {
		return nil, err
	}

	if err := c.checkResponse(resp); err != nil {
		return &Response{Response: resp}, err
	}

	return &Response{Response: resp}, nil
}

// checkResponse checks for API errors
func (c *Client) checkResponse(r *req.Response) error {
	if r.IsSuccessState() {
		return nil
	}

	errorResponse := &ErrorResponse{Response: r}

	// Try to parse error from response body
	var errorBody map[string]any
	if err := json.Unmarshal(r.Bytes(), &errorBody); err == nil {
		if message, ok := errorBody["message"].(string); ok {
			errorResponse.Message = message
		}
		if details, ok := errorBody["details"].(string); ok {
			errorResponse.Details = details
		}
	}

	if errorResponse.Message == "" {
		errorResponse.Message = fmt.Sprintf("HTTP %d: %s", r.StatusCode, http.StatusText(r.StatusCode))
	}

	return errorResponse
}

// buildFullURL constructs a full URL from the base URL and path, preserving URL encoding
func (c *Client) buildFullURL(path string) string {
	baseURL, _ := url.Parse(c.baseURL + apiVersionPath + "/")
	fullURL := baseURL.ResolveReference(&url.URL{Path: path})
	return fullURL.String()
}

// buildQueryParams is a helper function to build query parameters from ListOptions
func buildQueryParams(req *req.Request, opt *ListOptions) {
	if opt == nil {
		return
	}

	if opt.Page != nil {
		req.SetQueryParam("page", fmt.Sprintf("%d", *opt.Page))
	}
	if opt.Limit != nil {
		req.SetQueryParam("limit", fmt.Sprintf("%d", *opt.Limit))
	}
	if opt.Sort != nil {
		req.SetQueryParam("sort", *opt.Sort)
	}
	if opt.Order != nil {
		req.SetQueryParam("order", *opt.Order)
	}
	if opt.Query != nil {
		req.SetQueryParam("query", *opt.Query)
	}
}

// performListRequest is a helper function for making list requests with pagination support
func (c *Client) performListRequest(ctx context.Context, path string, opt *ListOptions, result any) (*Response, error) {
	fullURL := c.buildFullURL(path)
	req := c.client.R().SetContext(ctx)
	req.SetSuccessResult(result)

	// Add common query parameters
	buildQueryParams(req, opt)

	resp, err := req.Get(fullURL)
	if err != nil {
		return &Response{Response: resp}, err
	}

	if err := c.checkResponse(resp); err != nil {
		return &Response{Response: resp}, err
	}

	// Parse pagination headers
	response := &Response{Response: resp}
	c.parsePaginationHeaders(response)

	return response, nil
}

// parsePaginationHeaders parses pagination information from response headers
func (c *Client) parsePaginationHeaders(response *Response) {
	if response.Response == nil {
		return
	}

	headers := response.Response.Header

	// Parse x-page
	if page := headers.Get("x-page"); page != "" {
		if val, err := strconv.Atoi(page); err == nil {
			response.Page = &val
		}
	}

	// Parse x-per-page
	if perPage := headers.Get("x-per-page"); perPage != "" {
		if val, err := strconv.Atoi(perPage); err == nil {
			response.PerPage = &val
		}
	}

	// Parse x-next-page
	if nextPage := headers.Get("x-next-page"); nextPage != "" {
		if val, err := strconv.Atoi(nextPage); err == nil {
			response.NextPage = &val
		}
	}

	// Parse x-total
	if total := headers.Get("x-total"); total != "" {
		if val, err := strconv.Atoi(total); err == nil {
			response.Total = &val
		}
	}

	// Parse x-total-pages
	if totalPages := headers.Get("x-total-pages"); totalPages != "" {
		if val, err := strconv.Atoi(totalPages); err == nil {
			response.TotalPages = &val
		}
	}
}

// Ptr returns a pointer to the provided value
func Ptr[T any](v T) *T {
	return &v
}

// Time represents a time value that can be unmarshaled from a JSON string
type Time time.Time

// UnmarshalJSON implements the json.Unmarshaler interface
func (t *Time) UnmarshalJSON(data []byte) error {
	var timeStr string
	if err := json.Unmarshal(data, &timeStr); err != nil {
		return err
	}

	parsedTime, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		return err
	}

	*t = Time(parsedTime)
	return nil
}

// MarshalJSON implements the json.Marshaler interface
func (t Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(t).Format(time.RFC3339))
}

// String returns the time formatted as RFC3339
func (t Time) String() string {
	return time.Time(t).Format(time.RFC3339)
}

// ListOptions specifies general pagination options
type ListOptions struct {
	Page  *int    `json:"page,omitempty" url:"page,omitempty"`
	Limit *int    `json:"limit,omitempty" url:"limit,omitempty"` // Gitness uses 'limit' not 'per_page'
	Sort  *string `json:"sort,omitempty" url:"sort,omitempty"`
	Order *string `json:"order,omitempty" url:"order,omitempty"`
	Query *string `json:"query,omitempty" url:"query,omitempty"`
}
