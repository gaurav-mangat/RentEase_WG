// internal/domain/services/request_service.go
package services

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"rentease/internal/domain/entities"
	"rentease/internal/domain/interfaces"
	"time"
)

type RequestService struct {
	requestRepo interfaces.RequestRepo
}

func NewRequestService(requestRepo interfaces.RequestRepo) *RequestService {
	return &RequestService{
		requestRepo: requestRepo,
	}
}

func (rs *RequestService) CreateRentRequest(tenantName string, propertyID primitive.ObjectID, landlordName string) error {

	request := entities.Request{
		PropertyID:    propertyID,
		TenantName:    tenantName,
		LandlordName:  landlordName,
		RequestStatus: "pending",
		CreatedAt:     time.Now(),
	}

	return rs.requestRepo.SaveRequest(request)
}

// GetRentRequestsInfoForLandlord gives  all the rent requests for the landlord
func (rs *RequestService) GetRentRequestsInfoForLandlord(landlordName string) ([]entities.Request, error) {
	ctx := context.TODO()
	return rs.requestRepo.FindByLandlordName(ctx, landlordName)
}

func (rs *RequestService) UpdateRequestStatus(request entities.Request, status string) error {
	return rs.requestRepo.UpdateRequest(request, status)
}

// GetRentRequestsInfoForTenant gives all the rent requests for the tenant
func (rs *RequestService) GetRentRequestsInfoForTenant(tenantName string) ([]entities.Request, error) {
	ctx := context.TODO()
	return rs.requestRepo.FindByTenantUsername(ctx, tenantName)
}
