package main

import (
	"fmt"
	"rentease/config"
	"rentease/internal/app/repositories"
	"rentease/internal/app/services"
	"rentease/internal/ui"
)

func main() {

	// Initializing user repo and user service
	userRepo, _ := repositories.NewUserRepo(config.USER_URI, config.DATABASE, config.USER_COLLECTION)
	userService := services.NewUserService(userRepo)

	// Initializing property repo and property service
	propertyRepo, err := repositories.NewPropertyRepo(config.PROPERTIES_URI, config.DATABASE, config.PROPERTIES_COLLECTION)
	if err != nil {
		fmt.Println("Error initializing repository:", err)
		return
	}
	propertyService := services.NewPropertyService(propertyRepo)

	// Initializing rent request repo and rent request service
	rentRequestRepo, err := repositories.NewRequestRepo(config.RENT_REQUEST_URI, config.DATABASE, config.RENT_REQUEST_COLLECTION)
	rentRequestService := services.NewRequestService(rentRequestRepo)

	appUI := ui.NewUI(userService, propertyService, rentRequestService)

	// Calling the AppDashboard
	appUI.AppDashboard()

}
