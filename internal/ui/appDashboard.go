package ui

import (
	"fmt"
	"rentease/pkg/utils"
	"strconv"
)

func (ui *UI) AppDashboard() {

	// Decorative border with ANSI colors
	fmt.Println()
	fmt.Println()
	fmt.Println("\033[1;36m******************************************\033[0m")
	fmt.Println("\033[1;36m*                                        *\033[0m")
	fmt.Println("\033[1;36m*         \033[1;32mWelcome to RentEase\033[1;36m            *\033[0m")
	fmt.Println("\033[1;36m*  \033[1;32mYour one-stop solution for renting!\033[1;36m   *\033[0m")
	fmt.Println("\033[1;36m*                                        *\033[0m")
	fmt.Println("\033[1;36m******************************************\033[0m")

	// Main Menu
	fmt.Println("\n\033[1;33mPlease select an option to proceed:\033[0m")
	fmt.Println("\033[1;34m==================================\033[0m")
	fmt.Println("\033[1;34m 1. \033[1;33m🔑 Log In\033[0m")
	fmt.Println("\033[1;34m-----------------------------\033[0m")
	fmt.Println("\033[1;34m 2. \033[1;33m📝 Sign Up\033[0m")
	fmt.Println("\033[1;34m-----------------------------\033[0m")
	fmt.Println("\033[1;34m 3. \033[1;31m❌  Exit\033[0m")
	fmt.Println("\033[1;34m==================================\033[0m")
	fmt.Println()

	// Get user's choice
	choiceTemp := utils.ReadInput("\033[1;35mEnter your choice: \033[0m")
	choice, err := strconv.Atoi(choiceTemp)
	if err != nil {
		fmt.Println("\033[1;31m🚫 Invalid input. Please enter a valid number.\033[0m")
		ui.AppDashboard() // Retry if input is not a valid integer
		return
	}

	// Process the user's choice
	switch choice {
	case 1:
		ui.LoginDashboard()
		ui.AppDashboard() // Show the dashboard again after login
	case 2:
		ui.SignUpDashboard()
		ui.AppDashboard() // Show the dashboard again after signup
	case 3:
		fmt.Println("\n\033[1;32mThank you for using RentEase! See you next time. 👋\033[0m")
		return
	default:
		fmt.Println("\n\033[1;31m🚫 Invalid choice, please try again.\033[0m")
		ui.AppDashboard() // Retry if the choice is out of range
	}
}
