package phpipam

import (
	"fmt"
	"strconv"
)

// Subnet represents a phpIPAM subnet object
type Subnet struct {
	ID             int    `json:"id,omitempty"`
	Subnet         string `json:"subnet"`
	Mask           string `json:"mask"`
	SectionID      string `json:"sectionId,omitempty"`
	Description    string `json:"description,omitempty"`
	LinkedSubnet   string `json:"linked_subnet,omitempty"`
	VlanID         string `json:"vlanId,omitempty"`
	VrfID          string `json:"vrfId,omitempty"`
	MasterSubnetID string `json:"masterSubnetId,omitempty"`
	NameserverID   string `json:"nameserverId,omitempty"`
	ShowName       string `json:"showName,omitempty"`
	Permissions    string `json:"permissions,omitempty"`
	DNSRecursive   string `json:"DNSrecursive,omitempty"`
	DNSRecords     string `json:"DNSrecords,omitempty"`
	AllowRequests  string `json:"allowRequests,omitempty"`
	ScanAgent      string `json:"scanAgent,omitempty"`
	PingSubnet     string `json:"pingSubnet,omitempty"`
	DiscoverSubnet string `json:"discoverSubnet,omitempty"`
	IsFolder       string `json:"isFolder,omitempty"`
	IsFull         string `json:"isFull,omitempty"`
	State          string `json:"state,omitempty"`
	Threshold      string `json:"threshold,omitempty"`
	Location       string `json:"location,omitempty"`
	EditDate       string `json:"editDate,omitempty"`
}

// SubnetUsage represents usage statistics for a subnet
type SubnetUsage struct {
	Used             string  `json:"used"`
	MaxHosts         string  `json:"maxhosts"`
	Freehosts        string  `json:"freehosts"`
	FreehostsPercent float64 `json:"freehosts_percent"`
}

// SubnetsService handles communication with the subnets related methods of the API
type SubnetsService struct {
	client *Client
}

// NewSubnetsService creates a new subnets service with the provided client
func NewSubnetsService(client *Client) *SubnetsService {
	return &SubnetsService{client: client}
}

// List returns all subnets
func (s *SubnetsService) List() ([]Subnet, error) {
	var subnets []Subnet
	_, err := s.client.Request("GET", "subnets", nil, &subnets)
	return subnets, err
}

// Get returns a specific subnet by ID
func (s *SubnetsService) Get(id string) (*Subnet, error) {
	var subnet Subnet
	_, err := s.client.Request("GET", fmt.Sprintf("subnets/%s", id), nil, &subnet)
	return &subnet, err
}

// GetUsage returns usage statistics for a subnet
func (s *SubnetsService) GetUsage(id string) (*SubnetUsage, error) {
	var usage SubnetUsage
	_, err := s.client.Request("GET", fmt.Sprintf("subnets/%s/usage", id), nil, &usage)
	return &usage, err
}

// GetSlaves returns all immediate slave subnets
func (s *SubnetsService) GetSlaves(id string) ([]Subnet, error) {
	var subnets []Subnet
	_, err := s.client.Request("GET", fmt.Sprintf("subnets/%s/slaves", id), nil, &subnets)
	return subnets, err
}

// GetSlavesRecursive returns all slave subnets recursively
func (s *SubnetsService) GetSlavesRecursive(id string) ([]Subnet, error) {
	var subnets []Subnet
	_, err := s.client.Request("GET", fmt.Sprintf("subnets/%s/slaves_recursive", id), nil, &subnets)
	return subnets, err
}

// GetAddresses returns all addresses in a subnet
func (s *SubnetsService) GetAddresses(id string) ([]Address, error) {
	var addresses []Address
	_, err := s.client.Request("GET", fmt.Sprintf("subnets/%s/addresses", id), nil, &addresses)
	return addresses, err
}

// GetAddress returns a specific IP address from a subnet
func (s *SubnetsService) GetAddress(id, ip string) (*Address, error) {
	var address Address
	_, err := s.client.Request("GET", fmt.Sprintf("subnets/%s/addresses/%s", id, ip), nil, &address)
	return &address, err
}

// GetFirstFree returns the first available IP address in a subnet
func (s *SubnetsService) GetFirstFree(id string) (string, error) {
	var firstFree string
	_, err := s.client.Request("GET", fmt.Sprintf("subnets/%s/first_free", id), nil, &firstFree)
	return firstFree, err
}

// GetFirstSubnet returns the first available subnet within a given subnet for specified mask
func (s *SubnetsService) GetFirstSubnet(id string, mask int) (string, error) {
	var firstSubnet string
	_, err := s.client.Request("GET", fmt.Sprintf("subnets/%s/first_subnet/%d", id, mask), nil, &firstSubnet)
	return firstSubnet, err
}

// GetLastSubnet returns the last available subnet within a given subnet for specified mask
func (s *SubnetsService) GetLastSubnet(id string, mask int) (string, error) {
	var lastSubnet string
	_, err := s.client.Request("GET", fmt.Sprintf("subnets/%s/last_subnet/%d", id, mask), nil, &lastSubnet)
	return lastSubnet, err
}

// GetAllSubnets returns all available subnets within a given subnet for specified mask
func (s *SubnetsService) GetAllSubnets(id string, mask int) ([]string, error) {
	var allSubnets []string
	_, err := s.client.Request("GET", fmt.Sprintf("subnets/%s/all_subnets/%d", id, mask), nil, &allSubnets)
	return allSubnets, err
}

// GetCustomFields returns all subnet custom fields
func (s *SubnetsService) GetCustomFields() (map[string]CustomField, error) {
	var customFields map[string]CustomField
	_, err := s.client.Request("GET", "subnets/custom_fields", nil, &customFields)
	return customFields, err
}

// SearchBySubnet searches for a subnet in CIDR format
func (s *SubnetsService) SearchBySubnet(cidr string) ([]Subnet, error) {
	var subnets []Subnet
	_, err := s.client.Request("GET", fmt.Sprintf("subnets/cidr/%s", cidr), nil, &subnets)
	return subnets, err
}

// GetOverlapping returns all overlapping subnets for a given subnet
func (s *SubnetsService) GetOverlapping(cidr string) ([]Subnet, error) {
	var subnets []Subnet
	_, err := s.client.Request("GET", fmt.Sprintf("subnets/overlapping/%s", cidr), nil, &subnets)
	return subnets, err
}

// Create creates a new subnet
func (s *SubnetsService) Create(subnet *Subnet) (*Subnet, error) {
	var createdSubnet Subnet
	resp, err := s.client.Request("POST", "subnets", subnet, &createdSubnet)
	if err != nil {
		return nil, err
	}

	// If we got an ID in the response but not in the subnet data, retrieve the full subnet
	if resp.ID != 0 && createdSubnet.ID == "" {
		return s.Get(strconv.Itoa(resp.ID))
	}

	return &createdSubnet, nil
}

// CreateFirstSubnet creates a new child subnet inside a subnet with specified mask
func (s *SubnetsService) CreateFirstSubnet(id string, mask int, subnet *Subnet) (*Subnet, error) {
	var createdSubnet Subnet
	resp, err := s.client.Request("POST", fmt.Sprintf("subnets/%s/first_subnet/%d", id, mask), subnet, &createdSubnet)
	if err != nil {
		return nil, err
	}

	// If we got an ID in the response but not in the subnet data, retrieve the full subnet
	if resp.ID != 0 && createdSubnet.ID == "" {
		return s.Get(strconv.Itoa(resp.ID))
	}

	return &createdSubnet, nil
}

// Update updates an existing subnet
func (s *SubnetsService) Update(subnet *Subnet) (*Subnet, error) {
	if subnet.ID == "" {
		return nil, fmt.Errorf("subnet ID is required for update")
	}

	var updatedSubnet Subnet
	_, err := s.client.Request("PATCH", fmt.Sprintf("subnets/%s", subnet.ID), subnet, &updatedSubnet)
	return &updatedSubnet, err
}

// Resize resizes a subnet to a new mask
func (s *SubnetsService) Resize(id string, mask int) error {
	data := map[string]int{"mask": mask}
	_, err := s.client.Request("PATCH", fmt.Sprintf("subnets/%s/resize", id), data, nil)
	return err
}

// Split splits a subnet into smaller subnets
func (s *SubnetsService) Split(id string, subnets int) error {
	data := map[string]int{"number": subnets}
	_, err := s.client.Request("PATCH", fmt.Sprintf("subnets/%s/split", id), data, nil)
	return err
}

// SetPermissions sets subnet permissions
func (s *SubnetsService) SetPermissions(id string, permissions map[string]string) error {
	_, err := s.client.Request("PATCH", fmt.Sprintf("subnets/%s/permissions", id), permissions, nil)
	return err
}

// Delete deletes a subnet
func (s *SubnetsService) Delete(id string) error {
	_, err := s.client.Request("DELETE", fmt.Sprintf("subnets/%s", id), nil, nil)
	return err
}

// Truncate removes all addresses from a subnet
func (s *SubnetsService) Truncate(id string) error {
	_, err := s.client.Request("DELETE", fmt.Sprintf("subnets/%s/truncate", id), nil, nil)
	return err
}

// RemovePermissions removes all permissions from a subnet
func (s *SubnetsService) RemovePermissions(id string) error {
	_, err := s.client.Request("DELETE", fmt.Sprintf("subnets/%s/permissions", id), nil, nil)
	return err
}
