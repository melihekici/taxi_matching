package handlers

import (
	"auth/services"
	"errors"
	"fmt"
	"net/http"
	"time"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := r.Header["Email"]; !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Email is missing"))
		return
	}
	if _, ok := r.Header["Username"]; !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Username is missing"))
		return
	}
	if _, ok := r.Header["Passwordhash"]; !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Passwordhash is missing"))
		return
	}

	fmt.Println(r.Header)
	// validate and then add the user
	check := services.AddUserObject(
		r.Header["Email"][0], r.Header["Username"][0], r.Header["Passwordhash"][0])

	if !check {
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte("Email or Username is taken."))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User Created"))
}

func SigninHandler(w http.ResponseWriter, r *http.Request) {
	if _, ok := r.Header["Email"]; !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Email Missing"))
		return
	}
	if _, ok := r.Header["Passwordhash"]; !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Passwordhash Missing"))
		return
	}

	valid, err := validateUser(r.Header["Email"][0], r.Header["Passwordhash"][0])
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
		"aud": "frontend.knowsearch.ml",
		"iss": "knowsearch.ml",
		"exp": fmt.Sprint(time.Now().Add(time.Minute * 1).Unix()),
	}

	secret := "Secure_Random_String"
	header := "HS256"
	tokenString, err := services.GenerateToken(header, claimsMap, secret)
	if err != nil {
		return tokenString, err
	}

	return tokenString, nil
}

func validateUser(email string, passwordHash string) (bool, error) {
	usr, exists := services.GetUserObject(email)
	if !exists {
		return false, errors.New("user does not exist")
	}

	passwordCheck := usr.ValidatePasswordHash(passwordHash)
	if !passwordCheck {
		return false, nil
	}

	return true, nil
}
