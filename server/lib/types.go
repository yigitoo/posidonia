package lib

import "go.mongodb.org/mongo-driver/bson/primitive"

/// UserID is equal to BSON ObjectID but in Hex/String format.
type UserID string

/// Username = Ex: "yigit" / "admin".
type Username string

/// strftime string format keeper.
type TimeStr string

type Coordinates []float64

type Locations struct {
	LocationID primitive.ObjectID `bson:"_id"`
	Polygon    []Coordinates      `bson:"polygon"`
	AddedTime  string             `bson:"added_time"`
	AddedBy    string             `bson:"added_by"`
  IsInDanger bool               `bson:"is_in_danger"`
}

type User struct {
	UserID   primitive.ObjectID `bson:"_id"`
	Username string
	Password string
}

type AddPolygon_RequestPayload struct {
	Polygon   []string `json:"polygon"`
	AddedBy   UserID   `json:"addedBy"`
  AddedTime TimeStr  `json:"addedTime"`
  IsInDanger bool `json:"is_in_danger"`
}
