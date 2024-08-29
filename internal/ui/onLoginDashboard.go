package ui

import (
	"fmt"
	"rentease/pkg/utils"
	"strconv"
)

func init() {

}
func (ui *UI) onLoginDashboard() {
	for {
		// Display the dashboard
		fmt.Println()
		fmt.Println("\033[1;36m-----------------------------------------------\033[0m")       // Sky blue
		fmt.Println("\033[1;35m                DASHBOARD                            \033[0m") // Red bold
		fmt.Println("\033[1;36m-----------------------------------------------\033[0m")       // Sky blue

		fmt.Println("\033[1;32m	1. LandLord Section\033[0m") // Green
		fmt.Println("\033[1;32m	2. Tenant Section\033[0m")   // Green
		fmt.Println("\033[1;32m	3. View Profile\033[0m")     // Green
		fmt.Println("\033[1;31m	4. Logout\033[0m")           // Red

		var choice int
		choiceTemp := utils.ReadInput("\nEnter your choice: ")
		choice, err := strconv.Atoi(choiceTemp)
		if err != nil {
			fmt.Printf("\033[1;31mError reading input: %v\033[0m\n", err) // Red
			continue
		}

		switch choice {
		case 1:
			// Moving to Landlord Section from here
			ui.landlordDashboard()

		case 2:
			// Moving to Tenant Section from here
			ui.TenantDashboardUI()

		case 3:
			// View the logged in user profile
			ui.userProfile()

		case 4:
			// Logging out of the account
			fmt.Println("\033[1;32m\nYou have been logged out.\033[0m") // Green
			return

		default:
			fmt.Println("\033[1;31m\nInvalid choice. Please select a valid option.\033[0m") // Red
		}
	}
}
