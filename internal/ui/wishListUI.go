package ui

import (
	"fmt"
	"rentease/internal/domain/entities"
	"rentease/pkg/utils"
)

// ShowWishlist displays the properties in the currently active user's wishlist.
func (ui *UI) ShowWishlist() error {

	// Fetch the user from the service
	user, err := ui.userService.FindByUsername(utils.ActiveUser)
	if err != nil {
		return err
	}

	if len(user.Wishlist) == 0 {
		fmt.Println("Your wishlist is empty.")
		return nil
	}

	fmt.Println("\n\033[1;34mYour Wishlist\033[0m")          // Blue
	fmt.Println("\033[1;34m========================\033[0m") // Blue

	var wishListProperties []entities.Property
	// Fetch and display properties in the wishlist
	for _, propertyID := range user.Wishlist {
		prop, err := ui.propertyService.FindByID(propertyID)
		if err != nil {
			fmt.Printf("\033[1;31mError fetching property details: %v\033[0m\n", err) // Red
			continue
		}
		wishListProperties = append(wishListProperties, prop)
	}

	utils.DisplayPropertyshortInfo(wishListProperties)

	for {
		fmt.Print("Enter the property number to see more details (or 0 to exit, -1 to remove a property): ")
		var choice int
		fmt.Scan(&choice)

		if choice == 0 {
			break
		}

		if choice == -1 {
			fmt.Print("Enter the property number to remove from wishlist: ")
			fmt.Scan(&choice)

			if choice < 1 || choice > len(user.Wishlist) {
				fmt.Println("\033[1;31mInvalid property number.\033[0m") // Red
				continue
			}

			// Remove the property from the wishlist
			user.Wishlist = append(user.Wishlist[:choice-1], user.Wishlist[choice:]...)
			err := ui.userService.UpdateUser(user)

			if err != nil {
				fmt.Printf("\033[1;31mError removing property from wishlist: %v\033[0m\n", err) // Red
				continue
			}

			fmt.Println("\033[1;32mProperty removed from wishlist.\033[0m") // Green
			return nil
		}

		if choice < 1 || choice > len(wishListProperties) {
			fmt.Println("\033[1;31mInvalid property number.\033[0m") // Red
			continue
		}

		prop := wishListProperties[choice-1]
		utils.DisplayProperty(prop)
		fmt.Println("Landlord Details are:")

		landlord, err := ui.userService.FindByUsername(prop.LandlordUsername)
		if err != nil {
			fmt.Printf("\033[1;31mError fetching landlord details: %v\033[0m\n", err) // Red
			continue
		}

		fmt.Printf("  Landlord Name: %s\n", landlord.Name)
		fmt.Printf("  Landlord Phone: %s\n", landlord.PhoneNumber)
		fmt.Printf("  Landlord Email: %s\n", landlord.Email)
		fmt.Printf("  Landlord Address: %s\n", landlord.Address)
		choice = 0
	}

	return nil
}
