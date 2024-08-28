package service_test

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"rentease/internal/app/services"
	"rentease/internal/domain/entities"
	mocks_interfaces "rentease/test/mocks/repository"
	"testing"
	"time"
)

var (
	mockRentRequestRepo *mocks_interfaces.MockRequestRepo
	rentRequestService  *services.RequestService
)

func setup3(t *testing.T) func() {
	// Set up the gomock controller
	ctrl := gomock.NewController(t)

	// Create a mock PropertyRepo
	mockRentRequestRepo = mocks_interfaces.NewMockRequestRepo(ctrl)

	// Initialize the RentRequestService with the mock repository
	rentRequestService = services.NewRequestService(mockRentRequestRepo)

	// Return a cleanup function to be called at the end of the test
	return func() {
		ctrl.Finish()
	}
}

func TestRequestService_GetRentRequestsInfoForLandlord(t *testing.T) {
	cleanup := setup3(t)
	defer cleanup()

	landlordName := "landlord1"
	now := time.Now()
	requests := []entities.Request{
		{
			PropertyID:    primitive.NewObjectID(),
			TenantName:    "tenant1",
			LandlordName:  landlordName,
			RequestStatus: "pending",
			CreatedAt:     now,
		},
		{
			PropertyID:    primitive.NewObjectID(),
			TenantName:    "tenant2",
			LandlordName:  landlordName,
			RequestStatus: "pending",
			CreatedAt:     now,
		},
	}

	tests := []struct {
		name             string
		mockReturn       []entities.Request
		mockError        error
		expectedError    bool
		expectedRequests []entities.Request
	}{
		{
			name:             "Successful retrieval",
			mockReturn:       requests,
			mockError:        nil,
			expectedError:    false,
			expectedRequests: requests,
		},
		{
			name:             "Error during retrieval",
			mockReturn:       nil,
			mockError:        errors.New("find error"),
			expectedError:    true,
			expectedRequests: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mock expectation
			mockRentRequestRepo.EXPECT().
				FindByLandlordName(gomock.Any(), landlordName).
				Return(tt.mockReturn, tt.mockError).
				Times(1)

			// Call the method under test
			result, err := rentRequestService.GetRentRequestsInfoForLandlord(landlordName)

			// Assertions
			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				// Compare without CreatedAt field
				for i := range result {
					result[i].CreatedAt = time.Time{}
				}
				for i := range tt.expectedRequests {
					tt.expectedRequests[i].CreatedAt = time.Time{}
				}
				assert.ElementsMatch(t, tt.expectedRequests, result)
			}
		})
	}
}

func TestRequestService_UpdateRequestStatus(t *testing.T) {
	cleanup := setup3(t)
	defer cleanup()

	request := entities.Request{
		PropertyID:    primitive.NewObjectID(),
		TenantName:    "tenant1",
		LandlordName:  "landlord1",
		RequestStatus: "pending",
		CreatedAt:     time.Now(),
	}
	status := "approved"

	tests := []struct {
		name          string
		mockError     error
		expectedError bool
	}{
		{
			name:          "Successful update",
			mockError:     nil,
			expectedError: false,
		},
		{
			name:          "Error during update",
			mockError:     errors.New("update error"),
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mock the UpdateRequest call with the expected arguments
			mockRentRequestRepo.EXPECT().
				UpdateRequest(request, status).
				Return(tt.mockError).
				Times(1)

			err := rentRequestService.UpdateRequestStatus(request, status)

			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestRequestService_GetRentRequestsInfoForTenant(t *testing.T) {
	cleanup := setup3(t)
	defer cleanup()

	tenantName := "tenant1"
	requests := []entities.Request{
		{
			PropertyID:    primitive.NewObjectID(),
			TenantName:    tenantName,
			LandlordName:  "landlord1",
			RequestStatus: "pending",
			CreatedAt:     time.Now(),
		},
		{
			PropertyID:    primitive.NewObjectID(),
			TenantName:    tenantName,
			LandlordName:  "landlord2",
			RequestStatus: "approved",
			CreatedAt:     time.Now(),
		},
	}

	tests := []struct {
		name             string
		mockReturn       []entities.Request
		mockError        error
		expectedError    bool
		expectedRequests []entities.Request
	}{
		{
			name:             "Successful retrieval",
			mockReturn:       requests,
			mockError:        nil,
			expectedError:    false,
			expectedRequests: requests,
		},
		{
			name:             "Error during retrieval",
			mockReturn:       nil,
			mockError:        errors.New("find error"),
			expectedError:    true,
			expectedRequests: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRentRequestRepo.EXPECT().
				FindByTenantUsername(gomock.Any(), tenantName).
				Return(tt.mockReturn, tt.mockError).
				Times(1)

			result, err := rentRequestService.GetRentRequestsInfoForTenant(tenantName)

			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.ElementsMatch(t, tt.expectedRequests, result)
			}
		})
	}
}

func TestRequestService_CreateRentRequest(t *testing.T) {
	cleanup := setup3(t)
	defer cleanup()

	tenantName := "tenant1"
	propertyID := primitive.NewObjectID()
	landlordName := "landlord1"

	tests := []struct {
		name          string
		mockError     error
		expectedError bool
	}{
		{
			name:          "Successful creation",
			mockError:     nil,
			expectedError: false,
		},
		{
			name:          "Error during creation",
			mockError:     errors.New("save error"),
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expectedRequest := entities.Request{
				PropertyID:    propertyID,
				TenantName:    tenantName,
				LandlordName:  landlordName,
				RequestStatus: "pending",
				CreatedAt:     time.Now(), // We can't precisely match time, but we'll compare the other fields.
			}

			mockRentRequestRepo.EXPECT().
				SaveRequest(gomock.AssignableToTypeOf(expectedRequest)).
				Return(tt.mockError).
				Times(1)

			err := rentRequestService.CreateRentRequest(tenantName, propertyID, landlordName)

			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
