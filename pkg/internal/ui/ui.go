package ui

import (
	"rentease/internal/app/services"
)

// UI struct holds the UserService, bufio.Reader, and other dependencies
type UI struct {
	userService     *services.UserService
	propertyService *services.PropertyService
	requestService  *services.RequestService
}

// NewUI initializes the UI with the provided services and a bufio.Reader
func NewUI(userService *services.UserService, propertyService *services.PropertyService, requestService *services.RequestService) *UI {
	return &UI{
		userService:     userService,
		propertyService: propertyService,
		requestService:  requestService,
	}
}
