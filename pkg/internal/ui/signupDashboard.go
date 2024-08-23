package ui

import (
	"fmt"
	"rentease/internal/domain/entities"
	"rentease/pkg/utils"
)

func (ui *UI) SignUpDashboard() {

	// Signup form
	fmt.Println()
	fmt.Println("\033[1;36m----------------------------------------------------------------\033[0m")    // Sky blue
	fmt.Println("\033[1;31m                       SIGN UP FORM                                \033[0m") // Red bold
	fmt.Println("\033[1;36m----------------------------------------------------------------\033[0m")
	fmt.Println()

	// Get username
	var username string
	valid := false
	for !valid {
		username = utils.ReadInput("\033[1;34mEnter username (Username should only be a single word): \033[0m")
		if utils.IsValidInput(username) {
			valid = true
		} else {
			fmt.Println("\033[1;31mInvalid username.\nPlease enter a valid username.\n\033[0m")
		}
		user, _ := ui.userService.FindByUsername(username)
		if user.Username != "" {
			fmt.Println("This username already exists.")
			valid = false
		}

	}

	// Get password

	var password string
	valid = false
	for !valid {
		password = utils.ReadInput("\033[1;34m\nEnter password (min 9 chars, include lowercase, uppercase, numbers, special): \033[0m")
		if utils.IsValidInput(password) && utils.IsValidPassword(password) {
			valid = true
		} else {
			fmt.Println("\033[1;31m\nPassword does not meet complexity requirements.\nPlease enter a valid password.\033[0m")
		}
	}

	// Get full name

	var fullName string

	fullName = utils.ReadInput("\033[1;34m\nEnter full name: \033[0m")

	// Get age
	var age int
	valid = false
	for !valid {
		fmt.Print("\u001B[1;34m\nEnter you age : \033[0m")
		fmt.Scanln(&age)
		if age >= 18 && age <= 125 {
			valid = true
		} else if age > 0 && age < 18 {
			fmt.Println("\033[1;31m\nYou are not eligible to create account.\033[0m")
			return
		} else {
			fmt.Println("\033[1;31m\nInvalid age.\nPlease enter a valid age.\033[0m")

		}
	}

	// Get mobile number
	var mobileNumber string
	valid = false
	for !valid {
		mobileNumber = utils.ReadInput("\033[1;34m\nEnter mobile number: \033[0m")
		if utils.IsValidInput(mobileNumber) && utils.IsValidMobileNumber(mobileNumber) {
			valid = true
		} else {
			fmt.Println("\033[1;31m\nInvalid mobile number.\nPlease enter a 10-digit number starting with 6, 7, 8, or 9.\033[0m")
		}
	}

	// Get email
	var email string
	valid = false
	for !valid {
		email = utils.ReadInput("\033[1;34m\nEnter your email ID: \033[0m")
		if utils.IsValidInput(email) && utils.IsValidEmail(email) {
			valid = true
		} else {
			fmt.Println("\033[1;31m\nPlease enter a valid email.\033[0m")
		}
	}

	address := utils.ReadInput("Enter your address : ")

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		fmt.Printf("\033[1;31m\nError hashing password: %v\033[0m\n", err) // Red bold
		return
	}

	user := entities.User{
		Username:     username,
		PasswordHash: hashedPassword,
		Name:         fullName,
		PhoneNumber:  mobileNumber,
		Age:          age,
		Email:        email,
		Address:      address,
		Role:         "User",
	}

	// Adding user to database
	isSignupSuccessful := ui.userService.SignUp(user)

	if isSignupSuccessful {

		fmt.Println("\033[1;32m\n\nUser signed up successfully!\033[0m") // Green bold
		fmt.Println("\n\nPress 1 to Login \nPress 2 to Exit")
		var choice int
		fmt.Print("\033[1;34m\nEnter your choice: \033[0m") // Blue bold
		_, err = fmt.Scan(&choice)
		if err != nil {
			fmt.Printf("\033[1;31mError reading choice: %v\033[0m\n", err) // Red bold
			return
		}

		switch choice {
		case 1:
			ui.LoginDashboard()
		case 2:
			return
		default:
			fmt.Println("\033[1;31mInvalid choice\033[0m") // Red bold
		}
	} else {
		fmt.Println("Signup was unsuccessful. Please try again.")
	}
}
