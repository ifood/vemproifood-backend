package utils

import (
	"encoding/base64"
)

// ParseKelvinToCelsius : convert Kelvin to Celsius
func ParseKelvinToCelsius(kelvin float64) float64 {
	return kelvin - 273.15
}

// GenerateBasicAuthToken : generate a token to a basic auth schema
func GenerateBasicAuthToken(user, password string) string {
	preToken := user + ":" + password
	token := base64.StdEncoding.EncodeToString([]byte(preToken))
	return token
}
