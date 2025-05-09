package phpipam

// PHPIPAM represents the main PHPIPAM client that encapsulates all service clients
type PHPIPAM struct {
	Client            *Client
	Sections          *SectionsService
	Subnets           *SubnetsService
	Addresses         *AddressesService
	VLANs             *VLANsService
	L2Domains         *L2DomainsService
	VRFs              *VRFsService
	Devices           *DevicesService
	Tools             *ToolsService
	Prefix            *PrefixService
	Search            *SearchService
}

// New creates a new PHPIPAM client with all services
func New(baseURL, appID, username, password string) (*PHPIPAM, error) {
	client, err := NewClient(baseURL, appID, username, password)
	if err != nil {
		return nil, err
	}

	return &PHPIPAM{
		Client:            client,
		Sections:          NewSectionsService(client),
		Subnets:           NewSubnetsService(client),
		Addresses:         NewAddressesService(client),
		VLANs:             NewVLANsService(client),
		L2Domains:         NewL2DomainsService(client),
		VRFs:              NewVRFsService(client),
		Devices:           NewDevicesService(client),
		Tools:             NewToolsService(client),
		Prefix:            NewPrefixService(client),
		Search:            NewSearchService(client),
	}, nil
}

// Authenticate authenticates with the phpIPAM API
func (p *PHPIPAM) Authenticate() error {
	return p.Client.Authenticate()
}
