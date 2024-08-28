package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Property struct {
	ID                primitive.ObjectID `bson:"_id"`           // MongoDB unique ID
	PropertyType      int                `bson:"property_type"` // 1: Commercial, 2: House, 3: Flat
	Title             string             `bson:"title"`
	Address           Address            `bson:"address"`
	LandlordUsername  string             `bson:"landlord_username"`
	RentAmount        float64            `bson:"rent_amount"`
	Applications      []string           `bson:"applications"`
	IsApprovedByAdmin bool               `bson:"is_approved_by_admin"`
	IsRented          bool               `bson:"is_rented"`
	Details           interface{}        `bson:"details"` // Holds specific details based on property type
}

type Address struct {
	Area    string `bson:"area"`
	City    string `bson:"city"`
	State   string `bson:"state"`
	Pincode int    `bson:"pincode"`
}

type CommercialDetails struct {
	FloorArea string `bson:"floor_area"`
	SubType   string `bson:"sub_type"` // shop, factory, warehouse
}

type HouseDetails struct {
	NoOfRooms         int      `bson:"no_of_rooms"`
	FurnishedCategory string   `bson:"furnished_category"`
	Amenities         []string `bson:"amenities"`
}

type FlatDetails struct {
	FurnishedCategory string   `bson:"furnished_category"`
	Amenities         []string `bson:"amenities"`
	BHK               int      `bson:"bhk"`
}
