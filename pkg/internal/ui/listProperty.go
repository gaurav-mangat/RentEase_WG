package ui

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"rentease/internal/domain/entities"
	"rentease/pkg/utils"
	"strings"
)

// ListPropertyUI handles user input to create and list a property.
func (ui *UI) ListPropertyUI() {
	fmt.Print("Enter property type (1. Commercial, 2. House, 3. Flat): ")
	var propertyType int
	fmt.Scanf("%d", &propertyType)
	title := utils.ReadInput("\nEnter property title: ")

	// Collect common property details
	fmt.Println("Please provide the address details of your property :")
	area := utils.ReadInput("Enter locality: ")
	city := utils.ReadInput("Enter city: ")
	state := utils.ReadInput("Enter state: ")

	var pincode int
	pincode = utils.ReadPincode()
	landlordUsername := utils.ActiveUser

	fmt.Print("Enter your expected rent amount (in rupees): ")
	var rentAmount float64
	fmt.Scanf("%.2f", &rentAmount)

	var details interface{}
	switch propertyType {
	case 1:
		// Collect Commercial-specific details
		_ = utils.ReadInput("")
		floorArea := utils.ReadInput("Enter floor area: ")
		subType := utils.ReadInput("Enter subtype (shop, factory, warehouse): ")

		details = entities.CommercialDetails{
			FloorArea: floorArea,
			SubType:   subType,
		}
	case 2:
		// Collect House-specific details

		fmt.Print("Enter number of rooms: ")
		var noOfRooms int
		_, _ = fmt.Scanf("%d", &noOfRooms)

		furnishedCategory := utils.ReadInput("Enter furnished category: ")
		amenitiesStr := utils.ReadInput("Enter amenities (comma separated): ")
		amenities := strings.Split(amenitiesStr, ",")

		details = entities.HouseDetails{
			NoOfRooms:         noOfRooms,
			FurnishedCategory: furnishedCategory,
			Amenities:         amenities,
		}
	case 3:
		// Collect Flat-specific details
		_ = utils.ReadInput("")

		furnishedCategory := utils.ReadInput("Enter furnished category: ")
		amenitiesStr := utils.ReadInput("Enter amenities (comma separated): ")
		amenities := strings.Split(amenitiesStr, ",")

		fmt.Print("Enter BHK: ")
		var bhk int
		fmt.Scanf("%d", &bhk)

		details = entities.FlatDetails{
			FurnishedCategory: furnishedCategory,
			Amenities:         amenities,
			BHK:               bhk,
		}
	default:
		fmt.Println("Invalid property type")
		return
	}

	property := entities.Property{
		ID:                primitive.NewObjectID(), // Generate a new unique ID
		PropertyType:      propertyType,
		Title:             title,
		Address:           entities.Address{Area: area, City: city, State: state, Pincode: pincode},
		LandlordUsername:  landlordUsername,
		IsRented:          false,
		IsApprovedByAdmin: false,
		Details:           details,
	}

	// Save property to the repository
	err := ui.propertyService.ListProperty(property)
	if err != nil {
		fmt.Println("\nError listing property:", err)
	} else {
		fmt.Println("\nProperty listed successfully.")
	}
}
