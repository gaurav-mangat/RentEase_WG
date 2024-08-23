package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Request struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	TenantName    string             `bson:"tenantName"`
	PropertyID    primitive.ObjectID `bson:"propertyID"`
	LandlordName  string             `bson:"landlordName"`
	RequestStatus string             `bson:"requestStatus"` // e.g., "Pending", "Accepted", "Rejected"
	CreatedAt     time.Time          `bson:"created_at"`
}
