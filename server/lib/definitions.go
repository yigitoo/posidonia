package lib

import "go.mongodb.org/mongo-driver/bson/primitive"

type Coordinates struct {
	Latitude  float64
	longitude float64
}

type Locations struct {
	LocationID primitive.ObjectID `bson:"_id"`
	Polygon    Polygon
	AddedTime  string
	AddedBy    string
}

type Polygon struct {
	CoordinateList []Coordinates
	CornerSize     uint8
}

type User struct {
	UserID   primitive.ObjectID `bson:"_id"`
	Username string
	Password string
}
