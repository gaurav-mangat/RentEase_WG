package mock_service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"rentease/internal/domain/entities"
)

type MockPropertyService struct {
}

func NewMockPropertyService() *MockPropertyService {
	return &MockPropertyService{}
}

// ListProperty mock implementation .
func (ms *MockPropertyService) ListProperty(property entities.Property) error {
	return nil
}

// GetAllListedProperties mock implementation .
func (ms *MockPropertyService) GetAllListedProperties(activeUseronly bool) ([]entities.Property, error) {
	return []entities.Property{}, nil
}

// UpdateListedProperty mock implementation .
func (ms *MockPropertyService) UpdateListedProperty(property entities.Property) error {

	return nil
}

// DeleteListedProperty mock implementation
func (ms *MockPropertyService) DeleteListedProperty(propertyID string) error {
	return nil
}

// SearchProperties function's Mock implementation
func (ms *MockPropertyService) SearchProperties(area, city, state string, pincode, propertyType int) ([]entities.Property, error) {

	return []entities.Property{}, nil
}

// FindByID function's Mock implementation
func (ms *MockPropertyService) FindByID(id primitive.ObjectID) (entities.Property, error) {
	return entities.Property{}, nil
}

// DeleteAllListedPropertiesOfaUser function's Mock implementation
func (ms *MockPropertyService) DeleteAllListedPropertiesOfaUser(username string) error {
	return nil
}

// GetPendingProperties function's Mock implementation
func (ms *MockPropertyService) GetPendingProperties() ([]entities.Property, error) {
	return []entities.Property{}, nil
}

// ApproveProperty function's  Mock implementation
func (ms *MockPropertyService) ApproveProperty(propertyID primitive.ObjectID, adminUsername string) error {
	return nil
}
