package phpipam

import (
	"fmt"
	"strconv"
)

// VLAN represents a phpIPAM VLAN object
type VLAN struct {
	ID          string `json:"id,omitempty"`
	DomainID    string `json:"domainId,omitempty"`
	Name        string `json:"name,omitempty"`
	Number      string `json:"number,omitempty"`
	Description string `json:"description,omitempty"`
	EditDate    string `json:"editDate,omitempty"`
}

// VLANsService handles communication with the VLAN related methods of the API
type VLANsService struct {
	client *Client
}

// NewVLANsService creates a new VLANs service with the provided client
func NewVLANsService(client *Client) *VLANsService {
	return &VLANsService{client: client}
}

// List returns all VLANs
func (v *VLANsService) List() ([]VLAN, error) {
	var vlans []VLAN
	_, err := v.client.Request("GET", "vlan", nil, &vlans)
	return vlans, err
}

// GetAll returns all VLANs (alias)
func (v *VLANsService) GetAll() ([]VLAN, error) {
	var vlans []VLAN
	_, err := v.client.Request("GET", "vlan/all", nil, &vlans)
	return vlans, err
}

// Get returns a specific VLAN by ID
func (v *VLANsService) Get(id string) (*VLAN, error) {
	var vlan VLAN
	_, err := v.client.Request("GET", fmt.Sprintf("vlan/%s", id), nil, &vlan)
	return &vlan, err
}

// GetSubnets returns all subnets attached to a VLAN
func (v *VLANsService) GetSubnets(id string) ([]Subnet, error) {
	var subnets []Subnet
	_, err := v.client.Request("GET", fmt.Sprintf("vlan/%s/subnets", id), nil, &subnets)
	return subnets, err
}

// GetSubnetsInSection returns all subnets attached to a VLAN in a specific section
func (v *VLANsService) GetSubnetsInSection(id, sectionID string) ([]Subnet, error) {
	var subnets []Subnet
	_, err := v.client.Request("GET", fmt.Sprintf("vlan/%s/subnets/%s", id, sectionID), nil, &subnets)
	return subnets, err
}

// GetCustomFields returns custom VLAN fields
func (v *VLANsService) GetCustomFields() (map[string]CustomField, error) {
	var customFields map[string]CustomField
	_, err := v.client.Request("GET", "vlan/custom_fields", nil, &customFields)
	return customFields, err
}

// Search searches for a VLAN by number
func (v *VLANsService) Search(number string) ([]VLAN, error) {
	var vlans []VLAN
	_, err := v.client.Request("GET", fmt.Sprintf("vlan/search/%s", number), nil, &vlans)
	return vlans, err
}

// Create creates a new VLAN
func (v *VLANsService) Create(vlan *VLAN) (*VLAN, error) {
	var createdVLAN VLAN
	resp, err := v.client.Request("POST", "vlan", vlan, &createdVLAN)
	if err != nil {
		return nil, err
	}

	// If we got an ID in the response but not in the VLAN data, retrieve the full VLAN
	if resp.ID != 0 && createdVLAN.ID == "" {
		return v.Get(strconv.Itoa(resp.ID))
	}

	return &createdVLAN, nil
}

// Update updates a VLAN
func (v *VLANsService) Update(vlan *VLAN) (*VLAN, error) {
	if vlan.ID == "" {
		return nil, fmt.Errorf("VLAN ID is required for update")
	}

	var updatedVLAN VLAN
	_, err := v.client.Request("PATCH", "vlan", vlan, &updatedVLAN)
	return &updatedVLAN, err
}

// Delete deletes a VLAN
func (v *VLANsService) Delete(id string) error {
	_, err := v.client.Request("DELETE", fmt.Sprintf("vlan/%s", id), nil, nil)
	return err
}
