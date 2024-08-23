package ui

import (
	"fmt"
	"rentease/pkg/utils"
)

func (ui *UI) ShowNotifications() {
	requests, err := ui.requestService.GetRequestsForTenant(utils.ActiveUser)
	if err != nil {
		fmt.Printf("\033[1;31mError retrieving notifications: %v\033[0m\n", err) // Red
		return
	}

	if len(requests) == 0 {
		fmt.Println("\033[1;33mNo notifications.\033[0m") // Yellow
		return
	}

	fmt.Println("\n\033[1;34mYour Property Requests\033[0m") // Blue
	for _, req := range requests {
		fmt.Printf("Property ID: %s, Status: %s\n", req.PropertyID.Hex(), req.RequestStatus)
		if req.RequestStatus == "accepted" {
			fmt.Println("Congratulations! Your request has been accpeted!")
		}
	}
}
