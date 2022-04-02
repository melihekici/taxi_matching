package controllers

// Request body for driver finding service
// swagger:model findDriverRequest
type findDriverRequestSwagger struct {
	// User location
	// required: true
	Location LocationSwagger
	// Radius in meters
	// required:true
	Radius float64
}

// Geolocation object
// swagger:model Location
type LocationSwagger struct {
	// Type
	Type string `json:"type"`
	// Coordinates
	Coordinates [2]float64 `json:"coordinates" bson:"coordinates"`
}

// swagger:parameters FindDriver
type findDriverRequestWrapper struct {
	// FindDriver request body
	// in: body
	// required: true
	Body findDriverRequestSwagger
}
