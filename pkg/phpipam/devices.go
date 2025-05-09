package phpipam

import (
	"fmt"
	"strconv"
	"net/url"
)

// Device represents a phpIPAM device object
type Device struct {
	ID          string `json:"id,omitempty"`
	Hostname    string `json:"hostname,omitempty"`
	IPAddr      string `json:"ip_addr,omitempty"`
	Description string `json:"description,omitempty"`
	Sections    string `json:"sections,omitempty"`
	Rack        string `json:"rack,omitempty"`
	RackStart   string `json:"rack_start,omitempty"`
	RackSize    string `json:"rack_size,omitempty"`
	Location    string `json:"location,omitempty"`
	EditDate    string `json:"editDate,omitempty"`
}

// DevicesService handles communication with the devices related methods of the API
type DevicesService struct {
	client *Client
}

// NewDevicesService creates a new devices service with the provided client
func NewDevicesService(client *Client) *DevicesService {
	return &DevicesService{client: client}
}

// List returns all devices
func (d *DevicesService) List() ([]Device, error) {
	var devices []Device
	_, err := d.client.Request("GET", "devices", nil, &devices)
	return devices, err
}

// GetAll returns all devices (alias)
func (d *DevicesService) GetAll() ([]Device, error) {
	var devices []Device
	_, err := d.client.Request("GET", "devices/all", nil, &devices)
	return devices, err
}

// Get returns a specific device by ID
func (d *DevicesService) Get(id string) (*Device, error) {
	var device Device
	_, err := d.client.Request("GET", fmt.Sprintf("devices/%s", id), nil, &device)
	return &device, err
}

// GetSubnets returns all subnets within a device
func (d *DevicesService) GetSubnets(id string) ([]Subnet, error) {
	var subnets []Subnet
	_, err := d.client.Request("GET", fmt.Sprintf("devices/%s/subnets", id), nil, &subnets)
	return subnets, err
}

// GetAddresses returns all addresses within a device
func (d *DevicesService) GetAddresses(id string) ([]Address, error) {
	var addresses []Address
	_, err := d.client.Request("GET", fmt.Sprintf("devices/%s/addresses", id), nil, &addresses)
	return addresses, err
}

// Search searches for devices with search_string in any belonging field
func (d *DevicesService) Search(searchString string) ([]Device, error) {
	var devices []Device
	_, err := d.client.Request("GET", fmt.Sprintf("devices/search/%s", url.QueryEscape(searchString)), nil, &devices)
	return devices, err
}

// Create creates a new device
func (d *DevicesService) Create(device *Device) (*Device, error) {
	var createdDevice Device
	resp, err := d.client.Request("POST", "devices", device, &createdDevice)
	if err != nil {
		return nil, err
	}

	// If we got an ID in the response but not in the device data, retrieve the full device
	if resp.ID != 0 && createdDevice.ID == "" {
		return d.Get(strconv.Itoa(resp.ID))
	}

	return &createdDevice, nil
}

// Update updates a device
func (d *DevicesService) Update(device *Device) (*Device, error) {
	if device.ID == "" {
		return nil, fmt.Errorf("device ID is required for update")
	}

	var updatedDevice Device
	_, err := d.client.Request("PATCH", "devices", device, &updatedDevice)
	return &updatedDevice, err
}

// Delete deletes a device
func (d *DevicesService) Delete(id string) error {
	_, err := d.client.Request("DELETE", fmt.Sprintf("devices/%s", id), nil, nil)
	return err
}
