package ui

import (
	"fmt"
)

func (ui *UI) AppDashboard() {

	// Decorative border with ANSI colors
	fmt.Println("\033[1;36m******************************************\033[0m")
	fmt.Println("\033[1;36m*                                        *\033[0m")
	fmt.Println("\033[1;36m*            \033[1;32mWelcome to RentEase\033[1;36m         *\033[0m")
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
	fmt.Println("\033[1;34m 3. \033[1;31m❌ Exit\033[0m")
	fmt.Println("\033[1;34m==================================\033[0m")
	fmt.Println()

	var choice int

	fmt.Print("\033[1;35mEnter your choice: \033[0m")
	_, err := fmt.Scanf("%d", &choice)
	if err != nil {
		fmt.Println("\033[1;31m⚠️ Error reading input:\033[0m", err)
		return
	}

	switch choice {
	case 1:
		ui.LoginDashboard()
	case 2:
		ui.SignUpDashboard()
	case 3:
		fmt.Println("\n\033[1;32mThank you for using RentEase! See you next time. 👋\033[0m")
		return
	default:
		fmt.Println("\n\033[1;31m🚫 Invalid choice, please try again.\033[0m")
	}
}
