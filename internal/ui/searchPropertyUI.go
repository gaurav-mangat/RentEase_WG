package ui

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"rentease/internal/domain/entities"
	"rentease/pkg/utils"
)

// SearchPropertyUI handles the user interface for searching properties.
func (ui *UI) SearchPropertyUI() {
	fmt.Println("\n\033[1;34m========================\033[0m") // Blue
	fmt.Println("\033[1;34mSearch Property\033[0m")            // Blue
	fmt.Println("\033[1;34m========================\033[0m")   // Blue

	// Collect property type from the user
	propertyType := ui.promptForPropertyType()

	// Collect additional search criteria
	area := utils.ReadInput("Enter locality (leave blank to skip): ")
	city := utils.ReadInput("Enter city (leave blank to skip): ")
	state := utils.ReadInput("Enter state (leave blank to skip): ")
	pincode := utils.ReadPincode()

	// Search for properties based on the criteria
	properties, err := ui.PropertyService.SearchProperties(area, city, state, pincode, propertyType)
	if err != nil {
		fmt.Printf("\033[1;31mError searching properties: %v\033[0m\n", err) // Red
		return
	}

	// Display search results
	if len(properties) == 0 {
		fmt.Println("\033[1;33mNo properties found matching your criteria.\033[0m") // Yellow
		return
	}

	fmt.Println("\n\033[1;34mSearch Results\033[0m")         // Blue
	fmt.Println("\033[1;34m========================\033[0m") // Blue
	ui.DisplayPropertyShortInfo(properties, nil)

	// Allow the user to view property details and perform actions
	ui.handlePropertyActions(properties)
}

// promptForPropertyType collects and validates the property type from the user.
func (ui *UI) promptForPropertyType() int {
	var propertyType int
	for {
		fmt.Print("Enter property type (1. Commercial, 2. House, 3. Flat): ")
		_, err := fmt.Scanf("%d", &propertyType)
		if err != nil || propertyType < 1 || propertyType > 3 {
			fmt.Println("\033[1;31mInvalid input. Please enter a valid property type (1, 2, or 3).\033[0m") // Red
			continue
		}
		break
	}
	return propertyType
}

// handlePropertyActions allows the user to perform actions on a selected property.
func (ui *UI) handlePropertyActions(properties []entities.Property) {
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
		landlord, err := ui.UserService.FindByUsername(prop.LandlordUsername)
		if err != nil {
			fmt.Printf("\033[1;31mError fetching landlord details: %v\033[0m\n", err) // Red
			continue
		}

		// Display landlord details
		ui.displayLandlordDetails(landlord)

		// Allow user to perform actions on the selected property
		ch := ui.performPropertyAction(prop)
		if ch == "exiting" {
			return
		}
	}
}

// displayLandlordDetails shows details about the landlord of the property.
func (ui *UI) displayLandlordDetails(landlord entities.User) {
	fmt.Println("\033[1;34mLandlord Details\033[0m")
	fmt.Println("  Name: ", landlord.Name)
	fmt.Println("  Phone: ", landlord.PhoneNumber)
	fmt.Println("  Email: ", landlord.Email)
	fmt.Println("  Address: ", landlord.Address)
}

// performPropertyAction presents options to the user for the selected property.
func (ui *UI) performPropertyAction(prop entities.Property) string {
	for {
		fmt.Println("\n\033[1;36mWhat would you like to do?\033[0m")
		fmt.Println("1. Add to Wishlist")
		fmt.Println("2. Request Property")
		fmt.Println("3. View Another Property")
		fmt.Println("4. Back to Tenant Dashboard")

		var action int
		fmt.Print("\nEnter your choice: ")
		_, _ = fmt.Scanln(&action)

		switch action {
		case 1:
			ui.handleAddToWishlist(prop.ID)
		case 2:
			ui.handlePropertyRequest(prop)
		case 3:
			// Break inner loop to return to property list
			return "not exiting"
		case 4:
			// Exit the entire action loop and go back to the main menu or previous screen
			return "exiting"
		default:
			fmt.Println("\033[1;31mInvalid choice. Please select a valid option.\033[0m") // Red
		}
	}
}

// handleAddToWishlist adds the selected property to the user's wishlist.
func (ui *UI) handleAddToWishlist(propertyID primitive.ObjectID) {
	err := ui.UserService.AddToWishlist(utils.ActiveUser, propertyID)
	if err != nil {
		fmt.Printf("\033[1;31mError adding property to wishlist: %v\033[0m\n", err) // Red
	} else {
		fmt.Println("\033[1;32mProperty added to wishlist successfully.\033[0m") // Green
	}
}

// handlePropertyRequest sends a request to rent the selected property.
func (ui *UI) handlePropertyRequest(prop entities.Property) {
	if utils.ActiveUser != prop.LandlordUsername {
		err := ui.RequestService.CreateRentRequest(utils.ActiveUser, prop.ID, prop.LandlordUsername)
		if err != nil {
			fmt.Printf("\033[1;31mError requesting property: %v\033[0m\n", err) // Red
		} else {
			fmt.Println("\033[1;32mProperty request sent successfully.\033[0m") // Green
		}
	} else {
		fmt.Println("\033[1;31mYou cannot request your own property!\033[0m") // Red
	}
}
