package ui

import (
	"fmt"
	"rentease/pkg/utils"
)

func (ui *UI) SearchPropertyUI() {
	fmt.Println("\n\033[1;34m========================\033[0m") // Blue
	fmt.Println("\033[1;34mSearch Property\033[0m")            // Blue
	fmt.Println("\033[1;34m========================\033[0m")   // Blue

	// Collect search criteria from the tenant
	var propertyType int
	for {
		fmt.Print("Enter property type (1. Commercial, 2. House, 3. Flat): ")
		_, err := fmt.Scanf("%d", &propertyType)

		// Check if the input is valid
		if err != nil || propertyType < 1 || propertyType > 3 {
			fmt.Println("\033[1;31mInvalid input. Please enter a valid property type (1, 2, or 3).\033[0m") // Red
			continue
		}

		// Break the loop if the input is valid
		break
	}

	area := utils.ReadInput("Enter locality (leave blank to skip): ")
	city := utils.ReadInput("Enter city (leave blank to skip): ")
	state := utils.ReadInput("Enter state (leave blank to skip): ")
	pincode := utils.ReadPincode()

	// Call service to search for properties
	properties, err := ui.propertyService.SearchProperties(area, city, state, pincode, propertyType)
	if err != nil {
		fmt.Printf("\033[1;31mError searching properties: %v\033[0m\n", err) // Red
		return
	}

	// Display the search results
	if len(properties) == 0 {
		fmt.Println("\033[1;33mNo properties found matching your criteria.\033[0m") // Yellow
		return
	}

	fmt.Println("\n\033[1;34mSearch Results\033[0m")         // Blue
	fmt.Println("\033[1;34m========================\033[0m") // Blue
	utils.DisplayPropertyshortInfo(properties)

	for {
		fmt.Print("Enter the property number to see more details (or 0 to exit): ")
		var choice int
		fmt.Scan(&choice)

		if choice == 0 {
			break
		}

		if choice < 1 || choice > len(properties) {
			fmt.Println("\033[1;31mInvalid property number.\033[0m") // Red
			continue
		}

		prop := properties[choice-1]
		utils.DisplayProperty(prop)

		// Fetch landlord details
		landlord, err := ui.userService.FindByUsername(prop.LandlordUsername)
		if err != nil {
			fmt.Printf("\033[1;31mError fetching landlord details: %v\033[0m\n", err) // Red
			continue
		}

		fmt.Println("\n\033[1;34mLandlord Details\033[0m")
		fmt.Println("  Name: ", landlord.Name)
		fmt.Println("  Phone: ", landlord.PhoneNumber)
		fmt.Println("  Email: ", landlord.Email)
		fmt.Println("  Address: ", landlord.Address)

		for {
			fmt.Println("\n\033[1;36mWhat would you like to do?\033[0m")
			fmt.Println("1. Add to Wishlist")
			fmt.Println("2. Request Property")
			fmt.Println("3. View Another Property")
			fmt.Println("4. Exit")

			var action int
			fmt.Print("Enter your choice: ")
			fmt.Scan(&action)

			switch action {
			case 1:
				err := ui.userService.AddToWishlist(utils.ActiveUser, prop.ID)
				if err != nil {
					fmt.Printf("\033[1;31mError adding property to wishlist: %v\033[0m\n", err) // Red
				} else {
					fmt.Println("\033[1;32mProperty added to wishlist successfully.\033[0m") // Green
				}
			case 2:
				if utils.ActiveUser != prop.LandlordUsername {
					err = ui.requestService.CreatePropertyRequest(utils.ActiveUser, prop.ID, prop.LandlordUsername)
					if err != nil {
						fmt.Printf("\033[1;31mError requesting property: %v\033[0m\n", err) // Red
					} else {
						fmt.Println("\033[1;32mProperty request sent successfully.\033[0m") // Green
					}
				} else {
					fmt.Println("\033[1;31mYou cannot request your own property!\033[0m") // Red
				}
			case 3:
				// Break inner loop to return to property list
				break
			case 4:
				// Exit the entire search flow
				return
			default:
				fmt.Println("\033[1;31mInvalid choice. Please select a valid option.\033[0m") // Red
			}

			if action == 3 {
				break
			}
		}
	}
}
