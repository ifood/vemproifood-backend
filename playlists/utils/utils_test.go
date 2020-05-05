package utils_test

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/bgildson/ifood_backend_challenge/base"
	"github.com/bgildson/ifood_backend_challenge/utils"
)

func TestParseKelvinToCelsius(t *testing.T) {
	kelvin := rand.Float64() * 300
	expected := kelvin - 273.15

	result := utils.ParseKelvinToCelsius(kelvin)

	require.Equal(t, expected, result)
}

func TestGenerateBasicAuthToken(t *testing.T) {
	user := "user"
	password := "password"
	expected := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", user, password)))

	result := utils.GenerateBasicAuthToken(user, password)

	require.Equal(t, string(expected), result)
}

func TestPrepareHTTPErrorMessage(t *testing.T) {
	errorMessage := "occur an error"

	result := utils.PrepareHTTPErrorMessage(errorMessage)

	var data map[string]interface{}
	err := json.Unmarshal([]byte(result), &data)

	require.Nil(t, err)

	require.Equal(t, errorMessage, data["message"])
}

func TestParseTemperatureToGenre(t *testing.T) {
	tests := []struct {
		value    float64
		expected base.Genre
	}{
		{value: 31, expected: base.GenreParty},
		{value: 30, expected: base.GenrePop},
		{value: 15, expected: base.GenrePop},
		{value: 14, expected: base.GenreRock},
		{value: 10, expected: base.GenreRock},
		{value: 9, expected: base.GenreClassical},
	}

	for _, test := range tests {
		result := utils.ParseTemperatureToGenre(test.value)
		require.Equal(t, test.expected, result)
	}
}
