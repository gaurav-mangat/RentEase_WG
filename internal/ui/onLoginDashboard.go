package ui

import (
	"fmt"
)

func (ui *UI) onLoginDashboard(username string) {
	for {
		// Display the dashboard
		fmt.Println("\033[1;36m----------------------------------------------------------------\033[0m") // Sky blue
		fmt.Println("\033[1;31m                          DASHBOARD                            \033[0m")  // Red bold
		fmt.Println("\033[1;36m----------------------------------------------------------------\033[0m") // Sky blue

		fmt.Println("\033[1;32m1. LandLord Section\033[0m") // Green
		fmt.Println("\033[1;32m2. Tenant Section\033[0m")   // Green
		fmt.Println("\033[1;31m3. Logout\033[0m")           // Red
		fmt.Print("\nEnter your choice: ")

		var choice int
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Printf("\033[1;31mError reading input: %v\033[0m\n", err) // Red
			continue
		}

		switch choice {
		case 1:
			fmt.Println("\033[1;33m\nYou are now in the Landlord Section.\033[0m") // Yellow
			// Add functionality for Seller Section here
			ui.sellerDashboard()
		case 2:
			fmt.Println("\033[1;33m\nYou are now in the Tenant Section.\033[0m") // Yellow
			// Add functionality for Buyer Section here
			ui.TenantDashboardUI()
		//case 3:
		//	fmt.Println("\033[1;33m\nHere is your profile information.\033[0m") // Yellow
		// Add functionality to view profile here
		case 3:
			fmt.Println("\033[1;32m\nYou have been logged out.\033[0m") // Green
			return
		default:
			fmt.Println("\033[1;31m\nInvalid choice. Please select a valid option.\033[0m") // Red
		}
	}
}
