// Package phpipam provides a Go client for the phpIPAM API.
package phpipam

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	defaultTimeout = 30 * time.Second
)

// Client represents a phpIPAM API client
type Client struct {
	BaseURL     *url.URL
	AppID       string
	Username    string
	Password    string
	Token       string
	TokenExp    time.Time
	HTTPClient  *http.Client
	UserAgent   string
	InsecureTLS bool
}

// Response represents a phpIPAM API response
type Response struct {
	Code    int             `json:"code"`
	Success bool            `json:"success"`
	Message string          `json:"message,omitempty"`
	ID      int             `json:"id,omitempty"`
	Data    json.RawMessage `json:"data,omitempty"`
	Time    float64         `json:"time,omitempty"`
}

// TokenResponse represents the authentication token response
type TokenResponse struct {
	Token   string `json:"token"`
	Expires string `json:"expires"`
}

// NewClient creates a new phpIPAM API client
func NewClient(baseURL, appID, username, password string, insecureTLS bool) (*Client, error) {
	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	// Ensure URL ends with /api/
	path := parsedURL.Path
	if !strings.HasSuffix(path, "/") {
		path += "/"
	}
	if !strings.HasSuffix(path, "/api/") {
		path = strings.TrimSuffix(path, "/") + "/api/"
	}
	parsedURL.Path = path

	// Create a transport with the appropriate TLS configuration
	transport := http.DefaultTransport.(*http.Transport).Clone()
	if insecureTLS {
		// This is the critical part - create a new TLS config that skips verification
		transport.TLSClientConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	}

	// Create the HTTP client with the configured transport
	httpClient := &http.Client{
		Timeout:   defaultTimeout,
		Transport: transport,
	}

	return &Client{
		BaseURL:     parsedURL,
		AppID:       appID,
		Username:    username,
		Password:    password,
		HTTPClient:  httpClient,
		UserAgent:   "go-phpipam/1.0",
		InsecureTLS: insecureTLS,
	}, nil
}

// SetTimeout sets a custom timeout for the HTTP client
func (c *Client) SetTimeout(timeout time.Duration) {
	c.HTTPClient.Timeout = timeout
}

// Authenticate performs authentication with the phpIPAM API and retrieves a token
func (c *Client) Authenticate() error {
	req, err := c.newRequest("POST", "user", nil)
	if err != nil {
		return err
	}

	req.SetBasicAuth(c.Username, c.Password)

	var tokenResp TokenResponse
	resp, err := c.do(req, &tokenResp)
	if err != nil {
		return err
	}

	if !resp.Success {
		return fmt.Errorf("authentication failed: %s", resp.Message)
	}

	c.Token = tokenResp.Token

	// Parse expiration time
	expTime, err := time.Parse("2006-01-02 15:04:05", tokenResp.Expires)
	if err != nil {
		return fmt.Errorf("failed to parse token expiration time: %v", err)
	}
	c.TokenExp = expTime

	return nil
}

// newRequest creates a new HTTP request to the phpIPAM API
func (c *Client) newRequest(method, endpoint string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(fmt.Sprintf("%s/%s/", c.AppID, endpoint))
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Set("Accept", "application/json")

	// Add token to request if it exists
	if c.Token != "" {
		req.Header.Set("token", c.Token)
		req.Header.Set("phpipam-token", c.Token) // some implementations use this header
	}

	return req, nil
}

// do sends an HTTP request and returns an API response
func (c *Client) do(req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	apiResp := &Response{}
	err = json.NewDecoder(resp.Body).Decode(apiResp)
	if err != nil {
		return nil, err
	}

	if v != nil && apiResp.Data != nil {
		err = json.Unmarshal(apiResp.Data, v)
		if err != nil {
			return nil, err
		}
	}

	return apiResp, nil
}

// IsTokenValid checks if the current token is still valid
func (c *Client) IsTokenValid() bool {
	if c.Token == "" {
		return false
	}

	// Add 5 minute buffer to ensure we don't use a token that's about to expire
	return time.Now().Add(5 * time.Minute).Before(c.TokenExp)
}

// RefreshToken extends the validity of the current token
func (c *Client) RefreshToken() error {
	if c.Token == "" {
		return fmt.Errorf("no token to refresh, authenticate first")
	}

	req, err := c.newRequest("PATCH", "user", nil)
	if err != nil {
		return err
	}

	var tokenResp TokenResponse
	resp, err := c.do(req, &tokenResp)
	if err != nil {
		return err
	}

	if !resp.Success {
		return fmt.Errorf("token refresh failed: %s", resp.Message)
	}

	// Parse expiration time
	expTime, err := time.Parse("2006-01-02 15:04:05", tokenResp.Expires)
	if err != nil {
		return fmt.Errorf("failed to parse token expiration time: %v", err)
	}
	c.TokenExp = expTime

	return nil
}

// EnsureAuthenticated makes sure that the client has a valid authentication token
func (c *Client) EnsureAuthenticated() error {
	if !c.IsTokenValid() {
		return c.Authenticate()
	}
	return nil
}

// Request performs an API request ensuring the client is authenticated
func (c *Client) Request(method, endpoint string, body, result interface{}) (*Response, error) {
	err := c.EnsureAuthenticated()
	if err != nil {
		return nil, err
	}

	req, err := c.newRequest(method, endpoint, body)
	if err != nil {
		return nil, err
	}

	return c.do(req, result)
}
