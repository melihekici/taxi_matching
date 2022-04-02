package controllers

// swagger:parameters DeleteDriver
type DeleteDriverRequestWrapper struct {
	// driver id
	// in: path
	ID string `json:"id"`
}

// Driver Response Model
// swagger:model DriverResponse
type DriverResponseModelSwagger struct {
	// Driver ID
	Id string `json:"id"`
	// Driver Location
	Location LocationSwagger `json:"location"`
}

// Driver Request Model
// swagger:model DriverRequest
type DriverRequestModelSwagger struct {
	// Driver Location
	Location LocationSwagger `json:"location"`
}

// Create batch drivers request body
// swagger:model CreateDriversRequest
type CreateDriversInBatchRequest struct {
	// List of drivers to be created
	// required: true
	Drivers []DriverRequestModelSwagger `json:"drivers"`
}

// Create batch drivers request body
// swagger:parameters CreateDrivers
type CreateDriversInBatchRequestWrapper struct {
	// CreateDrivers Request Body
	// in: body
	// required: true
	Body CreateDriversInBatchRequest
}

// Create single driver request body
// swagger:parameters CreateDriver
type CreateDriverRequestWrapper struct {
	// CreateDriver Request Body
	// in: body
	// required: true
	Body DriverRequestModelSwagger
}

// Geolocation object
// swagger:model
type LocationSwagger struct {
	// Type
	Type string `json:"type"`
	// Coordinates
	Coordinates [2]float64 `json:"coordinates" bson:"coordinates"`
}
