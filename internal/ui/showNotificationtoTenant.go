package ui

import (
	"fmt"
	"log"
	"rentease/internal/domain/entities"
	"rentease/pkg/utils"
)

func (ui *UI) ShowNotifications() {
	requests, err := ui.RequestService.GetRentRequestsInfoForTenant(utils.ActiveUser)
	if err != nil {
		fmt.Printf("\033[1;31mError retrieving notifications: %v\033[0m\n", err) // Red
		return
	}

	if len(requests) == 0 {
		fmt.Println("\033[1;33mNo notifications.\033[0m") // Yellow
		return
	}

	fmt.Println("\n\033[1;34mYour Property Requests\033[0m") // Blue
	var properties []entities.Property
	for _, req := range requests {

		property, err := ui.PropertyService.FindByID(req.PropertyID)
		if err != nil {
			log.Println("\033[1;31mError finding property by id: \033[0m\n", err)
		}
		properties = append(properties, property)

	}
	ui.DisplayPropertyShortInfo(properties, requests)

}
