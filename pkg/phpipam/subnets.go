package phpipam

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// Subnet represents a phpIPAM subnet object with fields matching API response
type Subnet struct {
	ID                    int             `json:"id,omitempty"`
	Subnet                string          `json:"subnet,omitempty"`
	Mask                  string          `json:"mask,omitempty"`
	SectionID             int             `json:"sectionId,omitempty"`
	Description           string          `json:"description,omitempty"`
	LinkedSubnet          interface{}     `json:"linked_subnet,omitempty"`
	FirewallAddressObject interface{}     `json:"firewallAddressObject,omitempty"`
	VrfID                 interface{}     `json:"vrfId,omitempty"`
	MasterSubnetID        int             `json:"masterSubnetId,omitempty"`
	AllowRequests         int             `json:"allowRequests,omitempty"`
	VlanID                interface{}     `json:"vlanId,omitempty"`
	ShowName              int             `json:"showName,omitempty"`
	Device                interface{}     `json:"device,omitempty"`
	Permissions           json.RawMessage `json:"permissions,omitempty"` // Using raw json.RawMessage directly
	PingSubnet            int             `json:"pingSubnet,omitempty"`
	DiscoverSubnet        int             `json:"discoverSubnet,omitempty"`
	ResolveDNS            int             `json:"resolveDNS,omitempty"`
	DNSRecursive          int             `json:"DNSrecursive,omitempty"`
	DNSRecords            int             `json:"DNSrecords,omitempty"`
	NameserverID          int             `json:"nameserverId,omitempty"`
	ScanAgent             int             `json:"scanAgent,omitempty"`
	CustomerID            interface{}     `json:"customer_id,omitempty"`
	IsFolder              int             `json:"isFolder,omitempty"`
	IsFull                int             `json:"isFull,omitempty"`
	IsPool                int             `json:"isPool,omitempty"`
	Tag                   int             `json:"tag,omitempty"`
	Threshold             int             `json:"threshold,omitempty"`
	Location              interface{}     `json:"location,omitempty"`
	EditDate              interface{}     `json:"editDate,omitempty"`
	LastScan              interface{}     `json:"lastScan,omitempty"`
	LastDiscovery         interface{}     `json:"lastDiscovery,omitempty"`
	Calculation           interface{}     `json:"calculation,omitempty"`
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

// Helper methods for working with permissions

// SetPermissionsString sets permissions using a JSON string
func (s *Subnet) SetPermissionsString(permissions string) {
	// Check if the string is already JSON formatted
	if len(permissions) > 0 && permissions[0] == '{' {
		// It's already a JSON object, no need for quotes
		s.Permissions = json.RawMessage(permissions)
	} else {
		// For simple quoted string like "{\"3\":\"1\",\"2\":\"2\"}"
		// We need to ensure it's treated as a JSON string by wrapping in quotes
		jsonStr := fmt.Sprintf("%q", permissions)
		s.Permissions = json.RawMessage(jsonStr)
	}
}

// SetPermissionsObject sets permissions using a map or array
func (s *Subnet) SetPermissionsObject(permissions interface{}) error {
	bytes, err := json.Marshal(permissions)
	if err != nil {
		return err
	}
	s.Permissions = bytes
	return nil
}

// GetPermissionsAsString attempts to get permissions as a string
func (s *Subnet) GetPermissionsAsString() (string, error) {
	if s.Permissions == nil {
		return "", nil
	}

	// Try to unmarshal as a string first
	var strVal string
	err := json.Unmarshal(s.Permissions, &strVal)
	if err == nil {
		return strVal, nil
	}

	// If that fails, it might be a map or array that we need to stringify
	var objVal interface{}
	err = json.Unmarshal(s.Permissions, &objVal)
	if err != nil {
		return "", fmt.Errorf("failed to parse permissions: %w", err)
	}

	// Convert the object to a JSON string
	jsonBytes, err := json.Marshal(objVal)
	if err != nil {
		return "", fmt.Errorf("failed to stringify permissions: %w", err)
	}

	return string(jsonBytes), nil
}

// GetPermissionsAsMap attempts to get permissions as a map
func (s *Subnet) GetPermissionsAsMap() (map[string]string, error) {
	if s.Permissions == nil {
		return map[string]string{}, nil
	}

	// Try direct unmarshaling to map first
	var mapVal map[string]string
	err := json.Unmarshal(s.Permissions, &mapVal)
	if err == nil {
		return mapVal, nil
	}

	// If that fails, it might be a JSON string that contains a map
	var strVal string
	err = json.Unmarshal(s.Permissions, &strVal)
	if err == nil {
		// Try to unmarshal the string content as a map
		err = json.Unmarshal([]byte(strVal), &mapVal)
		if err == nil {
			return mapVal, nil
		}
	}

	return nil, fmt.Errorf("permissions not in expected format")
}

// Helper methods for nullable fields

// GetVrfID returns the VRF ID as an integer if not null
func (s *Subnet) GetVrfID() (int, bool) {
	if s.VrfID == nil {
		return 0, false
	}
	switch v := s.VrfID.(type) {
	case float64:
		return int(v), true
	case int:
		return v, true
	case string:
		// Try to parse as int
		intVal, err := strconv.Atoi(v)
		if err == nil {
			return intVal, true
		}
	}
	return 0, false
}

// GetVlanID returns the VLAN ID as an integer if not null
func (s *Subnet) GetVlanID() (int, bool) {
	if s.VlanID == nil {
		return 0, false
	}
	switch v := s.VlanID.(type) {
	case float64:
		return int(v), true
	case int:
		return v, true
	case string:
		// Try to parse as int
		intVal, err := strconv.Atoi(v)
		if err == nil {
			return intVal, true
		}
	}
	return 0, false
}

// GetLocationID returns the Location ID as an integer if not null
func (s *Subnet) GetLocationID() (int, bool) {
	if s.Location == nil {
		return 0, false
	}
	switch v := s.Location.(type) {
	case float64:
		return int(v), true
	case int:
		return v, true
	case string:
		// Try to parse as int
		intVal, err := strconv.Atoi(v)
		if err == nil {
			return intVal, true
		}
	}
	return 0, false
}

// SetVrfID sets the VRF ID (handles nil case)
func (s *Subnet) SetVrfID(id int) {
	if id == 0 {
		s.VrfID = nil
	} else {
		s.VrfID = id
	}
}

// SetVlanID sets the VLAN ID (handles nil case)
func (s *Subnet) SetVlanID(id int) {
	if id == 0 {
		s.VlanID = nil
	} else {
		s.VlanID = id
	}
}

// SetLocationID sets the Location ID (handles nil case)
func (s *Subnet) SetLocationID(id int) {
	if id == 0 {
		s.Location = nil
	} else {
		s.Location = id
	}
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
