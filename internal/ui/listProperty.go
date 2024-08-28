package ui

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"rentease/internal/domain/entities"
	"rentease/pkg/utils"
	"strconv"
	"strings"
)

// ListPropertyUI handles the input process for landlords to list a new property.
func (ui *UI) ListPropertyUI() {
	// Prompt user to select the property type
	fmt.Print("Enter property type (1. Commercial, 2. House, 3. Flat): ")
	var propertyType int
	fmt.Scanf("%d", &propertyType)

	// Collect property title from the user
	title := utils.ReadInput("\nEnter property title: ")

	// Collect the address details from the user
	fmt.Println("Please provide the address details of your property")
	area := utils.ReadInput("    Enter locality: ")
	city := utils.ReadInput("    Enter city: ")
	state := utils.ReadInput("    Enter state: ")
	pincode := utils.ReadPincode()

	// Retrieve active landlord's username
	landlordUsername := utils.ActiveUser

	// Collect the expected rent amount from the user
	input := utils.ReadInput("Enter your expected rent amount (in rupees): ")
	rentAmount, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatalf("Invalid input, please enter a valid number: %v", err)
	}

	// Determine property details based on the selected property type
	var details interface{}
	switch propertyType {
	case 1:
		// Collect Commercial-specific details
		details = ui.collectCommercialDetails()
	case 2:
		// Collect House-specific details
		details = ui.collectHouseDetails()
	case 3:
		// Collect Flat-specific details
		details = ui.collectFlatDetails()
	default:
		fmt.Println("Invalid property type")
		return
	}

	// Create a new property entity with collected details
	property := entities.Property{
		ID:                primitive.NewObjectID(), // Generate a new unique ID
		PropertyType:      propertyType,
		Title:             title,
		RentAmount:        rentAmount,
		Address:           entities.Address{Area: area, City: city, State: state, Pincode: pincode},
		LandlordUsername:  landlordUsername,
		IsRented:          false,
		IsApprovedByAdmin: false,
		Details:           details,
	}

	// Save the property to the repository
	err = ui.PropertyService.ListProperty(property)
	if err != nil {
		fmt.Println("\nError listing property:", err)
	} else {
		fmt.Println("\nProperty listed successfully.")
	}
}

// collectCommercialDetails collects details specific to a Commercial property type.
func (ui *UI) collectCommercialDetails() entities.CommercialDetails {
	// Collect floor area for the commercial property
	floorArea := utils.ReadInput("Enter floor area(in sq. ft): ")

	// Prompt user to select the commercial property subtype
	var subType string
	var subTypeInt int
	fmt.Print("Enter subtype (1. Shop, 2. Factory, 3. Warehouse): ")
	_, _ = fmt.Scanf("%d", &subTypeInt)
	switch subTypeInt {
	case 1:
		subType = "Shop"
	case 2:
		subType = "Factory"
	case 3:
		subType = "Warehouse"
	}

	return entities.CommercialDetails{
		FloorArea: floorArea,
		SubType:   subType,
	}
}

// collectHouseDetails collects details specific to a House property type.
func (ui *UI) collectHouseDetails() entities.HouseDetails {
	// Collect the number of rooms for the house
	noOfRoomsInput := utils.ReadInput("Enter number of rooms: ")
	noOfRooms, err := strconv.Atoi(noOfRoomsInput)
	if err != nil {
		fmt.Println(err)
		return entities.HouseDetails{}
	}

	// Prompt user to select the furnished category and collect amenities
	furnishedCategory := ui.FurnishedTypeInput()
	amenitiesStr := utils.ReadInput("Enter amenities (comma separated): ")
	amenities := strings.Split(amenitiesStr, ",")

	return entities.HouseDetails{
		NoOfRooms:         noOfRooms,
		FurnishedCategory: furnishedCategory,
		Amenities:         amenities,
	}
}

// collectFlatDetails collects details specific to a Flat property type.
func (ui *UI) collectFlatDetails() entities.FlatDetails {
	// Prompt user to select the furnished category and collect amenities
	furnishedCategory := ui.FurnishedTypeInput()
	amenitiesStr := utils.ReadInput("Enter amenities (comma separated): ")
	amenities := strings.Split(amenitiesStr, ",")

	// Collect the number of BHK for the flat
	bhk, _ := utils.ReadBHKInput()

	return entities.FlatDetails{
		FurnishedCategory: furnishedCategory,
		Amenities:         amenities,
		BHK:               bhk,
	}
}

// FurnishedTypeInput prompts the user to select a furnished category and returns the selection.
func (ui *UI) FurnishedTypeInput() string {
	var furnishedType string
	var furnishedTypeInt int
	for {
		fmt.Print("Enter Furnished type (1. Unfurnished, 2. Semi Furnished, 3. Fully Furnished): ")
		_, _ = fmt.Scanf("%d", &furnishedTypeInt)
		if furnishedTypeInt < 1 || furnishedTypeInt > 3 {
			fmt.Println("Invalid Furnished type")
			continue
		}
		switch furnishedTypeInt {
		case 1:
			furnishedType = "Unfurnished"
		case 2:
			furnishedType = "Semi Furnished"
		case 3:
			furnishedType = "Fully Furnished"
		}
		return furnishedType
	}
}
