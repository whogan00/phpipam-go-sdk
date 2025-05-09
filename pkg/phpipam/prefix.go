package phpipam

import (
	"fmt"
)

// IPVersion represents the IP version type
type IPVersion string

const (
	// IPv4 represents IPv4 address family
	IPv4 IPVersion = "v4"
	// IPv6 represents IPv6 address family
	IPv6 IPVersion = "v6"
)

// PrefixService handles communication with the prefix related methods of the API
type PrefixService struct {
	client *Client
}

// NewPrefixService creates a new prefix service with the provided client
func NewPrefixService(client *Client) *PrefixService {
	return &PrefixService{client: client}
}

// GetSubnets returns all subnets used to deliver new subnets
func (p *PrefixService) GetSubnets(customerType string) ([]Subnet, error) {
	var subnets []Subnet
	_, err := p.client.Request("GET", fmt.Sprintf("prefix/%s", customerType), nil, &subnets)
	return subnets, err
}

// GetSubnetsForIPVersion returns all subnets used to deliver new subnets for specific IP version
func (p *PrefixService) GetSubnetsForIPVersion(customerType string, addressType IPVersion) ([]Subnet, error) {
	var subnets []Subnet
	_, err := p.client.Request("GET", fmt.Sprintf("prefix/%s/%s", customerType, addressType), nil, &subnets)
	return subnets, err
}

// GetSubnetsForMask returns all subnets used to deliver new subnets for specific IP version and mask
func (p *PrefixService) GetSubnetsForMask(customerType string, addressType IPVersion, mask int) ([]Subnet, error) {
	var subnets []Subnet
	_, err := p.client.Request("GET", fmt.Sprintf("prefix/%s/%s/%d", customerType, addressType, mask), nil, &subnets)
	return subnets, err
}

// GetSubnetsByExternalID returns subnets by external identifier field
func (p *PrefixService) GetSubnetsByExternalID(externalID string) ([]Subnet, error) {
	var subnets []Subnet
	_, err := p.client.Request("GET", fmt.Sprintf("prefix/external_id/%s", externalID), nil, &subnets)
	return subnets, err
}

// GetFirstAvailableSubnet returns first available subnet for IP version and requested mask
func (p *PrefixService) GetFirstAvailableSubnet(customerType string, addressType IPVersion, mask int) (string, error) {
	var subnet string
	_, err := p.client.Request("GET", fmt.Sprintf("prefix/%s/%s/%d", customerType, addressType, mask), nil, &subnet)
	return subnet, err
}

// GetFirstAvailableAddress returns first available address for IP version
func (p *PrefixService) GetFirstAvailableAddress(customerType string, addressType IPVersion) (string, error) {
	var address string
	_, err := p.client.Request("GET", fmt.Sprintf("prefix/%s/%s/address", customerType, addressType), nil, &address)
	return address, err
}

// CreateFirstAvailableSubnet creates first available subnet for IP version and requested mask
func (p *PrefixService) CreateFirstAvailableSubnet(customerType string, addressType IPVersion, mask int, subnet *Subnet) (*Subnet, error) {
	var createdSubnet Subnet
	_, err := p.client.Request("POST", fmt.Sprintf("prefix/%s/%s/%d", customerType, addressType, mask), subnet, &createdSubnet)
	return &createdSubnet, err
}

// CreateFirstAvailableAddress creates first available address for IP version
func (p *PrefixService) CreateFirstAvailableAddress(customerType string, addressType IPVersion, address *Address) (*Address, error) {
	var createdAddress Address
	_, err := p.client.Request("POST", fmt.Sprintf("prefix/%s/%s/address", customerType, addressType), address, &createdAddress)
	return &createdAddress, err
}
