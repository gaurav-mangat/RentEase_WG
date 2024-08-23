package ui

import (
	"fmt"
	"rentease/pkg/utils"
)

func (ui *UI) sellerDashboard() {
	for {
		// Display the seller dashboard
		fmt.Println("\033[1;36m----------------------------------------------------------------\033[0m") // Sky blue
		fmt.Println("\033[1;31m                        SELLER DASHBOARD                        \033[0m") // Red bold
		fmt.Println("\033[1;36m----------------------------------------------------------------\033[0m") // Sky blue

		fmt.Println("\033[1;32m1. List Your Property\033[0m")              // Green
		fmt.Println("\033[1;32m2. View and Manage Listed Property\033[0m") // Green
		fmt.Println("\033[1;32m3. Managing rent request\033[0m")           // Green
		fmt.Println("\033[1;31m4. Back to Main Dashboard\033[0m")          // Red
		fmt.Print("\nEnter your choice: ")

		var choice int
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Printf("\033[1;31mError reading input: %v\033[0m\n", err) // Red
			continue
		}

		switch choice {
		case 1:
			fmt.Println("\033[1;33m\n\nListing a new property...\033[0m") // Yellow
			ui.ListPropertyUI()
		case 2:
			fmt.Println("\033[1;33m\nViewing and managing listed properties...\033[0m") // Yellow
			listedProperties, err := ui.propertyService.GetAllListedProperties(true)
			if err != nil {
				fmt.Printf("\033[1;31mError fetching listed properties: %v\033[0m\n", err)
				continue
			}

			if len(listedProperties) == 0 {
				fmt.Println("\033[1;31mNo listed properties found.\033[0m")
				continue
			}

			// Display properties with indexing
			for i, property := range listedProperties {
				fmt.Printf("\n%d. \n", i+1)
				utils.DisplayProperty(property)
			}

			fmt.Println("\033[1;32mSelect a property to update or delete (enter 0 to go back):\033[0m")
			var index int
			_, err = fmt.Scanln(&index)
			if err != nil || index < 0 || index > len(listedProperties) {
				fmt.Printf("\033[1;31mInvalid choice. Please select a valid option.\033[0m\n")
				continue
			}

			if index == 0 {
				continue
			}

			selectedProperty := listedProperties[index-1]
			fmt.Println("\033[1;33m\nSelected Property:\033[0m")
			utils.DisplayProperty(selectedProperty)

			fmt.Println("\033[1;32mDo you want to update or delete this property?\033[0m")
			fmt.Println("\033[1;32m1. Update\033[0m")
			fmt.Println("\033[1;32m2. Delete\033[0m")
			fmt.Println("\033[1;31m3. Go Back\033[0m")
			fmt.Print("\nEnter your choice: ")

			var action int
			_, err = fmt.Scanln(&action)
			if err != nil || action < 1 || action > 3 {
				fmt.Printf("\033[1;31mInvalid choice. Please select a valid option.\033[0m\n")
				continue
			}

			switch action {
			case 1:
				ui.UpdatePropertyUI(selectedProperty)
			case 2:
				err := ui.propertyService.DeleteListedProperty(selectedProperty.Title)
				if err != nil {
					fmt.Printf("\033[1;31mError deleting property: %v\033[0m\n", err)
				} else {
					fmt.Println("\033[1;32mProperty deleted successfully.\033[0m")
				}
			case 3:
				continue
			}

		case 3:
			ui.LandlordRequestsDashboard()
			// Add functionality to handle chats here
		case 4:
			return // Go back to the main dashboard
		default:
			fmt.Println("\033[1;31m\nInvalid choice. Please select a valid option.\033[0m") // Red
		}
	}
}
