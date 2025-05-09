package phpipam

import (
	"crypto/tls"
	"net/http"
	"time"
)

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
func New(baseURL, appID, username, password string, insecureTLS bool) (*PHPIPAM, error) {
	client, err := NewClient(baseURL, appID, username, password, insecureTLS)
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
// This method is useful when you have a static API key configured in phpIP
func NewTokenClient(baseURL, appID, token string, insecureTLS bool) (*PHPIPAM, error) {
	// Create a base client without username/password
	client, err := NewClient(baseURL, appID, "", "", insecureTLS)
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

// SetInsecureTLS configures the client to skip TLS certificate verification
func (c *Client) SetInsecureTLS(insecure bool) {
	c.InsecureTLS = insecure

	// Update the transport
	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.TLSClientConfig = &tls.Config{
		InsecureSkipVerify: insecure,
	}

	c.HTTPClient.Transport = transport
}

// Authenticate authenticates with the phpIPAM API
func (p *PHPIPAM) Authenticate() error {
	return p.Client.Authenticate()
}
