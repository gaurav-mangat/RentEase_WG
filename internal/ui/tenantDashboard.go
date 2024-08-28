package ui

import (
	"fmt"
	"log"
	"rentease/pkg/utils"
)

func (ui *UI) TenantDashboardUI() {
	for {
		fmt.Println("\033[1;34m\n========================\033[0m") // Blue
		fmt.Println("\033[1;34m   Tenant Dashboard\033[0m")        // Blue
		fmt.Println("\033[1;34m========================\033[0m")   // Blue
		fmt.Println("1. Search Property")
		fmt.Println("2. Your Wishlist")
		fmt.Println("3. Your Rent Requests' Status")
		fmt.Println("4. Go Back")

		choice := utils.ReadInput("\nEnter your choice: ")

		switch choice {
		case "1":
			ui.SearchPropertyUI()
		case "2":
			err := ui.ShowWishlist()
			if err != nil {
				log.Println("Error in showing Wishlist : ", err)
			}

		case "3":
			ui.ShowNotifications()

		case "4":
			fmt.Println("\033[1;32mLogging out...\033[0m") // Green
			return
		default:
			fmt.Println("\033[1;31mInvalid choice, please try again.\033[0m") // Red
		}
	}
}
