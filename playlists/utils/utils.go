package utils

import (
	"encoding/base64"
	"encoding/json"

	"github.com/bgildson/ifood_backend_challenge/base"
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

// PrepareHTTPErrorMessage : takes one message parameter and prepare to show as http error response text
func PrepareHTTPErrorMessage(message string) string {
	data := map[string]string{
		"message": message,
	}
	result, _ := json.Marshal(data)
	return string(result)
}

// ParseTemperatureToGenre : takes a temperature and return the properly genre
func ParseTemperatureToGenre(temperature float64) base.Genre {
	if temperature > 30 {
		return base.GenreParty
	} else if temperature >= 15 {
		return base.GenrePop
	} else if temperature >= 10 {
		return base.GenreRock
	} else {
		return base.GenreClassical
	}
}
