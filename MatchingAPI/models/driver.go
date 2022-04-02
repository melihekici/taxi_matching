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

func (d *Drivers) HasNilDriver() bool {
	for _, driver := range d.Drivers {
		if driver.IsNil() {
			return true
		}
	}
	return false
}

func (d *Drivers) IsNil() bool {
	if len(d.Drivers) == 0 {
		return true
	}
	for _, driver := range d.Drivers {
		if !driver.IsNil() {
			return false
		}
	}

	return true
}

func (d *Drivers) Equals(o *Drivers) bool {
	if len(d.Drivers) != len(o.Drivers) {
		return false
	}
	for i, currentDriver := range d.Drivers {
		if !currentDriver.Equals(&o.Drivers[i]) {
			return false
		}
	}
	return true
}
