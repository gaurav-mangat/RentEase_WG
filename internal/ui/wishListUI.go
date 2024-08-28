package ui

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"rentease/internal/domain/entities"
	"rentease/pkg/utils"
)

// ShowWishlist displays the properties in the currently active user's wishlist.
func (ui *UI) ShowWishlist() error {
	// Fetch the user details from the UserService
	user, err := ui.UserService.FindByUsername(utils.ActiveUser)
	if err != nil {
		return err
	}

	// Check if the wishlist is empty
	if len(user.Wishlist) == 0 {
		fmt.Println("Your wishlist is empty.")
		return nil
	}

	fmt.Println("\n\033[1;34m    Your Wishlist\033[0m")      // Blue
	fmt.Println("\033[1;34m========================\033[0m") // Blue
	fmt.Println()

	// Fetch and display properties in the wishlist
	wishListProperties, err := ui.getPropertiesFromWishlist(user.Wishlist)
	if err != nil {
		return err
	}

	ui.DisplayPropertyShortInfo(wishListProperties, nil)

	// Allow the user to view property details or perform actions
	return ui.handleWishlistActions(user, wishListProperties)
}

// getPropertiesFromWishlist retrieves property details for the given wishlist.
func (ui *UI) getPropertiesFromWishlist(wishlist []primitive.ObjectID) ([]entities.Property, error) {
	var properties []entities.Property
	for _, propertyID := range wishlist {
		prop, err := ui.PropertyService.FindByID(propertyID)
		if err != nil {
			fmt.Printf("\033[1;31mError fetching property details: %v\033[0m\n", err) // Red
			continue
		}
		properties = append(properties, prop)
	}
	return properties, nil
}

// handleWishlistActions handles user actions on properties in the wishlist.
func (ui *UI) handleWishlistActions(user entities.User, properties []entities.Property) error {
	for {
		fmt.Print("\nEnter the property number to see more details (or 0 to exit, -1 to remove a property, 1 to request a property): ")
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 0:
			return nil // Exit

		case 1:
			if err := ui.requestPropertyFromWishlist(user, properties); err != nil {
				return err
			}

		case -1:
			if err := ui.removePropertyFromWishlist(user, properties); err != nil {
				return err
			}

		default:
			if choice < 1 || choice > len(properties) {
				fmt.Println("\033[1;31mInvalid property number.\033[0m") // Red
				continue
			}
			ui.displayPropertyDetails(properties[choice-1])
		}
	}
}

// requestPropertyFromWishlist processes the request to rent a property from the wishlist.
func (ui *UI) requestPropertyFromWishlist(user entities.User, properties []entities.Property) error {
	fmt.Print("Enter the property number to request from wishlist: ")
	var choice int
	fmt.Scan(&choice)

	if choice < 1 || choice > len(properties) {
		fmt.Println("\033[1;31mInvalid property number.\033[0m") // Red
		return nil
	}

	prop := properties[choice-1]
	err := ui.RequestService.CreateRentRequest(utils.ActiveUser, prop.ID, prop.LandlordUsername)
	if err != nil {
		fmt.Printf("\033[1;31mError creating property request: %v\033[0m\n", err) // Red
		return err
	}

	// Remove the property from the wishlist
	user.Wishlist = removePropertyFromList(user.Wishlist, prop.ID)
	err = ui.UserService.UpdateUser(user)
	if err != nil {
		fmt.Printf("\033[1;31mError updating user wishlist: %v\033[0m\n", err) // Red
		return err
	}

	fmt.Println("\033[1;32mProperty removed from wishlist and request sent.\033[0m") // Green
	return nil
}

// removePropertyFromWishlist removes a property ID from the wishlist.
func removePropertyFromList(wishlist []primitive.ObjectID, propertyID primitive.ObjectID) []primitive.ObjectID {
	for i, id := range wishlist {
		if id == propertyID {
			return append(wishlist[:i], wishlist[i+1:]...)
		}
	}
	return wishlist
}

// removePropertyFromWishlist processes the removal of a property from the wishlist.
func (ui *UI) removePropertyFromWishlist(user entities.User, properties []entities.Property) error {
	fmt.Print("Enter the property number to remove from wishlist: ")
	var choice int
	fmt.Scan(&choice)

	if choice < 1 || choice > len(properties) {
		fmt.Println("\033[1;31mInvalid property number.\033[0m") // Red
		return nil
	}

	prop := properties[choice-1]
	user.Wishlist = removePropertyFromList(user.Wishlist, prop.ID)
	err := ui.UserService.UpdateUser(user)
	if err != nil {
		fmt.Printf("\033[1;31mError updating user wishlist: %v\033[0m\n", err) // Red
		return err
	}

	fmt.Println("\033[1;32mProperty removed from wishlist.\033[0m") // Green
	fmt.Println()
	return nil
}

// displayPropertyDetails displays detailed information about a selected property.
func (ui *UI) displayPropertyDetails(prop entities.Property) {
	utils.DisplayProperty(prop)
	fmt.Println("Landlord Details are:")

	landlord, err := ui.UserService.FindByUsername(prop.LandlordUsername)
	if err != nil {
		fmt.Printf("\033[1;31mError fetching landlord details: %v\033[0m\n", err) // Red
		return
	}

	fmt.Printf("  Landlord Name: %s\n", landlord.Name)
	fmt.Printf("  Landlord Phone: %s\n", landlord.PhoneNumber)
	fmt.Printf("  Landlord Email: %s\n", landlord.Email)
	fmt.Printf("  Landlord Address: %s\n", landlord.Address)
	fmt.Println()
}

// DisplayPropertyShortInfo displays a brief summary of each property and this is also used for showNotification for sent requests
func (ui *UI) DisplayPropertyShortInfo(properties []entities.Property, requests []entities.Request) {
	for i, property := range properties {
		// This if condition is useful when property for request is deleted by the landlord
		if property.Title != "" && property.Address.Pincode != 0 {
			fmt.Printf("Property #%d:\n", i+1)
			fmt.Printf("  Title: %s\n", property.Title)
			fmt.Printf("  Rent Amount: %.2f\n", property.RentAmount)
			fmt.Printf("  Address: %s, %s, %s, %d\n", property.Address.Area, property.Address.City, property.Address.State, property.Address.Pincode)
			if requests != nil {
				fmt.Printf("Status: %s\n", requests[i].RequestStatus)
				if requests[i].RequestStatus == "accepted" {
					fmt.Println("\033[32mCongratulations! Your request has been accepted!\033[0m")

					fmt.Println()
				}
			}
			fmt.Println()
		}
	}
}
