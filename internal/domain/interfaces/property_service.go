package interfaces

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"rentease/internal/domain/entities"
)

type PropertyService interface {
	ListProperty(property entities.Property) error

	GetAllListedProperties(activeUseronly bool) ([]entities.Property, error)

	UpdateListedProperty(property entities.Property) error

	DeleteListedProperty(propertyID string) error

	SearchProperties(area, city, state string, pincode, propertyType int) ([]entities.Property, error)

	FindByID(id primitive.ObjectID) (entities.Property, error)

	DeleteAllListedPropertiesOfaUser(username string) error

	GetPendingProperties() ([]entities.Property, error)

	ApproveProperty(propertyID primitive.ObjectID, adminUsername string) error
}
