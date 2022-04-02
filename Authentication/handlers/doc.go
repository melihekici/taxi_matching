package handlers

// A token returns in the response
// swagger:response signinResponse
type signinResponse struct {
	// JWT Token (15 mins)
	Token string `json:"token"`
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
