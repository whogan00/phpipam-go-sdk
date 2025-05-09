package main

import (
	"fmt"
	"log"
	"os"

	"github.com/whogan00/phpipam-go-sdk/pkg/phpipam"
)

func main() {
	// Create a new PHPIPAM client
	client, err := phpipam.New(
		os.Getenv("PHPIPAM_URL"),      // e.g., "https://ipam.example.com"
		os.Getenv("PHPIPAM_APP_ID"),   // Application ID configured in phpIPAM
		os.Getenv("PHPIPAM_USERNAME"), // phpIPAM username
		os.Getenv("PHPIPAM_PASSWORD"), // phpIPAM password
		true,
	)
	if err != nil {
		log.Fatalf("Failed to create PHPIPAM client: %v", err)
	}

	// Authenticate to the API to retrieve a token
	err = client.Authenticate()
	if err != nil {
		log.Fatalf("Authentication failed: %v", err)
	}

	fmt.Println("Authentication successful!")
	fmt.Printf("Token: %s (valid until: %s)\n", client.Client.Token, client.Client.TokenExp.Format("2006-01-02 15:04:05"))

	// Example: List all sections
	sections, err := client.Sections.List()
	if err != nil {
		log.Fatalf("Failed to list sections: %v", err)
	}

	fmt.Println("\nAvailable Sections:")
	for _, section := range sections {
		fmt.Printf("- ID: %s, Name: %s, Description: %s\n", section.ID, section.Name, section.Description)
	}

	// Example: Get subnets in a section
	if len(sections) > 0 {
		sectionID := sections[0].ID
		subnets, err := client.Sections.GetSubnets(sectionID)
		if err != nil {
			log.Fatalf("Failed to get subnets for section %s: %v", sectionID, err)
		}

		fmt.Printf("\nSubnets in Section %s (%s):\n", sectionID, sections[0].Name)
		for _, subnet := range subnets {
			fmt.Printf("- %s/%s - %s\n", subnet.Subnet, subnet.Mask, subnet.Description)
		}

		// Example: Get addresses in a subnet
		if len(subnets) > 0 {
			subnetID := subnets[0].ID
			addresses, err := client.Subnets.GetAddresses(subnetID)
			if err != nil {
				log.Fatalf("Failed to get addresses for subnet %s: %v", subnetID, err)
			}

			fmt.Printf("\nAddresses in Subnet %s (%s/%s):\n", subnetID, subnets[0].Subnet, subnets[0].Mask)
			for _, address := range addresses {
				hostname := address.Hostname
				if hostname == "" {
					hostname = "[No hostname]"
				}
				fmt.Printf("- %s - %s\n", address.IP, hostname)
			}
		}
	}

	// Example: Search for an IP address
	searchTerm := "192.168.1"
	fmt.Printf("\nSearching for IP addresses containing '%s':\n", searchTerm)
	results, err := client.Addresses.Search(searchTerm)
	if err != nil {
		log.Fatalf("Failed to search for addresses: %v", err)
	}

	for _, address := range results {
		fmt.Printf("- %s (Subnet ID: %s, Description: %s)\n", address.IP, address.SubnetID, address.Description)
	}

	// Example: Get first free IP address in a subnet
	if len(sections) > 0 && len(client.Sections.GetSubnets(sections[0].ID)) > 0 {
		subnets, _ := client.Sections.GetSubnets(sections[0].ID)
		if len(subnets) > 0 {
			subnetID := subnets[0].ID
			firstFree, err := client.Subnets.GetFirstFree(subnetID)
			if err != nil {
				log.Printf("Failed to get first free IP in subnet %s: %v", subnetID, err)
			} else {
				fmt.Printf("\nFirst free IP address in subnet %s/%s: %s\n", subnets[0].Subnet, subnets[0].Mask, firstFree)
			}
		}
	}

	// Example: Create a new subnet
	newSubnet := &phpipam.Subnet{
		Subnet:      "192.168.100.0",
		Mask:        "24",
		SectionID:   sections[0].ID,
		Description: "New subnet created via Go SDK",
	}

	createdSubnet, err := client.Subnets.Create(newSubnet)
	if err != nil {
		log.Printf("Failed to create subnet: %v", err)
	} else {
		fmt.Printf("\nCreated new subnet: %s/%s (ID: %s)\n", createdSubnet.Subnet, createdSubnet.Mask, createdSubnet.ID)
	}
}
