package ui

import (
	"fmt"
	"rentease/internal/domain/entities"
	"rentease/pkg/utils"
	"rentease/pkg/validation"
)

func (ui *UI) SignUpDashboard() {
	// Display signup form and collect user data
	fmt.Println()
	fmt.Println("\033[1;36m----------------------------------------------------------------\033[0m")    // Sky blue
	fmt.Println("\033[1;31m                       SIGN UP FORM                                \033[0m") // Red bold
	fmt.Println("\033[1;36m----------------------------------------------------------------\033[0m")
	fmt.Println()

	// Get username
	username := ui.promptForUsername()

	// Get password
	password := ui.promptForPassword()

	// Get full name
	fullName := utils.ReadInput("\033[1;34m\nEnter full name: \033[0m")

	// Get age
	age := ui.promptForAge()
	if age == 0 {
		return
	}

	// Get mobile number
	mobileNumber := ui.promptForMobileNumber()

	// Get email
	email := ui.promptForEmail()

	address := utils.ReadInput("\033[1;34m\nEnter your address : \033[0m")

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

	// Call userService to save user
	if err := ui.UserService.SignUp(user); err != nil {
		fmt.Printf("\033[1;31mError signing up: %v\033[0m\n", err)
		return
	}

	fmt.Println("\033[1;32m\n\nUser signed up successfully!\033[0m") // Green bold
	ui.promptPostSignupActions()
}

func (ui *UI) promptForUsername() string {
	var username string
	valid := false
	for !valid {
		username = utils.ReadInput("\033[1;34mEnter username (Username should only be a single word): \033[0m")
		if validation.IsInputSpaceFree(username) {
			valid = true
		} else {
			fmt.Println("\033[1;31mInvalid username.\nPlease enter a valid username.\n\033[0m")
		}
		user, _ := ui.UserService.FindByUsername(username)
		if user.Username != "" {
			fmt.Println("This username already exists.")
			valid = false
		}
	}
	return username
}

func (ui *UI) promptForPassword() string {
	var password string
	valid := false
	for !valid {
		password = utils.ReadInput("\033[1;34m\nEnter password (min 9 chars, include lowercase, uppercase, numbers, special): \033[0m")
		if validation.IsInputSpaceFree(password) && utils.IsValidPassword(password) {
			valid = true
		} else {
			fmt.Println("\033[1;31m\nPassword does not meet complexity requirements.\nPlease enter a valid password.\033[0m")
		}
	}
	return password
}

func (ui *UI) promptForAge() int {
	var age int
	valid := false
	for !valid {
		fmt.Print("\u001B[1;34m\nEnter your age: \033[0m")
		fmt.Scanln(&age)
		if age >= 18 && age <= 125 {
			valid = true
		} else if age > 0 && age < 18 {
			fmt.Println("\033[1;31m\nYou are not eligible to create an account.\033[0m")
			return 0
		} else {
			fmt.Println("\033[1;31m\nInvalid age.\nPlease enter a valid age.\033[0m")
		}
	}
	return age
}

func (ui *UI) promptForMobileNumber() string {
	var mobileNumber string
	valid := false
	for !valid {
		mobileNumber = utils.ReadInput("\033[1;34m\nEnter mobile number: \033[0m")
		if validation.IsInputSpaceFree(mobileNumber) && validation.IsValidMobileNumber(mobileNumber) {
			valid = true
		} else {
			fmt.Println("\033[1;31m\nInvalid mobile number.\nPlease enter a 10-digit number starting with 6, 7, 8, or 9.\033[0m")
		}
	}
	return mobileNumber
}

func (ui *UI) promptForEmail() string {
	var email string
	valid := false
	for !valid {
		email = utils.ReadInput("\033[1;34m\nEnter your email ID: \033[0m")
		if validation.IsInputSpaceFree(email) && validation.IsValidEmail(email) {
			valid = true
		} else {
			fmt.Println("\033[1;31m\nPlease enter a valid email.\033[0m")
		}
	}
	return email
}

func (ui *UI) promptPostSignupActions() {
	fmt.Println("\n\nPress 1 to Login \nPress 2 to Exit")
	var choice int
	fmt.Print("\033[1;34m\nEnter your choice: \033[0m") // Blue bold
	_, err := fmt.Scan(&choice)
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
}
