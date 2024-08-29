package ui

import (
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io/ioutil"
	"log"
	"net/http"
	"rentease/internal/domain/entities"
	"rentease/pkg/utils"
	"strconv"
	"strings"
)

func (ui *UI) ListPropertyUI() {
	// Prompt user to select the property type
	var propertyType int
	propertyTypeTemp := utils.ReadInput("Enter property type (1. Commercial, 2. House, 3. Flat): ")
	propertyType, err := strconv.Atoi(propertyTypeTemp)

	// Collect property title from the user
	title := utils.ReadInput("\nEnter property title: ")

	// Collect the address details from the user
	fmt.Println("\nPlease provide the address details of your property")

	address, err := ui.GetAddress()
	if err != nil {
		fmt.Println("Failed to retrieve address:", err)
		// Handle error appropriately
	}

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

	// Retrieve active landlord's username
	landlordUsername := utils.ActiveUser

	// Create a new property entity with collected details
	property := entities.Property{
		ID:                primitive.NewObjectID(), // Generate a new unique ID
		PropertyType:      propertyType,
		Title:             title,
		RentAmount:        rentAmount,
		Address:           address,
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

// getAddressFromPincode fetches address details from the API based on the provided pincode.
func (ui *UI) getAddressFromPincode(pincode int) (entities.Address, error) {
	url := fmt.Sprintf("https://api.postalpincode.in/pincode/%d", pincode)
	resp, err := http.Get(url)
	if err != nil {
		return entities.Address{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return entities.Address{}, err
	}

	var apiResponse []struct {
		Message    string `json:"Message"`
		Status     string `json:"Status"`
		PostOffice []struct {
			Name     string `json:"Name"`
			District string `json:"District"`
			State    string `json:"State"`
			Pincode  string `json:"Pincode"`
		} `json:"PostOffice"`
	}

	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return entities.Address{}, err
	}

	if len(apiResponse) == 0 || len(apiResponse[0].PostOffice) == 0 {
		return entities.Address{}, fmt.Errorf("no address details found for pincode %d", pincode)
	}

	postOffice := apiResponse[0].PostOffice[0]
	return entities.Address{
		Area:    postOffice.Name,
		City:    postOffice.District,
		State:   postOffice.State,
		Pincode: pincode,
	}, nil
}

// collectCommercialDetails collects details specific to a Commercial property type.
func (ui *UI) collectCommercialDetails() entities.CommercialDetails {
	// Collect floor area for the commercial property
	floorArea := utils.ReadInput("Enter floor area(in sq. ft): ")

	// Prompt user to select the commercial property subtype
	var subType string
	var subTypeInt int
	subTypeTemp := utils.ReadInput("Enter subtype (1. Shop, 2. Factory, 3. Warehouse): ")
	subTypeInt, _ = strconv.Atoi(subTypeTemp)
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
	var noOfRooms int
	for {
		noOfRoomsInput := utils.ReadInput("Enter number of rooms: ")
		var err error
		noOfRooms, err = strconv.Atoi(noOfRoomsInput)
		if noOfRooms <= 0 || noOfRooms > 11 {
			fmt.Println("Invalid number of rooms")
			continue
		}
		if err != nil {
			fmt.Println(err)
			return entities.HouseDetails{}
		}
		break
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

		furnishedTypeTemp := utils.ReadInput("Enter Furnished type (1. Unfurnished, 2. Semi Furnished, 3. Fully Furnished): ")
		furnishedTypeInt, _ = strconv.Atoi(furnishedTypeTemp)
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

// GetAddress prompts the user for a pincode, fetches the address, and allows the user to confirm or update the details.
func (ui *UI) GetAddress() (entities.Address, error) {
	// Prompt user to enter pincode
	pincode := utils.ReadPincode(0) // input 0 to indicate that the function is called for either listing or updating property

	// Retrieve address details based on pincode
	address, err := utils.GetAddressFromPincode(pincode)
	if err != nil {
		fmt.Println("\nError fetching address details:", err)
		fmt.Println("Please enter the full address details manually.")

		// Collect full address details from the user
		address = entities.Address{
			Area:    utils.ReadInput("    Enter locality: "),
			City:    utils.ReadInput("    Enter city: "),
			State:   utils.ReadInput("    Enter state: "),
			Pincode: pincode,
		}
	} else {
		// Display fetched address details to the user
		fmt.Printf("\nFetched Address Details:\n")
		fmt.Printf("    Locality: %s\n", address.Area)
		fmt.Printf("    City: %s\n", address.City)
		fmt.Printf("    State: %s\n", address.State)
		fmt.Printf("    Pincode: %d\n", address.Pincode)

		// Prompt the user to confirm or edit the address details
		fmt.Println("\nIf you want to update any field, enter the new value. Leave it blank to keep the current value.")

		// Collect updated address details from the user
		area := utils.ReadInput(fmt.Sprintf("    Enter locality (current: %s): ", address.Area))
		if area != "" {
			address.Area = area
		}

		city := utils.ReadInput(fmt.Sprintf("    Enter city (current: %s): ", address.City))
		if city != "" {
			address.City = city
		}

		state := utils.ReadInput(fmt.Sprintf("    Enter state (current: %s): ", address.State))
		if state != "" {
			address.State = state
		}

		// Pincode remains unchanged since it was used to fetch the address
		address.Pincode = pincode
	}

	// Return the finalized address and any potential error
	return address, nil
}
