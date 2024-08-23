package main

import (
	"fmt"
	"rentease/internal/app/repositories"
	"rentease/internal/app/services"
	"rentease/internal/ui"
)

func main() {
	userRepo, _ := repositories.NewUserRepo("mongodb://localhost:27017/users", "RentEase", "users")
	userService := services.NewUserService(userRepo)

	// Initializing property repoo and property service
	propertyRepo, err := repositories.NewPropertyRepo("mongodb://localhost:27017/properties", "RentEase", "properties")
	if err != nil {
		fmt.Println("Error initializing repository:", err)
		return
	}
	propertyService := services.NewPropertyService(propertyRepo)

	requestRepo, err := repositories.NewRequestRepo("mongodb://localhost:27017/request", "RentEase", "request")
	requestService := services.NewRequestService(requestRepo)
	appUI := ui.NewUI(userService, propertyService, requestService)
	appUI.AppDashboard()
	// Calling application dashboard

	//username := "johndoe"
	//user, err := userService.Findbyuname(username)
	//if err != nil {
	//	log.Fatalf("Failed to find user: %v", err)
	//}
	//
	//if (user == entities.User{}) {
	//	fmt.Println("User not found.")
	//} else {
	//	fmt.Printf("Found user: %+v\n", user)
	//}

}
