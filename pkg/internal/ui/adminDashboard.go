package ui

import (
	"fmt"
	"rentease/pkg/utils"
)

func (ui *UI) AdminDashboard() {
	for {
		fmt.Println("\n\033[1;34mAdmin Dashboard\033[0m") // Blue
		fmt.Println("1. View all users")
		fmt.Println("2. Delete a user")
		fmt.Println("3. Approve properties")
		fmt.Println("4. Exit")

		fmt.Print("Enter your choice: ")
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			ui.ViewAllUsers()
		case 2:
			ui.DeleteUser()
		case 3:
			ui.ApproveProperties()
		case 4:
			return
		default:
			fmt.Println("\033[1;31mInvalid choice. Please try again.\033[0m") // Red
		}
	}
}

func (ui *UI) ViewAllUsers() {
	users, err := ui.userService.GetAllUsers()
	if err != nil {
		fmt.Printf("\033[1;31mError retrieving users: %v\033[0m\n", err) // Red
		return
	}

	fmt.Println("\n\033[1;34mAll Users\033[0m") // Blue
	for _, user := range users {
		fmt.Printf("Username: %s, Email: %s, Role: %s\n", user.Username, user.Email, user.Role)
	}
}

func (ui *UI) DeleteUser() {
	fmt.Print("Enter the username of the user to delete: ")
	var username string
	fmt.Scan(&username)

	err := ui.userService.DeleteUser(username)
	if err != nil {
		fmt.Printf("\033[1;31mError deleting user: %v\033[0m\n", err) // Red
	} else {
		fmt.Println("\033[1;32mUser deleted successfully.\033[0m") // Green
	}
}
func (ui *UI) ApproveProperties() {
	properties, err := ui.propertyService.GetPendingProperties()
	if err != nil {
		fmt.Printf("\033[1;31mError retrieving properties: %v\033[0m\n", err) // Red
		return
	}

	if len(properties) == 0 {
		fmt.Println("\033[1;33mNo properties to approve.\033[0m") // Yellow
		return
	}

	fmt.Println("\n\033[1;34mPending Properties\033[0m") // Blue
	for i, property := range properties {
		fmt.Printf("%d. Title: %s, Address: %s ,%s ,%s ,%d\n", i+1, property.Title, property.Address.Area, property.Address.City,
			property.Address.State, property.Address.Pincode)
	}

	fmt.Print("\nEnter the number of the property to approve (or 0 to exit): ")
	var choice int
	fmt.Scan(&choice)

	if choice == 0 {
		return
	}

	if choice < 1 || choice > len(properties) {
		fmt.Println("\033[1;31mInvalid property number.\033[0m") // Red
		return
	}

	selectedProperty := properties[choice-1]
	err = ui.propertyService.ApproveProperty(selectedProperty.ID, utils.ActiveUser) // `ActiveUser` is the admin who approves the property
	if err != nil {
		fmt.Printf("\033[1;31mError approving property: %v\033[0m\n", err) // Red
	} else {
		fmt.Println("\033[1;32mProperty approved successfully.\033[0m") // Green
	}
}
