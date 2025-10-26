package auth

import (
	"errors"
	"net/http"
	"strings"
)

//GetAPIKey extracts and API key from the headers of an HTTP request
//Example:
// Authorization: APIKey(insert apiKey here)

func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no authentication info found ")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed auth header")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("malformed auth header")
	}

	return vals[1], nil
}
