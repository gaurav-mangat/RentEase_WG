package ui

import (
	"fmt"
	"rentease/internal/domain/entities"
	"rentease/pkg/utils"
)

func (ui *UI) userProfile() {

	// Display the dashboard
	fmt.Println("\033[1;36m----------------------------------------------------------------\033[0m")  // Sky blue
	fmt.Println("\033[1;31m                          YOUR PROFILE                           \033[0m") // Red bold
	fmt.Println("\033[1;36m----------------------------------------------------------------\033[0m")
	fmt.Println()
	ui.DisplayUserInfo(utils.ActiveUserobject)
	fmt.Println("\n\033[1;36m----------------------------------------------------------------\033[0m")
	fmt.Println("\nYour listed properties are :")
	ui.viewAndManageListedProperties(true)

}

func (ui *UI) DisplayUserInfo(ActiveUserobject entities.User) {
	fmt.Println("			Username     : ", ActiveUserobject.Username)
	fmt.Println("			Full Name    : ", ActiveUserobject.Name)
	fmt.Println("			Age          : ", ActiveUserobject.Age)
	fmt.Println("			Address      : ", ActiveUserobject.Address)
	fmt.Println("			Phone Number : ", ActiveUserobject.PhoneNumber)
	fmt.Println("			Email        : ", ActiveUserobject.Email)

}
