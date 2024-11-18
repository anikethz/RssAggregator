package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Extract API Key from Headers
// Authorization: ApiKey {key}
func GetApiKey(headers http.Header) (string, error) {
	authorization := headers.Get("Authorization")
	if authorization == "" {
		return "", errors.New("No authentication info found")
	}

	values := strings.Split(authorization, " ")
	if len(values) != 2 {
		return "", errors.New("Malformed header authorization")
	}

	if values[0] != "ApiKey" {
		return "", errors.New("Malformed header authorization key first part")
	}

	return values[1], nil
}
