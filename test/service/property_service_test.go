package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"rentease/internal/app/services"
	"rentease/internal/domain/entities"
	mocks_interfaces "rentease/test/mocks/repository"
)

var (
	mockPropertyRepo *mocks_interfaces.MockPropertyRepo
	propertyService  *services.PropertyService
)

func setup2(t *testing.T) func() {
	// Set up the gomock controller
	ctrl := gomock.NewController(t)

	// Create a mock PropertyRepo
	mockPropertyRepo = mocks_interfaces.NewMockPropertyRepo(ctrl)

	// Initialize the PropertyService with the mock repository
	propertyService = services.NewPropertyService(mockPropertyRepo)

	// Return a cleanup function to be called at the end of the test
	return func() {
		ctrl.Finish()
	}
}

func TestPropertyService_ListProperty(t *testing.T) {
	// Setup and defer cleanup
	cleanup := setup2(t)
	defer cleanup()

	tests := []struct {
		name          string
		property      entities.Property
		mockError     error
		expectedError bool
	}{
		{
			name: "Successful Save Commercial Property",
			property: entities.Property{
				ID:               primitive.NewObjectID(),
				PropertyType:     1, // Commercial
				Title:            "Commercial Space",
				Address:          entities.Address{Area: "Downtown", City: "Metropolis", State: "NY", Pincode: 10001},
				LandlordUsername: "landlord1",
				RentAmount:       2000.00,
				Details: entities.CommercialDetails{
					FloorArea: "5000 sq ft",
					SubType:   "warehouse",
				},
			},
			mockError:     nil,
			expectedError: false,
		},
		{
			name: "Successful Save House Property",
			property: entities.Property{
				ID:               primitive.NewObjectID(),
				PropertyType:     2, // House
				Title:            "Family House",
				Address:          entities.Address{Area: "Suburb", City: "Smalltown", State: "TX", Pincode: 75001},
				LandlordUsername: "landlord2",
				RentAmount:       1500.00,
				Details: entities.HouseDetails{
					NoOfRooms:         4,
					FurnishedCategory: "semi-furnished",
					Amenities:         []string{"garden", "garage"},
				},
			},
			mockError:     nil,
			expectedError: false,
		},
		{
			name: "Successful Save Flat Property",
			property: entities.Property{
				ID:               primitive.NewObjectID(),
				PropertyType:     3, // Flat
				Title:            "Luxury Flat",
				Address:          entities.Address{Area: "Uptown", City: "Metropolis", State: "NY", Pincode: 10002},
				LandlordUsername: "landlord3",
				RentAmount:       2500.00,
				Details: entities.FlatDetails{
					FurnishedCategory: "furnished",
					Amenities:         []string{"gym", "pool"},
					BHK:               3,
				},
			},
			mockError:     nil,
			expectedError: false,
		},
		{
			name: "Error Saving Property",
			property: entities.Property{
				ID:    primitive.NewObjectID(),
				Title: "Error Property",
			},
			mockError:     errors.New("save property error"),
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set up the expected behavior of the mock repository
			mockPropertyRepo.EXPECT().SaveProperty(tt.property).Return(tt.mockError).Times(1)

			// Call the ListProperty method and capture the result
			err := propertyService.ListProperty(tt.property)

			// Validate the result
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestPropertyService_GetAllListedProperties(t *testing.T) {
	cleanup := setup2(t)
	defer cleanup()

	predefinedID1 := primitive.NewObjectID()
	predefinedID2 := primitive.NewObjectID()

	tests := []struct {
		name             string
		activeUserOnly   bool
		mockProperties   []entities.Property
		mockError        error
		expectedError    bool
		expectedResponse []entities.Property
	}{
		{
			name:           "Successful retrieval of all listed properties",
			activeUserOnly: false,
			mockProperties: []entities.Property{
				{
					ID:               predefinedID1,
					PropertyType:     1, // Commercial
					Title:            "Commercial Space",
					Address:          entities.Address{Area: "Downtown", City: "Metropolis", State: "NY", Pincode: 10001},
					LandlordUsername: "landlord1",
					RentAmount:       2000.00,
				},
				{
					ID:               predefinedID2,
					PropertyType:     2, // House
					Title:            "Family House",
					Address:          entities.Address{Area: "Suburb", City: "Smalltown", State: "TX", Pincode: 75001},
					LandlordUsername: "landlord2",
					RentAmount:       1500.00,
				},
			},
			mockError:     nil,
			expectedError: false,
			expectedResponse: []entities.Property{
				{
					ID:               predefinedID1,
					PropertyType:     1, // Commercial
					Title:            "Commercial Space",
					Address:          entities.Address{Area: "Downtown", City: "Metropolis", State: "NY", Pincode: 10001},
					LandlordUsername: "landlord1",
					RentAmount:       2000.00,
				},
				{
					ID:               predefinedID2,
					PropertyType:     2, // House
					Title:            "Family House",
					Address:          entities.Address{Area: "Suburb", City: "Smalltown", State: "TX", Pincode: 75001},
					LandlordUsername: "landlord2",
					RentAmount:       1500.00,
				},
			},
		},
		{
			name:             "Error retrieving properties",
			activeUserOnly:   true,
			mockProperties:   nil,
			mockError:        errors.New("error fetching properties"),
			expectedError:    true,
			expectedResponse: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockPropertyRepo.EXPECT().GetAllListedProperties(tt.activeUserOnly).Return(tt.mockProperties, tt.mockError).Times(1)

			properties, err := propertyService.GetAllListedProperties(tt.activeUserOnly)

			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResponse, properties)
			}
		})
	}
}

func TestPropertyService_UpdateListedProperty(t *testing.T) {
	cleanup := setup2(t) // Assuming setup2 initializes the mock and service
	defer cleanup()

	predefinedID := primitive.NewObjectID()

	tests := []struct {
		name          string
		property      entities.Property
		mockError     error
		expectedError bool
		expectedCalls int
	}{
		{
			name: "Successful update of title only",
			property: entities.Property{
				ID:                predefinedID,
				Title:             "Updated Commercial Space",
				PropertyType:      1, // Commercial
				Address:           entities.Address{Area: "Downtown", City: "Metropolis", State: "NY", Pincode: 10001},
				RentAmount:        2000.00,
				IsApprovedByAdmin: false, // After update, this should be false
				IsRented:          false,
				Details: entities.CommercialDetails{
					FloorArea: "5000 sq ft",
					SubType:   "warehouse",
				},
			},
			mockError:     nil,
			expectedError: false,
			expectedCalls: 1,
		},
		{
			name: "Successful update of address only",
			property: entities.Property{
				ID:                predefinedID,
				Title:             "Commercial Space",
				PropertyType:      1, // Commercial
				Address:           entities.Address{Area: "Uptown", City: "Metropolis", State: "NY", Pincode: 10002},
				RentAmount:        2000.00,
				IsApprovedByAdmin: false, // After update, this should be false
				IsRented:          false,
				Details: entities.CommercialDetails{
					FloorArea: "5000 sq ft",
					SubType:   "warehouse",
				},
			},
			mockError:     nil,
			expectedError: false,
			expectedCalls: 1,
		},
		{
			name: "Successful update of rent amount",
			property: entities.Property{
				ID:                predefinedID,
				Title:             "Commercial Space",
				PropertyType:      1, // Commercial
				Address:           entities.Address{Area: "Downtown", City: "Metropolis", State: "NY", Pincode: 10001},
				RentAmount:        2500.00,
				IsApprovedByAdmin: false, // After update, this should be false
				IsRented:          false,
				Details: entities.CommercialDetails{
					FloorArea: "5000 sq ft",
					SubType:   "warehouse",
				},
			},
			mockError:     nil,
			expectedError: false,
			expectedCalls: 1,
		},
		{
			name: "Successful update of details",
			property: entities.Property{
				ID:                predefinedID,
				Title:             "Commercial Space",
				PropertyType:      1, // Commercial
				Address:           entities.Address{Area: "Downtown", City: "Metropolis", State: "NY", Pincode: 10001},
				RentAmount:        2000.00,
				IsApprovedByAdmin: false, // After update, this should be false
				IsRented:          false,
				Details: entities.CommercialDetails{
					FloorArea: "6000 sq ft", // Updated floor area
					SubType:   "office",     // Updated subtype
				},
			},
			mockError:     nil,
			expectedError: false,
			expectedCalls: 1,
		},
		{
			name: "Successful update of approved and unrented property with all fields",
			property: entities.Property{
				ID:                predefinedID,
				PropertyType:      1, // Commercial
				Title:             "Commercial Space",
				Address:           entities.Address{Area: "Downtown", City: "Metropolis", State: "NY", Pincode: 10001},
				LandlordUsername:  "landlord1",
				RentAmount:        2000.00,
				IsApprovedByAdmin: false, // After update, this should be false
				IsRented:          false,
				Details: entities.CommercialDetails{
					FloorArea: "5000 sq ft",
					SubType:   "warehouse",
				},
			},
			mockError:     nil,
			expectedError: false,
			expectedCalls: 1,
		},
		{
			name: "Successful update of rented property with all fields",
			property: entities.Property{
				ID:                predefinedID,
				PropertyType:      2, // House
				Title:             "Family House",
				Address:           entities.Address{Area: "Suburb", City: "Smalltown", State: "TX", Pincode: 75001},
				LandlordUsername:  "landlord2",
				RentAmount:        1500.00,
				IsApprovedByAdmin: true,
				IsRented:          true,
				Details: entities.HouseDetails{
					NoOfRooms:         4,
					FurnishedCategory: "semi-furnished",
					Amenities:         []string{"garden", "garage"},
				},
			},
			mockError:     nil,
			expectedError: false,
			expectedCalls: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Expect the UpdateListedProperty method to be called with the modified property
			mockPropertyRepo.EXPECT().
				UpdateListedProperty(tt.property).
				Return(tt.mockError).
				Times(tt.expectedCalls)

			err := propertyService.UpdateListedProperty(tt.property)

			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestPropertyService_DeleteListedProperty(t *testing.T) {
	cleanup := setup2(t) // Assuming setup2 initializes the mock and service
	defer cleanup()

	predefinedID := "somePropertyID"

	tests := []struct {
		name          string
		propertyID    string
		mockError     error
		expectedError bool
		expectedCalls int
	}{
		{
			name:          "Successful deletion",
			propertyID:    predefinedID,
			mockError:     nil,
			expectedError: false,
			expectedCalls: 1,
		},
		{
			name:          "Error during deletion",
			propertyID:    predefinedID,
			mockError:     errors.New("delete error"),
			expectedError: true,
			expectedCalls: 1,
		},
		{
			name:          "Empty property ID",
			propertyID:    "",
			mockError:     nil,
			expectedError: false,
			expectedCalls: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set up the expected call only if the property ID is not empty
			if tt.expectedCalls > 0 {
				mockPropertyRepo.EXPECT().
					DeleteListedProperty(tt.propertyID).
					Return(tt.mockError).
					Times(tt.expectedCalls)
			}

			err := propertyService.DeleteListedProperty(tt.propertyID)

			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestPropertyService_SearchProperties(t *testing.T) {
	cleanup := setup2(t) // Assuming setup2 initializes the mock and service
	defer cleanup()

	// Create mock properties
	properties := []entities.Property{
		{
			ID:                primitive.NewObjectID(),
			Title:             "Commercial Space 1",
			PropertyType:      1, // Commercial
			Address:           entities.Address{Area: "Downtown", City: "Metropolis", State: "NY", Pincode: 10001},
			IsApprovedByAdmin: true,
			IsRented:          false,
		},
		{
			ID:                primitive.NewObjectID(),
			Title:             "House in Suburb",
			PropertyType:      2, // House
			Address:           entities.Address{Area: "Suburb", City: "Smalltown", State: "TX", Pincode: 75001},
			IsApprovedByAdmin: true,
			IsRented:          false,
		},
		{
			ID:                primitive.NewObjectID(),
			Title:             "Flat in Uptown",
			PropertyType:      3, // Flat
			Address:           entities.Address{Area: "Uptown", City: "Metropolis", State: "NY", Pincode: 10002},
			IsApprovedByAdmin: true,
			IsRented:          false,
		},
	}

	tests := []struct {
		name           string
		area           string
		city           string
		state          string
		pincode        int
		propertyType   int
		mockProperties []entities.Property
		expectedResult []entities.Property
		expectedError  bool
	}{
		{
			name:           "Successful search with full criteria",
			area:           "Downtown",
			city:           "Metropolis",
			state:          "NY",
			pincode:        10001,
			propertyType:   1, // Commercial
			mockProperties: properties,
			expectedResult: []entities.Property{properties[0]},
			expectedError:  false,
		},
		{
			name:           "Successful search with partial criteria",
			area:           "Suburb",
			city:           "",
			state:          "TX",
			pincode:        75001,
			propertyType:   2, // House
			mockProperties: properties,
			expectedResult: []entities.Property{properties[1]},
			expectedError:  false,
		},
		{
			name:           "No matching properties",
			area:           "Nonexistent",
			city:           "Nowhere",
			state:          "ZZ",
			pincode:        99999,
			propertyType:   3, // Flat
			mockProperties: properties,
			expectedResult: []entities.Property{},
			expectedError:  false,
		},
		{
			name:           "Error from repository",
			area:           "Downtown",
			city:           "Metropolis",
			state:          "NY",
			pincode:        10001,
			propertyType:   1,   // Commercial
			mockProperties: nil, // No properties
			expectedResult: []entities.Property{},
			expectedError:  false,
		},
		{
			name:           "Case insensitivity",
			area:           "downtown",
			city:           "metropolis",
			state:          "ny",
			pincode:        10001,
			propertyType:   1, // Commercial
			mockProperties: properties,
			expectedResult: []entities.Property{properties[0]},
			expectedError:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// Set up the mock to return the predefined properties or error
			mockPropertyRepo.EXPECT().
				GetAllListedProperties(false).
				Return(tt.mockProperties, nil).
				Times(1)

			result, err := propertyService.SearchProperties(tt.area, tt.city, tt.state, tt.pincode, tt.propertyType)

			if tt.expectedError {
				assert.Error(t, err)
				assert.Empty(t, result)
			} else {
				assert.NoError(t, err)
				assert.ElementsMatch(t, tt.expectedResult, result)
			}
		})
	}
}

func TestPropertyService_FindByID(t *testing.T) {
	cleanup := setup2(t) // Assuming setup2 initializes the mock and service
	defer cleanup()

	// Create mock property
	propertyID := primitive.NewObjectID()
	property := &entities.Property{
		ID:                propertyID,
		Title:             "Commercial Space",
		PropertyType:      1, // Commercial
		Address:           entities.Address{Area: "Downtown", City: "Metropolis", State: "NY", Pincode: 10001},
		IsApprovedByAdmin: true,
		IsRented:          false,
	}

	tests := []struct {
		name           string
		propertyID     primitive.ObjectID
		mockProperty   *entities.Property
		mockError      error
		expectedResult entities.Property
		expectedError  bool
	}{
		{
			name:           "Successful retrieval",
			propertyID:     propertyID,
			mockProperty:   property,
			mockError:      nil,
			expectedResult: *property,
			expectedError:  false,
		},
		{
			name:           "Property not found",
			propertyID:     propertyID,
			mockProperty:   nil,
			mockError:      nil,
			expectedResult: entities.Property{},
			expectedError:  false,
		},
		{
			name:           "Error from repository",
			propertyID:     propertyID,
			mockProperty:   nil,
			mockError:      assert.AnError, // Using assert.AnError as a general error
			expectedResult: entities.Property{},
			expectedError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// Set up the mock to return the predefined property or error
			mockPropertyRepo.EXPECT().
				FindByID(context.TODO(), tt.propertyID). // Match any context
				Return(tt.mockProperty, tt.mockError).
				Times(1)

			result, err := propertyService.FindByID(tt.propertyID)

			if tt.expectedError {
				assert.Error(t, err)
				assert.Empty(t, result)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResult, result)
			}
		})
	}
}

func TestPropertyService_DeleteAllListedPropertiesOfaUser(t *testing.T) {
	cleanup := setup2(t) // Assuming setup2 initializes the mock and service
	defer cleanup()

	tests := []struct {
		name          string
		username      string
		mockError     error
		expectedError bool
	}{
		{
			name:          "Successful deletion",
			username:      "testuser",
			mockError:     nil,
			expectedError: false,
		},
		{
			name:          "Error during deletion",
			username:      "testuser",
			mockError:     assert.AnError, // Using assert.AnError as a general error
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// Set up the mock to return the predefined error
			mockPropertyRepo.EXPECT().
				DeleteAllListedPropertiesOfaUser(tt.username).
				Return(tt.mockError).
				Times(1)

			err := propertyService.DeleteAllListedPropertiesOfaUser(tt.username)

			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestPropertyService_GetPendingProperties(t *testing.T) {
	cleanup := setup2(t) // Assuming setup2 initializes the mock and service
	defer cleanup()

	// Define a fixed ObjectID for consistency
	fixedID := primitive.NewObjectID()

	tests := []struct {
		name           string
		mockProperties []entities.Property
		mockError      error
		expectedResult []entities.Property
		expectedError  bool
	}{
		{
			name: "Successful retrieval",
			mockProperties: []entities.Property{
				{
					ID:                fixedID,
					Title:             "Pending Commercial Space",
					PropertyType:      1,
					Address:           entities.Address{Area: "Downtown", City: "Metropolis", State: "NY", Pincode: 10001},
					IsApprovedByAdmin: false,
					IsRented:          false,
				},
			},
			mockError: nil,
			expectedResult: []entities.Property{
				{
					ID:                fixedID,
					Title:             "Pending Commercial Space",
					PropertyType:      1,
					Address:           entities.Address{Area: "Downtown", City: "Metropolis", State: "NY", Pincode: 10001},
					IsApprovedByAdmin: false,
					IsRented:          false,
				},
			},
			expectedError: false,
		},
		{
			name:           "Error from repository",
			mockProperties: nil,
			mockError:      assert.AnError, // Using assert.AnError as a general error
			expectedResult: []entities.Property{},
			expectedError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// Set up the mock to return the predefined properties or error
			mockPropertyRepo.EXPECT().
				FindPendingProperties().
				Return(tt.mockProperties, tt.mockError).
				Times(1)

			result, err := propertyService.GetPendingProperties()

			if tt.expectedError {
				assert.Error(t, err)
				assert.Empty(t, result)
			} else {
				assert.NoError(t, err)
				// Custom comparison for IDs to handle potential differences
				for i, resProperty := range result {
					assert.True(t, i < len(tt.expectedResult), "Unexpected result length")
					expProperty := tt.expectedResult[i]
					assert.True(t, resProperty.ID == expProperty.ID, "ID does not match")
					assert.Equal(t, expProperty.Title, resProperty.Title)
					assert.Equal(t, expProperty.PropertyType, resProperty.PropertyType)
					assert.Equal(t, expProperty.Address, resProperty.Address)
					assert.Equal(t, expProperty.IsApprovedByAdmin, resProperty.IsApprovedByAdmin)
					assert.Equal(t, expProperty.IsRented, resProperty.IsRented)
				}
			}
		})
	}
}

func TestPropertyService_ApproveProperty(t *testing.T) {
	cleanup := setup2(t) // Assuming setup2 initializes the mock and service
	defer cleanup()

	tests := []struct {
		name          string
		propertyID    primitive.ObjectID
		adminUsername string
		mockError     error
		expectedError bool
	}{
		{
			name:          "Successful approval",
			propertyID:    primitive.NewObjectID(), // Use a new ObjectID for the test
			adminUsername: "adminUser",
			mockError:     nil,
			expectedError: false,
		},
		{
			name:          "Error from repository",
			propertyID:    primitive.NewObjectID(),
			adminUsername: "adminUser",
			mockError:     assert.AnError, // Simulate an error
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set up the mock to expect the correct call
			mockPropertyRepo.EXPECT().
				UpdateApprovalStatus(tt.propertyID, true, tt.adminUsername).
				Return(tt.mockError).
				Times(1)

			err := propertyService.ApproveProperty(tt.propertyID, tt.adminUsername)

			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
