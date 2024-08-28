package interfaces

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"rentease/internal/domain/entities"
)

type RentRequestService interface {
	CreateRentRequest(tenantName string, propertyID primitive.ObjectID, landlordName string) error
	GetRentRequestsInfoForLandlord(landlordName string) ([]entities.Request, error)
	UpdateRequestStatus(request entities.Request, status string) error
	GetRentRequestsInfoForTenant(tenantName string) ([]entities.Request, error)
}
