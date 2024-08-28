package ui

import (
	"fmt"
	"rentease/internal/domain/entities"
	"rentease/pkg/utils"
)

func (ui *UI) landlordDashboard() {
	for {
		// Display the landlord dashboard header
		fmt.Println("\n\n\n\033[1;36m----------------------------------------------------------\033[0m") // Sky blue
		fmt.Println("\033[1;31m              LANDLORD DASHBOARD                        \033[0m")         // Red bold
		fmt.Println("\033[1;36m-----------------------------------------------------------\033[0m")      // Sky blue

		// Display the options available to the landlord
		fmt.Println("		\033[1;32m1. List Your Property\033[0m")              // Green
		fmt.Println("		\033[1;32m2. View and Manage Listed Property\033[0m") // Green
		fmt.Println("		\033[1;32m3. Manage Rent Requests\033[0m")            // Green
		fmt.Println("		\033[1;31m4. Back to Main Dashboard\033[0m")          // Red
		fmt.Print("\nEnter your choice: ")

		// Read user input for the selected option
		var choice int
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Printf("\033[1;31mError reading input: %v\033[0m\n", err) // Red
			continue
		}

		// Handle the user's choice with a switch case
		switch choice {
		case 1:
			fmt.Println("\033[1;33m\n\nListing a new property...\033[0m") // Yellow
			ui.ListPropertyUI()

		case 2:
			// View and manage listed properties
			ui.viewAndManageListedProperties(false)

		case 3:
			// Manage rent requests
			ui.RentRequestsDashboardForLandlord()

		case 4:
			// Go back to the main dashboard
			return

		default:
			fmt.Println("\033[1;31m\nInvalid choice. Please select a valid option.\033[0m") // Red
		}
	}
}

// viewAndManageListedProperties handles the process of viewing and managing listed properties.
func (ui *UI) viewAndManageListedProperties(isViewingProfile bool) {

	// Fetch all listed properties
	listedProperties, err := ui.PropertyService.GetAllListedProperties(true)
	if err != nil {
		// Display error if fetching properties fails
		ui.displayError("fetching listed properties", err)
		return
	}

	// Check if there are no properties listed
	if len(listedProperties) == 0 {
		fmt.Println("\033[1;31mNo listed properties found.\033[0m")
		return
	}

	// Display properties with their respective index
	ui.displayPropertiesWithIndex(listedProperties)

	if !isViewingProfile {
		// Allow user to select a property by index
		index := ui.getPropertySelection(len(listedProperties))
		if index == 0 {
			return // Go back if the user selects 0
		}

		// Handle action on the selected property (update/delete)
		selectedProperty := listedProperties[index-1]
		ui.handlePropertyAction(selectedProperty)
	}
}

// displayError shows an error message with context about the action that failed.
func (ui *UI) displayError(action string, err error) {
	fmt.Printf("\033[1;31mError %s: %v\033[0m\n", action, err)
}

// displayPropertiesWithIndex shows a list of properties with index numbers for selection.
func (ui *UI) displayPropertiesWithIndex(properties []entities.Property) {
	for i, property := range properties {
		fmt.Printf("\n%d. \n", i+1)
		utils.DisplayProperty(property)
	}
}

// getPropertySelection prompts the user to select a property by index.
func (ui *UI) getPropertySelection(propertyCount int) int {
	fmt.Print("\033[1;32mSelect a property to update or delete (enter 0 to go back):  \033[0m")
	var index int
	_, err := fmt.Scanln(&index)
	if err != nil || index < 0 || index > propertyCount {
		// Handle invalid input by asking again
		fmt.Printf("\033[1;31mInvalid choice. Please select a valid option.\033[0m\n")
		return ui.getPropertySelection(propertyCount)
	}
	return index
}

// handlePropertyAction allows the landlord to either update or delete the selected property.
func (ui *UI) handlePropertyAction(property entities.Property) {
	fmt.Println("\033[1;33m\nSelected Property:\033[0m")
	utils.DisplayProperty(property)

	// Prompt the user to choose an action (update/delete) on the selected property
	fmt.Println("\033[1;32mDo you want to update or delete this property?\033[0m")
	fmt.Println("\033[1;32m1. Update\033[0m")
	fmt.Println("\033[1;32m2. Delete\033[0m")
	fmt.Println("\033[1;31m3. Go Back\033[0m")
	fmt.Print("\nEnter your choice: ")

	var action int
	_, err := fmt.Scanln(&action)
	if err != nil || action < 1 || action > 3 {
		// Handle invalid action by asking again
		fmt.Printf("\033[1;31mInvalid choice. Please select a valid option.\033[0m\n")
		ui.handlePropertyAction(property)
		return
	}

	// Perform the selected action
	switch action {
	case 1:
		// Update the selected property
		ui.UpdatePropertyUI(property)
	case 2:
		// Delete the selected property
		err := ui.PropertyService.DeleteListedProperty(property.Title)
		if err != nil {
			ui.displayError("deleting property :", err)
		} else {
			fmt.Println("\033[1;32mProperty deleted successfully.\033[0m")
		}
	case 3:
		// Go back without doing anything
		return
	}
}
