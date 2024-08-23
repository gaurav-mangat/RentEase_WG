package ui

import (
	"fmt"
	"rentease/internal/domain/entities"
	"rentease/pkg/utils"
	"strings"
)

// UpdatePropertyUI handles the property update user interface.
func (ui *UI) UpdatePropertyUI(property entities.Property) {
	fmt.Println("\033[1;33m\nUpdating property...\033[0m") // Yellow

	// Create a copy of the existing property
	updatedProperty := property

	// Update Title
	newTitle := utils.ReadInput("Current Title: " + property.Title + "\nEnter new title (leave blank to skip): ")
	if newTitle != "" {
		updatedProperty.Title = newTitle
	}

	// Update Address
	fmt.Printf("Current Address: %s, %s, %s, %d\n", property.Address.Area, property.Address.City, property.Address.State, property.Address.Pincode)
	newArea := utils.ReadInput("Enter new area (leave blank to skip): ")
	if newArea != "" {
		updatedProperty.Address.Area = newArea
	}

	newCity := utils.ReadInput("Enter new city (leave blank to skip): ")
	if newCity != "" {
		updatedProperty.Address.City = newCity
	}

	newState := utils.ReadInput("Enter new state (leave blank to skip): ")
	if newState != "" {
		updatedProperty.Address.State = newState
	}

	newPincode := utils.ReadInput("Enter new pincode (leave blank to skip): ")
	if newPincode != "" {
		var pincode int
		_, err := fmt.Sscanf(newPincode, "%d", &pincode)
		if err == nil {
			updatedProperty.Address.Pincode = pincode
		} else {
			fmt.Println("\033[1;31mInvalid pincode format.\033[0m") // Red
			return
		}
	}

	// Update Rent Amount
	fmt.Println("Current expected rent amount:", property.RentAmount)
	newRentAmountStr := utils.ReadInput("Enter new expected rent amount (leave blank to skip): ")
	if newRentAmountStr != "" {
		var newRentAmount float64
		_, err := fmt.Sscanf(newRentAmountStr, "%f", &newRentAmount)
		if err == nil {
			fmt.Printf("Rent amount %T: ", newRentAmount) //fsgsgsgsg
			updatedProperty.RentAmount = newRentAmount
		} else {
			fmt.Println("\033[1;31mInvalid rent amount format.\033[0m") // Red
			return
		}
	}

	// Update Details based on Property Type
	switch property.Details.(type) {
	case entities.CommercialDetails:
		// Handle Commercial-specific updates
		fmt.Println("\033[1;31mUpdated details\033[0m", updatedProperty) //gsgsgsgsg
		if utils.ReadInput("Update Other details (floor area, subtype)? (yes/no): ") == "yes" {
			newFloorArea := utils.ReadInput("Enter new floor area in sq. feet (leave blank to skip): ")
			if newFloorArea != "" {
				updatedProperty.Details = entities.CommercialDetails{
					FloorArea: newFloorArea,
					SubType:   updatedProperty.Details.(entities.CommercialDetails).SubType,
				}
			}

			newSubType := utils.ReadInput("Enter new subtype (leave blank to skip): ")
			if newSubType != "" {
				updatedProperty.Details = entities.CommercialDetails{
					FloorArea: updatedProperty.Details.(entities.CommercialDetails).FloorArea,
					SubType:   newSubType,
				}
			}
		}

	case entities.HouseDetails:
		// Handle House-specific updates
		if utils.ReadInput("Update house details (number of rooms, furnished category, amenities)? (yes/no): ") == "yes" {
			newNoOfRooms := utils.ReadInput("Enter new number of rooms (leave blank to skip): ")
			if newNoOfRooms != "" {
				var noOfRooms int
				_, err := fmt.Sscanf(newNoOfRooms, "%d", &noOfRooms)
				if err == nil {
					updatedProperty.Details = entities.HouseDetails{
						NoOfRooms:         noOfRooms,
						FurnishedCategory: updatedProperty.Details.(entities.HouseDetails).FurnishedCategory,
						Amenities:         updatedProperty.Details.(entities.HouseDetails).Amenities,
					}
				} else {
					fmt.Println("\033[1;31mInvalid number of rooms format.\033[0m") // Red
					return
				}
			}

			newFurnishedCategory := utils.ReadInput("Enter new furnished category (leave blank to skip): ")
			if newFurnishedCategory != "" {
				updatedProperty.Details = entities.HouseDetails{
					NoOfRooms:         updatedProperty.Details.(entities.HouseDetails).NoOfRooms,
					FurnishedCategory: newFurnishedCategory,
					Amenities:         updatedProperty.Details.(entities.HouseDetails).Amenities,
				}
			}

			newAmenitiesStr := utils.ReadInput("Enter new amenities (comma separated, leave blank to skip): ")
			if newAmenitiesStr != "" {
				newAmenities := strings.Split(newAmenitiesStr, ",")
				updatedProperty.Details = entities.HouseDetails{
					NoOfRooms:         updatedProperty.Details.(entities.HouseDetails).NoOfRooms,
					FurnishedCategory: updatedProperty.Details.(entities.HouseDetails).FurnishedCategory,
					Amenities:         newAmenities,
				}
			}
		}

	case entities.FlatDetails:
		// Handle Flat-specific updates
		if utils.ReadInput("Update flat details (furnished category, amenities, BHK)? (yes/no): ") == "yes" {
			newFurnishedCategory := utils.ReadInput("Enter new furnished category (leave blank to skip): ")
			if newFurnishedCategory != "" {
				updatedProperty.Details = entities.FlatDetails{
					FurnishedCategory: newFurnishedCategory,
					Amenities:         updatedProperty.Details.(entities.FlatDetails).Amenities,
					BHK:               updatedProperty.Details.(entities.FlatDetails).BHK,
				}
			}

			newAmenitiesStr := utils.ReadInput("Enter new amenities (comma separated, leave blank to skip): ")
			if newAmenitiesStr != "" {
				newAmenities := strings.Split(newAmenitiesStr, ",")
				updatedProperty.Details = entities.FlatDetails{
					FurnishedCategory: updatedProperty.Details.(entities.FlatDetails).FurnishedCategory,
					Amenities:         newAmenities,
					BHK:               updatedProperty.Details.(entities.FlatDetails).BHK,
				}
			}

			newBHK := utils.ReadInput("Enter new BHK (leave blank to skip): ")
			if newBHK != "" {
				var bhk int
				_, err := fmt.Sscanf(newBHK, "%d", &bhk)
				if err == nil {
					updatedProperty.Details = entities.FlatDetails{
						FurnishedCategory: updatedProperty.Details.(entities.FlatDetails).FurnishedCategory,
						Amenities:         updatedProperty.Details.(entities.FlatDetails).Amenities,
						BHK:               bhk,
					}
				} else {
					fmt.Println("\033[1;31mInvalid BHK format.\033[0m") // Red
					return
				}
			}
		}

	default:
		fmt.Println("Unknown property details type")
		return
	}

	// Reset approval status if the property was approved
	if property.IsApprovedByAdmin {
		updatedProperty.IsApprovedByAdmin = false
	}

	// Save updated property
	err := ui.propertyService.UpdateListedProperty(updatedProperty)
	if err != nil {
		fmt.Printf("\033[1;31mError updating property: %v\033[0m\n", err)
	} else {
		fmt.Println("\033[1;32mProperty updated successfully.\033[0m")
	}
}
