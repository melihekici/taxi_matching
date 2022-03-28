package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Driver struct {
	Id       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Location Location           `json:"location" bson:"location"`
}

type Drivers struct {
	Drivers []Driver `json:"drivers"`
}

func (d *Driver) IsNil() bool {
	return d.Equals(&Driver{})
}

func (d *Driver) Equals(o *Driver) bool {
	if d.Id == o.Id && d.Location.Coordinates == o.Location.Coordinates {
		return true
	}
	return false
}
