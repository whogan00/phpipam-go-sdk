package phpipam

import (
	"fmt"
	"strconv"
)

// L2Domain represents a phpIPAM VLAN domain (L2 domain) object
type L2Domain struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Permissions string `json:"permissions,omitempty"`
}

// L2DomainsService handles communication with the L2 domains related methods of the API
type L2DomainsService struct {
	client *Client
}

// NewL2DomainsService creates a new L2 domains service with the provided client
func NewL2DomainsService(client *Client) *L2DomainsService {
	return &L2DomainsService{client: client}
}

// List returns all L2 domains
func (l *L2DomainsService) List() ([]L2Domain, error) {
	var domains []L2Domain
	_, err := l.client.Request("GET", "l2domains", nil, &domains)
	return domains, err
}

// GetAll returns all L2 domains (alias)
func (l *L2DomainsService) GetAll() ([]L2Domain, error) {
	var domains []L2Domain
	_, err := l.client.Request("GET", "l2domains/all", nil, &domains)
	return domains, err
}

// Get returns a specific L2 domain by ID
func (l *L2DomainsService) Get(id string) (*L2Domain, error) {
	var domain L2Domain
	_, err := l.client.Request("GET", fmt.Sprintf("l2domains/%s", id), nil, &domain)
	return &domain, err
}

// GetVLANs returns all VLANs within a L2 domain
func (l *L2DomainsService) GetVLANs(id string) ([]VLAN, error) {
	var vlans []VLAN
	_, err := l.client.Request("GET", fmt.Sprintf("l2domains/%s/vlans", id), nil, &vlans)
	return vlans, err
}

// GetCustomFields returns all custom fields for L2 domains
func (l *L2DomainsService) GetCustomFields() (map[string]CustomField, error) {
	var customFields map[string]CustomField
	_, err := l.client.Request("GET", "l2domains/custom_fields", nil, &customFields)
	return customFields, err
}

// Create creates a new L2 domain
func (l *L2DomainsService) Create(domain *L2Domain) (*L2Domain, error) {
	var createdDomain L2Domain
	resp, err := l.client.Request("POST", "l2domains", domain, &createdDomain)
	if err != nil {
		return nil, err
	}

	// If we got an ID in the response but not in the domain data, retrieve the full domain
	if resp.ID != 0 && createdDomain.ID == "" {
		return l.Get(strconv.Itoa(resp.ID.Int()))
	}

	return &createdDomain, nil
}

// Update updates a L2 domain
func (l *L2DomainsService) Update(domain *L2Domain) (*L2Domain, error) {
	if domain.ID == "" {
		return nil, fmt.Errorf("L2 domain ID is required for update")
	}

	var updatedDomain L2Domain
	_, err := l.client.Request("PATCH", "l2domains", domain, &updatedDomain)
	return &updatedDomain, err
}

// Delete deletes a L2 domain
func (l *L2DomainsService) Delete(id string) error {
	_, err := l.client.Request("DELETE", fmt.Sprintf("l2domains/%s", id), nil, nil)
	return err
}
