package ui

import (
	"bufio"
	"fmt"
	"os"
	"rentease/pkg/utils"
	"rentease/pkg/validation"
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

		// Clear the input buffer
		reader := bufio.NewReader(os.Stdin)

		for {
			username = utils.ReadInput("\n             Enter username: ")
			if validation.IsSingleWordUsername(username) {
				break
			}
		}

		// Read and validate password
		password = utils.ReadInput("             Enter password: ")

		fmt.Println()

		// Check credentials
		loginSuccessful, _ := ui.UserService.Login(username, password)

		if loginSuccessful {
			// Successful login
			fmt.Println("\033[1;32mLogin successful!\n\n\033[0m") // Green

			//Checking for admin
			user, err := ui.UserService.FindByUsername(utils.ActiveUser)
			if err != nil {
				fmt.Println("\033[1;31mError finding user :\033[0m", err)
			}

			// It contains the active user struct
			utils.ActiveUserobject = user

			if user.Role == "Admin" {
				ui.AdminDashboard()
				return
			} else {
				ui.onLoginDashboard()
				return
			}
			// ye check karna hai

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
			_, err := fmt.Fscanf(reader, "%d\n", &choice) // Use Fscanf with "\n" to clear the buffer
			if err != nil {
				fmt.Println("\033[1;31mInvalid input.\033[0m")
				continue // Continue the loop to retry input
			}

			switch choice {
			case 1:
				// Retry login
				continue
			case 2:
				// Call the SignUp function
				ui.SignUpDashboard()
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
