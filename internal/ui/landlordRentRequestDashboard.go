package ui

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
	"rentease/internal/domain/entities"
	"rentease/pkg/utils"
	"strconv"
)

// RentRequestsDashboardForLandlord handles the dashboard for landlords to manage property rental requests.
func (ui *UI) RentRequestsDashboardForLandlord() {

	// Fetch the landlord's details based on the active user
	landlord, err := ui.fetchLandlordDetails()
	if err != nil {
		fmt.Printf("\033[1;31mError fetching landlord details: %v\033[0m\n", err) // Red
		return
	}

	// Fetch rental requests associated with the landlord
	requests, err := ui.fetchRequestsForLandlord(landlord.Username)
	if err != nil {
		fmt.Printf("\033[1;31mError retrieving requests: %v\033[0m\n", err) // Red
		return
	}

	// Inform the landlord if there are no new requests
	if len(requests) == 0 {
		fmt.Println("\033[1;33mNo new requests.\033[0m") // Yellow
		return
	}

	// Display all requests with details
	ui.displayRequests(requests)

	// Get user's choice for request action
	choice := ui.getRequestActionChoice(len(requests))
	if choice == 0 {
		return
	}

	// Validate the request number
	if choice < 1 || choice > len(requests) {
		fmt.Println("\033[1;31mInvalid request number.\033[0m") // Red
		return
	}

	req := requests[choice-1]

	// Get the new status for the selected request
	status := ui.getRequestStatusChoice()
	if status == "" {
		return
	}

	// Update the status of the selected request
	err = ui.RequestService.UpdateRequestStatus(req, status)
	if err != nil {
		fmt.Printf("\033[1;31mError updating request status: %v\033[0m\n", err) // Red
	} else {
		fmt.Println("\033[1;32mRequest status updated successfully.\033[0m") // Green
		ui.updatePropertyRentalStatus(req, status)
	}
}

// fetchLandlordDetails retrieves the details of the landlord from the user service.
func (ui *UI) fetchLandlordDetails() (entities.User, error) {
	return ui.UserService.FindByUsername(utils.ActiveUser)
}

// fetchRequestsForLandlord retrieves all property rental requests for a given landlord.
func (ui *UI) fetchRequestsForLandlord(username string) ([]entities.Request, error) {
	return ui.RequestService.GetRentRequestsInfoForLandlord(username)
}

// displayRequests prints the details of all rental requests to the console.
func (ui *UI) displayRequests(requests []entities.Request) {
	fmt.Println("\n\033[1;34mProperty Requests\033[0m") // Blue

	// Create a new tablewriter table
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"#", "Property Title", "Tenant Name", "Phone", "Email", "Address", "Status"})
	table.SetBorder(true)        // Enable border
	table.SetRowLine(true)       // Enable row separator
	table.SetColMinWidth(5, 50)  // Set a minimum width for the address column
	table.SetAutoWrapText(false) // Disable text wrapping

	for i, req := range requests {
		// Fetch property details using PropertyID
		property, err := ui.PropertyService.FindByID(req.PropertyID)
		if err != nil {
			fmt.Printf("\033[1;31mError fetching property details for request %d: %v\033[0m\n", i+1, err) // Red
			continue
		}

		// Fetch tenant details using TenantName
		tenant, err := ui.UserService.FindByUsername(req.TenantName)
		if err != nil {
			fmt.Printf("\033[1;31mError fetching tenant details for request %d: %v\033[0m\n", i+1, err) // Red
			continue
		}

		// Concatenate the address into a single line, ensuring it fits the column
		address := fmt.Sprintf("%s", tenant.Address)

		// Add the row to the table
		table.Append([]string{
			fmt.Sprintf("%d", i+1),
			property.Title,
			tenant.Name,
			tenant.PhoneNumber,
			tenant.Email,
			address,
			req.RequestStatus,
		})
	}

	// Render the table
	table.Render()
}

// getRequestActionChoice prompts the user to enter the number of the request they want to act on.
func (ui *UI) getRequestActionChoice(maxOptions int) int {

	var choice int
	choiceTemp := utils.ReadInput("\nEnter the request number to act on (or 0 to exit): ")
	choice, _ = strconv.Atoi(choiceTemp)
	return choice
}

// getRequestStatusChoice prompts the user to select the new status for the request.
func (ui *UI) getRequestStatusChoice() string {

	var statusChoice int
	choiceTemp := utils.ReadInput("Enter new status (1 for Accepted, 2 for Rejected): ")
	statusChoice, _ = strconv.Atoi(choiceTemp)

	switch statusChoice {
	case 1:
		return "accepted"
	case 2:
		return "rejected"
	default:
		fmt.Println("\033[1;31mInvalid choice.\033[0m") // Red
		return ""
	}
}

// updatePropertyRentalStatus updates the property status to rented if the request is accepted.
func (ui *UI) updatePropertyRentalStatus(req entities.Request, status string) {
	if status == "accepted" {
		prop, _ := ui.PropertyService.FindByID(req.PropertyID)
		prop.IsRented = true
		_ = ui.PropertyService.UpdateListedProperty(prop)
	}
}
