package phpipam

import (
	"fmt"
	"strconv"
)

// VRF represents a phpIPAM VRF object
type VRF struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	RD          string `json:"rd,omitempty"`
	Description string `json:"description,omitempty"`
	Sections    string `json:"sections,omitempty"`
	EditDate    string `json:"editDate,omitempty"`
}

// VRFsService handles communication with the VRF related methods of the API
type VRFsService struct {
	client *Client
}

// NewVRFsService creates a new VRFs service with the provided client
func NewVRFsService(client *Client) *VRFsService {
	return &VRFsService{client: client}
}

// List returns all VRFs
func (v *VRFsService) List() ([]VRF, error) {
	var vrfs []VRF
	_, err := v.client.Request("GET", "vrf", nil, &vrfs)
	return vrfs, err
}

// GetAll returns all VRFs (alias)
func (v *VRFsService) GetAll() ([]VRF, error) {
	var vrfs []VRF
	_, err := v.client.Request("GET", "vrf/all", nil, &vrfs)
	return vrfs, err
}

// Get returns a specific VRF by ID
func (v *VRFsService) Get(id string) (*VRF, error) {
	var vrf VRF
	_, err := v.client.Request("GET", fmt.Sprintf("vrf/%s", id), nil, &vrf)
	return &vrf, err
}

// GetSubnets returns all subnets within a VRF
func (v *VRFsService) GetSubnets(id string) ([]Subnet, error) {
	var subnets []Subnet
	_, err := v.client.Request("GET", fmt.Sprintf("vrf/%s/subnets", id), nil, &subnets)
	return subnets, err
}

// GetCustomFields returns all custom fields for VRFs
func (v *VRFsService) GetCustomFields() (map[string]CustomField, error) {
	var customFields map[string]CustomField
	_, err := v.client.Request("GET", "vrf/custom_fields", nil, &customFields)
	return customFields, err
}

// Create creates a new VRF
func (v *VRFsService) Create(vrf *VRF) (*VRF, error) {
	var createdVRF VRF
	resp, err := v.client.Request("POST", "vrf", vrf, &createdVRF)
	if err != nil {
		return nil, err
	}

	// If we got an ID in the response but not in the VRF data, retrieve the full VRF
	if resp.ID != 0 && createdVRF.ID == "" {
		return v.Get(strconv.Itoa(resp.ID))
	}

	return &createdVRF, nil
}

// Update updates a VRF
func (v *VRFsService) Update(vrf *VRF) (*VRF, error) {
	if vrf.ID == "" {
		return nil, fmt.Errorf("VRF ID is required for update")
	}

	var updatedVRF VRF
	_, err := v.client.Request("PATCH", "vrf", vrf, &updatedVRF)
	return &updatedVRF, err
}

// Delete deletes a VRF
func (v *VRFsService) Delete(id string) error {
	_, err := v.client.Request("DELETE", fmt.Sprintf("vrf/%s", id), nil, nil)
	return err
}
