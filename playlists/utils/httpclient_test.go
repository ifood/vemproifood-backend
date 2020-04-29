package utils_test

import (
	"bytes"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/bgildson/ifood_backend_challenge/utils"
)

func TestHTTPClientMock(t *testing.T) {
	// certify that the Do really uses funcDo
	expected := strconv.FormatFloat(rand.Float64(), 'e', -1, 64)

	mock := utils.NewHTTPClientMock(func(*http.Request) (*http.Response, error) {
		return &http.Response{
			Body: ioutil.NopCloser(bytes.NewBuffer([]byte(expected))),
		}, nil
	})

	result, _ := mock.Do(nil)
	defer result.Body.Close()

	content, _ := ioutil.ReadAll(result.Body)

	require.Equal(t, expected, string(content))
}
