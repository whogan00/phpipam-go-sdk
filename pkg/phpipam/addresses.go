package phpipam

import (
	"fmt"
	"strconv"
	"net/url"
)

// Address represents a phpIPAM address object
type Address struct {
	ID          string `json:"id,omitempty"`
	SubnetID    string `json:"subnetId,omitempty"`
	IP          string `json:"ip,omitempty"`
	IsGateway   string `json:"is_gateway,omitempty"`
	Description string `json:"description,omitempty"`
	Hostname    string `json:"hostname,omitempty"`
	Mac         string `json:"mac,omitempty"`
	Owner       string `json:"owner,omitempty"`
	Tag         string `json:"tag,omitempty"`
	PTRIgnore   string `json:"PTRignore,omitempty"`
	PTR         string `json:"PTR,omitempty"`
	DeviceID    string `json:"deviceId,omitempty"`
	Port        string `json:"port,omitempty"`
	Note        string `json:"note,omitempty"`
	LastSeen    string `json:"lastSeen,omitempty"`
	ExcludePing string `json:"excludePing,omitempty"`
	EditDate    string `json:"editDate,omitempty"`
}

// Tag represents an IP address tag
type Tag struct {
	ID          string `json:"id,omitempty"`
	Type        string `json:"type,omitempty"`
	ShowTag     string `json:"showtag,omitempty"`
	BgColor     string `json:"bgcolor,omitempty"`
	FgColor     string `json:"fgcolor,omitempty"`
	DisplayName string `json:"displayname,omitempty"`
	Description string `json:"description,omitempty"`
}

// AddressesService handles communication with the addresses related methods of the API
type AddressesService struct {
	client *Client
}

// NewAddressesService creates a new addresses service with the provided client
func NewAddressesService(client *Client) *AddressesService {
	return &AddressesService{client: client}
}

// Get returns a specific address by ID
func (a *AddressesService) Get(id string) (*Address, error) {
	var address Address
	_, err := a.client.Request("GET", fmt.Sprintf("addresses/%s", id), nil, &address)
	return &address, err
}

// GetAll returns all addresses in all sections
func (a *AddressesService) GetAll() ([]Address, error) {
	var addresses []Address
	_, err := a.client.Request("GET", "addresses/all", nil, &addresses)
	return addresses, err
}

// Ping checks the status of an address
func (a *AddressesService) Ping(id string) (map[string]interface{}, error) {
	var result map[string]interface{}
	_, err := a.client.Request("GET", fmt.Sprintf("addresses/%s/ping", id), nil, &result)
	return result, err
}

// GetByIPAndSubnet returns an address from a subnet by IP address
func (a *AddressesService) GetByIPAndSubnet(ip string, subnetID string) (*Address, error) {
	var address Address
	_, err := a.client.Request("GET", fmt.Sprintf("addresses/%s/%s", ip, subnetID), nil, &address)
	return &address, err
}

// Search searches for addresses in database by IP
func (a *AddressesService) Search(ip string) ([]Address, error) {
	var addresses []Address
	_, err := a.client.Request("GET", fmt.Sprintf("addresses/search/%s", ip), nil, &addresses)
	return addresses, err
}

// SearchByHostname searches for addresses in database by hostname
func (a *AddressesService) SearchByHostname(hostname string) ([]Address, error) {
	var addresses []Address
	_, err := a.client.Request("GET", fmt.Sprintf("addresses/search_hostname/%s", url.QueryEscape(hostname)), nil, &addresses)
	return addresses, err
}

// SearchByLinkedValue searches for addresses linked by custom "Link addresses" field
func (a *AddressesService) SearchByLinkedValue(value string) ([]Address, error) {
	var addresses []Address
	_, err := a.client.Request("GET", fmt.Sprintf("addresses/search_linked/%s", url.QueryEscape(value)), nil, &addresses)
	return addresses, err
}

// SearchByHostbase searches for addresses by leading substring (base) of hostname
func (a *AddressesService) SearchByHostbase(hostbase string) ([]Address, error) {
	var addresses []Address
	_, err := a.client.Request("GET", fmt.Sprintf("addresses/search_hostbase/%s", url.QueryEscape(hostbase)), nil, &addresses)
	return addresses, err
}

// SearchByMAC searches for addresses by MAC address
func (a *AddressesService) SearchByMAC(mac string) ([]Address, error) {
	var addresses []Address
	_, err := a.client.Request("GET", fmt.Sprintf("addresses/search_mac/%s", url.QueryEscape(mac)), nil, &addresses)
	return addresses, err
}

// GetFirstFree returns the first available address in a subnet
func (a *AddressesService) GetFirstFree(subnetID string) (string, error) {
	var firstFree string
	_, err := a.client.Request("GET", fmt.Sprintf("addresses/first_free/%s", subnetID), nil, &firstFree)
	return firstFree, err
}

// GetCustomFields returns custom fields for addresses
func (a *AddressesService) GetCustomFields() (map[string]CustomField, error) {
	var customFields map[string]CustomField
	_, err := a.client.Request("GET", "addresses/custom_fields", nil, &customFields)
	return customFields, err
}

// GetTags returns all address tags
func (a *AddressesService) GetTags() ([]Tag, error) {
	var tags []Tag
	_, err := a.client.Request("GET", "addresses/tags", nil, &tags)
	return tags, err
}

// GetTag returns a specific address tag
func (a *AddressesService) GetTag(id string) (*Tag, error) {
	var tag Tag
	_, err := a.client.Request("GET", fmt.Sprintf("addresses/tags/%s", id), nil, &tag)
	return &tag, err
}

// GetAddressesByTag returns addresses for a specific tag
func (a *AddressesService) GetAddressesByTag(id string) ([]Address, error) {
	var addresses []Address
	_, err := a.client.Request("GET", fmt.Sprintf("addresses/tags/%s/addresses", id), nil, &addresses)
	return addresses, err
}

// Create creates a new address
func (a *AddressesService) Create(address *Address) (*Address, error) {
	var createdAddress Address
	resp, err := a.client.Request("POST", "addresses", address, &createdAddress)
	if err != nil {
		return nil, err
	}

	// If we got an ID in the response but not in the address data, retrieve the full address
	if resp.ID != 0 && createdAddress.ID == "" {
		return a.Get(strconv.Itoa(resp.ID))
	}

	return &createdAddress, nil
}

// CreateFirstFree creates a new address in a subnet - first available
func (a *AddressesService) CreateFirstFree(subnetID string, address *Address) (*Address, error) {
	var createdAddress Address
	resp, err := a.client.Request("POST", fmt.Sprintf("addresses/first_free/%s", subnetID), address, &createdAddress)
	if err != nil {
		return nil, err
	}

	// If we got an ID in the response but not in the address data, retrieve the full address
	if resp.ID != 0 && createdAddress.ID == "" {
		return a.Get(strconv.Itoa(resp.ID))
	}

	return &createdAddress, nil
}

// Update updates an address
func (a *AddressesService) Update(address *Address) (*Address, error) {
	if address.ID == "" {
		return nil, fmt.Errorf("address ID is required for update")
	}

	var updatedAddress Address
	_, err := a.client.Request("PATCH", fmt.Sprintf("addresses/%s", address.ID), address, &updatedAddress)
	return &updatedAddress, err
}

// Delete deletes an address
func (a *AddressesService) Delete(id string) error {
	_, err := a.client.Request("DELETE", fmt.Sprintf("addresses/%s", id), nil, nil)
	return err
}

// DeleteWithRemoveDNS deletes an address and removes all related DNS records
func (a *AddressesService) DeleteWithRemoveDNS(id string) error {
	params := map[string]string{"remove_dns": "1"}
	_, err := a.client.Request("DELETE", fmt.Sprintf("addresses/%s", id), params, nil)
	return err
}

// DeleteByIPAndSubnet deletes an address by IP in a specific subnet
func (a *AddressesService) DeleteByIPAndSubnet(ip string, subnetID string) error {
	_, err := a.client.Request("DELETE", fmt.Sprintf("addresses/%s/%s", ip, subnetID), nil, nil)
	return err
}
