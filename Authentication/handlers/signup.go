// Package classification Authentication API
//
// Documentation for Authentication API
//
// Schemes: http
// Host: localhost:9090
// BasePath: /
// Version: 1.0.0
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
// swagger:meta
package handlers

import (
	"auth/config"
	"auth/services"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type signupHandler struct{}
type signinHandler struct{}

var SignupHandler = &signupHandler{}
var SigninHandler = &signinHandler{}

// swagger:route POST /auth/signup Signup
// Creates a new user
// responses:
//  201: noContent
//  400: Bad Request
//  409: Conflict

// Creates a new user
func (s *signupHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var request signupRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}

	if request.Email == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Email is missing"))
		return
	} else if request.Username == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Username is missing"))
		return
	} else if request.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Password is missing"))
		return
	}

	passwordHash, err := services.HashPassword(request.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error hashing password"))
		return
	}

	// validate and then add the user
	check := services.AddUserObject(
		request.Email, request.Username, passwordHash)

	if !check {
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte("Email or Username is taken."))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User Created"))
}

// swagger:route POST /auth/signin Signin
// Returns a jwt token token for authentication
// responses:
//  200: signinResponse
//  400:
//  401:
//  500:

// Signin the user and return a jwt token
func (s *signinHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var request signupRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}

	if request.Email == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Email is missing"))
		return
	} else if request.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Password is missing"))
		return
	}

	valid, err := validateUser(request.Email, request.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("User does not exist"))
		return
	}

	if !valid {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Wrong username or password"))
		return
	}

	tokenString, err := getSignedToken()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error " + err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(tokenString))
}

func getSignedToken() (string, error) {
	claimsMap := map[string]string{
		"authenticated": "true",
		"expiration":    fmt.Sprint(time.Now().Add(time.Minute * 15).Unix()),
	}

	header := "HS256"
	tokenString, err := services.GenerateToken(header, claimsMap, config.SECRET)
	if err != nil {
		return tokenString, err
	}

	return tokenString, nil
}

func validateUser(email string, password string) (bool, error) {
	usr, exists := services.GetUserObject(email)
	if !exists {
		return false, errors.New("user does not exist")
	}

	passwordCheck := services.ValidatePasswordHash(usr.Password, password)
	if !passwordCheck {
		return false, nil
	}

	return true, nil
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
