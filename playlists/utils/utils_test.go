package utils_test

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"

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
