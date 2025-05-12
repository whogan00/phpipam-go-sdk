package phpipam

import (
	"fmt"
	"strconv"
)

// Section represents a phpIPAM section object
type Section struct {
	ID               string `json:"id,omitempty"`
	Name             string `json:"name"`
	Description      string `json:"description,omitempty"`
	MasterSection    int    `json:"masterSection,omitempty"`
	Permissions      string `json:"permissions,omitempty"`
	StrictMode       int    `json:"strictMode,omitempty"`
	SubnetOrdering   string `json:"subnetOrdering,omitempty"`
	Order            int    `json:"order,omitempty"`
	EditDate         string `json:"editDate,omitempty"`
	ShowVLAN         int    `json:"showVLAN,omitempty"`
	ShowVRF          int    `json:"showVRF,omitempty"`
	ShowSupernetOnly int    `json:"showSupernetOnly,omitempty"`
	DNS              string `json:"DNS,omitempty"`
}

// CustomField represents a custom field definition
type CustomField struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Comment     string `json:"comment"`
	Default     string `json:"default"`
	Required    int    `json:"required"`
	Permissions string `json:"permissions"`
}

// SectionsService handles communication with the sections related methods of the API
type SectionsService struct {
	client *Client
}

// NewSectionsService creates a new sections service with the provided client
func NewSectionsService(client *Client) *SectionsService {
	return &SectionsService{client: client}
}

// List returns all sections
func (s *SectionsService) List() ([]Section, error) {
	var sections []Section
	_, err := s.client.Request("GET", "sections", nil, &sections)
	return sections, err
}

// Get returns a specific section by ID
func (s *SectionsService) Get(id string) (*Section, error) {
	var section Section
	_, err := s.client.Request("GET", fmt.Sprintf("sections/%s", id), nil, &section)
	return &section, err
}

// GetByName returns a specific section by name
func (s *SectionsService) GetByName(name string) (*Section, error) {
	var section Section
	_, err := s.client.Request("GET", fmt.Sprintf("sections/%s", name), nil, &section)
	return &section, err
}

// Create creates a new section
func (s *SectionsService) Create(section *Section) (*Section, error) {
	var createdSection Section
	resp, err := s.client.Request("POST", "sections", section, &createdSection)
	if err != nil {
		return nil, err
	}

	// If we got an ID in the response but not in the section data, retrieve the full section
	if resp.ID != 0 && createdSection.ID == "" {
		return s.Get(strconv.Itoa(resp.ID.Int()))
	}

	return &createdSection, nil
}

// Update updates an existing section
func (s *SectionsService) Update(section *Section) (*Section, error) {
	if section.ID == "" {
		return nil, fmt.Errorf("section ID is required for update")
	}

	var updatedSection Section
	_, err := s.client.Request("PATCH", fmt.Sprintf("sections/%s", section.ID), section, &updatedSection)
	return &updatedSection, err
}

// Delete deletes a section
func (s *SectionsService) Delete(id string) error {
	_, err := s.client.Request("DELETE", fmt.Sprintf("sections/%s", id), nil, nil)
	return err
}

// GetSubnets returns all subnets in a section
func (s *SectionsService) GetSubnets(id string) ([]Subnet, error) {
	var subnets []Subnet
	_, err := s.client.Request("GET", fmt.Sprintf("sections/%s/subnets", id), nil, &subnets)
	return subnets, err
}

// GetSubnetAddresses returns all subnets with addresses in a section
func (s *SectionsService) GetSubnetAddresses(id string) ([]Subnet, error) {
	var subnets []Subnet
	_, err := s.client.Request("GET", fmt.Sprintf("sections/%s/subnets/addresses", id), nil, &subnets)
	return subnets, err
}

// GetCustomFields returns custom section fields
func (s *SectionsService) GetCustomFields() (map[string]CustomField, error) {
	var customFields map[string]CustomField
	_, err := s.client.Request("GET", "sections/custom_fields", nil, &customFields)
	return customFields, err
}
