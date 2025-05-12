package phpipam

import (
	"fmt"
	"net/url"
	"strconv"
)

// Address represents a phpIPAM address object
type Address struct {
	ID          int    `json:"id,omitempty"`
	SubnetID    int    `json:"subnetId,omitempty"`
	IP          string `json:"ip,omitempty"`
	IsGateway   int    `json:"is_gateway,omitempty"`
	Description string `json:"description,omitempty"`
	Hostname    string `json:"hostname,omitempty"`
	Mac         string `json:"mac,omitempty"`
	Owner       string `json:"owner,omitempty"`
	Tag         int    `json:"tag,omitempty"`
	PTRIgnore   int    `json:"PTRignore,omitempty"`
	PTR         int    `json:"PTR,omitempty"`
	DeviceID    int    `json:"deviceId,omitempty"`
	Port        string `json:"port,omitempty"`
	Note        string `json:"note,omitempty"`
	LastSeen    string `json:"lastSeen,omitempty"`
	ExcludePing int    `json:"excludePing,omitempty"`
	EditDate    string `json:"editDate,omitempty"`
}

// Tag represents an IP address tag
type Tag struct {
	ID          int    `json:"id,omitempty"`
	Type        string `json:"type,omitempty"`
	ShowTag     int    `json:"showtag,omitempty"`
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
func (a *AddressesService) Get(id int) (*Address, error) {
	var address Address
	_, err := a.client.Request("GET", fmt.Sprintf("addresses/%d", id), nil, &address)
	return &address, err
}

// GetByStringID returns a specific address by string ID (convenience method)
func (a *AddressesService) GetByStringID(id string) (*Address, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return nil, fmt.Errorf("invalid address ID: %w", err)
	}
	return a.Get(idInt)
}

// GetAll returns all addresses in all sections
func (a *AddressesService) GetAll() ([]Address, error) {
	var addresses []Address
	_, err := a.client.Request("GET", "addresses/all", nil, &addresses)
	return addresses, err
}

// Ping checks the status of an address
func (a *AddressesService) Ping(id int) (map[string]interface{}, error) {
	var result map[string]interface{}
	_, err := a.client.Request("GET", fmt.Sprintf("addresses/%d/ping", id), nil, &result)
	return result, err
}

// GetByIPAndSubnet returns an address from a subnet by IP address
func (a *AddressesService) GetByIPAndSubnet(ip string, subnetID int) (*Address, error) {
	var address Address
	_, err := a.client.Request("GET", fmt.Sprintf("addresses/%s/%d", ip, subnetID), nil, &address)
	return &address, err
}

// GetByIPAndSubnetString returns an address using a string subnet ID (convenience method)
func (a *AddressesService) GetByIPAndSubnetString(ip string, subnetID string) (*Address, error) {
	subnetIDInt, err := strconv.Atoi(subnetID)
	if err != nil {
		return nil, fmt.Errorf("invalid subnet ID: %w", err)
	}
	return a.GetByIPAndSubnet(ip, subnetIDInt)
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
func (a *AddressesService) GetFirstFree(subnetID int) (string, error) {
	var firstFree string
	_, err := a.client.Request("GET", fmt.Sprintf("addresses/first_free/%d", subnetID), nil, &firstFree)
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
func (a *AddressesService) GetTag(id int) (*Tag, error) {
	var tag Tag
	_, err := a.client.Request("GET", fmt.Sprintf("addresses/tags/%d", id), nil, &tag)
	return &tag, err
}

// GetAddressesByTag returns addresses for a specific tag
func (a *AddressesService) GetAddressesByTag(id int) ([]Address, error) {
	var addresses []Address
	_, err := a.client.Request("GET", fmt.Sprintf("addresses/tags/%d/addresses", id), nil, &addresses)
	return addresses, err
}

// Create creates a new address
func (a *AddressesService) Create(address *Address) (*Address, error) {
	var createdAddress Address
	resp, err := a.client.Request("POST", "addresses", address, &createdAddress)
	fmt.Printf("response: %v\n", resp)
	if err != nil {
		return nil, err
	}

	// If we got an ID in the response but not in the address data, retrieve the full address
	if resp.ID != 0 && createdAddress.ID == 0 {
		return a.Get(resp.ID)
	}

	return &createdAddress, nil
}

// CreateFirstFree creates a new address in a subnet - first available
func (a *AddressesService) CreateFirstFree(subnetID int, address *Address) (*Address, error) {
	var createdAddress Address
	resp, err := a.client.Request("POST", fmt.Sprintf("addresses/first_free/%d", subnetID), address, &createdAddress)
	if err != nil {
		return nil, err
	}

	// If we got an ID in the response but not in the address data, retrieve the full address
	if resp.ID != 0 && createdAddress.ID == 0 {
		return a.Get(resp.ID)
	}

	return &createdAddress, nil
}

// Update updates an address
func (a *AddressesService) Update(address *Address) (*Address, error) {
	if address.ID == 0 {
		return nil, fmt.Errorf("address ID is required for update")
	}

	var updatedAddress Address
	_, err := a.client.Request("PATCH", fmt.Sprintf("addresses/%d", address.ID), address, &updatedAddress)
	return &updatedAddress, err
}

// Delete deletes an address
func (a *AddressesService) Delete(id int) error {
	_, err := a.client.Request("DELETE", fmt.Sprintf("addresses/%d", id), nil, nil)
	return err
}

// DeleteWithRemoveDNS deletes an address and removes all related DNS records
func (a *AddressesService) DeleteWithRemoveDNS(id int) error {
	params := map[string]string{"remove_dns": "1"}
	_, err := a.client.Request("DELETE", fmt.Sprintf("addresses/%d", id), params, nil)
	return err
}

// DeleteByIPAndSubnet deletes an address by IP in a specific subnet
func (a *AddressesService) DeleteByIPAndSubnet(ip string, subnetID int) error {
	_, err := a.client.Request("DELETE", fmt.Sprintf("addresses/%s/%d", ip, subnetID), nil, nil)
	return err
}
