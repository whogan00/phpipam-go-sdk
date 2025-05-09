package phpipam

import "time"

// PHPIPAM represents the main PHPIPAM client that encapsulates all service clients
type PHPIPAM struct {
	Client    *Client
	Sections  *SectionsService
	Subnets   *SubnetsService
	Addresses *AddressesService
	VLANs     *VLANsService
	L2Domains *L2DomainsService
	VRFs      *VRFsService
	Devices   *DevicesService
	Tools     *ToolsService
	Prefix    *PrefixService
	Search    *SearchService
}

// New creates a new PHPIPAM client with all services
func New(baseURL, appID, username, password string) (*PHPIPAM, error) {
	client, err := NewClient(baseURL, appID, username, password)
	if err != nil {
		return nil, err
	}

	return &PHPIPAM{
		Client:    client,
		Sections:  NewSectionsService(client),
		Subnets:   NewSubnetsService(client),
		Addresses: NewAddressesService(client),
		VLANs:     NewVLANsService(client),
		L2Domains: NewL2DomainsService(client),
		VRFs:      NewVRFsService(client),
		Devices:   NewDevicesService(client),
		Tools:     NewToolsService(client),
		Prefix:    NewPrefixService(client),
		Search:    NewSearchService(client),
	}, nil
}

// NewTokenClient creates a new phpIPAM API client using App Code (token) authentication
// This method is useful when you have a static API key configured in phpIPAM
func NewTokenClient(baseURL, appID, token string) (*PHPIPAM, error) {
	// Create a base client without username/password
	client, err := NewClient(baseURL, appID, "", "")
	if err != nil {
		return nil, err
	}

	// Set the token directly instead of obtaining it via authentication
	client.Token = token

	// Set a far future expiration time since API keys don't typically expire
	// If you're using phpIPAM with token expiration enabled, you might need to adjust this
	client.TokenExp = time.Now().AddDate(1, 0, 0) // 1 year in the future

	// Create a new PHPIPAM client with all services
	return &PHPIPAM{
		Client:    client,
		Sections:  NewSectionsService(client),
		Subnets:   NewSubnetsService(client),
		Addresses: NewAddressesService(client),
		VLANs:     NewVLANsService(client),
		L2Domains: NewL2DomainsService(client),
		VRFs:      NewVRFsService(client),
		Devices:   NewDevicesService(client),
		Tools:     NewToolsService(client),
		Prefix:    NewPrefixService(client),
		Search:    NewSearchService(client),
	}, nil
}

// Authenticate authenticates with the phpIPAM API
func (p *PHPIPAM) Authenticate() error {
	return p.Client.Authenticate()
}
