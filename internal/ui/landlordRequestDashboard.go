package ui

import (
	"fmt"
	"rentease/pkg/utils"
)

func (ui *UI) LandlordRequestsDashboard() {
	// Fetch landlord details
	landlord, err := ui.userService.FindByUsername(utils.ActiveUser)
	if err != nil {
		fmt.Printf("\033[1;31mError fetching landlord details: %v\033[0m\n", err) // Red
		return
	}

	// Fetch requests for the landlord

	requests, err := ui.requestService.GetRequestsForLandlord(landlord.Username)
	if err != nil {
		fmt.Printf("\033[1;31mError retrieving requests: %v\033[0m\n", err) // Red
		return
	}

	if len(requests) == 0 {
		fmt.Println("\033[1;33mNo new requests.\033[0m") // Yellow
		return
	}

	// Display the requests with additional details
	fmt.Println("\n\033[1;34mProperty Requests\033[0m") // Blue
	for i, req := range requests {
		// Fetch property details based on PropertyID
		property, err := ui.propertyService.FindByID(req.PropertyID)
		if err != nil {
			fmt.Printf("\033[1;31mError fetching property details for request %d: %v\033[0m\n", i+1, err) // Red
			continue
		}

		// Fetch tenant details
		tenant, err := ui.userService.FindByUsername(req.TenantName)
		if err != nil {
			fmt.Printf("\033[1;31mError fetching tenant details for request %d: %v\033[0m\n", i+1, err) // Red
			continue
		}

		// Display the request details
		fmt.Printf("%d. Property: %s (ID: %s)\n", i+1, property.Title, req.PropertyID.Hex())
		fmt.Printf("   Tenant: %s\n", tenant.Name)
		fmt.Printf("   Phone: %s\n", tenant.PhoneNumber)
		fmt.Printf("   Email: %s\n", tenant.Email)
		fmt.Printf("   Address: %s\n", tenant.Address)
		fmt.Printf("   Status: %s\n", req.RequestStatus)
		fmt.Println()
	}

	// Handle request actions
	fmt.Print("\nEnter the request number to act on (or 0 to exit): ")
	var choice int
	fmt.Scan(&choice)

	if choice == 0 {
		return
	}

	if choice < 1 || choice > len(requests) {
		fmt.Println("\033[1;31mInvalid request number.\033[0m") // Red
		return
	}

	req := requests[choice-1]

	fmt.Print("Enter new status (1 for Accepted, 2 for Rejected): ")
	var statusChoice int
	fmt.Scan(&statusChoice)

	var status string
	switch statusChoice {
	case 1:
		status = "accepted"
	case 2:
		status = "rejected"
	default:
		fmt.Println("\033[1;31mInvalid choice.\033[0m") // Red
		return
	}

	// Update the request status
	err = ui.requestService.UpdateRequestStatus(req, status)
	if err != nil {
		fmt.Printf("\033[1;31mError updating request status: %v\033[0m\n", err) // Red
	} else {
		fmt.Println("\033[1;32mRequest status updated successfully.\033[0m") // Green

		// Updating our IsRented fiels as request is approved
		fmt.Println("Status: ", status)
		if status == "accepted" {
			prop, _ := ui.propertyService.FindByID(req.PropertyID)
			prop.IsRented = true
			_ = ui.propertyService.UpdateListedProperty(prop)
		}
	}
}
