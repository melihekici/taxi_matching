package services

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type jwtPayload struct {
	Authentication bool      `json:"authentication"`
	Expiration     time.Time `json:"expiration"`
}

func GenerateToken(header string, payload map[string]string, secret string) (string, error) {
	// create a new hash of type sha256
	h := hmac.New(sha256.New, []byte(secret))
	header64 := base64.StdEncoding.EncodeToString([]byte(header))

	// we then marshal the payload which is a map
	payloadstr, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error generating token")
		return string(payloadstr), err
	}

	payload64 := base64.StdEncoding.EncodeToString(payloadstr)

	// Now add the encoded string
	message := header64 + "." + payload64

	unsignedStr := header + string(payloadstr)

	h.Write([]byte(unsignedStr))
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	tokenStr := message + "." + signature
	return tokenStr, nil
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
