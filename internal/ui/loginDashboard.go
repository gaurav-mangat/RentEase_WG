package ui

import (
	"fmt"
	"rentease/pkg/utils"
)

// LoginDashboard manages user login attempts with options to retry, sign up, or exit.
func (ui *UI) LoginDashboard() {
	const maxAttempts = 3
	attemptsLeft := maxAttempts

	for attemptsLeft > 0 {
		var username, password string

		fmt.Println()
		fmt.Println("\033[1;36m----------------------------------------------------------------\033[0m") // Sky blue
		fmt.Println("\033[1;31m                          LOG IN                                \033[0m") // Red bold
		fmt.Println("\033[1;36m----------------------------------------------------------------\033[0m")

		for {
			username = utils.ReadInput("\n             Enter username: ")
			if utils.IsValidInput2(username) {
				break
			}
		}

		// Read and validate password
		password = utils.ReadInput("             Enter password: ")

		fmt.Println()

		// Check credentials (this part needs to be implemented)
		loginSuccessful, _ := ui.userService.Login(username, password) // Placeholder for actual authentication logic

		if loginSuccessful {
			// Successful login
			fmt.Println("\033[1;32mLogin successful!\033[0m") // Green

			//Checking for admin
			user, err := ui.userService.FindByUsername(utils.ActiveUser)
			if err != nil {
				fmt.Println("\033[1;31mError finding user :\033[0m", err)
			}
			if user.Role == "Admin" {
				ui.AdminDashboard()
			} else {
				ui.onLoginDashboard(username)
			}
			ui.AppDashboard()
		} else {
			// Failed login, decrement attempts left
			attemptsLeft--
			if attemptsLeft == 0 {
				fmt.Println("\033[1;31mLogin failed. You have exhausted all attempts.\033[0m") // Red bold
				return
			}
			fmt.Printf("Login failed. You have %d attempt(s) left.\n", attemptsLeft)

			fmt.Println("\nWhat would you like to do next?")
			fmt.Println("1. Retry Login")
			fmt.Println("2. Sign up")
			fmt.Println("3. Exit")

			var choice int
			fmt.Print("Enter your choice: ")
			_, err := fmt.Scan(&choice)
			if err != nil {
				fmt.Println("\033[1;31mInvalid input.\033[0m")
				return
			}

			switch choice {
			case 1:
				// Retry login
				continue
			case 2:
				// Call the SignUp function
				//SignUp()
				return // Return to avoid retrying after sign up
			case 3:
				// Exit
				fmt.Println("Exiting...")
				return
			default:
				// Invalid choice
				fmt.Println("Invalid choice. Exiting...")
				return
			}
		}
	}
}
