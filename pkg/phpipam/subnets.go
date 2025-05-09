package phpipam

import (
	"fmt"
	"strconv"
)

// Subnet represents a phpIPAM subnet object
type Subnet struct {
	ID             int      `json:"id,omitempty"`
	Subnet         string   `json:"subnet"`
	Mask           string   `json:"mask"`
	SectionID      int      `json:"sectionId,omitempty"`
	Description    string   `json:"description,omitempty"`
	LinkedSubnet   int      `json:"linked_subnet,omitempty"`
	VlanID         int      `json:"vlanId,omitempty"`
	VrfID          int      `json:"vrfId,omitempty"`
	MasterSubnetID int      `json:"masterSubnetId,omitempty"`
	NameserverID   int      `json:"nameserverId,omitempty"`
	ShowName       int      `json:"showName,omitempty"`
	Permissions    []string `json:"permissions,omitempty"`
	DNSRecursive   int      `json:"DNSrecursive,omitempty"`
	DNSRecords     int      `json:"DNSrecords,omitempty"`
	AllowRequests  int      `json:"allowRequests,omitempty"`
	ScanAgent      int      `json:"scanAgent,omitempty"`
	PingSubnet     int      `json:"pingSubnet,omitempty"`
	DiscoverSubnet int      `json:"discoverSubnet,omitempty"`
	IsFolder       int      `json:"isFolder,omitempty"`
	IsFull         int      `json:"isFull,omitempty"`
	State          int      `json:"state,omitempty"`
	Threshold      int      `json:"threshold,omitempty"`
	Location       int      `json:"location,omitempty"`
	EditDate       string   `json:"editDate,omitempty"`
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
func (s *SubnetsService) Get(id int) (*Subnet, error) {
	var subnet Subnet
	_, err := s.client.Request("GET", fmt.Sprintf("subnets/%d", id), nil, &subnet)
	return &subnet, err
}

// GetByStringID returns a specific subnet by string ID (convenience method)
func (s *SubnetsService) GetByStringID(id string) (*Subnet, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return nil, fmt.Errorf("invalid subnet ID: %w", err)
	}
	return s.Get(idInt)
}

// GetUsage returns usage statistics for a subnet
func (s *SubnetsService) GetUsage(id int) (*SubnetUsage, error) {
	var usage SubnetUsage
	_, err := s.client.Request("GET", fmt.Sprintf("subnets/%d/usage", id), nil, &usage)
	return &usage, err
}

// GetSlaves returns all immediate slave subnets
func (s *SubnetsService) GetSlaves(id int) ([]Subnet, error) {
	var subnets []Subnet
	_, err := s.client.Request("GET", fmt.Sprintf("subnets/%d/slaves", id), nil, &subnets)
	return subnets, err
}

// GetSlavesRecursive returns all slave subnets recursively
func (s *SubnetsService) GetSlavesRecursive(id int) ([]Subnet, error) {
	var subnets []Subnet
	_, err := s.client.Request("GET", fmt.Sprintf("subnets/%d/slaves_recursive", id), nil, &subnets)
	return subnets, err
}

// GetAddresses returns all addresses in a subnet
func (s *SubnetsService) GetAddresses(id int) ([]Address, error) {
	var addresses []Address
	_, err := s.client.Request("GET", fmt.Sprintf("subnets/%d/addresses", id), nil, &addresses)
	return addresses, err
}

// GetAddress returns a specific IP address from a subnet
func (s *SubnetsService) GetAddress(id int, ip string) (*Address, error) {
	var address Address
	_, err := s.client.Request("GET", fmt.Sprintf("subnets/%d/addresses/%s", id, ip), nil, &address)
	return &address, err
}

// GetFirstFree returns the first available IP address in a subnet
func (s *SubnetsService) GetFirstFree(id int) (string, error) {
	var firstFree string
	_, err := s.client.Request("GET", fmt.Sprintf("subnets/%d/first_free", id), nil, &firstFree)
	return firstFree, err
}

// GetFirstSubnet returns the first available subnet within a given subnet for specified mask
func (s *SubnetsService) GetFirstSubnet(id int, mask int) (string, error) {
	var firstSubnet string
	_, err := s.client.Request("GET", fmt.Sprintf("subnets/%d/first_subnet/%d", id, mask), nil, &firstSubnet)
	return firstSubnet, err
}

// GetLastSubnet returns the last available subnet within a given subnet for specified mask
func (s *SubnetsService) GetLastSubnet(id int, mask int) (string, error) {
	var lastSubnet string
	_, err := s.client.Request("GET", fmt.Sprintf("subnets/%d/last_subnet/%d", id, mask), nil, &lastSubnet)
	return lastSubnet, err
}

// GetAllSubnets returns all available subnets within a given subnet for specified mask
func (s *SubnetsService) GetAllSubnets(id int, mask int) ([]string, error) {
	var allSubnets []string
	_, err := s.client.Request("GET", fmt.Sprintf("subnets/%d/all_subnets/%d", id, mask), nil, &allSubnets)
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
	if resp.ID != 0 && createdSubnet.ID == 0 {
		return s.Get(resp.ID)
	}

	return &createdSubnet, nil
}

// CreateFirstSubnet creates a new child subnet inside a subnet with specified mask
func (s *SubnetsService) CreateFirstSubnet(id int, mask int, subnet *Subnet) (*Subnet, error) {
	var createdSubnet Subnet
	resp, err := s.client.Request("POST", fmt.Sprintf("subnets/%d/first_subnet/%d", id, mask), subnet, &createdSubnet)
	if err != nil {
		return nil, err
	}

	// If we got an ID in the response but not in the subnet data, retrieve the full subnet
	if resp.ID != 0 && createdSubnet.ID == 0 {
		return s.Get(resp.ID)
	}

	return &createdSubnet, nil
}

// Update updates an existing subnet
func (s *SubnetsService) Update(subnet *Subnet) (*Subnet, error) {
	if subnet.ID == 0 {
		return nil, fmt.Errorf("subnet ID is required for update")
	}

	var updatedSubnet Subnet
	_, err := s.client.Request("PATCH", fmt.Sprintf("subnets/%d", subnet.ID), subnet, &updatedSubnet)
	return &updatedSubnet, err
}

// Resize resizes a subnet to a new mask
func (s *SubnetsService) Resize(id int, mask int) error {
	data := map[string]int{"mask": mask}
	_, err := s.client.Request("PATCH", fmt.Sprintf("subnets/%d/resize", id), data, nil)
	return err
}

// Split splits a subnet into smaller subnets
func (s *SubnetsService) Split(id int, subnets int) error {
	data := map[string]int{"number": subnets}
	_, err := s.client.Request("PATCH", fmt.Sprintf("subnets/%d/split", id), data, nil)
	return err
}

// SetPermissions sets subnet permissions
func (s *SubnetsService) SetPermissions(id int, permissions map[string]string) error {
	_, err := s.client.Request("PATCH", fmt.Sprintf("subnets/%d/permissions", id), permissions, nil)
	return err
}

// Delete deletes a subnet
func (s *SubnetsService) Delete(id int) error {
	_, err := s.client.Request("DELETE", fmt.Sprintf("subnets/%d", id), nil, nil)
	return err
}

// Truncate removes all addresses from a subnet
func (s *SubnetsService) Truncate(id int) error {
	_, err := s.client.Request("DELETE", fmt.Sprintf("subnets/%d/truncate", id), nil, nil)
	return err
}

// RemovePermissions removes all permissions from a subnet
func (s *SubnetsService) RemovePermissions(id int) error {
	_, err := s.client.Request("DELETE", fmt.Sprintf("subnets/%d/permissions", id), nil, nil)
	return err
}
