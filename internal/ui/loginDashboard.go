package ui

import (
	"fmt"
	"rentease/pkg/utils"
	"rentease/pkg/validation"
	"strconv"
)

// LoginDashboard manages user login attempts with options to retry, sign up, or exit.
func (ui *UI) LoginDashboard() {

	const maxAttempts = 3
	attemptsLeft := maxAttempts

	for attemptsLeft > 0 {

		fmt.Println()
		fmt.Println("\033[1;36m----------------------------------------------------------------\033[0m") // Sky blue
		fmt.Println("\033[1;35m                          LOG IN                                \033[0m") // Red bold
		fmt.Println("\033[1;36m----------------------------------------------------------------\033[0m")

		var username string
		for {
			username = utils.ReadInput("             Enter username: ")
			if validation.IsSingleWordUsername(username) {
				break
			}
		}

		password, err := utils.GetHiddenInput("             Enter password: ")
		if err != nil {
			fmt.Println("\033[1;31mError during getting hidden password:\033[0m", err)
		}

		// Check credentials
		loginSuccessful, _ := ui.UserService.Login(username, password)

		if loginSuccessful {
			fmt.Println("\033[1;32mLogin successful!\n\n\033[0m") // Green

			// Checking for admin
			user, err := ui.UserService.FindByUsername(username)
			if err != nil {
				fmt.Println("\033[1;31mError finding user :\033[0m", err)
				return
			}

			utils.ActiveUserobject = user

			if user.Role == "Admin" {
				ui.AdminDashboard()
				return
			} else {
				ui.onLoginDashboard()
				return
			}
		} else {
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
			choiceTemp := utils.ReadInput("Enter your choice: ")
			choice, _ = strconv.Atoi(choiceTemp)
			switch choice {
			case 1:
				continue
			case 2:
				ui.SignUpDashboard()
				return
			case 3:
				fmt.Println("Exiting...")
				return
			default:
				fmt.Println("Invalid choice. Exiting...")
				return
			}
		}
	}
}
