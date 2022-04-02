package handlers

// A token returns in the response
// swagger:response signinResponse
type signinResponse struct {
	// JWT Token (15 mins)
	Token string `json:"token"`
}

// Request body for signin service
// swagger:model
type signinRequest struct {
	// User email
	// required: true
	Email string `json:"email"`
	// User password
	// required: true
	Password string `json:"password"`
}

// Request body for signup service
// swagger:model
type signupRequest struct {
	// User email
	// required: true
	Email string `json:"email"`
	// Username
	// required: true
	Username string `json:"username"`
	// User password
	// required: true
	Password string `json:"password"`
}

// swagger:parameters Signin
type signinRequestWrapper struct {
	// Signin Request Body
	// in: body
	// required: true
	Body signinRequest
}

// swagger:parameters Signup
type signupRequestWrapper struct {
	// Signup Request Body
	// in: body
	// required: true
	Body signupRequest
}
