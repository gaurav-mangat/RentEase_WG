package ui

import (
	"rentease/internal/app/services"
)

// UI struct holds the UserService, PropertyService and RequestService
type UI struct {
	UserService     *services.UserService
	PropertyService *services.PropertyService
	RequestService  *services.RequestService
}

// NewUI initializes the UI with the provided services
func NewUI(userService *services.UserService, propertyService *services.PropertyService, requestService *services.RequestService) *UI {
	return &UI{
		UserService:     userService,
		PropertyService: propertyService,
		RequestService:  requestService,
	}
}
