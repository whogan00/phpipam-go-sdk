package phpipam

import (
	"fmt"
	"strconv"
)

// ToolsService handles communication with the tools related methods of the API
type ToolsService struct {
	client *Client
}

// NewToolsService creates a new tools service with the provided client
func NewToolsService(client *Client) *ToolsService {
	return &ToolsService{client: client}
}

// Subcontroller types for the Tools controller
const (
	SubcontrollerTags        = "tags"
	SubcontrollerDevices     = "devices"
	SubcontrollerDeviceTypes = "device_types"
	SubcontrollerVLANs       = "vlans"
	SubcontrollerVRFs        = "vrfs"
	SubcontrollerNameservers = "nameservers"
	SubcontrollerScanagents  = "scanagents"
	SubcontrollerLocations   = "locations"
	SubcontrollerNAT         = "nat"
	SubcontrollerRacks       = "racks"
)

// Tag represents a phpIPAM IP address tag
type IPTag struct {
	ID          string `json:"id,omitempty"`
	Type        string `json:"type,omitempty"`
	ShowTag     string `json:"showtag,omitempty"`
	BgColor     string `json:"bgcolor,omitempty"`
	FgColor     string `json:"fgcolor,omitempty"`
	DisplayName string `json:"displayname,omitempty"`
	Description string `json:"description,omitempty"`
}

// DeviceType represents a phpIPAM device type
type DeviceType struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// Nameserver represents a phpIPAM nameserver
type Nameserver struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Permissions string `json:"permissions,omitempty"`
	Namesrv1    string `json:"namesrv1,omitempty"`
	Namesrv2    string `json:"namesrv2,omitempty"`
	Namesrv3    string `json:"namesrv3,omitempty"`
}

// ScanAgent represents a phpIPAM scan agent
type ScanAgent struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Type        string `json:"type,omitempty"`
	Code        string `json:"code,omitempty"`
	Status      string `json:"status,omitempty"`
}

// Location represents a phpIPAM location
type Location struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Address     string `json:"address,omitempty"`
	Lat         string `json:"lat,omitempty"`
	Long        string `json:"long,omitempty"`
}

// NAT represents a phpIPAM NAT object
type NAT struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Type        string `json:"type,omitempty"`
	Device      string `json:"device,omitempty"`
	Src         string `json:"src,omitempty"`
	Dst         string `json:"dst,omitempty"`
	Description string `json:"description,omitempty"`
	Policy      string `json:"policy,omitempty"`
}

// Rack represents a phpIPAM rack
type Rack struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Location    string `json:"location,omitempty"`
	Size        string `json:"size,omitempty"`
	Description string `json:"description,omitempty"`
}

// GetIPTags returns all IP tags
func (t *ToolsService) GetIPTags() ([]IPTag, error) {
	var tags []IPTag
	_, err := t.client.Request("GET", "tools/tags", nil, &tags)
	return tags, err
}

// GetIPTag returns a specific IP tag by ID
func (t *ToolsService) GetIPTag(id string) (*IPTag, error) {
	var tag IPTag
	_, err := t.client.Request("GET", fmt.Sprintf("tools/tags/%s", id), nil, &tag)
	return &tag, err
}

// CreateIPTag creates a new IP tag
func (t *ToolsService) CreateIPTag(tag *IPTag) (*IPTag, error) {
	var createdTag IPTag
	resp, err := t.client.Request("POST", "tools/tags", tag, &createdTag)
	if err != nil {
		return nil, err
	}

	// If we got an ID in the response but not in the tag data, retrieve the full tag
	if resp.ID != 0 && createdTag.ID == "" {
		return t.GetIPTag(strconv.Itoa(resp.ID.Int()))
	}

	return &createdTag, nil
}

// UpdateIPTag updates an IP tag
func (t *ToolsService) UpdateIPTag(tag *IPTag) (*IPTag, error) {
	if tag.ID == "" {
		return nil, fmt.Errorf("tag ID is required for update")
	}

	var updatedTag IPTag
	_, err := t.client.Request("PATCH", fmt.Sprintf("tools/tags/%s", tag.ID), tag, &updatedTag)
	return &updatedTag, err
}

// DeleteIPTag deletes an IP tag
func (t *ToolsService) DeleteIPTag(id string) error {
	_, err := t.client.Request("DELETE", fmt.Sprintf("tools/tags/%s", id), nil, nil)
	return err
}

// GetDeviceTypes returns all device types
func (t *ToolsService) GetDeviceTypes() ([]DeviceType, error) {
	var deviceTypes []DeviceType
	_, err := t.client.Request("GET", "tools/device_types", nil, &deviceTypes)
	return deviceTypes, err
}

// GetDeviceType returns a specific device type by ID
func (t *ToolsService) GetDeviceType(id string) (*DeviceType, error) {
	var deviceType DeviceType
	_, err := t.client.Request("GET", fmt.Sprintf("tools/device_types/%s", id), nil, &deviceType)
	return &deviceType, err
}

// GetDevicesByType returns all devices belonging to device type
func (t *ToolsService) GetDevicesByType(id string) ([]Device, error) {
	var devices []Device
	_, err := t.client.Request("GET", fmt.Sprintf("tools/device_types/%s/devices", id), nil, &devices)
	return devices, err
}

// CreateDeviceType creates a new device type
func (t *ToolsService) CreateDeviceType(deviceType *DeviceType) (*DeviceType, error) {
	var createdDeviceType DeviceType
	resp, err := t.client.Request("POST", "tools/device_types", deviceType, &createdDeviceType)
	if err != nil {
		return nil, err
	}

	// If we got an ID in the response but not in the device type data, retrieve the full device type
	if resp.ID != 0 && createdDeviceType.ID == "" {
		return t.GetDeviceType(strconv.Itoa(resp.ID.Int()))
	}

	return &createdDeviceType, nil
}

// UpdateDeviceType updates a device type
func (t *ToolsService) UpdateDeviceType(deviceType *DeviceType) (*DeviceType, error) {
	if deviceType.ID == "" {
		return nil, fmt.Errorf("device type ID is required for update")
	}

	var updatedDeviceType DeviceType
	_, err := t.client.Request("PATCH", fmt.Sprintf("tools/device_types/%s", deviceType.ID), deviceType, &updatedDeviceType)
	return &updatedDeviceType, err
}

// DeleteDeviceType deletes a device type
func (t *ToolsService) DeleteDeviceType(id string) error {
	_, err := t.client.Request("DELETE", fmt.Sprintf("tools/device_types/%s", id), nil, nil)
	return err
}

// GetVLANsByToolsController returns all VLANs using tools controller
func (t *ToolsService) GetVLANsByToolsController() ([]VLAN, error) {
	var vlans []VLAN
	_, err := t.client.Request("GET", "tools/vlans", nil, &vlans)
	return vlans, err
}

// GetVLANByToolsController returns a specific VLAN by ID using tools controller
func (t *ToolsService) GetVLANByToolsController(id string) (*VLAN, error) {
	var vlan VLAN
	_, err := t.client.Request("GET", fmt.Sprintf("tools/vlans/%s", id), nil, &vlan)
	return &vlan, err
}

// GetSubnetsByVLAN returns all subnets belonging to VLAN
func (t *ToolsService) GetSubnetsByVLAN(id string) ([]Subnet, error) {
	var subnets []Subnet
	_, err := t.client.Request("GET", fmt.Sprintf("tools/vlans/%s/subnets", id), nil, &subnets)
	return subnets, err
}

// GetVRFsByToolsController returns all VRFs using tools controller
func (t *ToolsService) GetVRFsByToolsController() ([]VRF, error) {
	var vrfs []VRF
	_, err := t.client.Request("GET", "tools/vrfs", nil, &vrfs)
	return vrfs, err
}

// GetVRFByToolsController returns a specific VRF by ID using tools controller
func (t *ToolsService) GetVRFByToolsController(id string) (*VRF, error) {
	var vrf VRF
	_, err := t.client.Request("GET", fmt.Sprintf("tools/vrfs/%s", id), nil, &vrf)
	return &vrf, err
}

// GetSubnetsByVRF returns all subnets belonging to VRF
func (t *ToolsService) GetSubnetsByVRF(id string) ([]Subnet, error) {
	var subnets []Subnet
	_, err := t.client.Request("GET", fmt.Sprintf("tools/vrfs/%s/subnets", id), nil, &subnets)
	return subnets, err
}

// GetNameservers returns all nameservers
func (t *ToolsService) GetNameservers() ([]Nameserver, error) {
	var nameservers []Nameserver
	_, err := t.client.Request("GET", "tools/nameservers", nil, &nameservers)
	return nameservers, err
}

// GetNameserver returns a specific nameserver by ID
func (t *ToolsService) GetNameserver(id string) (*Nameserver, error) {
	var nameserver Nameserver
	_, err := t.client.Request("GET", fmt.Sprintf("tools/nameservers/%s", id), nil, &nameserver)
	return &nameserver, err
}

// CreateNameserver creates a new nameserver
func (t *ToolsService) CreateNameserver(nameserver *Nameserver) (*Nameserver, error) {
	var createdNameserver Nameserver
	resp, err := t.client.Request("POST", "tools/nameservers", nameserver, &createdNameserver)
	if err != nil {
		return nil, err
	}

	// If we got an ID in the response but not in the nameserver data, retrieve the full nameserver
	if resp.ID != 0 && createdNameserver.ID == "" {
		return t.GetNameserver(strconv.Itoa(resp.ID.Int()))
	}

	return &createdNameserver, nil
}

// UpdateNameserver updates a nameserver
func (t *ToolsService) UpdateNameserver(nameserver *Nameserver) (*Nameserver, error) {
	if nameserver.ID == "" {
		return nil, fmt.Errorf("nameserver ID is required for update")
	}

	var updatedNameserver Nameserver
	_, err := t.client.Request("PATCH", fmt.Sprintf("tools/nameservers/%s", nameserver.ID), nameserver, &updatedNameserver)
	return &updatedNameserver, err
}

// DeleteNameserver deletes a nameserver
func (t *ToolsService) DeleteNameserver(id string) error {
	_, err := t.client.Request("DELETE", fmt.Sprintf("tools/nameservers/%s", id), nil, nil)
	return err
}

// GetScanagents returns all scanagents
func (t *ToolsService) GetScanagents() ([]ScanAgent, error) {
	var scanagents []ScanAgent
	_, err := t.client.Request("GET", "tools/scanagents", nil, &scanagents)
	return scanagents, err
}

// GetScanagent returns a specific scanagent by ID
func (t *ToolsService) GetScanagent(id string) (*ScanAgent, error) {
	var scanagent ScanAgent
	_, err := t.client.Request("GET", fmt.Sprintf("tools/scanagents/%s", id), nil, &scanagent)
	return &scanagent, err
}

// GetLocations returns all locations
func (t *ToolsService) GetLocations() ([]Location, error) {
	var locations []Location
	_, err := t.client.Request("GET", "tools/locations", nil, &locations)
	return locations, err
}

// GetLocation returns a specific location by ID
func (t *ToolsService) GetLocation(id string) (*Location, error) {
	var location Location
	_, err := t.client.Request("GET", fmt.Sprintf("tools/locations/%s", id), nil, &location)
	return &location, err
}

// GetSubnetsByLocation returns all subnets belonging to a location
func (t *ToolsService) GetSubnetsByLocation(id string) ([]Subnet, error) {
	var subnets []Subnet
	_, err := t.client.Request("GET", fmt.Sprintf("tools/locations/%s/subnets", id), nil, &subnets)
	return subnets, err
}

// GetDevicesByLocation returns all devices belonging to a location
func (t *ToolsService) GetDevicesByLocation(id string) ([]Device, error) {
	var devices []Device
	_, err := t.client.Request("GET", fmt.Sprintf("tools/locations/%s/devices", id), nil, &devices)
	return devices, err
}

// GetRacksByLocation returns all racks belonging to a location
func (t *ToolsService) GetRacksByLocation(id string) ([]Rack, error) {
	var racks []Rack
	_, err := t.client.Request("GET", fmt.Sprintf("tools/locations/%s/racks", id), nil, &racks)
	return racks, err
}

// GetAddressesByLocation returns all IP addresses belonging to a location
func (t *ToolsService) GetAddressesByLocation(id string) ([]Address, error) {
	var addresses []Address
	_, err := t.client.Request("GET", fmt.Sprintf("tools/locations/%s/ipaddresses", id), nil, &addresses)
	return addresses, err
}

// CreateLocation creates a new location
func (t *ToolsService) CreateLocation(location *Location) (*Location, error) {
	var createdLocation Location
	resp, err := t.client.Request("POST", "tools/locations", location, &createdLocation)
	if err != nil {
		return nil, err
	}

	// If we got an ID in the response but not in the location data, retrieve the full location
	if resp.ID != 0 && createdLocation.ID == "" {
		return t.GetLocation(strconv.Itoa(resp.ID.Int()))
	}

	return &createdLocation, nil
}

// UpdateLocation updates a location
func (t *ToolsService) UpdateLocation(location *Location) (*Location, error) {
	if location.ID == "" {
		return nil, fmt.Errorf("location ID is required for update")
	}

	var updatedLocation Location
	_, err := t.client.Request("PATCH", fmt.Sprintf("tools/locations/%s", location.ID), location, &updatedLocation)
	return &updatedLocation, err
}

// DeleteLocation deletes a location
func (t *ToolsService) DeleteLocation(id string) error {
	_, err := t.client.Request("DELETE", fmt.Sprintf("tools/locations/%s", id), nil, nil)
	return err
}

// GetRacks returns all racks
func (t *ToolsService) GetRacks() ([]Rack, error) {
	var racks []Rack
	_, err := t.client.Request("GET", "tools/racks", nil, &racks)
	return racks, err
}

// GetRack returns a specific rack by ID
func (t *ToolsService) GetRack(id string) (*Rack, error) {
	var rack Rack
	_, err := t.client.Request("GET", fmt.Sprintf("tools/racks/%s", id), nil, &rack)
	return &rack, err
}

// GetDevicesByRack returns all devices belonging to rack
func (t *ToolsService) GetDevicesByRack(id string) ([]Device, error) {
	var devices []Device
	_, err := t.client.Request("GET", fmt.Sprintf("tools/racks/%s/devices", id), nil, &devices)
	return devices, err
}

// CreateRack creates a new rack
func (t *ToolsService) CreateRack(rack *Rack) (*Rack, error) {
	var createdRack Rack
	resp, err := t.client.Request("POST", "tools/racks", rack, &createdRack)
	if err != nil {
		return nil, err
	}

	// If we got an ID in the response but not in the rack data, retrieve the full rack
	if resp.ID != 0 && createdRack.ID == "" {
		return t.GetRack(strconv.Itoa(resp.ID.Int()))
	}

	return &createdRack, nil
}

// UpdateRack updates a rack
func (t *ToolsService) UpdateRack(rack *Rack) (*Rack, error) {
	if rack.ID == "" {
		return nil, fmt.Errorf("rack ID is required for update")
	}

	var updatedRack Rack
	_, err := t.client.Request("PATCH", fmt.Sprintf("tools/racks/%s", rack.ID), rack, &updatedRack)
	return &updatedRack, err
}

// DeleteRack deletes a rack
func (t *ToolsService) DeleteRack(id string) error {
	_, err := t.client.Request("DELETE", fmt.Sprintf("tools/racks/%s", id), nil, nil)
	return err
}

// GetNATs returns all NATs
func (t *ToolsService) GetNATs() ([]NAT, error) {
	var nats []NAT
	_, err := t.client.Request("GET", "tools/nat", nil, &nats)
	return nats, err
}

// GetNAT returns a specific NAT by ID
func (t *ToolsService) GetNAT(id string) (*NAT, error) {
	var nat NAT
	_, err := t.client.Request("GET", fmt.Sprintf("tools/nat/%s", id), nil, &nat)
	return &nat, err
}

// GetNATObjects returns all objects belonging to NAT
func (t *ToolsService) GetNATObjects(id string) ([]interface{}, error) {
	var objects []interface{}
	_, err := t.client.Request("GET", fmt.Sprintf("tools/nat/%s/objects", id), nil, &objects)
	return objects, err
}

// GetNATObjectsFull returns all objects with all parameters belonging to NAT
func (t *ToolsService) GetNATObjectsFull(id string) (map[string]interface{}, error) {
	var objects map[string]interface{}
	_, err := t.client.Request("GET", fmt.Sprintf("tools/nat/%s/objects_full", id), nil, &objects)
	return objects, err
}

// CreateNAT creates a new NAT
func (t *ToolsService) CreateNAT(nat *NAT) (*NAT, error) {
	var createdNAT NAT
	resp, err := t.client.Request("POST", "tools/nat", nat, &createdNAT)
	if err != nil {
		return nil, err
	}

	// If we got an ID in the response but not in the NAT data, retrieve the full NAT
	if resp.ID != 0 && createdNAT.ID == "" {
		return t.GetNAT(strconv.Itoa(resp.ID.Int()))
	}

	return &createdNAT, nil
}

// UpdateNAT updates a NAT
func (t *ToolsService) UpdateNAT(nat *NAT) (*NAT, error) {
	if nat.ID == "" {
		return nil, fmt.Errorf("NAT ID is required for update")
	}

	var updatedNAT NAT
	_, err := t.client.Request("PATCH", fmt.Sprintf("tools/nat/%s", nat.ID), nat, &updatedNAT)
	return &updatedNAT, err
}

// DeleteNAT deletes a NAT
func (t *ToolsService) DeleteNAT(id string) error {
	_, err := t.client.Request("DELETE", fmt.Sprintf("tools/nat/%s", id), nil, nil)
	return err
}
