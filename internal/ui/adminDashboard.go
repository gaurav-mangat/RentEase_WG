package ui

import (
	"fmt"
	"rentease/internal/domain/entities"
	"rentease/pkg/utils"
)

var allUsers []entities.User

func (ui *UI) AdminDashboard() {
	users, err := ui.UserService.GetAllUsers()
	if err != nil {
		fmt.Printf("\033[1;31mError retrieving users: %v\033[0m\n", err) // Red
		return
	}

	allUsers = users

	for {
		fmt.Println("\n\033[1;34m╔════════════════════════════════════════╗\033[0m") // Blue border
		fmt.Println("\033[1;34m║            Admin Dashboard             ║\033[0m")   // Blue header
		fmt.Println("\033[1;34m╚════════════════════════════════════════╝\033[0m")   // Blue border
		fmt.Println()
		fmt.Println("		1. \033[1;36mView all users\033[0m")     // Cyan
		fmt.Println("		2. \033[1;36mDelete a user\033[0m")      // Cyan
		fmt.Println("		3. \033[1;36mApprove properties\033[0m") // Cyan
		fmt.Println("		4. \033[1;31mLogout\033[0m")             // Red for Logout

		fmt.Print("\n\033[1;33mEnter your choice: \033[0m") // Yellow
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
			fmt.Println("\033[1;32mLogout successful.\033[0m") // Green
			return
		default:
			fmt.Println("\033[1;31mInvalid choice. Please try again.\033[0m") // Red
		}
	}
}

func (ui *UI) ViewAllUsers() {
	fmt.Printf("\n\033[1;36mTotal Users: %d\033[0m\n", len(allUsers)-1) // Cyan

	saveAdminUsername := utils.ActiveUser
	i := 0
	for _, user := range allUsers {
		if user.Role != "Admin" {
			fmt.Printf("\n\033[1;34mUser #%d:\033[0m\n", i+1) // Blue
			ui.DisplayUserInfo(user)
			fmt.Println("\n\033[1;36mListed Properties:\033[0m") // Cyan

			utils.ActiveUser = user.Username
			ui.viewAndManageListedProperties(true)
			fmt.Println("\n\033[1;36m----------------------------------------------------------------\033[0m") // Cyan separator
			i++
		}
	}
	utils.ActiveUser = saveAdminUsername
}

func (ui *UI) DeleteUser() {
	for {
		username := utils.ReadInput("\n\033[1;33mEnter the username of the user to delete (or 0 to exit): \033[0m") // Yellow input
		if username == "0" {
			break
		}

		userFound := false
		for _, user := range allUsers {
			if user.Username == username && user.Role != "Admin" {
				userFound = true
				break
			}
		}

		if !userFound {
			fmt.Println("\033[1;31mUser with this username doesn't exist.\033[0m") // Red
			continue
		}

		err := ui.UserService.DeleteUser(username)
		if err != nil {
			fmt.Printf("\033[1;31mError deleting user: %v\033[0m\n", err) // Red
		} else {
			fmt.Println("\033[1;32mUser deleted successfully.\033[0m") // Green
			err = ui.PropertyService.DeleteAllListedPropertiesOfaUser(username)
			if err != nil {
				fmt.Printf("\033[1;31mError in deleting the properties of this user: %v\033[0m\n", err) // Red
			} else {
				fmt.Println("\033[1;32mProperties of this user also deleted successfully.\033[0m") // Green
			}
		}
	}
}

func (ui *UI) ApproveProperties() {
	properties, err := ui.PropertyService.GetPendingProperties()
	if err != nil {
		fmt.Printf("\033[1;31mError retrieving properties: %v\033[0m\n", err) // Red
		return
	}

	if len(properties) == 0 {
		fmt.Println("\033[1;33mNo properties to approve.\033[0m") // Yellow
		return
	}

	for {
		fmt.Println("\n\033[1;34m╔══════════════════════════════════════════════════╗\033[0m") // Blue border
		fmt.Println("\033[1;34m║         Pending Properties for Approval          ║\033[0m")   // Blue header
		fmt.Println("\033[1;34m╚══════════════════════════════════════════════════╝\033[0m")   // Blue border
		fmt.Println()
		ui.DisplayPropertyShortInfo(properties, nil)

		fmt.Print("\n\033[1;33mEnter 0 to go back, 1 to see more details, 2 to approve a property: \033[0m") // Yellow
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 0:
			return // Go back to the previous menu
		case 1:
			fmt.Print("\033[1;33mEnter the property number to see more details: \033[0m") // Yellow
			var propertyIndex int
			fmt.Scan(&propertyIndex)

			if propertyIndex < 1 || propertyIndex > len(properties) {
				fmt.Println("\033[1;31mInvalid property number.\033[0m") // Red
				continue
			}
			ui.displayPropertyDetails(properties[propertyIndex-1])

		case 2:
			fmt.Print("\033[1;33mEnter the property number to approve: \033[0m") // Yellow
			var propertyIndex int
			fmt.Scan(&propertyIndex)

			if propertyIndex < 1 || propertyIndex > len(properties) {
				fmt.Println("\033[1;31mInvalid property number.\033[0m") // Red
				continue
			}

			selectedProperty := properties[propertyIndex-1]
			err = ui.PropertyService.ApproveProperty(selectedProperty.ID, utils.ActiveUser) // `ActiveUser` is the admin who approves the property
			if err != nil {
				fmt.Printf("\033[1;31mError approving property: %v\033[0m\n", err) // Red
			} else {
				fmt.Println("\033[1;32mProperty approved successfully.\033[0m") // Green
				return                                                          // After approving, return to the previous menu
			}

		default:
			fmt.Println("\033[1;31mInvalid choice. Please try again.\033[0m") // Red
		}
	}
}
