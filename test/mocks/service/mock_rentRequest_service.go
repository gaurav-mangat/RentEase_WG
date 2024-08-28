package mock_service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"rentease/internal/domain/entities"
)

type MockRentRequestService struct {
}

func NewMockRentRequestService() *MockRentRequestService {
	return &MockRentRequestService{}
}

func (ms *MockRentRequestService) CreateRentRequest(tenantName string, propertyID primitive.ObjectID, landlordName string) error {

	return nil

}

func (ms *MockRentRequestService) GetRentRequestsInfoForLandlord(landlordName string) ([]entities.Request, error) {

	return []entities.Request{}, nil
}

func (ms *MockRentRequestService) UpdateRequestStatus(request entities.Request, status string) error {

	return nil

}

func (ms *MockUserService) GetRentRequestsInfoForTenant(tenantName string) ([]entities.Request, error) {
	return []entities.Request{}, nil
}
