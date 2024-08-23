package interfaces

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"rentease/internal/domain/entities"
)

type PropertyRepo interface {
	SaveProperty(property entities.Property) error
	GetAllListedProperties(activerUseronly bool) ([]entities.Property, error)
	UpdateListedProperty(property entities.Property) error
	DeleteListedProperty(propertyID string) error
	//SearchProperties(area, city, state string, pincode int) ([]entities.Property, error)
	FindByID(ctx context.Context, id primitive.ObjectID) (*entities.Property, error)
	UpdateApprovalStatus(propertyID primitive.ObjectID, approved bool, adminUsername string) error
	FindPendingProperties() ([]entities.Property, error)
}

//type PropertyService interface {
//	ListProperty(property entities.Property) error
//	GetAllListedProperties(username string) ([]entities.Property, error)
//}
