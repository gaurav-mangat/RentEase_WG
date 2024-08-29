package ui

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"rentease/internal/domain/entities"
	"rentease/pkg/utils"
	"strconv"
)

// SearchPropertyUI handles the user interface for searching properties.
func (ui *UI) SearchPropertyUI() {
	fmt.Println("\n\033[1;34m========================\033[0m") // Blue
	fmt.Println("\033[1;34mSearch Property\033[0m")            // Blue
	fmt.Println("\033[1;34m========================\033[0m")   // Blue

	// Collect property type from the user
	propertyType := ui.promptForPropertyType()

	fmt.Println("\nProvide the address ")
	// Prompt user to enter pincode
	pincode := utils.ReadPincode(1)

	// Retrieve address details based on pincode
	address, err := utils.GetAddressFromPincode(pincode)
	if err != nil {
		if pincode != 555555 {
			fmt.Println("\nError fetching address details:", err)
		}
		fmt.Println("Please enter the full address details manually.")

		// Collect full address details from the user
		address = entities.Address{
			Area:    utils.ReadInput("    Enter locality: "),
			City:    utils.ReadInput("    Enter city: "),
			State:   utils.ReadInput("    Enter state: "),
			Pincode: pincode,
		}
	} else {
		// Display fetched address details to the user
		fmt.Printf("\nFetched Address Details:\n")
		fmt.Printf("    Locality: %s\n", address.Area)
		fmt.Printf("    City: %s\n", address.City)
		fmt.Printf("    State: %s\n", address.State)
		fmt.Printf("    Pincode: %d\n", address.Pincode)

		// Prompt the user to confirm or edit the address details
		fmt.Println("\nIf you want to update any field, enter the new value. Leave it blank to keep the current value.")

		// Collect updated address details from the user
		area := utils.ReadInput(fmt.Sprintf("    Enter locality (current: %s): ", address.Area))
		if area != "" {
			address.Area = area
		}

		city := utils.ReadInput(fmt.Sprintf("    Enter city (current: %s): ", address.City))
		if city != "" {
			address.City = city
		}

		state := utils.ReadInput(fmt.Sprintf("    Enter state (current: %s): ", address.State))
		if state != "" {
			address.State = state
		}

		// Pincode remains unchanged since it was used to fetch the address
		address.Pincode = pincode
	}

	// Search for properties based on the criteria
	properties, err := ui.PropertyService.SearchProperties(address.Area, address.City, address.State, pincode, propertyType)
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
	ui.DisplayPropertyShortInfo(properties)

	// Allow the user to view property details and perform actions
	ui.handlePropertyActions(properties)
}

// promptForPropertyType collects and validates the property type from the user.
func (ui *UI) promptForPropertyType() int {
	var propertyType int
	for {
		propertyTypeTemp := utils.ReadInput("Enter property type (1. Commercial, 2. House, 3. Flat): ")
		var err error
		propertyType, err = strconv.Atoi(propertyTypeTemp)
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
		var choice int
		choiceTemp := utils.ReadInput("Enter the property number to see more details (or 0 to exit): ")
		choice, err := strconv.Atoi(choiceTemp)

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
		actionTemp := utils.ReadInput("\nEnter your choice: ")
		action, _ = strconv.Atoi(actionTemp)

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
