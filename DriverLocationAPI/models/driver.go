package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Driver struct {
	Id       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Location Location           `json:"location" bson:"location"`
}

type Drivers struct {
	Drivers []Driver `json:"drivers"`
}
