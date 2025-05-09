package phpipam

import (
	"fmt"
	"net/url"
)

// SearchResult represents a phpIPAM search result
type SearchResult struct {
	Subnets   []Subnet   `json:"subnets,omitempty"`
	Addresses []Address  `json:"addresses,omitempty"`
	VLANs     []VLAN     `json:"vlans,omitempty"`
	VRFs      []VRF      `json:"vrfs,omitempty"`
}

// SearchOptions represents search options for phpIPAM search
type SearchOptions struct {
	IncludeAddresses bool
	IncludeSubnets   bool
	IncludeVLANs     bool
	IncludeVRFs      bool
}

// DefaultSearchOptions provides default search options
var DefaultSearchOptions = SearchOptions{
	IncludeAddresses: true,
	IncludeSubnets:   true,
	IncludeVLANs:     false,
	IncludeVRFs:      false,
}

// SearchService handles communication with the search related methods of the API
type SearchService struct {
	client *Client
}

// NewSearchService creates a new search service with the provided client
func NewSearchService(client *Client) *SearchService {
	return &SearchService{client: client}
}

// Search searches phpipam database for the required string with default options
func (s *SearchService) Search(searchString string) (*SearchResult, error) {
	return s.SearchWithOptions(searchString, DefaultSearchOptions)
}

// SearchWithOptions searches phpipam database for the required string with custom options
func (s *SearchService) SearchWithOptions(searchString string, options SearchOptions) (*SearchResult, error) {
	// Build query parameters
	query := url.Values{}
	query.Set("addresses", boolToStr(options.IncludeAddresses))
	query.Set("subnets", boolToStr(options.IncludeSubnets))
	query.Set("vlan", boolToStr(options.IncludeVLANs))
	query.Set("vrf", boolToStr(options.IncludeVRFs))

	// URL escape the search string
	escapedSearch := url.QueryEscape(searchString)
	
	var result SearchResult
	_, err := s.client.Request("GET", fmt.Sprintf("search/%s?%s", escapedSearch, query.Encode()), nil, &result)
	return &result, err
}

// Helper function to convert bool to string "1" or "0"
func boolToStr(b bool) string {
	if b {
		return "1"
	}
	return "0"
}
