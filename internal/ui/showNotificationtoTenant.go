package ui

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"log"
	"os"
	"rentease/internal/domain/entities"
	"rentease/pkg/utils"
)

func (ui *UI) ShowNotifications() {
	requests, err := ui.RequestService.GetRentRequestsInfoForTenant(utils.ActiveUser)
	if err != nil {
		fmt.Printf("\033[1;31mError retrieving notifications: %v\033[0m\n", err) // Red
		return
	}

	if len(requests) == 0 {
		fmt.Println("\033[1;33mNo notifications.\033[0m") // Yellow
		return
	}

	fmt.Println("\n\033[1;34mYour Property Requests\033[0m") // Blue
	var properties []entities.Property
	for _, req := range requests {

		property, err := ui.PropertyService.FindByID(req.PropertyID)
		if err != nil {
			log.Println("\033[1;31mError finding property by id: \033[0m\n", err)
		}
		properties = append(properties, property)

	}
	ui.DisplayRentRequestStatusToTenant(properties, requests)

}
func (ui *UI) DisplayRentRequestStatusToTenant(properties []entities.Property, requests []entities.Request) {
	// Create a new table writer
	table := tablewriter.NewWriter(os.Stdout)

	// Set the header for the table
	table.SetHeader([]string{"No.", "Title", "Rent Amount", "Address", "Request Status"})

	// Set column width and auto-wrap
	table.SetColMinWidth(3, 50) // Minimum width for "Address" column
	table.SetAutoWrapText(false)

	// Populate the table with property data
	for i, property := range properties {
		if property.Address.Pincode != 0 && property.Title != "" {
			address := fmt.Sprintf("%s, %s, %s, %d", property.Address.Area, property.Address.City, property.Address.State, property.Address.Pincode)
			requestStatus := "N/A"
			if requests != nil && i < len(requests) {
				requestStatus = requests[i].RequestStatus
			}
			index := 1

			// Append data to the table
			table.Append([]string{
				fmt.Sprintf("%d", index),
				property.Title,
				fmt.Sprintf("%.2f", property.RentAmount),
				address,
				requestStatus,
			})
			index++
		}
	}

	// Render the table
	table.SetBorder(true)
	table.Render()
}
