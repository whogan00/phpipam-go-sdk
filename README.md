# phpipam-go-sdk

A Go client library for the phpIPAM API.

## Features

- Complete implementation of phpIPAM API controllers
- Authentication handling with automatic token renewal
- Strongly typed models for all phpIPAM objects
- Comprehensive documentation and examples
- Simple and easy-to-use interface

## Installation

```bash
go get github.com/yourusername/phpipam-go-sdk
```

## Usage

```go
package main

import (
    "fmt"
    "log"
    "os"

    "github.com/yourusername/phpipam-go-sdk"
)

func main() {
    // Create a new PHPIPAM client
    client, err := phpipam.New(
        os.Getenv("PHPIPAM_URL"),       // e.g., "https://ipam.example.com"
        os.Getenv("PHPIPAM_APP_ID"),    // Application ID configured in phpIPAM
        os.Getenv("PHPIPAM_USERNAME"),  // phpIPAM username
        os.Getenv("PHPIPAM_PASSWORD"),  // phpIPAM password
    )
    if err != nil {
        log.Fatalf("Failed to create PHPIPAM client: %v", err)
    }

    // Authenticate to the API to retrieve a token
    err = client.Authenticate()
    if err != nil {
        log.Fatalf("Authentication failed: %v", err)
    }

    // List all sections
    sections, err := client.Sections.List()
    if err != nil {
        log.Fatalf("Failed to list sections: %v", err)
    }

    fmt.Println("Available Sections:")
    for _, section := range sections {
        fmt.Printf("- ID: %s, Name: %s\n", section.ID, section.Name)
    }

    // Create a new subnet
    newSubnet := &phpipam.Subnet{
        Subnet:      "192.168.100.0",
        Mask:        "24",
        SectionID:   "1",  // Section ID where the subnet should be created
        Description: "New subnet created via Go SDK",
    }

    createdSubnet, err := client.Subnets.Create(newSubnet)
    if err != nil {
        log.Fatalf("Failed to create subnet: %v", err)
    }

    fmt.Printf("Created subnet: %s/%s (ID: %s)\n", 
        createdSubnet.Subnet, 
        createdSubnet.Mask, 
        createdSubnet.ID)
}
```

## API Overview

The SDK provides intuitive access to all phpIPAM API controllers:

### Authentication

```go
// Authenticate and get token
err := client.Authenticate()

// Check if token is valid
isValid := client.Client.IsTokenValid()

// Refresh token
err := client.Client.RefreshToken()
```

### Sections

```go
// List all sections
sections, err := client.Sections.List()

// Get a specific section
section, err := client.Sections.Get("1")

// Get section by name
section, err := client.Sections.GetByName("Production")

// Create a new section
newSection := &phpipam.Section{
    Name:        "Test Section",
    Description: "Section created via Go SDK",
}
section, err := client.Sections.Create(newSection)

// Update a section
section.Description = "Updated description"
updatedSection, err := client.Sections.Update(section)

// Delete a section
err := client.Sections.Delete("1")

// Get subnets in a section
subnets, err := client.Sections.GetSubnets("1")

// Get custom fields for sections
customFields, err := client.Sections.GetCustomFields()
```

### Subnets

```go
// List all subnets
subnets, err := client.Subnets.List()

// Get a specific subnet
subnet, err := client.Subnets.Get("1")

// Get subnet usage
usage, err := client.Subnets.GetUsage("1")

// Get addresses in a subnet
addresses, err := client.Subnets.GetAddresses("1")

// Get the first free IP address in a subnet
firstFree, err := client.Subnets.GetFirstFree("1")

// Create a new subnet
newSubnet := &phpipam.Subnet{
    Subnet:      "192.168.200.0",
    Mask:        "24",
    SectionID:   "1",
    Description: "New subnet created via Go SDK",
}
subnet, err := client.Subnets.Create(newSubnet)

// Update a subnet
subnet.Description = "Updated description"
updatedSubnet, err := client.Subnets.Update(subnet)

// Resize a subnet
err := client.Subnets.Resize("1", 25)  // Resize to /25

// Split a subnet
err := client.Subnets.Split("1", 2)  // Split into 2 subnets

// Delete a subnet
err := client.Subnets.Delete("1")
```

### Addresses

```go
// Get a specific address
address, err := client.Addresses.Get("1")

// Get all addresses
addresses, err := client.Addresses.GetAll()

// Search for addresses
results, err := client.Addresses.Search("192.168.1")

// Search by hostname
results, err := client.Addresses.SearchByHostname("server")

// Get the first free address in a subnet
firstFree, err := client.Addresses.GetFirstFree("1")

// Create a new address
newAddress := &phpipam.Address{
    SubnetID:    "1",
    IP:          "192.168.1.10",
    Hostname:    "test-server",
    Description: "Address created via Go SDK",
}
address, err := client.Addresses.Create(newAddress)

// Update an address
address.Description = "Updated description"
updatedAddress, err := client.Addresses.Update(address)

// Delete an address
err := client.Addresses.Delete("1")
```

### VLANs

```go
// List all VLANs
vlans, err := client.VLANs.List()

// Get a specific VLAN
vlan, err := client.VLANs.Get("1")

// Get subnets in a VLAN
subnets, err := client.VLANs.GetSubnets("1")

// Create a new VLAN
newVLAN := &phpipam.VLAN{
    Name:        "Test VLAN",
    Number:      "100",
    Description: "VLAN created via Go SDK",
}
vlan, err := client.VLANs.Create(newVLAN)

// Update a VLAN
vlan.Description = "Updated description"
updatedVLAN, err := client.VLANs.Update(vlan)

// Delete a VLAN
err := client.VLANs.Delete("1")
```

### L2 Domains

```go
// List all L2 domains
domains, err := client.L2Domains.List()

// Get a specific L2 domain
domain, err := client.L2Domains.Get("1")

// Get VLANs in a L2 domain
vlans, err := client.L2Domains.GetVLANs("1")
```

### VRFs

```go
// List all VRFs
vrfs, err := client.VRFs.List()

// Get a specific VRF
vrf, err := client.VRFs.Get("1")

// Get subnets in a VRF
subnets, err := client.VRFs.GetSubnets("1")
```

### Devices

```go
// List all devices
devices, err := client.Devices.List()

// Get a specific device
device, err := client.Devices.Get("1")

// Get subnets in a device
subnets, err := client.Devices.GetSubnets("1")

// Get addresses in a device
addresses, err := client.Devices.GetAddresses("1")

// Search for devices
results, err := client.Devices.Search("server")
```

### Tools

The Tools controller provides access to various utility objects in phpIPAM:

```go
// Get all IP tags
tags, err := client.Tools.GetIPTags()

// Get device types
deviceTypes, err := client.Tools.GetDeviceTypes()

// Get nameservers
nameservers, err := client.Tools.GetNameservers()

// Get locations
locations, err := client.Tools.GetLocations()

// Get racks
racks, err := client.Tools.GetRacks()
```

### Prefix

The Prefix controller is used for automatic subnet/address provisioning:

```go
// Get subnets for a customer type
subnets, err := client.Prefix.GetSubnets("CustomerA")

// Get first available subnet
subnet, err := client.Prefix.GetFirstAvailableSubnet("CustomerA", phpipam.IPv4, 24)

// Get first available address
address, err := client.Prefix.GetFirstAvailableAddress("CustomerA", phpipam.IPv4)
```

### Search

```go
// Search with default options
results, err := client.Search.Search("server")

// Search with custom options
options := phpipam.SearchOptions{
    IncludeAddresses: true,
    IncludeSubnets:   true,
    IncludeVLANs:     true,
    IncludeVRFs:      true,
}
results, err := client.Search.SearchWithOptions("server", options)
```

## Configuration in phpIPAM

Before using this SDK, you need to configure an API app in phpIPAM:

1. Enable the API in phpIPAM: Administration > phpIPAM Settings > Feature Settings > API
2. Create an API app: Administration > API > Create API key
   - Set a unique App ID (this will be used in your code)
   - Set appropriate permissions (usually read/write for all controllers)
   - Select "User token" for security
3. Use the App ID along with phpIPAM user credentials in your Go code

## License

This SDK is released under the MIT License. See the LICENSE file for details.
