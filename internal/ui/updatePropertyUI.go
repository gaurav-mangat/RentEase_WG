package ui

import (
	"fmt"
	"rentease/internal/domain/entities"
	"rentease/pkg/utils"
)

// UpdatePropertyUI handles the property update user interface.
func (ui *UI) UpdatePropertyUI(property entities.Property) {
	fmt.Println("\033[1;33m\nUpdating property...\033[0m") // Yellow

	// Create a copy of the existing property
	updatedProperty := property

	// Update Title
	ui.updateTitle(&updatedProperty)

	// Update Address
	ui.updateAddress(&updatedProperty)

	// Update Rent Amount
	ui.updateRentAmount(&updatedProperty)

	// Update Details based on Property Type
	ui.updateDetails(&updatedProperty)

	// Reset approval status if the property was approved
	if property.IsApprovedByAdmin {
		updatedProperty.IsApprovedByAdmin = false
	}

	// Save updated property
	if err := ui.PropertyService.UpdateListedProperty(updatedProperty); err != nil {
		fmt.Printf("\033[1;31mError updating property: %v\033[0m\n", err)
	} else {
		fmt.Println("\033[1;32mProperty updated successfully.\033[0m")
	}
}

// updateTitle updates the title of the property.
func (ui *UI) updateTitle(property *entities.Property) {
	newTitle := utils.ReadInput("\nCurrent Title: " + property.Title + "\nEnter new title (leave blank to skip): ")
	if newTitle != "" {
		property.Title = newTitle
	}
}

// updateAddress updates the address of the property.
func (ui *UI) updateAddress(property *entities.Property) {
	fmt.Printf("\nCurrent Address: %s, %s, %s, %d\n", property.Address.Area, property.Address.City, property.Address.State, property.Address.Pincode)

	newArea := utils.ReadInput("Enter new area (leave blank to skip): ")
	if newArea != "" {
		property.Address.Area = newArea
	}

	newCity := utils.ReadInput("Enter new city (leave blank to skip): ")
	if newCity != "" {
		property.Address.City = newCity
	}

	newState := utils.ReadInput("Enter new state (leave blank to skip): ")
	if newState != "" {
		property.Address.State = newState
	}

	newPincode := utils.ReadPincode()
	property.Address.Pincode = newPincode

}

// updateRentAmount updates the rent amount of the property.
func (ui *UI) updateRentAmount(property *entities.Property) {
	fmt.Println("\nCurrent expected rent amount:", property.RentAmount)
	newRentAmountStr := utils.ReadInput("Enter new expected rent amount (leave blank to skip): ")
	if newRentAmountStr != "" {
		var newRentAmount float64
		if _, err := fmt.Sscanf(newRentAmountStr, "%f", &newRentAmount); err == nil {
			property.RentAmount = newRentAmount
		} else {
			fmt.Println("\033[1;31mInvalid rent amount format.\033[0m") // Red
		}
	}
}

// updateDetails updates the details of the property based on its type.
func (ui *UI) updateDetails(property *entities.Property) {
	switch property.Details.(type) {
	case entities.CommercialDetails:
		if utils.ReadInput("\nUpdate commercial property details (floor area, subtype)? (yes/no): ") == "yes" {
			fmt.Println("Current details:", property.Details.(entities.CommercialDetails))
			fmt.Println()
			property.Details = ui.collectCommercialDetails()
		}

	case entities.HouseDetails:
		if utils.ReadInput("\nUpdate house details (number of rooms, furnished category, amenities)? (yes/no): ") == "yes" {
			fmt.Println("Current details:", property.Details.(entities.HouseDetails))
			fmt.Println()
			property.Details = ui.collectHouseDetails()
		}

	case entities.FlatDetails:
		if utils.ReadInput("\nUpdate flat details (furnished category, amenities, BHK)? (yes/no): ") == "yes" {
			fmt.Println("Current details:", property.Details.(entities.FlatDetails))
			fmt.Println()
			property.Details = ui.collectFlatDetails()
		}

	default:
		fmt.Println("\nUnknown property details type")
	}
}
