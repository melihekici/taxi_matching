package services

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"matching/config"
	"matching/customErrors"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func ValidateAuthentication(token string) (bool, *customErrors.HttpError) {
	if token == "" {
		return false, &customErrors.HttpError{
			StatusCode:  http.StatusUnauthorized,
			ErrorString: "Token is missing",
		}
	}

	valid, err := ValidateToken(token, config.SECRET)
	if !valid {
		return false, &customErrors.HttpError{
			StatusCode:  http.StatusUnauthorized,
			ErrorString: err.Error(),
		}
	}

	payloadString := strings.Split(token, ".")[1]
	payload, _ := base64.StdEncoding.DecodeString(payloadString)
	fmt.Println(payload)

	return true, nil
}

func ValidateToken(token string, secret string) (bool, error) {
	splitToken := strings.Split(token, ".")

	if len(splitToken) != 3 {
		return false, errors.New("invalid token")
	}

	// decode the header and payload back to strings
	header, err := base64.StdEncoding.DecodeString(splitToken[0])
	if err != nil {
		return false, errors.New("invalid token")
	}

	payload, err := base64.StdEncoding.DecodeString(splitToken[1])
	if err != nil {
		return false, errors.New("invalid token")
	}

	jwtPayload := make(map[string]string)
	err = json.Unmarshal(payload, &jwtPayload)
	if err != nil {
		return false, errors.New("invalid token")
	}

	if expiration, ok := jwtPayload["expiration"]; ok {
		expirationInt, err := strconv.Atoi(expiration)
		if err != nil {
			return false, errors.New("invalid token")
		}

		if int(time.Now().Unix()) > expirationInt {
			return false, errors.New("token expired")
		}
	} else {
		return false, errors.New("invalid token")
	}

	// again create the signature
	unsignedStr := string(header) + string(payload)

	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(unsignedStr))

	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))
	// check if signature is equal to passed signature
	if signature != splitToken[2] {
		return false, errors.New("invalid token")
	}

	return true, nil
}
