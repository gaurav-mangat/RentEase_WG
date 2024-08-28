package main

import (
	"fmt"
	"rentease/internal/app/repositories"
	"rentease/internal/app/services"
	"rentease/internal/ui"
)

func main() {

	// Initializing user repo and user service
	userRepo, _ := repositories.NewUserRepo("mongodb://localhost:27017/users", "RentEase", "users")
	userService := services.NewUserService(userRepo)

	// Initializing property repo and property service
	propertyRepo, err := repositories.NewPropertyRepo("mongodb://localhost:27017/properties", "RentEase", "properties")
	if err != nil {
		fmt.Println("Error initializing repository:", err)
		return
	}
	propertyService := services.NewPropertyService(propertyRepo)

	// Initializing rent request repo and rent request service
	requestRepo, err := repositories.NewRequestRepo("mongodb://localhost:27017/request", "RentEase", "request")
	requestService := services.NewRequestService(requestRepo)

	appUI := ui.NewUI(userService, propertyService, requestService)

	// Calling the AppDashboard
	appUI.AppDashboard()

}
