package ui

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
	"rentease/internal/domain/entities"
	"rentease/pkg/utils"
	"strconv"
)

var allUsers []entities.User

func (ui *UI) AdminDashboard() {
	for {
		fmt.Println("\n\033[1;34m╔════════════════════════════════════════╗\033[0m") // Blue border
		fmt.Println("\033[1;34m║            Admin Dashboard             ║\033[0m")   // Blue header
		fmt.Println("\033[1;34m╚════════════════════════════════════════╝\033[0m")   // Blue border
		fmt.Println()
		fmt.Println("	1. \033[1;36mView all users\033[0m")     // Cyan
		fmt.Println("	2. \033[1;36mDelete a user\033[0m")      // Cyan
		fmt.Println("	3. \033[1;36mApprove properties\033[0m") // Cyan
		fmt.Println("	4. \033[1;36mLogout\033[0m")             // Red for Logout

		// Read and convert choice input
		choiceTemp := utils.ReadInput("\n\033[1;33mEnter your choice: \033[0m")
		choice, err := strconv.Atoi(choiceTemp)
		if err != nil {
			fmt.Println("\033[1;31mInvalid input, please enter a number.\033[0m")
			continue
		}

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

	//Fetching all users from the database
	users, err := ui.UserService.GetAllUsers()
	if err != nil {
		fmt.Printf("\033[1;31mError retrieving users: %v\033[0m\n", err) // Red
		return
	}

	allUsers = users

	fmt.Printf("\n\033[1;36m                                                        Total Users: %d\033[0m\n", len(allUsers)-1) // Cyan
	fmt.Println()
	// Create a new table
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Username", "Full Name", "Age", "Address", "Phone Number", "Email"})

	table.SetColWidth(40)

	// Enable auto wrapping of text within the cells
	table.SetAutoWrapText(false)
	// Populate the table with user data
	for _, user := range allUsers {
		if user.Role != "Admin" {
			table.Append([]string{user.Username, user.Name, fmt.Sprintf("%d", user.Age), user.Address, user.PhoneNumber, user.Email})
		}
	}

	// Render the table
	table.SetBorder(true)
	table.Render()

	// Gather all properties for all users
	var allProperties []entities.Property
	activeUserRecord := utils.ActiveUser
	for _, user := range allUsers {
		if user.Role != "Admin" {
			// Get properties for each user
			utils.ActiveUser = user.Username
			properties, err := ui.PropertyService.GetAllListedProperties(true)
			if err != nil {
				fmt.Printf("\033[1;31mError retrieving properties for user %s: %v\033[0m\n", user.Username, err)
				continue
			}
			// Append to the list of all properties
			allProperties = append(allProperties, properties...)
		}
	}
	utils.ActiveUser = activeUserRecord

	fmt.Println()
	// Display all properties in a single table
	fmt.Println("			\033[1;34m╔════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╗\033[0m") // Blue border
	fmt.Println("			\033[1;34m║                                               All Properties                                                               ║\033[0m") // Blue header
	fmt.Println("			\033[1;34m╚════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╝\033[0m") // Blue border

	fmt.Println()
	// Create a new table for properties
	propertyTable := tablewriter.NewWriter(os.Stdout)
	propertyTable.SetHeader([]string{"Username", "Type", "Title", "Address", "Rent Amount", "Is Approved By Admin", "Details"})

	// Set minimum width for the "Address" and "Details" columns
	propertyTable.SetColMinWidth(3, 50) // Set minimum width for the "Address" column
	propertyTable.SetColMinWidth(6, 50) // Set minimum width for the "Details" column

	// Set the column alignment
	propertyTable.SetColumnAlignment([]int{
		tablewriter.ALIGN_LEFT,  // Username
		tablewriter.ALIGN_LEFT,  // Type
		tablewriter.ALIGN_LEFT,  // Title
		tablewriter.ALIGN_LEFT,  // Address
		tablewriter.ALIGN_RIGHT, // Rent Amount
		tablewriter.ALIGN_LEFT,  // Is Approved By Admin
		tablewriter.ALIGN_LEFT,  // Details
	})

	// Enable auto-wrapping of text within the cells
	propertyTable.SetAutoWrapText(false)

	// Populate the table with property data
	for _, property := range allProperties {
		address := fmt.Sprintf("%s, %s, %s - %d", property.Address.Area, property.Address.City, property.Address.State, property.Address.Pincode)
		propertyTable.Append([]string{
			property.LandlordUsername,
			fmt.Sprintf("%d", property.PropertyType),
			property.Title,
			address,
			fmt.Sprintf("%.2f", property.RentAmount),
			fmt.Sprintf("%t", property.IsApprovedByAdmin),
			fmt.Sprintf("%v", property.Details),
		})
	}

	// Render the property table
	propertyTable.SetBorder(true)
	propertyTable.Render()
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
		ui.DisplayPropertyShortInfo(properties)

		choiceTemp := utils.ReadInput("\n\033[1;33mEnter 0 to go back, 1 to see more details, 2 to approve a property: \033[0m")
		choice, err := strconv.Atoi(choiceTemp)
		if err != nil {
			fmt.Println("\033[1;31mInvalid input, please enter a number.\033[0m")
			continue
		}

		switch choice {
		case 0:
			return // Go back to the previous menu
		case 1:
			var propertyIndex int
			propertyIndexTemp := utils.ReadInput("\033[1;33mEnter the property number to see more details: \033[0m")
			propertyIndex, _ = strconv.Atoi(propertyIndexTemp)

			if propertyIndex < 1 || propertyIndex > len(properties) {
				fmt.Println("\033[1;31mInvalid property number.\033[0m") // Red
				continue
			}
			ui.displayPropertyDetails(properties[propertyIndex-1])

		case 2:
			var propertyIndex int
			propertyIndexTemp := utils.ReadInput("\033[1;33mEnter the property number to approve: \033[0m")
			propertyIndex, _ = strconv.Atoi(propertyIndexTemp)

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
				properties, _ = ui.PropertyService.GetPendingProperties()
			}
		}
	}
}
