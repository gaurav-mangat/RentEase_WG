package services

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"rentease/internal/domain/entities"
	"rentease/internal/domain/interfaces"
	"strings"
)

type PropertyService struct {
	propertyRepo interfaces.PropertyRepo
}

func NewPropertyService(propertyRepo interfaces.PropertyRepo) *PropertyService {
	return &PropertyService{
		propertyRepo: propertyRepo,
	}
}

// ListProperty saves a property to the repository.
func (ps *PropertyService) ListProperty(property entities.Property) error {
	return ps.propertyRepo.SaveProperty(property)
}

// GetAllListedProperties retrieves all listed properties from the repository.
func (ps *PropertyService) GetAllListedProperties(activeUseronly bool) ([]entities.Property, error) {
	return ps.propertyRepo.GetAllListedProperties(activeUseronly)
}

// UpdateListedProperty updates a property in the repository.
func (ps *PropertyService) UpdateListedProperty(property entities.Property) error {

	fmt.Println("In UpdateListedProperty Function ", property.IsRented)
	// Check if the property is approved before updating
	if property.IsApprovedByAdmin && !property.IsRented {
		// Reset approval status if the property was approved
		property.IsApprovedByAdmin = false
	}
	return ps.propertyRepo.UpdateListedProperty(property)
}

// DeleteListedProperty deletes a property from the repository by ID.
func (ps *PropertyService) DeleteListedProperty(propertyID string) error {
	return ps.propertyRepo.DeleteListedProperty(propertyID)
}

// SearchProperties searches for properties based on the given criteria.
func (ps *PropertyService) SearchProperties(area, city, state string, pincode, propertyType int) ([]entities.Property, error) {
	properties, err := ps.propertyRepo.GetAllListedProperties(false)
	if err != nil {
		return nil, err
	}

	// Normalize the input strings
	area = strings.TrimSpace(strings.ToLower(area))
	city = strings.TrimSpace(strings.ToLower(city))
	state = strings.TrimSpace(strings.ToLower(state))

	var results []entities.Property
	for _, property := range properties {
		if property.PropertyType == propertyType {
			// Normalize the property address fields
			propArea := strings.TrimSpace(strings.ToLower(property.Address.Area))
			propCity := strings.TrimSpace(strings.ToLower(property.Address.City))
			propState := strings.TrimSpace(strings.ToLower(property.Address.State))

			if (strings.Contains(propArea, area) && strings.Contains(propCity, city) && property.Address.Pincode == pincode) ||
				(strings.Contains(propCity, city) && property.Address.Pincode == pincode) ||
				(property.Address.Pincode == pincode) ||
				(strings.Contains(propState, state)) {
				results = append(results, property)
			}
		}
	}

	return results, nil
}

// FindByID retrieves a property by its ID.
func (ps *PropertyService) FindByID(id primitive.ObjectID) (entities.Property, error) {
	// Use context in a real application
	ctx := context.TODO()

	property, err := ps.propertyRepo.FindByID(ctx, id)
	if err != nil {
		return entities.Property{}, err
	}
	if property == nil {
		return entities.Property{}, nil // Property not found
	}

	return *property, nil
}

// Admin

func (ps *PropertyService) GetPendingProperties() ([]entities.Property, error) {
	return ps.propertyRepo.FindPendingProperties()
}

func (ps *PropertyService) ApproveProperty(propertyID primitive.ObjectID, adminUsername string) error {
	return ps.propertyRepo.UpdateApprovalStatus(propertyID, true, adminUsername)
}
